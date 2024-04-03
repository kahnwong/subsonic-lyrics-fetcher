package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/go-querystring/query"
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

// get now playing
type NowPlaying struct {
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

type NowPlayingSong struct {
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
}

func getNowPlaying(baseURL string, authPayload *AuthPayload) ([]NowPlayingSong, error) {
	authParams, _ := query.Values(authPayload)
	url := fmt.Sprintf("%s/rest/getNowPlaying?%s", baseURL, authParams.Encode())
	resp, err := http.Get(url)

	if err != nil {
		log.Println("No response from request")
		return []NowPlayingSong{}, err
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
		return []NowPlayingSong{}, err
	}

	var result NowPlaying
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		return []NowPlayingSong{}, err
	}

	// parse response
	tracksRaw := result.SubsonicResponse.NowPlaying.Entry

	if len(tracksRaw) == 1 { // if has recently played tracks
		tracks := []NowPlayingSong{}

		for _, rec := range tracksRaw {
			tracks = append(tracks, rec)
		}

		return tracks, nil
	} else {
		return []NowPlayingSong{}, nil
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
	if len(nowPlaying) == 0 {
		log.Println("Currently nothing is playing")
		os.Exit(1)
	}

	// print lyrics

	// print result
	fmt.Println(nowPlaying)
}
