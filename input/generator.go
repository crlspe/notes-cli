package input

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateId() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 8)
	if _, err := rand.Read(randomBytes); err != nil {
		return fmt.Sprintf("%d", timestamp)
	}
	randomHex := hex.EncodeToString(randomBytes)
	id := fmt.Sprintf("%d%s", timestamp, randomHex)
	return id
}
