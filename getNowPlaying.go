package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type NowPlayingFull struct {
	SubsonicResponse struct {
		Status        string `json:"status"`
		Version       string `json:"version"`
		Type          string `json:"type"`
		ServerVersion string `json:"serverVersion"`
		NowPlaying    struct {
			Entry []struct {
				ID          string    `json:"id"`
				Parent      string    `json:"parent"`
				IsDir       bool      `json:"isDir"`
				Title       string    `json:"title"`
				Album       string    `json:"album"`
				Artist      string    `json:"artist"`
				Track       int       `json:"track"`
				Year        int       `json:"year"`
				Genre       string    `json:"genre"`
				CoverArt    string    `json:"coverArt"`
				Size        int       `json:"size"`
				ContentType string    `json:"contentType"`
				Suffix      string    `json:"suffix"`
				Duration    int       `json:"duration"`
				BitRate     int       `json:"bitRate"`
				Path        string    `json:"path"`
				DiscNumber  int       `json:"discNumber"`
				Created     time.Time `json:"created"`
				AlbumID     string    `json:"albumId"`
				ArtistID    string    `json:"artistId"`
				Type        string    `json:"type"`
				IsVideo     bool      `json:"isVideo"`
				Username    string    `json:"username"`
				MinutesAgo  int       `json:"minutesAgo"`
				PlayerID    int       `json:"playerId"`
				PlayerName  string    `json:"playerName"`
			} `json:"entry"`
		} `json:"nowPlaying"`
	} `json:"subsonic-response"`
}

type NowPlaying struct {
	Title  string
	Artist string
}

func getNowPlaying(baseURL string, authPayload *AuthPayload) (*NowPlaying, error) {
	authParams, _ := query.Values(authPayload)
	url := fmt.Sprintf("%s/rest/getNowPlaying?%s", baseURL, authParams.Encode())
	resp, err := http.Get(url)

	if err != nil {
		log.Println("No response from request")
		return nil, err
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
		return nil, err
	}

	var result NowPlayingFull
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		return nil, err
	}

	// parse response
	tracksRaw := result.SubsonicResponse.NowPlaying.Entry

	if len(tracksRaw) == 1 { // if has recently played tracks
		nowPlayingTrack := tracksRaw[0]
		nowPlaying := &NowPlaying{tracksRaw[0].Title, nowPlayingTrack.Artist}

		return nowPlaying, nil
	} else {
		return nil, nil
	}
}
