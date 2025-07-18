package auth

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	originalKey := "secreto123"
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", originalKey))

	answKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("GetAPIKey devolvió un error inesperado: %v", err)
	}
	if answKey != originalKey {
		t.Errorf("La clave extraída fue '%s', pero se esperaba '%s'", answKey, originalKey)
	}
}
