package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rafael84/shortener/handler"
	"github.com/rafael84/shortener/persistence"
	"github.com/rafael84/shortener/service"
)

func TestRecover(t *testing.T) {
	service.Storage = persistence.NewMemory()
	// a  b  c  ba  bb  bc  ca  cb  cc  baa  bab
	// 0  1  2  3   4   5   6   7   8   9    10
	service.Alphabet = "abc"

	if err := service.Storage.Set("v1", "http://valid.com"); err != nil {
		t.Error(err)
	}
	if err := service.Storage.Set("v2", "http://valid.com/2"); err != nil {
		t.Error(err)
	}
	if err := service.Storage.Set("v3", "http://valid.com/3"); err != nil {
		t.Error(err)
	}

	for _, tc := range []struct {
		Scenario     string
		GivenAlias   string
		WantStatus   int
		WantLocation string
		WantBody     string
	}{
		{
			Scenario:     "Valid 1",
			GivenAlias:   "v1",
			WantStatus:   302,
			WantLocation: "http://valid.com",
		},
		{
			Scenario:     "Valid 2",
			GivenAlias:   "v2",
			WantStatus:   302,
			WantLocation: "http://valid.com/2",
		},
		{
			Scenario:     "Valid 3",
			GivenAlias:   "v3",
			WantStatus:   302,
			WantLocation: "http://valid.com/3",
		},
		{
			Scenario:   "Invalid 1",
			GivenAlias: "inv1",
			WantStatus: 404,
			WantBody:   `{"err":"Not Found"}`,
		},
		{
			Scenario:   "Invalid 2",
			GivenAlias: "inv2",
			WantStatus: 404,
			WantBody:   `{"err":"Not Found"}`,
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			path := "/a/" + tc.GivenAlias
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			hr := http.HandlerFunc(handler.Recover)
			hr.ServeHTTP(rr, req)

			if rr.Code != tc.WantStatus {
				t.Fatalf("unexpected status code\nwant:\t[%d]\ngot:\t[%d]", tc.WantStatus, rr.Code)
			}

			if tc.WantLocation != "" {
				location := rr.Header().Get("Location")
				if location != tc.WantLocation {
					t.Fatalf("unexpected location\nwant:\t[%v]\ngot:\t[%v]", tc.WantLocation, location)
				}
			}

			contentType := rr.Header().Get("Content-Type")
			if !strings.HasPrefix(contentType, "application/json") {
				t.Fatalf("unexpected content-type\nwant:\t[%v]\ngot:\t[%v]", "application/json", contentType)
			}

			if tc.WantBody != "" {
				if rr.Body.String() != tc.WantBody {
					t.Fatalf("unexpected body\nwant:\t[%s]\ngot:\t[%s]", tc.WantBody, rr.Body.String())
				}
			}
		})
	}
}
