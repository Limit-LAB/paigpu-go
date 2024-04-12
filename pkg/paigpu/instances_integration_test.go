package paigpu

import (
	"context"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestClient_Instances(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	instances, err := client.Instances(ctx, "", 100, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("instances", instances)
}

func TestClient_Instance(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	instance, err := client.Instance(ctx, os.Getenv("PAIGPU_INSTANCE_ID"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("instance", instance)
}

func TestClient_StartInstance(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	result, err := client.StartInstance(ctx, os.Getenv("PAIGPU_INSTANCE_ID"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("start instance", result)
}

func TestClient_StopInstance(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	result, err := client.StopInstance(ctx, os.Getenv("PAIGPU_INSTANCE_ID"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("stop instance", result)
}

func TestClient_CreateInstance(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	result, err := client.CreateInstance(ctx,
		"paigpu-go integration test",
		"23",
		"75",
		1,
		10,
		"afterusage",
		"image.ppio.cloud/prod-gpucloudpublic/cuda:v11.8",
		"",
		"",
		"",
		[]int{8888},
		[]Env{},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("create instance", result)
}
