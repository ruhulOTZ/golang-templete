package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"` // Subject (user ID)
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func CreateJWT(secretKey string, data Payload) (string, error) {
	// Implement JWT creation logic here
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrayHeader, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	headerBase64 := base64UrlEncode(byteArrayHeader)

	byteArrayPayload, err := json.Marshal(data)

	if err != nil {
		return "", err
	}
	payloadBase64 := base64UrlEncode(byteArrayPayload)

	message := headerBase64 + "." + payloadBase64

	byteArraySecret := []byte(secretKey)
	byteArrayMessage := []byte(message)

	h := hmac.New(sha256.New, byteArraySecret)
	h.Write(byteArrayMessage)

	signature := h.Sum(nil)
	signatureBase64 := base64UrlEncode(signature)

	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	// Implement base64 URL encoding here
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
