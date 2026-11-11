package main

import (
	"encoding/json/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerJSON(t *testing.T) {
	t.Setenv("GOENVECHO_TEST_VALUE", "<value>&\"\n")
	for _, pretty := range []string{"false", "true"} {
		t.Run(pretty, func(t *testing.T) {
			t.Setenv("PRETTY_PRINT", pretty)
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest(http.MethodGet, "/", nil))
			if w.Code != http.StatusOK || w.Header().Get("Content-Type") != "application/json" || w.Header().Get("X-App-Version") != Version {
				t.Fatalf("unexpected response: status=%d headers=%v", w.Code, w.Header())
			}
			var env map[string]string
			if err := json.Unmarshal(w.Body.Bytes(), &env); err != nil {
				t.Fatal(err)
			}
			if env["GOENVECHO_TEST_VALUE"] != "<value>&\"\n" {
				t.Fatalf("environment value = %q", env["GOENVECHO_TEST_VALUE"])
			}
			body := w.Body.String()
			if !strings.HasSuffix(body, "\n") || strings.Contains(body, "\n  ") != (pretty == "true") {
				t.Fatalf("unexpected formatting: %q", body)
			}
		})
	}
}

func TestHandlerRejectsInvalidUTF8(t *testing.T) {
	t.Setenv("GOENVECHO_INVALID_UTF8", string([]byte{0xff}))
	w := httptest.NewRecorder()
	handler(w, httptest.NewRequest(http.MethodGet, "/", nil))
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, body = %q", w.Code, w.Body.String())
	}
}
