package service_test

import (
	"testing"

	"github.com/rafael84/shortener/service"
)

func TestRecover(t *testing.T) {
	for _, tc := range []struct {
		Scenario string
		ReqAlias string
		ResURL   string
		ResErr   string
	}{
		{
			Scenario: "Bad 1",
			ReqAlias: "",
			ResURL:   "",
			ResErr:   "invalid alias",
		},
		{
			Scenario: "Bad 2",
			ReqAlias: " ",
			ResURL:   "",
			ResErr:   "invalid alias",
		},
		{
			Scenario: "Not Found 1",
			ReqAlias: "a",
			ResURL:   "",
			ResErr:   "alias not found",
		},
		{
			Scenario: "Not Found 2",
			ReqAlias: "b",
			ResURL:   "",
			ResErr:   "alias not found",
		},
		{
			Scenario: "Not Found 3",
			ReqAlias: "c",
			ResURL:   "",
			ResErr:   "alias not found",
		},
		{
			Scenario: "Found 1",
			ReqAlias: "f1",
			ResURL:   "http://valid.com",
			ResErr:   "",
		},
		{
			Scenario: "Found 2",
			ReqAlias: "f2",
			ResURL:   "http://valid.com/2",
			ResErr:   "",
		},
		{
			Scenario: "Found 3",
			ReqAlias: "f3",
			ResURL:   "http://valid.com/3",
			ResErr:   "",
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			resErr := ""
			resURL, err := service.Recover(tc.ReqAlias)
			if err != nil {
				resErr = err.Error()
			}
			if resURL != tc.ResURL {
				t.Errorf("unexpected url\nwant:\t[%v]\ngot:\t[%v]", tc.ResURL, resURL)
			}
			if resErr != tc.ResErr {
				t.Errorf("unexpected error\nwant:\t[%v]\ngot:\t[%v]", tc.ResErr, resErr)
			}
		})
	}
}
