package facebook

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// Credentials stores google client-ids.
type Credentials struct {
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"secret"`
}

// User is a retrieved and authenticated user.
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Hd            string `json:"hd"`
}

var conf *oauth2.Config
var state string

// Setup the Provider
func Setup(redirectURL string, cred Credentials, scopes []string, secret []byte) {
	// store = sessions.NewCookieStore(secret)

	conf = &oauth2.Config{
		ClientID:     cred.ClientID,
		ClientSecret: cred.ClientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     facebook.Endpoint,
	}
}

func Auth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		state := "getfromstate"
		retrievedState := q.Get("state")
		if state != retrievedState {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Error: %s -- %s", w.Header(), "Unauthorized access")
			return
		}

		tok, err := conf.Exchange(oauth2.NoContext, q.Get("code"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %s -- %v", w.Header(), err)
			return
		}

		client := conf.Client(oauth2.NoContext, tok)
		email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %s -- %v", w.Header(), err)
			return
		}
		defer email.Body.Close()
		data, _ := ioutil.ReadAll(email.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
}

/*
data, err := ioutil.ReadAll(email.Body)
if err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: %s -- %v", w.Header(), "Could not read body")
	return
}

var user User
err = json.Unmarshal(data, &user)
if err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: %s -- %v", w.Header(), "Unmarshal userinfo failed")
	return
}
ctx := context.Background()
r.
--------------------------------
func(ctx *gin.Context) {
// Handle the exchange code to initiate a transport.
session := sessions.Default(ctx)
retrievedState := session.Get("state")
if retrievedState != ctx.Query("state") {
	ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
	return
}

tok, err := conf.Exchange(oauth2.NoContext, ctx.Query("code"))
if err != nil {
	ctx.AbortWithError(http.StatusBadRequest, err)
	return
}

client := conf.Client(oauth2.NoContext, tok)
email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
if err != nil {
	ctx.AbortWithError(http.StatusBadRequest, err)
	return
}
defer email.Body.Close()
data, err := ioutil.ReadAll(email.Body)
if err != nil {
	glog.Errorf("[Gin-OAuth] Could not read Body: %s", err)
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
}

var user User
err = json.Unmarshal(data, &user)
if err != nil {
	glog.Errorf("[Gin-OAuth] Unmarshal userinfo failed: %s", err)
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
}
// save userinfo, which could be used in Handlers
ctx.Set("user", user)
}
*/
