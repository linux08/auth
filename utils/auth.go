package auth

import (
	"context"
	"encoding/json"
	"expense/models"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//Exception struct
type Exception models.Exception

// JwtVerify Middleware function, which will be called for each request
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api", "/api/register", "/api/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                                  //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		var authHdr = r.Header.Get("x-access-token") //Grab the token from the header

		authHdr = strings.TrimSpace(authHdr)

		if authHdr == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(authHdr, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		// if claims, ok := token.Claims.(*models.Token); ok && token.Valid {
		// 	fmt.Printf("%v %v", claims.UserID, claims.StandardClaims)
		// } else {
		// 	fmt.Println(err)
		// }
		// fmt.Println(token.Claims.(*models.Token))
		ctx := context.WithValue(r.Context(), "user", tk)
		// ctx := context.WithValue(r.Context(), "user", tk.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
