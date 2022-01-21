package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T){
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	index(w,r)
	if w.Code != http.StatusOK{
		t.Errorf("Codigo de estado esperado %d se obtuvo: %d", http.StatusOK, w.Code)
	}
	defer r.Body.Close()
}
