package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUsername() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("User_%d", rand.Intn(9999))
}
