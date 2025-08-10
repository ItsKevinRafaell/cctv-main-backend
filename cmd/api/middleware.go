package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const UserClaimsKey = ContextKey("userClaims")

var jwtSecret = []byte("kunci-rahasia-yang-sangat-aman-dan-panjang")

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Header Authorization tidak ditemukan", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Format token tidak valid", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token tidak valid", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println("Claims in middleware:", claims)
			ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
			next(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Token tidak valid", http.StatusUnauthorized)
		}
	}
}
