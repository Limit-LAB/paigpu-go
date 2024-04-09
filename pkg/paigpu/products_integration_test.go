package paigpu

import (
	"context"
	"os"
	"testing"
)
import _ "github.com/joho/godotenv/autoload"

func TestClient_Products(t *testing.T) {
	ctx := context.Background()
	client := NewClient(os.Getenv("PAIGPU_APP_ID"), os.Getenv("PAIGPU_APP_SECRET"))
	products, err := client.Products(ctx, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("products", products)
}
