package middleware

import (
	"context"
	"net/http"
	"strings"
	"example/moviecrud/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "user"

func Auth(requiredAdmin bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		{
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization")
				if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
					utils.JSONError(w, "missing or invalid token", http.StatusUnauthorized)
					return
				}

				tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
				claims := &utils.Claims{}

				token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
					return utils.AccessSecret, nil
				})

				if err != nil || !token.Valid {
					utils.JSONError(w, "invalid or expired token", http.StatusUnauthorized)
					return
				}

				// اگر نیاز به ادمین باشه و کاربر ادمین نباشه
				if requiredAdmin && claims.LevelName != "admin" {
					utils.JSONError(w, "admin access required", http.StatusForbidden)
					return
				}

				// کاربر رو توی context می‌ذاریم تا کنترلرها استفاده کنن
				ctx := context.WithValue(r.Context(), UserKey, claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		}
	}
}

func GetCurrentUser(r *http.Request) *utils.Claims {
	if user := r.Context().Value(UserKey); user != nil {
		return user.(*utils.Claims)
	}
	return nil
}
