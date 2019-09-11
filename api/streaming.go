package api // import "github.com/SevereCloud/vksdk/api"

// StreamingGetServerURLResponse struct
type StreamingGetServerURLResponse struct {
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
}

// StreamingGetServerURL allows to receive data for the connection to Streaming API.
//
// https://vk.com/dev/streaming.getServerUrl
func (vk *VK) StreamingGetServerURL(params map[string]string) (response StreamingGetServerURLResponse, err error) {
	err = vk.RequestUnmarshal("streaming.getServerUrl", params, &response)
	return
}

// StreamingGetSettingsResponse struct
type StreamingGetSettingsResponse struct {
	MonthlyLimit string `json:"monthly_limit"`
}

// StreamingGetSettings allows to receive monthly tier for Streaming API.
//
// https://vk.com/dev/streaming.getSettings
func (vk *VK) StreamingGetSettings(params map[string]string) (response StreamingGetSettingsResponse, err error) {
	err = vk.RequestUnmarshal("streaming.getSettings", params, &response)
	return
}

// StreamingGetStatsResponse struct
type StreamingGetStatsResponse []struct {
	EventType string `json:"event_type"`
	Stats     []struct {
		Timestamp int `json:"timestamp"`
		Value     int `json:"value"`
	} `json:"stats"`
}

// StreamingGetStats allows to receive statistics for prepared and received events in Streaming API.
//
// https://vk.com/dev/streaming.getStats
func (vk *VK) StreamingGetStats(params map[string]string) (response StreamingGetStatsResponse, err error) {
	err = vk.RequestUnmarshal("streaming.getStats", params, &response)
	return
}

// StreamingGetStemResponse struct
type StreamingGetStemResponse struct {
	Stem string `json:"stem"`
}

// StreamingGetStem allows to receive the stem of the word.
//
// https://vk.com/dev/streaming.getStem
func (vk *VK) StreamingGetStem(params map[string]string) (response StreamingGetStemResponse, err error) {
	err = vk.RequestUnmarshal("streaming.getStem", params, &response)
	return
}

// StreamingSetSettings allows to set monthly tier for Streaming API.
//
// https://vk.com/dev/streaming.setSettings
func (vk *VK) StreamingSetSettings(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("streaming.setSettings", params, &response)

	return
}
