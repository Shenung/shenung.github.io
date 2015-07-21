package main

import (
	"net/http"
	"text/template"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

//making a type called tweet that has attributes of message as an array of string
//and has a time stamp
type tweet struct {
	message   string
	timestamp time.Time
}

//the main page data with attributes of Tweets type tweet
//auth as a bool for login check
//users email as a string
//Login URL as a string
type mpData struct {
	Tweets   []tweet
	Auth     bool
	Email    string
	LoginURL string
}

//type profile with attributes of Username thats  string
//contains the users Username
//contains the users Email
type profile struct {
	Username string
	Email    string
}

type ProfileData struct {
	Tweets  []tweet
	Profile profile
}

//assigns a new template to the variable tpl
var tpl = template.New("")

//Routers to handle routes for the pages
func init() {
	//Parse the templates
	_, err := tpl.ParseFiles("template/index.gohtml", "template/createProfile.gohtml", "template/profile.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/createProfile", createProfile)
}

//Function to handle the main page taht people will be on
func handler(res http.ResponseWriter, req *http.Request) {
	//Checks to make sure you have the page before executing rest of code
	if req.URL.Path != "/" {
		getProfile(res, req)
		return
	}

	//creates new context and passes it to user to get the context of user
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	//Check if the user is not equal to nil
	if u != nil {
		_, err := getProfileByEmail(ctx, u.Email)
		if err == datastore.ErrNoSuchEntity {
			http.Redirect(res, req, "/CreateProfile", http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(res, "server Error!", http.StatusInternalServerError)
			log.Errorf(ctx, "Datastore get Error %s\n", err.Error())
			return
		}
	}

	//Get the recent Tweets
	tweets, err := getTweets(ctx, "")
	if err != nil {
		http.Error(res, "Server Error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Query Error: %s\n", err.Error())
		return
	}

	//Creating template
	loginURL, err := user.LoginURL(ctx, "/")
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Login URL Error: %s\n", err.Error())
		return
	}
	//variable data given the mainpage data
	data := mpData{
		Tweets:   tweets,
		Auth:     u != nil,
		LoginURL: loginURL,
	}
	//checks is the the data variable with the auth attribute is true
	if data.Auth {
		data.Email = u.Email
	}

	//Serves template defaults to file name
	err = tpl.ExecuteTemplate(res, "index.gohtml", data)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}

//checks to make sure that the username is in the correct length
//formating as well as in format
func confirmCreateProfile(username string) bool {
	return len(username) > 5
}

func createProfile(res http.ResponseWriter, req *http.Request) {
	//create a new context
	ctx := appengine.NewContext(req)
	//create a new user with the the user as the current context
	u := user.Current(ctx)

	//check if the the request method is POST
	if req.Method == "POST" {
		//create variable with the name of the form field
		username := req.FormValue("username")
		//check the username from the formfield if it meets requirements
		//errors if false
		if !confirmCreateProfile(username) {
			http.Error(res, "Invalid input!", http.StatusBadRequest)
			log.Warningf(ctx, "Invalid profile information from %s\n", req.RemoteAddr)
			return
		}

		//assigns attributes its values
		key := datastore.NewKey(ctx, "profile", u.Email, 0, nil)
		p := profile{
			Username: username,
			Email:    u.Email,
		}
		//Puts the key into datastore of p to be accesssable
		_, err := datastore.Put(ctx, key, &p)
		if err != nil {
			http.Error(res, "Server error!", http.StatusInternalServerError)
			log.Errorf(ctx, "Create profile Error: %s\n", err.Error())
			return
		}
	}
	err := tpl.ExecuteTemplate(res, "createProfile.gohtml", nil)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}

func getProfile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
}
