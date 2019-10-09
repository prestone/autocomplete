package autocomplete

type item struct {
	id    int
	words map[int]bool
}

func (a *item) has(words []int) bool {
	for _, i := range words {
		if !a.words[i] {
			return false
		}
	}
	return true
}
