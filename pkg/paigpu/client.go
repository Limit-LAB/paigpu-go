package paigpu

import (
	"net/http"
)

const HeaderAppID = "X-Appid"
const HeaderNonce = "X-Nonce"
const HeaderTimestamp = "X-Timestamp"
const HeaderSignature = "X-Signature"

const DefaultBaseURL = "https://openapi.paigpu.com/openapi"

type Client struct {
	httpClient *http.Client
	baseURL    string
	appID      string
	appKey     string
}

func NewClient(appID string, appKey string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		baseURL:    DefaultBaseURL,
		appID:      appID,
		appKey:     appKey,
	}
}
