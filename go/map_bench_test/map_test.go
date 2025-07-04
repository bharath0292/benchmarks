package benchmarks

import (
	"strconv"
	"sync"
	"testing"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/puzpuzpuz/xsync/v4"
)

type Bar struct {
	Symbol     string
	Timestamp  int64
	Open       float64
	High       float64
	Low        float64
	Close      float64
	Volume     float64
	Indicators map[string]map[string]float64
}

var (
	numOps = 100000
	keys   = make([]string, numOps)
	value  = Bar{
		Symbol:    "AAPL",
		Timestamp: 1729231293,
		Open:      100,
		High:      120,
		Low:       95,
		Close:     110,
		Volume:    1_000_000,
		Indicators: map[string]map[string]float64{
			"MACD": {"value": 1.2},
		},
	}
)

func init() {
	for i := range numOps {
		keys[i] = "key" + strconv.Itoa(i)
	}
}

func BenchmarkSyncMap_Bar(b *testing.B) {
	var m sync.Map
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Store(keys[i%numOps], value)
			_, _ = m.Load(keys[i%numOps])
			i++
		}
	})
}

func BenchmarkMutexMap_Bar(b *testing.B) {
	m := make(map[string]Bar)
	var mu sync.RWMutex
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			mu.Lock()
			m[keys[i%numOps]] = value
			mu.Unlock()

			mu.RLock()
			_, _ = m[keys[i%numOps]]
			mu.RUnlock()
			i++
		}
	})
}

func BenchmarkCMap_Bar(b *testing.B) {
	m := cmap.New()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Set(keys[i%numOps], value)
			_, _ = m.Get(keys[i%numOps])
			i++
		}
	})
}

func BenchmarkXSyncMap_Bar(b *testing.B) {
	m := xsync.NewMap[string, Bar]()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Store(keys[i%numOps], value)
			_, _ = m.Load(keys[i%numOps])
			i++
		}
	})
}
