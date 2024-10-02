package environmentpathappender_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/portofrotterdam/environmentpathappender"
)

func TestMissingEnvVariable(t *testing.T) {
	envVar := "TEST_ENV"
	requestUrl := "https://test.com/ship"
	expectedPath := "/ship"

	cfg := environmentpathappender.CreateConfig()
	cfg.Env = envVar

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := environmentpathappender.New(ctx, next, cfg, "environmentpathappender")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, req)

	assertRequestPathAdded(t, req, expectedPath)
}

func TestMissingEnvParamater(t *testing.T) {
	envVar := "TEST_ENV"
	envValue := "FOO"

	err := os.Setenv(envVar, envValue)
	if err != nil {
		t.Fatal(err)
	}

	cfg := environmentpathappender.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	_, err = environmentpathappender.New(ctx, next, cfg, "environmentpathappender")
	if err == nil {
		t.Fatal(err)
	}
}

func TestNoSlashAllowed(t *testing.T) {
	envVar := "TEST_ENV"
	envValue := "FO/O"

	err := os.Setenv(envVar, envValue)
	if err != nil {
		t.Fatal(err)
	}

	cfg := environmentpathappender.CreateConfig()
	cfg.Env = envVar

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	_, err = environmentpathappender.New(ctx, next, cfg, "environmentpathappender")
	if err == nil {
		t.Fatal(err)
	}
}

func TestEnvironmentPathAppend(t *testing.T) {
	envVar := "TEST_ENV"
	envValue := "FOO"
	requestUrl := "https://test.com/ship"
	expectedPath := "/ship/FOO"

	err := os.Setenv(envVar, envValue)
	if err != nil {
		t.Fatal(err)
	}

	cfg := environmentpathappender.CreateConfig()
	cfg.Env = envVar

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := environmentpathappender.New(ctx, next, cfg, "environmentpathappender")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, req)

	assertRequestPathAdded(t, req, expectedPath)
}

func assertRequestPathAdded(t *testing.T, req *http.Request, expectedValue string) {
	t.Helper()

	if req.URL.Path != expectedValue {
		t.Errorf("Expected path to be '%s', but was '%s'", expectedValue, req.URL.Path)
	}
}
