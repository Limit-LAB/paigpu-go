package paigpu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Signature(path string, appID string, appKey string, nonce string, timestamp int64) string {
	data := fmt.Sprintf("%s%s%d%s", appID, nonce, timestamp, path)
	hash := hmac.New(sha256.New, []byte(appKey))
	hash.Write([]byte(data))
	signature := hex.EncodeToString(hash.Sum(nil))
	return signature
}

func RandomNonce(n int) string {
	b := make([]byte, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

func Timestamp() int64 {
	return time.Now().UnixMilli()
}
