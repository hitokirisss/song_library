package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"song_library/models"
)

func GetSongInfo(group, song string) (models.Song, error) {
	url := fmt.Sprintf("%s?group=%s&song=%s", os.Getenv("MUSIC_INFO_API_URL"), group, song)
	resp, err := http.Get(url)
	if err != nil {
		return models.Song{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Song{}, fmt.Errorf("failed to get song info: %s", resp.Status)
	}

	var songDetail models.Song
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return models.Song{}, err
	}
	return songDetail, nil
}
