package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Super Secret", mySigningKey)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Invalid Signing Method")
				}
				aud := "billing.jwtgo.io"
				checkAudiance := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)

				if !checkAudiance {
					return nil, fmt.Errorf("Invalid aud")
				}

				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return nil, fmt.Errorf("Invalid Iss")
				}

				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "No Authorization")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {

	fmt.Println("Server")
	handleRequests()

}
