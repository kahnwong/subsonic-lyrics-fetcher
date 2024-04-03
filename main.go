package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// init auth payload
type AuthPayload struct {
	u string
	t string
	s string
	v string
	c string
	f string
}

func initAuthPayload() *AuthPayload {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading env from env var instead...")
	}

	return &AuthPayload{os.Getenv("USERNAME"), os.Getenv("TOKEN"), os.Getenv("SALT"), "1.16.1", "github-readme", "json"}
}

// get now playing
//func getNowPlaying() string {
//	return "foo"
//}

func main() {
	authPayload := initAuthPayload()
	fmt.Println("hello")
	fmt.Println(authPayload)
}
