package random

import (
	"math/rand"
	"strings"
	"time"
)

const (
	UpperCaseLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCaseLetter = "abcdefghijklmnopqrstuvwxyz"
	Number          = "1234567890"
	Characters      = "!@$%^&*-+="
)

func Rand(length int, available ...string) (res string) {
	randIN := ""
	if len(available) < 1 {
		randIN = Number
	}
	for _, c := range available {
		randIN += c
	}
	randList := strings.Split(randIN, "")
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(randList) - 1)
		res += randList[index]
	}
	return res
}
