package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID 			string 		`json:"id"`
	NAME 		string 		`json:"name"`
	ISBN		string 		`json:"isbn"`
	TITLE 		string 		`json:"title"`
	DIRECTOR 	*Director 	`json:"director"`
}

type Director struct {
	FIRSTNAME 	string 		`json:"first_name"`
	LASTNAME 	string 		`json:"last_name"`
}

type Response struct {
    MSG 		string 		`json:"msg"`
}

var movies []Movie;

// simple server testing route handle
// returning response in json form
func testing(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	res := Response { MSG: "testing api" };
	json.NewEncoder(w).Encode(res);
}

// fetching all the movies and return them into 
// json form
func getMovies(w http.ResponseWriter, r *http.Request){
	// Set(key string, value string)
	w.Header().Set("Content-Type", "application/json");
	
	// func json.NewEncoder(w io.Writer) *json.Encoder
	// Encode(v any) error
	json.NewEncoder(w).Encode(movies);
}

// get the id from params of request url
// find the movie with that id
// return json form of that movie
func getMovie(w http.ResponseWriter, r *http.Request){
	// setting up the response header type
	w.Header().Set("Content-Type", "application/json");
	// when using the Gorilla Mux router, mux.Vars(r) is used to 
	// extract URL parameters from an incoming HTTP request.
	params := mux.Vars(r);
	
	// we use json.NewEncoder(w).Encode(response) to convert a Go 
	// struct or map into a JSON response and write it directly 
	// to the http.ResponseWriter.	
	// json.NewEncoder(w).Encode(item), automatically write the res into 
	// header and returns;
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item);
			break
		}
	}	
}

// creating the movie with the incomming data
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	var newMovieData Movie 
	err := json.NewDecoder(r.Body).Decode(&newMovieData);
	if err != nil {
		http.Error(w, `{"error": "Invalid JSON data"}`, http.StatusBadRequest)
		return
	}

	// conversion of the int to the string from 0 to 100
	// here how we can address struct inside struct
	// appeding the movie and restuns in response as a json format
	newMovie := Movie{
		ID: strconv.Itoa(rand.Intn(100)),
		NAME: newMovieData.NAME,
		ISBN: newMovieData.ISBN,
		TITLE: newMovieData.TITLE,
		DIRECTOR: &Director{
			FIRSTNAME: newMovieData.DIRECTOR.FIRSTNAME,
			LASTNAME: newMovieData.DIRECTOR.LASTNAME,
		},
	}

	movies = append(movies, newMovie);
	json.NewEncoder(w).Encode(newMovie);
}

// edit the exisitng movies present the database
func editMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);

	// makign an object updatedMovieData as a type Movie
	// decoding the incomming json data into struct
	// decodng them into the address of updatedMovieData
	// "&updatedMovieData" address of the variable updatedMovieData
    var updatedMovieData Movie 
	_ = json.NewDecoder(r.Body).Decode(&updatedMovieData) 

	for index, item := range movies {
		if item.ID == params["id"] {
			// updating the specific (id) movie
			movies[index].ID = updatedMovieData.ID;
			movies[index].ISBN = updatedMovieData.ISBN;
			movies[index].NAME = updatedMovieData.NAME;
			movies[index].TITLE = updatedMovieData.TITLE;
			movies[index].DIRECTOR.FIRSTNAME = updatedMovieData.DIRECTOR.FIRSTNAME;
			movies[index].DIRECTOR.LASTNAME = updatedMovieData.DIRECTOR.LASTNAME;
			json.NewEncoder(w).Encode(movies[index])
 			break;
		}
	}	
}

// slicing the movies to remove the specific (id) movie
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	
	for index, item := range movies {
		// this index is moving from 0 ..... 
		if item.ID == params["id"] {
			// now we have to remove this
			// we are slicing slice [:index] to the [index+1:] full
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}	
}


func main(){

	movies = append(movies, Movie{
		ID: "1",
		NAME: "movie1",
		ISBN: "1234",
		TITLE: "movie1_title",
		// here we want the refrence of the director address
		// & gives the address
		// * used to access the value at that address
		DIRECTOR: &Director{
			FIRSTNAME: "director1",
			LASTNAME: "director_last_1",
		},
	})

    r := mux.NewRouter()
	r.HandleFunc("/", testing)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
    r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", editMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	http.Handle("/", r)

	fmt.Printf("starting server at: 8080\n");
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error", err);
	}
}

