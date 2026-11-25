package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSprint1MVP(t *testing.T) {
	h := newMVPMux(&MVPStore{})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	rec2 := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/tickets", bytes.NewReader([]byte(`{"title":"MVP测试工单","severity":"P1"}`)))
	h.ServeHTTP(rec2, req)
	if rec2.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec2.Code)
	}
}
