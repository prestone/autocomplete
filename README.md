# 🐸 Autocomplete
This is simple memroy based autocomplete for your data. It use hash, shards and int ID and it do not store data. It is useful for lookup lists and so on. Its very fast with O(1) access. I made it because some times I need a real simple autocomplete but there is no it.

### Install
```go get github.com/prestone/autocomplete```

### Example
```go
//create a simple search engine
search := autocomplete.New()

//add some data for index
search.Add(1, "text I love")
search.Add(2, "text engine")
search.Add(3, "text nice")
search.Add(4, "xt")

//search
search.Search(10, "te") // [1,2,3]
search.Search(10, "xt") // [4]
```

### Benchmarks
```
BenchmarkSearch_100k_Limit_10-8          2356702               445 ns/op             344 B/op          7 allocs/op
BenchmarkSearch_100k_Limit_100-8         1045707              1114 ns/op            2136 B/op         10 allocs/op
BenchmarkSearch_1m_Limit_10-8            2309048               493 ns/op             344 B/op          7 allocs/op
BenchmarkSearch_1m_Limit_100-8           1000000              1462 ns/op            2136 B/op         10 allocs/op
```

### Thanks
