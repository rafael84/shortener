package service_test

import (
	"testing"

	"github.com/rafael84/shortener/persistence"
	"github.com/rafael84/shortener/service"
)

func TestCreate(t *testing.T) {
	service.Storage = persistence.NewMemory()
	// a  b  c  ba  bb  bc  ca  cb  cc  baa  bab
	// 0  1  2  3   4   5   6   7   8   9    10
	service.Alphabet = "abc"

	for _, tc := range []struct {
		Scenario string
		ReqURL   string
		ReqAlias string
		ResAlias string
		ResErr   string
	}{
		{
			Scenario: "Bad 1",
			ReqURL:   "",
			ReqAlias: "",
			ResErr:   "url is required",
		},
		{
			Scenario: "Bad 2",
			ReqURL:   "invalid",
			ReqAlias: "",
			ResAlias: "",
			ResErr:   "url is invalid",
		},
		{
			Scenario: "Bad 3",
			ReqURL:   " ",
			ReqAlias: "",
			ResAlias: "",
			ResErr:   "url is invalid",
		},
		{
			Scenario: "With Alias 1",
			ReqURL:   "http://valid.com",
			ReqAlias: "a1",
			ResAlias: "a1",
			ResErr:   "",
		},
		{
			Scenario: "With Alias 2",
			ReqURL:   "http://valid.com",
			ReqAlias: "a2",
			ResAlias: "a2",
			ResErr:   "",
		},
		{
			Scenario: "No Alias 1",
			ReqURL:   "http://valid.com",
			ReqAlias: "",
			ResAlias: "a",
			ResErr:   "",
		},
		{
			Scenario: "No Alias 2",
			ReqURL:   "http://valid.com",
			ReqAlias: "",
			ResAlias: "b",
			ResErr:   "",
		},
		{
			Scenario: "No Alias 3",
			ReqURL:   "http://valid.com",
			ReqAlias: "",
			ResAlias: "c",
			ResErr:   "",
		},
		{
			Scenario: "No Alias 4",
			ReqURL:   "http://valid.com",
			ReqAlias: "",
			ResAlias: "ba",
			ResErr:   "",
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			resErr := ""
			resAlias, err := service.Create(tc.ReqURL, tc.ReqAlias)
			if err != nil {
				resErr = err.Error()
			}
			if resAlias != tc.ResAlias {
				t.Errorf("unexpected alias\nwant:\t[%v]\ngot:\t[%v]", tc.ResAlias, resAlias)
			}
			if resErr != tc.ResErr {
				t.Errorf("unexpected error\nwant:\t[%v]\ngot:\t[%v]", tc.ResErr, resErr)
			}
		})
	}
}
