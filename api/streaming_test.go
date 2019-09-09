package api

import (
	"testing"
)

func TestVK_StreamingGetServerURL(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	t.Run("StreamingGetServerURL not empty", func(t *testing.T) {
		response, vkErr := vkService.StreamingGetServerURL(map[string]string{})
		if vkErr.Code != 0 {
			t.Errorf("%d %s", vkErr.Code, vkErr.Message)
		}

		if response.Endpoint == "" {
			t.Error("Empty endpoint field")
		}
		if response.Key == "" {
			t.Error("Empty key field")
		}
	})
}

func TestVK_StreamingGetSettings(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	t.Run("StreamingGetSettings not empty", func(t *testing.T) {
		response, vkErr := vkService.StreamingGetSettings(map[string]string{})
		if vkErr.Code != 0 {
			t.Errorf("%d %s", vkErr.Code, vkErr.Message)
		}

		if response.MonthlyLimit == "" {
			t.Error("Empty monthly_limit field")
		}
	})
}

func TestVK_StreamingGetStats(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := make(map[string]string)
	params["type"] = "received"
	params["interval"] = "1h"

	t.Run("StreamingGetStats not empty", func(t *testing.T) {
		// TODO: check it
		_, vkErr := vkService.StreamingGetStats(params)
		if vkErr.Code != 0 {
			t.Errorf("%d %s", vkErr.Code, vkErr.Message)
		}

		// if len(response) == 0 {
		// 	t.Error("Empty array")
		// }
	})
}

func TestVK_StreamingGetStem(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := make(map[string]string)
	params["word"] = "собака"
	t.Run("StreamingGetStem not empty", func(t *testing.T) {
		response, vkErr := vkService.StreamingGetStem(params)
		if vkErr.Code != 0 {
			t.Errorf("%d %s", vkErr.Code, vkErr.Message)
		}

		if response.Stem != "собак" {
			t.Error("Sterm wrong")
		}
	})
}

func TestVK_StreamingSetSettings(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := make(map[string]string)
	params["monthly_tier"] = "unlimited"
	t.Run("StreamingSetSettings not empty", func(t *testing.T) {
		// TODO: check StreamingSetSettings test
		_, vkErr := vkService.StreamingSetSettings(params)
		if vkErr.Code != 0 {
			t.Errorf("%d %s", vkErr.Code, vkErr.Message)
		}
	})
}

func TestVK_StreamingError(t *testing.T) {
	vk := Init("")

	t.Run("StreamingGetServerURL error", func(t *testing.T) {
		_, vkErr := vk.StreamingGetServerURL(map[string]string{})
		if vkErr.Code != 5 {
			t.Errorf("StreamingGetServerURL error bad %d %s", vkErr.Code, vkErr.Message)
		}
	})

	t.Run("StreamingGetSettings error", func(t *testing.T) {
		_, vkErr := vk.StreamingGetSettings(map[string]string{})
		if vkErr.Code != 5 {
			t.Errorf("StreamingGetSettings error bad %d %s", vkErr.Code, vkErr.Message)
		}
	})

	t.Run("StreamingGetStats error", func(t *testing.T) {
		_, vkErr := vk.StreamingGetStats(map[string]string{})
		if vkErr.Code != 5 {
			t.Errorf("StreamingGetStats error bad %d %s", vkErr.Code, vkErr.Message)
		}
	})

	t.Run("StreamingGetStem error", func(t *testing.T) {
		_, vkErr := vk.StreamingGetStem(map[string]string{})
		if vkErr.Code != 5 {
			t.Errorf("StreamingGetStem error bad %d %s", vkErr.Code, vkErr.Message)
		}
	})

	t.Run("StreamingSetSettings error", func(t *testing.T) {
		_, vkErr := vk.StreamingSetSettings(map[string]string{})
		if vkErr.Code != 5 {
			t.Errorf("StreamingSetSettings error bad %d %s", vkErr.Code, vkErr.Message)
		}
	})
}
