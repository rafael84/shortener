package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafael84/shortener/handler"
	"github.com/rafael84/shortener/persistence"
	"github.com/rafael84/shortener/service"
)

func TestRecover(t *testing.T) {
	service.Storage = persistence.NewMemory()

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
			WantBody:   "Not Found\n",
		},
		{
			Scenario:   "Invalid 2",
			GivenAlias: "inv2",
			WantStatus: 404,
			WantBody:   "Not Found\n",
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

			if tc.WantBody != "" {
				if rr.Body.String() != tc.WantBody {
					t.Fatalf("unexpected body\nwant:\t[%s]\ngot:\t[%s]", tc.WantBody, rr.Body.String())
				}
			}
		})
	}
}
