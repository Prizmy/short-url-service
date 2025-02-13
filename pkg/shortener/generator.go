package shortener

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

var usedURLs = make(map[string]struct{})
var mu sync.Mutex

func Generate() string {
	mu.Lock()
	defer mu.Unlock()

	var shortURL string
	for {
		var sb strings.Builder
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			sb.WriteByte(chars[rand.Intn(len(chars))])
		}
		shortURL = sb.String()
		if _, exists := usedURLs[shortURL]; !exists {
			usedURLs[shortURL] = struct{}{}
			break
		}
	}
	return shortURL
}
