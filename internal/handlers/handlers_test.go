package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{
		name:               "home",
		url:                "/",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "about",
		url:                "/about",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "gq",
		url:                "/generals-quarters",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "ms",
		url:                "/majors-suite",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "sa",
		url:                "/search-availability",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "contact",
		url:                "/contact",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "reservation",
		url:                "/make-reservation",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "post-search-availability",
		url:    "/search-availability",
		method: "POST",
		params: []postData{
			{
				key:   "start",
				value: "2020-01-01",
			},
			{
				key:   "end",
				value: "2022-01-02",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "post-search-availability-json",
		url:    "/search-availability-json",
		method: "POST",
		params: []postData{
			{
				key:   "start",
				value: "2020-01-01",
			},
			{
				key:   "end",
				value: "2022-01-02",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "make reservation post",
		url:    "/make-reservation",
		method: "POST",
		params: []postData{
			{
				key:   "first_name",
				value: "John",
			},
			{
				key:   "last_name",
				value: "Smith",
			},
			{
				key:   "email",
				value: "me@here.com",
			},
			{
				key:   "phone",
				value: "555-555-5555",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		} else {

			values := url.Values{}
			for _, p := range e.params {
				values.Add(p.key, p.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		}
	}
}
