package paigpu

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	CPUPerGpu       int    `json:"cpuPerGpu"`
	MemoryPerGpu    int    `json:"memoryPerGpu"`
	DiskPerGpu      int    `json:"diskPerGpu"`
	AvailableDeploy bool   `json:"availableDeploy"`
	Prices          []struct {
		BillingMode string `json:"billingMode"`
		Price       string `json:"price"`
	} `json:"prices"`
}

type ProductsResponse struct {
	Data []Product `json:"data"`
}

func (c *Client) Products(
	ctx context.Context,
	clusterID string,
	productName string,
) (*ProductsResponse, error) {
	url := fmt.Sprintf("%s/v1/products", c.baseURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(HeaderAppID, c.appID)
	request.Header.Set(HeaderNonce, RandomNonce(16))
	timestamp := Timestamp()
	request.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signature := Signature("/openapi/v1/products", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
	request.Header.Set(HeaderSignature, signature)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result := ProductsResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
