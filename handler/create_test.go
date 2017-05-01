package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/rafael84/shortener/handler"
	"github.com/rafael84/shortener/persistence"
	"github.com/rafael84/shortener/service"
)

func TestCreate(t *testing.T) {
	service.Storage = persistence.NewMemory()
	// a  b  c  ba  bb  bc  ca  cb  cc  baa  bab
	// 0  1  2  3   4   5   6   7   8   9    10
	service.Alphabet = "abc"

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
				t.Fatalf("unexpected status code\nwant:\t[%d]\ngot:\t[%d]", tc.WantStatus, rr.Code)
			}

			contentType := rr.Header().Get("Content-Type")
			if !strings.HasPrefix(contentType, "application/json") {
				t.Fatalf("unexpected content-type\nwant:\t[%v]\ngot:\t[%v]", "application/json", contentType)
			}

			if rr.Body.String() != tc.WantBody {
				t.Fatalf("unexpected body\nwant:\t[%s]\ngot:\t[%s]", tc.WantBody, rr.Body.String())
			}

		})
	}
}
