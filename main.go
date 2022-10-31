package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

var movies []Movie

func main() {
	movies = append(movies, Movie{
		Id:       "13",
		Name:     "krisss",
		Year:     2000,
		Director: "joe",
	})
	r := mux.NewRouter()
	r.HandleFunc("/sayhi", func(w http.ResponseWriter, r *http.Request) {
		// 		io.WriteString(w, "hello world!")
		fmt.Fprintln(w, "thisiscontentfromfprintln")
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	log.Fatal(srv.ListenAndServe())
	fmt.Println("listenting......")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		fmt.Println(err)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			fmt.Println(movies[index])
			json.NewEncoder(w).Encode(movies[index])
			return
		}

	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "{success:false, msg:\"no movie with given id\"}")
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		fmt.Println(err)
		return
	}
	movies = append(movies, m)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	var m Movie
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	for index, item := range movies {
		if item.Id == m.Id {
			movies[index].Director = m.Director
			movies[index].Name = m.Name
			movies[index].Year = m.Year
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "{success:true}")
			return
		}

	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "{success:false,message:\"movie with id not found\"}")

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(204)
}
