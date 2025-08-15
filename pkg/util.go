package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/oklog/ulid"
)

func NegativeOf[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](value T) T {
	return -value
}

func PointerTime(p time.Time) *time.Time {
	return &p
}

func GeneratePrefixString(prefix string) (id string) {
	entropy := ulid.Monotonic(rand.Reader, 0)
	ms := ulid.Timestamp(time.Now())
	UID, _ := ulid.New(ms, entropy)

	id = fmt.Sprintf("%s-%s", prefix, UID.String())
	return id
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func SetBasicAuth(username, password string) string {
	return "Basic " + basicAuth(username, password)
}
