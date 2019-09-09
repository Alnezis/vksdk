package api

import (
	"testing"
)

func TestVK_StatusGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.StatusGet(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.StatusGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_StatusSet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.StatusSet(map[string]string{
		"text": "Hello world",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.StatusSet() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
