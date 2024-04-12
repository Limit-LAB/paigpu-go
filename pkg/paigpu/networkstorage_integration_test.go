package paigpu

import (
	"context"
	"os"
	"testing"
)

func TestClient_ListStorages(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	storages, err := client.ListStorages(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("storages", storages)
}

func TestClient_CreateStorage(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	response, err := client.CreateStorage(ctx, "22", "paigpu-go-integration-test", 22)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("response", response)
}
