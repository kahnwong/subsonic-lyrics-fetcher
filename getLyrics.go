package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

type LyricsFull struct {
	SubsonicResponse struct {
		Status        string `json:"status"`
		Version       string `json:"version"`
		Type          string `json:"type"`
		ServerVersion string `json:"serverVersion"`
		OpenSubsonic  bool   `json:"openSubsonic"`
		Lyrics        struct {
			Artist string `json:"artist"`
			Title  string `json:"title"`
			Value  string `json:"value"`
		} `json:"lyrics"`
	} `json:"subsonic-response"`
}

func getLyrics(baseURL string, authPayload AuthPayload, nowPlaying NowPlaying) (string, error) {
	getLyricsPayload := struct {
		AuthPayload
		NowPlaying
	}{authPayload, nowPlaying}

	params, _ := query.Values(getLyricsPayload)
	url := fmt.Sprintf("%s/rest/getLyrics?%s", baseURL, params.Encode())
	resp, err := http.Get(url)

	if err != nil {
		log.Println("No response from request")
		return "", err
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
		return "", err
	}

	var result LyricsFull
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		return "", err
	}

	// parse response
	lyrics := result.SubsonicResponse.Lyrics.Value

	return lyrics, nil
}
