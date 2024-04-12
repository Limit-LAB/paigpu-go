package paigpu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Storage struct {
	StorageID   string `json:"storageId"`
	StorageName string `json:"storageName"`
	StorageSize int    `json:"storageSize"`
	ClusterID   string `json:"clusterId"`
	ClusterName string `json:"clusterName"`
	Price       string `json:"price"`
}

type ListStoragesResponse struct {
	Data  []Storage `json:"data"`
	Total int       `json:"total"`
}

func (c *Client) ListStorages(ctx context.Context) (*ListStoragesResponse, error) {
	url := fmt.Sprintf("%s/v1/networkstorages/list", c.baseURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(HeaderAppID, c.appID)
	request.Header.Set(HeaderNonce, RandomNonce(16))
	timestamp := Timestamp()
	request.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signature := Signature("/openapi/v1/networkstorages/list", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
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
	result := ListStoragesResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateStorageRequest struct {
	ClusterID   string `json:"clusterId"`
	StorageName string `json:"storageName"`
	StorageSize int    `json:"storageSize"`
}

type CreateStorageResponse struct {
}

func (c *Client) CreateStorage(ctx context.Context,
	clusterID string,
	storageName string,
	storageSize int,
) (*CreateStorageResponse, error) {
	url := fmt.Sprintf("%s/v1/networkstorage/create", c.baseURL)

	requestBody := CreateStorageRequest{
		ClusterID:   clusterID,
		StorageName: storageName,
		StorageSize: storageSize,
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	request.Header.Set(HeaderAppID, c.appID)
	request.Header.Set(HeaderNonce, RandomNonce(16))
	timestamp := Timestamp()
	request.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signature := Signature("/openapi/v1/networkstorage/create", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
	request.Header.Set(HeaderSignature, signature)

	response, err := c.httpClient.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := CreateStorageResponse{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
