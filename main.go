package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct{
  ID string `json:"id"`
  Title string `json:"title"`
  ISBN string `json:"isbn"`
  Director *Director `json:"director"`

}

type Director struct{
  FirstName string `json:"firstname"`
  LastName string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){ //r is request you'll send through postman and w is response you'll get
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(movies)

//Printing to the terminal
for index,movie := range movies{
	jsonData, err := json.Marshal(movie) //convert movies slice in to json
	if err !=nil {
	http.Error(w,"Internal Server Error", http.StatusInternalServerError )
	return
	}

	fmt.Println(index, string(jsonData))
}
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
params:=mux.Vars(r)
for index, movie := range movies{
	if movie.ID == params["id"]{
		movies = append(movies[:index], movies[index+1:]...)
		break
	}
}
json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)

	for _, movie:= range movies{
		if(movie.ID == params["id"]){
			json.NewEncoder(w).Encode(movie) //returns the movie
			break
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
newMovie:= Movie{
	ID: "3",
	Title: "consensus algorithm such as Proof of Work or Proof of Stake.",
	ISBN: "124444",
	Director: &Director{
    FirstName: "Vicky",
	LastName: "Tyagi",
	},
}

movies = append(movies, newMovie)

json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
params:=mux.Vars(r);
for idx, movie := range movies{
	if(movie.ID == params["id"]){
     movies[idx].Title="Modular Blockchain in Go"
	 break
	}
}
json.NewEncoder(w).Encode(movies)
}

func main(){
r:=mux.NewRouter() //mux router initialised !!

movies = append(movies, Movie{
	ID:"1",
	Title: "Blockchain in Go",
	ISBN:"144232",
	Director: &Director{
		FirstName: "Vivek",
		LastName: "Tyagi",
	},
})

movies = append(movies, Movie{
	ID:"2",
	Title: "Implementing Lightning wallet in Go",
	ISBN:"66632",
	Director: &Director{
		FirstName: "Vivek",
		LastName: "T",
	},
})

r.HandleFunc("/",getMovies).Methods("GET")  //get movies
r.HandleFunc("/movies/{id}",getMovie).Methods("GET") //get a specific movie
r.HandleFunc("/movies", addMovie).Methods("POST") //add a movie
r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")// update a movie
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") //delete a movie

fmt.Printf("Starting server at port:8000\n")
log.Fatal(http.ListenAndServe(":8000",r))//first arg setup the port at 8000 and serve it, and then r is used to handle incoming request
//log.fatal() is used to handle all type of error

}
