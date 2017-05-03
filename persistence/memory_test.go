package persistence

import (
	"strconv"
	"testing"
)

func TestMemory(t *testing.T) {
	memory := NewMemory()
	type KeyValue struct {
		Key   string
		Value string
	}
	for _, tc := range []struct {
		Scenario string
		Set      KeyValue
		Get      KeyValue
		Count    int
	}{
		{
			Scenario: "Set [A] 1 Get [A] 1",
			Set:      KeyValue{"A", "1"},
			Get:      KeyValue{"A", "1"},
			Count:    1,
		},
		{
			Scenario: "Set [ ] 2 Get [ ] 2",
			Set:      KeyValue{" ", "2"},
			Get:      KeyValue{" ", "2"},
			Count:    2,
		},
		{
			Scenario: "Set [C] 3 Get [D] ''",
			Set:      KeyValue{"C", "3"},
			Get:      KeyValue{"D", ""},
			Count:    3,
		},
		{
			Scenario: "Set [C] 4 Get [C] 4",
			Set:      KeyValue{"C", "4"},
			Get:      KeyValue{"C", "4"},
			Count:    3,
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			if err := memory.Set(tc.Set.Key, tc.Set.Value); err != nil {
				t.Error(err)
			}
			url, _ := memory.Get(tc.Get.Key)
			if url != tc.Get.Value {
				t.Errorf("unexpected value\nwant\t[%v]\ngot\t[%v]",
					tc.Get.Value, url)
			}
			count := memory.Count()
			if count != tc.Count {
				t.Errorf("unexpected count\nwant\t[%v]\ngot\t[%v]",
					tc.Count, count)
			}
		})
	}
}

func benchmarkMemorySet(keyCount int, b *testing.B) {
	memory := NewMemory()
	aliases := []string{}
	for n := 0; n < keyCount; n++ {
		aliases = append(aliases, strconv.Itoa(n))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		memory.Set(aliases[n%keyCount], "url")
	}
}

func BenchmarkMemorySet1(b *testing.B)       { benchmarkMemorySet(1, b) }
func BenchmarkMemorySet10(b *testing.B)      { benchmarkMemorySet(10, b) }
func BenchmarkMemorySet100(b *testing.B)     { benchmarkMemorySet(100, b) }
func BenchmarkMemorySet1000(b *testing.B)    { benchmarkMemorySet(1000, b) }
func BenchmarkMemorySet10000(b *testing.B)   { benchmarkMemorySet(10000, b) }
func BenchmarkMemorySet100000(b *testing.B)  { benchmarkMemorySet(100000, b) }
func BenchmarkMemorySet1000000(b *testing.B) { benchmarkMemorySet(1000000, b) }

func benchmarkMemoryGet(keyCount int, b *testing.B) {
	memory := NewMemory()
	for n := 0; n < keyCount; n++ {
		memory.Set(strconv.Itoa(n), "url")
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		memory.Get(strconv.Itoa(n % keyCount))
	}
}

func BenchmarkMemoryGet1(b *testing.B)       { benchmarkMemoryGet(1, b) }
func BenchmarkMemoryGet10(b *testing.B)      { benchmarkMemoryGet(10, b) }
func BenchmarkMemoryGet100(b *testing.B)     { benchmarkMemoryGet(100, b) }
func BenchmarkMemoryGet1000(b *testing.B)    { benchmarkMemoryGet(1000, b) }
func BenchmarkMemoryGet10000(b *testing.B)   { benchmarkMemoryGet(10000, b) }
func BenchmarkMemoryGet100000(b *testing.B)  { benchmarkMemoryGet(100000, b) }
func BenchmarkMemoryGet1000000(b *testing.B) { benchmarkMemoryGet(1000000, b) }
