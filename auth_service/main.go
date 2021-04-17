package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("No client file found. Exiting\n")
	}
	fileName := args[1]

	config, err := readConfig(fileName)
	if err != nil {
		log.Fatal("Unable to parse config")
	}
	log.Printf("clients: %+v", config)

	manager := manage.NewDefaultManager()
	// token memory store
	tokenStore, err := store.NewMemoryTokenStore()
	manager.MustTokenStorage(tokenStore, err)
	cfg := &manage.Config{
		AccessTokenExp:    time.Hour * 1,
		RefreshTokenExp:   time.Hour * 24,
		IsGenerateRefresh: true,
	}
	manager.SetClientTokenCfg(cfg)

	// client memory store
	clientStore := store.NewClientStore()
	createClients(clientStore, config)
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/introspect", createIntrospectHandler(tokenStore))

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}
