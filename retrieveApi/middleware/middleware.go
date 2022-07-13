package middleware

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

const USERNAME = "lusia"
const PASSWORD = "$2a$10$M1zkOqA4de6TSTt090KzeumqBT2ZY8myC5gPBSvWUqoIKc5/rczOy" //"123456"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == USERNAME)
		if !isValid {
			w.Write([]byte(`wrong username`))
			return
		}

		/*	// Hashing the password with the default cost of 10
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(PASSWORD), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(hashedPassword))
		*/

		// Comparing the password with the hash
		err := bcrypt.CompareHashAndPassword([]byte(PASSWORD), []byte(password))
		if err != nil {
			w.Write([]byte(`wrong password`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
