package main

import (
	"testing"

	"github.com/zmb3/spotify"
)

type mockSpotifyClient struct{}

func (m *mockSpotifyClient) GetPlaylist(playlistId spotify.ID) (*spotify.FullPlaylist, error) {
	return &spotify.FullPlaylist{
		SimplePlaylist: spotify.SimplePlaylist{
			Name: "Whatever",
		},
	}, nil
}

func Text_NewGetPlaylistName(t *testing.T) {
	client := &mockSpotifyClient{}
	name := getPlaylistName(client, "whatever")

	if name != "whatever" {
		t.Errorf("Expected %s, got %s", "whatever", name)
	}
}