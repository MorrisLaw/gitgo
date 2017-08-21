package main

import (
	"encoding/json"
	"time"
	"io/ioutil"
	"log"
	"net/http"
)

// Constants.
const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/" 
)

/* User struct represents the JSON data from GitHub API: 
   https://api.github.com/users/defunct this struct was
   generated via a JSON-to-GO utility by Matt Holt: 
   https://mholt.github.io/json-to-go/ */
type User struct {
	Login				string		`json:"login"`
	ID					int			`json:"id"`
	AvatarURL			string		`json:"avatar_url"`
	GravatarID			string 		`json:"gravatar_id"`
	URL					string		`json:"url"`
	HTMLURL				string 		`json:"html_url"`
	FollowersURL		string		`json:"followers_url"`
	FollowingURL		string		`json:"following_url"`
	GistsURL			string		`json:"gists_url"`
	StarredURL			string 		`json:"starred_url"`
	SubscriptionsURL	string		`json:"subscriptions_url"`
	OrganizationsURL 	string		`json:"organizations_url"`
	ReposURL			string		`json:"repos_url"`
	EventsURL			string		`json:"events_url"`
	ReceivedEventsURL	string		`json:"received_events_url"`
	Type				string		`json:"type"`
	SiteAdmin			bool 		`json:"site_admin"`
	Name				string		`json:"name"`
	Company				string		`json:"company"`
	Blog				string 		`json:"blog"`
	Location			string		`json:"location"`
	Email 				string		`json:"email"`
	Hireable			interface{}	`json:"hireable"`
	Bio					string		`json:"bio"`
	PublicRepos			int			`json:"public_repos"`
	PublicGists			int			`json:"public_gists"`
	Followers			int			`json:"followers"`
	Following			int			`json:"following"`
	CreatedAt			time.Time	`json:"create_at"`
	UpdatedAt			time.Time	`json:"updated_at"`
}

// getUsers queries GitHub API for a given user.
func getUsers(name string) User {
	// Send GET request to GitHub API with the requested user "name".
	resp, err := http.Get(apiURL + userEndpoint + name)

	/* If err occurs during GET request, then throw error and quit 
	   application. */
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	/* Always good practice to defer closing the response body.
	   If application crashes or function finishes successfully, 
	   GO will always execute this "defer" statement. */
	defer resp.Body.Close()

	// Read the response body and handle any errors during reading.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	/* Create a user variable of type "User" struct to store the
	   "Unmarshal"-ed (a.k.a. parsed JSON) data, then return the
	   user. */
	var user User
	json.Unmarshal(body, &user)
	return user
}