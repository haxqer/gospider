package mgtv

import (
	"fmt"
	"math/rand"
	"time"
)

func GenJsonp() string {
	nowTS := time.Now().UnixNano() / int64(time.Millisecond)
	jpRand := rand.Intn(9000) + 90000
	return fmt.Sprintf("jsonp_%d_%d", nowTS, jpRand)
}
