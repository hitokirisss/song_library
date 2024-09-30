package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"song_library/models"
	"song_library/services"
)

// GetSongs получает список песен
// @Summary Получить список песен
// @Description Получает список песен с фильтрацией и пагинацией
// @Tags songs
// @Produce json
// @Success 200 {array} models.Song
// @Router /songs [get]
func GetSongs(db *sql.DB) http.HandlerFunc {
	log.Println("GetSongs function called")

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec(`insert into songs(id, group_name, title, release_date, lyrics, link) values (1, 'qwe','ewq','rty','zzxc','cxz')`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rows, err := db.Query("SELECT * FROM songs")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		fmt.Println(rows)
		var songs []models.Song

		for rows.Next() {
			var song models.Song
			if err := rows.Scan(&song.ID, &song.Group, &song.Title, &song.ReleaseDate, &song.Lyrics, &song.Link); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println(song)
			songs = append(songs, song)
		}

		log.Printf("Retrieved songs: %+v\n", songs) // Логируем полученные песни

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)

	}
}

// GetSongText получает текст песни
// @Summary Получить текст песни
// @Description Получает текст песни по ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {string} string
// @Router /songs/{id}/text [get]
func GetSongText(w http.ResponseWriter, r *http.Request) {
	// Реализовать получение текста песни по куплетам
}

// AddSong добавляет новую песню
// @Summary Добавить новую песню
// @Description Добавляет новую песню в библиотеку
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Song object"
// @Success 201 {object} models.Song
// @Router /songs [post]
func AddSong(w http.ResponseWriter, r *http.Request) {
	var newSong models.Song
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := services.GetSongInfo(newSong.Group, newSong.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Сохранить enrichedSong в БД
}

// UpdateSong обновляет данные песни
// @Summary Обновить данные песни
// @Description Обновляет данные песни по ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body models.Song true "Song object"
// @Success 200 {object} models.Song
// @Router /songs/{id} [put]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	// Реализовать изменение данных песни
}

// DeleteSong удаляет песню
// @Summary Удалить песню
// @Description Удаляет песню по ID
// @Tags songs
// @Param id path int true "Song ID"
// @Success 204
// @Router /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	// Реализовать удаление песни
}
