package service_test

import (
	"testing"

	"github.com/rafael84/shortener/service"
)

func TestRecover(t *testing.T) {
	if err := service.Storage.Set("f1", "http://valid.com"); err != nil {
		t.Error(err)
	}
	if err := service.Storage.Set("f2", "http://valid.com/2"); err != nil {
		t.Error(err)
	}
	if err := service.Storage.Set("f3", "http://valid.com/3"); err != nil {
		t.Error(err)
	}

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
			ReqAlias: "nf1",
			ResURL:   "",
			ResErr:   "alias not found",
		},
		{
			Scenario: "Not Found 2",
			ReqAlias: "nf2",
			ResURL:   "",
			ResErr:   "alias not found",
		},
		{
			Scenario: "Not Found 3",
			ReqAlias: "nf3",
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
				return
			}
			if resErr != tc.ResErr {
				t.Errorf("unexpected error\nwant:\t[%v]\ngot:\t[%v]", tc.ResErr, resErr)
				return
			}
		})
	}
}

func TestCreateRecover(t *testing.T) {
	for _, tc := range []struct {
		Scenario string
		Alias    string
		URL      string
	}{
		{
			Scenario: "Custom Alias 1",
			Alias:    "ca1",
			URL:      "http://custom.alias/1",
		},
		{
			Scenario: "Custom Alias 2",
			Alias:    "ca2",
			URL:      "http://custom.alias/2",
		},
		{
			Scenario: "Custom Alias 3",
			Alias:    "ca3",
			URL:      "http://custom.alias/3",
		},
		{
			Scenario: "Generated Alias 1",
			URL:      "http://generated.alias/1",
		},
		{
			Scenario: "Generated Alias 2",
			URL:      "http://generated.alias/2",
		},
		{
			Scenario: "Generated Alias 3",
			URL:      "http://generated.alias/3",
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			created, err := service.Create(tc.URL, tc.Alias)
			if err != nil {
				t.Error(err)
				return
			}
			recovered, err := service.Recover(created)
			if err != nil {
				t.Error(err)
				return
			}
			if recovered != tc.URL {
				t.Errorf("unexpected url\nwant:\t[%v]\ngot:\t[%v]", tc.URL, recovered)
				return
			}
		})
	}
}
