package api

import (
	"reflect"
	"testing"
)

func TestVK_StoriesDelete(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesDelete error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.StoriesDelete(tt.argParams); !reflect.DeepEqual(gotVkErr, tt.wantVkErr) {
				t.Errorf("VK.StoriesDelete() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_StoriesGet(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGet error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGet(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGet() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGet() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetExtended(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetExtended error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetExtended(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetBanned(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetBanned error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetBanned(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetBanned() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetBanned() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetBannedExtended(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetBannedExtended error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetBannedExtended(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetBannedExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetBannedExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetByID(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetByID error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetByID(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetByID() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetByID() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetByIDExtended(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetByIDExtended error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetByIDExtended(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetByIDExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetByIDExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetPhotoUploadServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetPhotoUploadServer error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetPhotoUploadServer(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetPhotoUploadServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetPhotoUploadServer() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetReplies(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetReplies error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetReplies(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetReplies() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetReplies() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetRepliesExtended(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetRepliesExtended error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetRepliesExtended(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetRepliesExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetRepliesExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetStats(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetStats error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetStats(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetStats() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetStats() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetVideoUploadServer(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetVideoUploadServer error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetVideoUploadServer(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetVideoUploadServer() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetVideoUploadServer() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetViewers(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetViewers error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetViewers(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetViewers() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetViewers() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesGetViewersExtended(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name         string
		argParams    map[string]string
		wantResponse GroupsAddLinkResponse
		wantVkErr    Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesGetViewersExtended error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, gotVkErr := vkGroup.StoriesGetViewersExtended(tt.argParams)
			if gotVkErr.Code != tt.wantVkErr.Code {
				t.Errorf("VK.StoriesGetViewersExtended() gotVkErr = %v, want %v", gotVkErr, tt.wantVkErr)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("VK.StoriesGetViewersExtended() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestVK_StoriesHideAllReplies(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesHideAllReplies error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.StoriesHideAllReplies(tt.argParams); !reflect.DeepEqual(gotVkErr, tt.wantVkErr) {
				t.Errorf("VK.StoriesHideAllReplies() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}

func TestVK_StoriesHideReply(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	tests := []struct {
		name      string
		argParams map[string]string
		wantVkErr Error
	}{
		// TODO: Add test cases.
		// {
		// 	name:      "StoriesHideReply error",
		// 	argParams: map[string]string{},
		// 	wantVkErr: Error{Code: 100},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, gotVkErr := vkGroup.StoriesHideReply(tt.argParams); !reflect.DeepEqual(gotVkErr, tt.wantVkErr) {
				t.Errorf("VK.StoriesHideReply() = %v, want %v", gotVkErr, tt.wantVkErr)
			}
		})
	}
}
