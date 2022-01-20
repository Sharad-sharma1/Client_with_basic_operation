package main

import (
	"encoding/json" //for encoding data to json for sending to our server
	"fmt"           //for printing
	"log"           //for log out the errors
	"math/rand"     //for giving random number
	"net/http"      //this will help us to hit server
	"strconv"       //the id we will create using math.rand it will be in int form str conv will convert that into sting after that Itoa

	//will format int into string
	"github.com/gorilla/mux" //It will handle http rquest and help to create go web server
)

//struct(work as key,value kind of method) is  help to group data and from records
type Clients struct {
	ID          string `json:"id"`
	Client_name string `json:"client_name"`
	Created_at  string `json:"created_at"`
	Created_by  string `json:"created_by"`
}

type Projects struct {
	ID           string `json:"id"`
	Project_name string `json:"project_name"`
	Created_at   string `json:"created_at"`
	Created_by   string `json:"created_by"`
}

type Users struct {
	ID          string    `json:"id"`
	Client_name string    `json:"client_name"`
	Projects    *Projects `json:"project"` //* is pointer which give access and the address is given by &
	Created_at  string    `json:"created_at"`
	Created_by  string    `json:"created_by"`
	Updated_at  string    `json:"Updated_at"`
}

var clients []Clients

//slicer it helps to get data

func getClients(w http.ResponseWriter, r *http.Request) {
	//r pointer will be senting request and w will response
	w.Header().Set("Content-Type", "application/json") //We have to convert struct into json
	//(the requested file will be converted into json format and given as a )
	json.NewEncoder(w).Encode(clients) //json.NewEncoder(w) it will convert into json and wrtitten (w) clients value and then encode clients give output
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)              //r is request sent by user or us
	for index, item := range clients { //it will loop through clients present in func main
		if item.ID == params["id"] { //here if condition will compares item of for loop with
			//params mux requested id(which will be entered by us or user), till here it will only give access to values
			clients = append(clients[:index], clients[index+1:]...) // the first part(clients[:index]) of append
			//will get replaced with all('...' all) upcoming clients indexes(clients[index+1:])
			break
		}
	}
	json.NewEncoder(w).Encode(clients) //once client it deleted remaining will displayed
}

func getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range clients { //we will not use index but we have to because it works as pair so we will use blank identifier
		//if we use index and dont use it golang will throw error
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var client Clients
	_ = json.NewDecoder(r.Body).Decode(&client) //r.body will take entire body from json and decode it and add into client variable defined above
	client.ID = strconv.Itoa(rand.Intn(10))     //here the ID present in client will be given a random value to it
	clients = append(clients, client)           // here everything will be appended into clients table from client(r.body)
	json.NewEncoder(w).Encode(client)           //once new client is a created it will get displayed
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clients {
		if item.ID == params["id"] {
			clients = append(clients[:index], clients[index+1:]...) //deleting the id we have sent in the server
			var client Clients
			_ = json.NewDecoder(r.Body).Decode(&client)
			client.ID = params["id"]
			clients = append(clients, client) //add a new client - the client we will send in server
			json.NewEncoder(w).Encode(client) // in this update what we have done is deleted the requested id and append with the passed data in server
		}
	}
}

var projects []Projects

func updateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range projects {
		if item.ID == params["id"] {
			projects = append(projects[:index], projects[index+1:]...)
			var project Projects
			_ = json.NewDecoder(r.Body).Decode(&project)
			project.ID = params["id"]
			projects = append(projects, project)
			json.NewEncoder(w).Encode(project)
		}
	}
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

var users []Users

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	r := mux.NewRouter() //it will handle incoming http request

	clients = append(clients, Clients{ID: "1", Client_name: "Microsoft", Created_at: "2021-12-24T11:03:55.931739+05:30", Created_by: "Roy"})
	//here what ever we add using append function it will showed in server
	clients = append(clients, Clients{ID: "2", Client_name: "Google", Created_at: "2021-12-23T11:03:55.931739+05:31", Created_by: "Sharad"})

	projects = append(projects, Projects{ID: "1", Project_name: "Azure", Created_at: "2021-11-22T11:03:55.931739+05:30", Created_by: "Sam"})
	projects = append(projects, Projects{ID: "2", Project_name: "Google Cloud", Created_at: "2021-11-21T11:03:55.931739+05:31", Created_by: "John"})

	users = append(users, Users{ID: "1", Client_name: "Microsoft", Projects: &Projects{ID: "2", Project_name: "Google Cloud"},
		Created_at: "2021-12-24T11:03:55.931739+05:30", Created_by: "Roy",
		Updated_at: "202-12-28T11:03:55.931739+05:45"})

	r.HandleFunc("/clients", getClients).Methods("GET") //handlefunc handles the functions and sent to mux.newrouter
	r.HandleFunc("/clients/{id}", getClient).Methods("GET")
	r.HandleFunc("/clients", createClient).Methods("POST")
	r.HandleFunc("/clients/{id}", updateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", deleteClient).Methods("DELETE")

	r.HandleFunc("/projects/{id}", updateProject).Methods("PUT")
	r.HandleFunc("/projects", getProjects).Methods("GET")

	r.HandleFunc("/users/{id}", getUsers).Methods("GET")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
