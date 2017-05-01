package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rafael84/shortener/handler"
)

func TestCreate(t *testing.T) {
	for _, tc := range []struct {
		Scenario    string
		GivenParams url.Values
		WantStatus  int
		WantBody    string
	}{
		{
			Scenario:    "No Alias 1",
			GivenParams: url.Values{"url": {"http://valid.com"}},
			WantStatus:  200,
			WantBody:    `{"alias":"a","url":"http://valid.com","timeTaken":"1ms"}`,
		},
		{
			Scenario:    "No Alias 2",
			GivenParams: url.Values{"url": {"http://valid.com"}},
			WantStatus:  200,
			WantBody:    `{"alias":"b","url":"http://valid.com","timeTaken":"1ms"}`,
		},
		{
			Scenario:    "No Alias 3",
			GivenParams: url.Values{"url": {"http://valid.com/3"}},
			WantStatus:  200,
			WantBody:    `{"alias":"c","url":"http://valid.com/3","timeTaken":"1ms"}`,
		},
		{
			Scenario: "Custom Alias 1",
			GivenParams: url.Values{
				"url":   {"http://custom.alias/1"},
				"alias": {"ca1"},
			},
			WantStatus: 200,
			WantBody:   `{"alias":"ca1","url":"http://custom.alias/1","timeTaken":"1ms"}`,
		},
		{
			Scenario: "Custom Alias 2",
			GivenParams: url.Values{
				"url":   {"http://custom.alias/2"},
				"alias": {"ca2"},
			},
			WantStatus: 200,
			WantBody:   `{"alias":"ca2","url":"http://custom.alias/2","timeTaken":"1ms"}`,
		},
		{
			Scenario:    "Invalid 1",
			GivenParams: url.Values{},
			WantStatus:  200,
			WantBody:    `{"err":"url is required"}`,
		},
		{
			Scenario: "Invalid 2",
			GivenParams: url.Values{
				"url": {"invalid"},
			},
			WantStatus: 200,
			WantBody:   `{"err":"url is invalid"}`,
		},
		{
			Scenario: "Alias Exist 1",
			GivenParams: url.Values{
				"url":   {"http://valid.com"},
				"alias": {"ca1"},
			},
			WantStatus: 400,
			WantBody:   `{"err":"alias already taken"}`,
		},
		{
			Scenario: "Alias Exist 2",
			GivenParams: url.Values{
				"url":   {"http://valid.com"},
				"alias": {"ca2"},
			},
			WantStatus: 400,
			WantBody:   `{"err":"alias already taken"}`,
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			path := "/create?" + tc.GivenParams.Encode()
			req, err := http.NewRequest("PUT", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			hc := http.HandlerFunc(handler.Create)

			hc.ServeHTTP(rr, req)

			if rr.Code != tc.WantStatus {
				t.Errorf("unexpected status code\nwant:\t[%d]\ngot:\t[%d]", tc.WantStatus, rr.Code)
			}

			if rr.Body.String() != tc.WantBody {
				t.Errorf("unexpected body\nwant:\t[%s]\ngot:\t[%s]", tc.WantBody, rr.Body.String())
			}
		})
	}
}
