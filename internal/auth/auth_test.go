package auth

import (
	"net/http"
	"strings"
	"testing"
)


func TestGetApiKey(t *testing.T) {
	headerGood :=http.Header{}
	headerNoAuth := http.Header{}
	headerNoApiKey := http.Header{}

	headerGood.Set("Authorization", "ApiKey 123THATSMYKEY123")
	headerNoAuth.Set("NotAuthorization", "ApiKey 123THATSMYKEY123")
	headerNoApiKey.Set("Authorization", "NotApiKey 123THATSMYKEY123")

	res, err := GetAPIKey(headerGood)
	if res != "123THATSMYKEY123" || err != nil {
		t.Errorf("GetApiKey failed on happy path")
	}

	res, err = GetAPIKey(headerNoApiKey)
	if res != "" || !strings.Contains(err.Error(), "malformed") {
		t.Errorf("GetApiKey failed for malformed Auth header")
	}

	res, err = GetAPIKey(headerNoAuth)
	if res != "" || err != ErrNoAuthHeaderIncluded {
		t.Errorf("GetApiKey failed for missing authorization header")
	}
}