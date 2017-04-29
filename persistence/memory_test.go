package persistence

import "testing"

func TestMemory(t *testing.T) {
	type KeyValue struct {
		Key   string
		Value string
	}
	for _, tc := range []struct {
		Scenario string
		Set      KeyValue
		Get      KeyValue
	}{
		{
			Scenario: "Set [A] 1 Get [A] 1",
			Set:      KeyValue{"A", "1"},
			Get:      KeyValue{"A", "1"},
		},
		{
			Scenario: "Set [ ] 2 Get [ ] 2",
			Set:      KeyValue{" ", "2"},
			Get:      KeyValue{" ", "2"},
		},
		{
			Scenario: "Set [C] 3 Get [D] ''",
			Set:      KeyValue{"C", "3"},
			Get:      KeyValue{"D", ""},
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			memory := NewMemory()
			if err := memory.Set(tc.Set.Key, tc.Set.Value); err != nil {
				t.Error(err)
			}
			url, _ := memory.Get(tc.Get.Key)
			if url != tc.Get.Value {
				t.Errorf("unexpected value\nwant\t[%v]\ngot\t[%v]",
					tc.Get.Value, url)
			}
		})
	}
}
