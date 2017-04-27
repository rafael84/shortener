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
