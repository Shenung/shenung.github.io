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
	message   []string
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
type profile struct {
	Username string
}

//assigns a new template to the variable tpl
var tpl = template.New("")

//Routers to handle routes for the pages
func init() {
	http.HandleFunc("/", handler)
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
	if data.Logged {
		data.Email = u.Email
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", data)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}

func getProfile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
}
