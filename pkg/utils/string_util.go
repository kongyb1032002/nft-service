package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomString(length int) string {
	randomUUID := uuid.New()
	return randomUUID.String()[:length]
}

func GenerateTraceNo() string {
	unixMilli := time.Now().UnixMilli()

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(10000) + 1

	traceNo := fmt.Sprintf("%d%d", unixMilli, randomNum)

	return traceNo
}
