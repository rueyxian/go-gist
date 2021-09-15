package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	helloHandler(recorder, req)
	// hf := http.HandlerFunc(handler)
	// hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello World!"
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

// ================================================================================

func TestRouter(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	res, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code: got %v want %v", res.StatusCode, http.StatusOK)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello World!"
	actual := string(b)

	if actual != expected {
		t.Errorf("response body: got %v want %v", actual, expected)
	}

}

// ================================================================================

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	res, err := http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status code: got %v want %v", res.StatusCode, http.StatusMethodNotAllowed)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := ""
	actual := string(b)

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

// ================================================================================

func TestStaticFileServer(t *testing.T) {

	r := newRouter()
	mockServer := httptest.NewServer(r)

	res, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code: got %v want %v", res.StatusCode, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	actualContentType := res.Header.Get("Content-Type")

	if actualContentType != expectedContentType {
		t.Errorf("Content-Type: got %v want %v", actualContentType, expectedContentType)
	}

}
