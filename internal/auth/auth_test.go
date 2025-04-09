package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	key, err := GetAPIKey(h)
	if key != "" || err == nil {
		t.Errorf("Does not return '' key or non-nil err for empty header. Returns: key - '%v', err - '%v'", key, err)
	}

	if err.Error() != "no authorization header included" {
		t.Errorf("Does not return 'no auth' message for empty header. Returns: '%v'", err.Error())
	}

	h.Add("Authorization", "ApiKey")
	key, err = GetAPIKey(h)
	if key != "" || err == nil {
		t.Errorf("Does not return '' key or non-nil err for malformed header. Returns: key - '%v', err - '%v'", key, err)
	}

	h.Del("Authorization")
	h.Add("Authorization", "ApiKey 1234567890")
	key, err = GetAPIKey(h)
	if key != "1234567890" {
		t.Errorf("Does not return key. Returns: '%v'", key)
	}
}
