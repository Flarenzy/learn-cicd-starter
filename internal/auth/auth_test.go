package auth

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoApiKey(t *testing.T) {
	t.Parallel()
	panic("implement me")
	req := httptest.NewRequest(http.MethodGet, "/example", nil)
	req.Header.Set("Authorization", "")
	_, err := GetAPIKey(req.Header)
	if err == nil {
		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Fatal("expected ErrNoAuthHeaderIncluded, got ", "err", err.Error())
		}
	}
}

func TestApiKey(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/example", nil)
	randomUUID, _ := uuid.NewRandom()
	req.Header.Set("Authorization", "ApiKey "+randomUUID.String())
	id, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatal(err)
	}
	if id != randomUUID.String() {
		t.Fatal("expected uuid, got ", randomUUID, id)
	}
}
