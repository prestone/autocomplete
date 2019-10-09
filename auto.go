package autocomplete

import (
	"bytes"
	"sync"

	pool "github.com/valyala/bytebufferpool"
)

func New() (a *DB) {
	a = &DB{
		index: make([]map[int][]*item, shards),
		items: make([][]*item, shardsid),
	}
	return
}

type DB struct {
	sync.RWMutex
	index []map[int][]*item // shard, word hash, list
	items [][]*item         // shard, list
}

func (a *DB) Delete(id int) {
	shard := id % shardsid
	var p int
	for _, i := range a.items[shard] {
		if id == i.id {
			for word, _ := range i.words {
				var pos int
				for _, i := range a.index[word%shards][word] {
					if i.id == id {
						a.index[word%shards][word] = append(a.index[word%shards][word][:pos], a.index[word%shards][word][pos+1:]...)
						if pos > 0 {
							pos = pos - 1
						}
						continue
					}
					pos++
				}
			}
			a.items[shard] = append(a.items[shard][:p], a.items[shard][p+1:]...)
		}
	}
}

func (a *DB) Add(id int, text string) {
	a.Lock()
	a.Delete(id)
	i := &item{id, make(map[int]bool)}
	vid := id % shards
	a.items[vid] = append(a.items[vid], i)
	var count uint8 = 1
	for _, word := range bytes.Fields(clear(text)) {
		b := pool.Get()
		for _, letter := range word {
			b.WriteByte(letter)
			a.set(count, b.Bytes(), i)
		}
		pool.Put(b)
		count++
	}
	a.Unlock()
}

func (a *DB) Search(limit int, q string) (res []int) {

	//clear
	query := clear(q)

	//divide
	words := bytes.Fields(query)

	//count
	wordslen := len(words)

	switch wordslen {
	case 0:
		//empty
		return
	case 1:
		//one word
		h := hash(words[0])
		for _, i := range a.index[h%shards][h] {
			res = append(res, i.id)
			if len(res) == limit {
				return
			}
		}
	default:
		var hh []int
		var lens int
		var bestpos int
		for pos, word := range words {
			h := hash(word)
			hh = append(hh, h)
			l := len(a.index[h%shards][h])
			if lens == 0 {
				lens = l
			}
			if l < lens {
				lens = l
				bestpos = pos
			}
		}

		best := hh[bestpos]
		bestshard := best % shards
		for _, i := range a.index[bestshard][best] {
			if !i.has(hh) {
				continue
			}
			res = append(res, i.id)
			if len(res) == limit {
				break
			}
		}
	}

	return
}

func (a *DB) set(count uint8, key []byte, i *item) {
	h := hash(key)
	shard := h % shards
	if a.index[shard] == nil {
		a.index[shard] = make(map[int][]*item)
	}
	a.index[shard][h] = append(a.index[shard][h], i)
	i.words[h] = true
}
