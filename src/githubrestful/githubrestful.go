package githubrestful

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

)

var (
	// you need to generate personal access token at
	// https://github.com/settings/applications#personal-access-tokens
	/* It will be generated after entering my login credentials at githug with above url */
	personalAccessToken = "0532ce602d329fe50310668afd04af29f14f507d"
	client              *github.Client
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

//GitaccessDemo Git Access using GIT REST API's
func GitaccessDemo() {

	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client = github.NewClient(oauthClient)

	//http.HandleFunc("/mygitdetails", func(w http.ResponseWriter, r *http.Request) ){ to define here itself..}
	fmt.Println("List of functions:")
	fmt.Printf("/getmydetails , ")

	fmt.Println()
	http.HandleFunc("/getmydetails", getdetails)

	fmt.Println("Listening on : localhost:9096")
	log.Fatal(http.ListenAndServe("localhost:9096", nil))

}

func getdetails(w http.ResponseWriter, r *http.Request) {
	user, _, err := client.Users.Get(context.Background(), "")
	/* "" specifies, use authenticated user for get operation */

	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		return
	}

	d, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("json.MarshlIndent() failed with %s\n", err)
		return
	}
	//w.Write([]byte(string(d)))

	fmt.Fprintln(w, " ------- GIT HUG GET DETAILS -----------", string(d))

}
