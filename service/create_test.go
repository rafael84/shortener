package service_test

import (
	"testing"

	"github.com/rafael84/shortener/service"
)

func TestCreate(t *testing.T) {
	for _, tc := range []struct {
		Scenario string
		ReqURL   string
		ReqAlias string
		ResAlias string
		ResErr   string
	}{
		{
			Scenario: "Empty URL",
			ReqURL:   "",
			ReqAlias: "",
			ResErr:   "url is required",
		},
		{
			Scenario: "URL is Incomplete",
			ReqURL:   "invalid",
			ReqAlias: "",
			ResAlias: "",
			ResErr:   "url is invalid",
		},
		{
			Scenario: "URL is Whitespace",
			ReqURL:   " ",
			ReqAlias: "",
			ResAlias: "",
			ResErr:   "url is invalid",
		},
		{
			Scenario: "URL is Valid",
			ReqURL:   "http://valid.com",
			ReqAlias: "a",
			ResAlias: "a",
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
