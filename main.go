package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

type Venue struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Region   string `json:"region"`
	Link     string `json:"link"`
}
type Label struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Region   string `json:"region"`
	Genre    string `json:"genre"`
	Link     string `json:"link"`
}
type Artist struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	DJName   string `json:"dj_name"`
	Label    string `json:"label"`
	Location string `json:"location"`
	Region   string `json:"region"`
	Genre    string `json:"genre"`
	Link     string `json:"link"`
}
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type GenreLink struct {
	ID      int    `json:"id"`
	GenreID int    `json:"genre_id"`
	URL     string `json:"url"`
}

func main() {
	db, err := sql.Open("sqlite", "./local-hub.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/venues", handleVenues(db))

	http.HandleFunc("/api/artists", handleArtists(db))

	http.HandleFunc("/api/labels", handleLabels(db))

	log.Println("Backlend API running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleVenues(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//Initialise variable region to hold and URL Query's for "region"
		venueChoice := r.URL.Query().Get("venue")

		var rows *sql.Rows
		var err error

		//Check if region input is not empty
		//If not empty Query the sqlite db for id name location and region from venues if matching URL Query
		//If no URL Query is inputted select from venues as default
		if venueChoice != "" {
			rows, err = db.Query("SELECT id, name, location, region, link FROM venues WHERE venue = ?", venueChoice)
		} else {
			rows, err = db.Query("SELECT id, name, location, region, link FROM venues")
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var venues []Venue
		for rows.Next() {
			var v Venue
			if err := rows.Scan(&v.ID, &v.Name, &v.Location, &v.Region, &v.Link); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			venues = append(venues, v)
		}
		writeJSON(w, venues)

	}
}

func handleArtists(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		artistChoice := r.URL.Query().Get("artist")

		var rows *sql.Rows
		var err error

		if artistChoice != "" {
			rows, err = db.Query("SELECT id, name, dj_name, label, location, region, genre, link FROM artists WHERE name = ?", artistChoice)
		} else {
			rows, err = db.Query("SELECT id, name, dj_name, label, location, region, genre, link FROM artists")
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var allArtists []Artist
		for rows.Next() {
			var a Artist
			if err = rows.Scan(&a.ID, &a.Name, &a.DJName, &a.Label, &a.Location, &a.Region, &a.Genre, &a.Link); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			allArtists = append(allArtists, a)
		}
		writeJSON(w, allArtists)

	}
}

func handleLabels(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		labelChoice := r.URL.Query().Get("label")

		var rows *sql.Rows
		var err error

		if labelChoice != "" {
			rows, err := db.Query("SELECT id, name, location, region, genre, link FROM labels WHERE label = ?", labelChoice)
		} else {
			rows, err := db.Query("SELECT id, name, location, region, genre, link FROM labels")
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var labels []Label
		for rows.Next() {
			var l Label
			if err = rows.Scan(&l.ID, &l.Name, &l.Location, &l.Region, &l.Genre, &l.Link); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			labels = append(labels, l)
		}
		writeJSON(w, labels)
	}

}

func handleGenres(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		genreChoice := r.URL.Query().Get("genre")

		var rows *sql.Rows
		var err error

		if genreChoice != "" {
			rows, err = db.Query("SELECT id, name WHERE genre = ?", genreChoice)
		} else {
			rows, err = db.Query("SELECT id, name FROM genres")
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var genres []Genre
		for rows.Next() {
			var g Genre
			if err = rows.Scan(&g.ID, &g.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			genres = append(genre, g)
		}
		writeJSON(w, genre)

	}

}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
