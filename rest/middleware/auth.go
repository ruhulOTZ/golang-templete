package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"expenseTracker/config"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(bearerToken, " ")

		if len(headerArr) != 2 || headerArr[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload
		byteArraySecret := []byte(config.GetConfig().JwtSecretKey)
		byteArrayMessage := []byte(message)

		h := hmac.New(sha256.New, byteArraySecret)
		h.Write(byteArrayMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)

		if newSignature != jwtSignature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	// Implement base64 URL encoding here
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
