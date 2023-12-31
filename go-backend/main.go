package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ISBN     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
}

var movies []Movie

// To avoid CORS error by allowing request from any origin
// func corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

func getMovies(w http.ResponseWriter, r *http.Request) { //r is request you'll send through postman and w is response you'll get
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

	//Printing to the terminal
	for index, movie := range movies {
		jsonData, err := json.Marshal(movie) //convert movies slice in to json
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println(index, string(jsonData))
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie) //returns the movie
			break
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, movie := range movies {
		if movie.ID == params["id"] {
			movies[idx].Title = "Modular Blockchain in Go"
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter() //mux router initialised !!

	// CORS middleware
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	movies = append(movies, Movie{
		ID:    "1",
		Title: "Blockchain in Go",
		ISBN:  "144232",
		Director: &Director{
			FirstName: "Vivek",
			LastName:  "Tyagi",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Title: "Implementing Lightning wallet in Go",
		ISBN:  "66632",
		Director: &Director{
			FirstName: "Vivek",
			LastName:  "T",
		},
	})

	r.HandleFunc("/", getMovies).Methods("GET")                 //get movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       //get a specific movie
	r.HandleFunc("/movies", addMovie).Methods("POST")           //add a movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // update a movie
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") //delete a movie

	fmt.Printf("Starting server at port:8080\n")
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)) //first arg setup the port at 8000 and serve it, and then r is used to handle incoming request
	//log.fatal() is used to handle all type of error

}
