package restapiserver

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"

)

//Demo show how the restapi works
func Demo() {
	fmt.Println("Package REST API DEMO...")
}

//ProtectedFunc Accessing Protected Function
func ProtectedFunc() {

	fmt.Println("Calling Protected API...")

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	/* Handle credentials and token functions */

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})
	/*
		http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
				clientId := uuid.New().String() // uuid.NewV4().String()[:8]
			clientSecret := uuid.New().String()[:8]
			err := clientStore.Set(clientId, &models.Client{
				ID:     clientId,
				Secret: clientSecret,
				Domain: "http://localhost:9094",
			})
			if err != nil {
				fmt.Println(err.Error())
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
		})
	*/
	/* End of credentials and token functions */

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I am protected, Need to use oAuth2 and token system in order to reach me. "))
	})
	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	})

	log.Fatal(http.ListenAndServe("localhost:9096", nil))

}
