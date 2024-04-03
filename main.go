package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// init auth payload
type AuthPayload struct {
	Username string `url:"u"`
	Token    string `url:"t"`
	Salt     string `url:"s"`
	Version  string `url:"v"`
	Client   string `url:"c"`
	Format   string `url:"f"`
}

func initAuthPayload() *AuthPayload {
	return &AuthPayload{
		Username: os.Getenv("USERNAME"),
		Token:    os.Getenv("TOKEN"),
		Salt:     os.Getenv("SALT"),
		Version:  "1.16.1",
		Client:   "github-readme",
		Format:   "json",
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading env from env var instead...")
	}

	baseURL := os.Getenv("BASE_URL")
	authPayload := initAuthPayload()

	// get now playing
	nowPlaying, err := getNowPlaying(baseURL, authPayload)
	if err != nil {
		fmt.Println("Error getting now playing")
	}
	if nowPlaying == nil {
		fmt.Println("Currently nothing is playing")
		os.Exit(1)
	} else {
		// print lyrics
		fmt.Println(nowPlaying)
		fmt.Println("lyrics")
	}
}
