package paigpu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Instance struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	NodeID        string `json:"nodeId"`
	NodeName      string `json:"nodeName"`
	Region        string `json:"region"`
	RegionName    string `json:"regionName"`
	Status        string `json:"status"`
	SSHCommand    string `json:"sshCommand"`
	GPUType       string `json:"GPUType"`
	RootPassword  string `json:"rootPassword"`
	ImageName     string `json:"imageName"`
	CPULimit      string `json:"cpuLimit"`
	MemLimit      string `json:"memLimit"`
	GPULimit      string `json:"GPULimit"`
	CreatedAt     string `json:"createdAt"`
	LastStartedAt string `json:"LastStartedAt"`
	LastStoppedAt string `json:"LastStoppedAt"`
	UseTime       string `json:"useTime"`
	PortMappings  []struct {
		Port     int    `json:"port"`
		Endpoint string `json:"endpoint"`
		Type     string `json:"type"`
	} `json:"portMappings"`
	BillingType string `json:"billingType"`
	NodeGPUNum  string `json:"nodeGPUNum"`
	NodeGPUFree string `json:"nodeGPUFree"`
	NodeStatus  string `json:"nodeStatus"`
	DataDisk    struct {
		Limit  string `json:"limit"`
		Sum    string `json:"sum"`
		Status string `json:"status"`
	} `json:"dataDisk"`
	JupyterAddress   string `json:"jupyterAddress"`
	ExpiredAt        string `json:"expiredAt"`
	NetworkStorageID string `json:"networkStorageId"`
	ProductID        string `json:"productId"`
	StatusError      struct {
		State   string `json:"state"`
		Message string `json:"message"`
	} `json:"statusError"`

	Envs []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"envs"`
	DiskSize int `json:"diskSize"`
}

type InstancesResponse struct {
	Instances []Instance `json:"instances"`
	PageSize  int        `json:"pageSize"`
	PageNum   int        `json:"pageNum"`
	Total     int        `json:"total"`
}

type InstanceResponse struct {
	Instance
}

func (c *Client) Instances(ctx context.Context, name string, pageSize int, pageNumber int) (*InstancesResponse, error) {
	url := fmt.Sprintf("%s/v1/gpu/instances", c.baseURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(HeaderAppID, c.appID)
	request.Header.Set(HeaderNonce, RandomNonce(16))
	timestamp := Timestamp()
	request.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signature := Signature("/openapi/v1/gpu/instances", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
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
	result := InstancesResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) Instance(ctx context.Context, instanceID string) (*InstanceResponse, error) {
	url := fmt.Sprintf("%s/v1/gpu/instance?instanceId="+instanceID, c.baseURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(HeaderAppID, c.appID)
	request.Header.Set(HeaderNonce, RandomNonce(16))
	timestamp := Timestamp()
	request.Header.Set(HeaderTimestamp, fmt.Sprintf("%d", timestamp))
	signature := Signature("/openapi/v1/gpu/instance", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
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
	result := InstanceResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type StartInstanceRequest struct {
	InstanceID string `json:"instanceId"`
}

type StartInstanceResponse struct {
}

func (c *Client) StartInstance(ctx context.Context, instanceID string) (*StartInstanceResponse, error) {
	url := fmt.Sprintf("%s/v1/gpu/instance/start", c.baseURL)

	requestBody := StartInstanceRequest{
		InstanceID: instanceID,
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
	signature := Signature("/openapi/v1/gpu/instance/start", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
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
	result := StartInstanceResponse{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type StopInstanceRequest struct {
	InstanceID string `json:"instanceId"`
}

type StopInstanceResponse struct{}

func (c *Client) StopInstance(ctx context.Context, instanceID string) (*StopInstanceResponse, error) {
	url := fmt.Sprintf("%s/v1/gpu/instance/stop", c.baseURL)

	requestBody := StopInstanceRequest{
		InstanceID: instanceID,
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
	signature := Signature("/openapi/v1/gpu/instance/stop", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
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
	result := StopInstanceResponse{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type Env struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type CreateInstanceRequest struct {
	Name                     string `json:"name"`
	ProductId                string `json:"productId"`
	GpuNum                   int    `json:"gpuNum"`
	DiskSize                 int    `json:"diskSize"`
	BillingMode              string `json:"billingMode"`
	Duration                 int    `json:"duration"`
	ImageUrl                 string `json:"imageUrl"`
	ImageAuth                string `json:"imageAuth"`
	Ports                    string `json:"ports"`
	Envs                     []Env  `json:"envs"`
	Command                  string `json:"command"`
	ClusterId                string `json:"clusterId"`
	NetworkStorageId         string `json:"networkStorageId"`
	LocalStorageMountPoint   string `json:"localStorageMountPoint"`
	NetworkStorageMountPoint string `json:"networkStorageMountPoint"`
}

type CreateInstanceResponse struct {
	ID string `json:"id"`
}

func (c *Client) CreateInstance(ctx context.Context,
	name string,
	productId string,
	gpuNum int,
	diskSize int,
	billingMode string,
	imageUrl string,
	networkStorageID string,
	ports []int,
	envs []Env,
) (*CreateInstanceResponse, error) {
	url := fmt.Sprintf("%s/v1/gpu/instance/create", c.baseURL)

	var portsString []string
	for _, port := range ports {
		portsString = append(portsString, fmt.Sprintf("%d", port))
	}
	requestBody := CreateInstanceRequest{
		Name:             name,
		ProductId:        productId,
		GpuNum:           gpuNum,
		DiskSize:         diskSize,
		BillingMode:      billingMode,
		ImageUrl:         imageUrl,
		Ports:            strings.Join(portsString, ","),
		Envs:             envs,
		NetworkStorageId: networkStorageID,
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
	signature := Signature("/openapi/v1/gpu/instance/create", c.appID, c.appKey, request.Header.Get(HeaderNonce), timestamp)
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
	result := CreateInstanceResponse{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
