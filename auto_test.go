package autocomplete

//testing

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/monopolly/console"
)

func TestNew(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	var log = console.New()
	log.OK(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	s := New()
	s.Add(1, "text")
	s.Add(2, "text")
	s.Add(3, "text")
	s.Add(4, "xt")

	log.Info(s.Search(10, "te"))
	log.Info(s.Search(10, "xt"))

}

func BenchmarkSearch_100k_Limit_10(b *testing.B) {
	s := New()
	//100k
	for x := 1; x < 100001; x++ {
		s.Add(x, fmt.Sprintf("text nice %d", x))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Search(10, "te")
	}
}

func BenchmarkSearch_100k_Limit_100(b *testing.B) {
	s := New()
	//100k
	for x := 1; x < 100001; x++ {
		s.Add(x, fmt.Sprintf("text nice %d", x))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Search(100, "te")
	}
}

func BenchmarkSearch_1m_Limit_10(b *testing.B) {
	s := New()
	//100k
	for x := 1; x < 1000001; x++ {
		s.Add(x, fmt.Sprintf("text nice %d", x))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Search(10, "te")
	}
}

func BenchmarkSearch_1m_Limit_100(b *testing.B) {
	s := New()
	//100k
	for x := 1; x < 1000001; x++ {
		s.Add(x, fmt.Sprintf("text nice %d", x))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Search(100, "te")
	}
}
