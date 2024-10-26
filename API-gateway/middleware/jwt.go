package middleware

import (
    "net/http"
    "github.com/dgrijalva/jwt-go"
   
)

func JwtAuthMiddleware(jwtSecret string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            tokenString := r.Header.Get("Authorization")
            if tokenString == "" {
                http.Error(w, "Missing token", http.StatusUnauthorized)
                return
            }

            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                return []byte(jwtSecret), nil
            })

            if err != nil || !token.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
