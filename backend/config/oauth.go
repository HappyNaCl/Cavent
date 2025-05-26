package config

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func SetupOAuth(){
	key := os.Getenv("SESSION_SECRET")
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if key == "" || clientId == "" || clientSecret == "" {
		panic("Missing environment variables for OAuth")
	}

    store := sessions.NewCookieStore([]byte(key))
    gothic.Store = store
	
	goth.UseProviders(
		google.New(clientId, clientSecret, "http://localhost:8080/api/v1/auth/google/callback"),	
	)
}