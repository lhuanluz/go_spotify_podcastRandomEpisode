package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"github.com/joho/godotenv"
)

const apiBaseURL = "https://api.spotify.com/v1"

type episode struct {
	URI string `json:"uri"`
}

type episodesResponse struct {
	Items []episode `json:"items"`
}

var (
	clientID     string
	clientSecret string
	podcastID    string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	podcastID = os.Getenv("PODCAST_ID")
	token, err := getToken()
	if err != nil {
		fmt.Printf("Failed to get token: %v\n", err)
		return
	}

	episodes, err := getAllEpisodes(token)
	if err != nil {
		fmt.Printf("Failed to get episodes: %v\n", err)
		return
	}

	if len(episodes) == 0 {
		fmt.Println("No episodes found.")
		return
	}

	randomEpisode := episodes[rand.Intn(len(episodes))]
	openInSpotifyApp(randomEpisode.URI)
}

func getToken() (string, error) {
	url := "https://accounts.spotify.com/api/token"
	data := "grant_type=client_credentials"

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	req.Header.Add("Authorization", "Basic "+encodedAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	if err = json.Unmarshal(body, &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func getAllEpisodes(token string) ([]episode, error) {
	var allEpisodes []episode
	offset := 0
	limit := 50

	for {
		url := fmt.Sprintf("%s/shows/%s/episodes?market=US&limit=%d&offset=%d", apiBaseURL, podcastID, limit, offset)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var episodesResp episodesResponse
		if err = json.Unmarshal(body, &episodesResp); err != nil {
			return nil, err
		}

		allEpisodes = append(allEpisodes, episodesResp.Items...)

		// If less than the limit is returned, we've fetched all episodes
		if len(episodesResp.Items) < limit {
			break
		}

		offset += limit
	}

	return allEpisodes, nil
}

func openInSpotifyApp(uri string) {
	trimmedURI := strings.TrimPrefix(uri, "spotify:")
	httpURL := "https://open.spotify.com/" + strings.ReplaceAll(trimmedURI, ":", "/")
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", httpURL)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", httpURL)
	case "darwin":
		cmd = exec.Command("open", httpURL)
	default:
		fmt.Printf("Unsupported OS: %s\n", runtime.GOOS)
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to open Spotify app: %v\n", err)
	}
}

