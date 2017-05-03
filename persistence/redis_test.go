package persistence_test

import (
	"testing"

	"github.com/alicebob/miniredis"

	"github.com/rafael84/shortener/persistence"
)

func TestRedis(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	redis := persistence.NewRedis(s.Addr(), "", 0)

	type KeyValue struct {
		Key   string
		Value string
	}
	for _, tc := range []struct {
		Scenario  string
		Set       KeyValue
		Get       KeyValue
		Increment bool
		Count     int
		SetErr    string
	}{
		{
			Scenario:  "Set [A] 1 Get [A] 1",
			Set:       KeyValue{"A", "1"},
			Get:       KeyValue{"A", "1"},
			Increment: true,
			Count:     1,
		},
		{
			Scenario:  "Set [ ] 2 Get [ ] 2",
			Set:       KeyValue{" ", "2"},
			Get:       KeyValue{" ", "2"},
			Increment: true,
			Count:     2,
		},
		{
			Scenario:  "Set [C] 3 Get [D] ''",
			Set:       KeyValue{"C", "3"},
			Get:       KeyValue{"D", ""},
			Increment: true,
			Count:     3,
		},
		{
			Scenario: "Set [C] 4 Get [C] 4",
			Set:      KeyValue{"C", "4"},
			SetErr:   "could not set alias[C] url[4]",
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			// set
			setErr := ""
			err := redis.Set(tc.Set.Key, tc.Set.Value)
			if err != nil {
				setErr = err.Error()
			}
			if setErr != tc.SetErr {
				t.Fatal(setErr)
			}
			// increment
			if tc.Increment {
				if err := redis.Increment(); err != nil {
					t.Fatal(err)
				}
			}
			// get
			url, _ := redis.Get(tc.Get.Key)
			if url != tc.Get.Value {
				t.Fatalf("unexpected value\nwant\t[%v]\ngot\t[%v]",
					tc.Get.Value, url)
			}
			// count
			if tc.Count > 0 {
				count := redis.Count()
				if count != tc.Count {
					t.Fatalf("unexpected count\nwant\t[%v]\ngot\t[%v]",
						tc.Count, count)
				}
			}
		})
	}
}
