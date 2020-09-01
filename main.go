package main

import (
	"fmt"
	"context"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type spotifyClient interface {
	GetPlaylist(playlist spotify.ID) (*spotify.FullPlaylist, error)
}

func main() {  
	const playlistId spotify.ID = "4OyKDT6cLw96G7bd8nTfxD"

	client := newSpotifyClient()
	name := getPlaylistName(client, playlistId)

	fmt.Println(name)
} 

func newSpotifyClient() *spotify.Client {
	config := &clientcredentials.Config{
		ClientID: os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL: spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("Couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	return &client
}

func getPlaylistName(client spotifyClient, playlistId spotify.ID) string {
	result, err := client.GetPlaylist(playlistId)
	if err != nil {
		log.Fatalf("Couldn't get playlist: %v", err)
	}

	return result.Name
}