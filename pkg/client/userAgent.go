package client

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

func UserAgent() string {
	content, err := os.ReadFile("user-agents.txt")
	if err != nil {
		log.Fatal("error opening file for reading:", err)
	}

	userAgentList := strings.Split(string(content), "\n")
	randIndex := rand.Intn(len(userAgentList) + 1)
	return userAgentList[randIndex]
}
