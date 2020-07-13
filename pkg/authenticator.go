package sbb_api

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

const NotSoSecretKey = "c3eAd3eC3a7845dE98f73942b3d5f9c0"

func GetAuthorization(path string, date string) string {
	hmacFunc := hmac.New(sha1.New, []byte(getKey()))
	hmacFunc.Write([]byte(fmt.Sprintf("%s%s", path, date)))
	result := hmacFunc.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}

func getKey() string {
	certHash, err := getCertificateHash()
	if err != nil {
		return ""
	}
	return retrieveKey(certHash, NotSoSecretKey)
}

func retrieveKey(certHash string, key string) string {
	clearText := fmt.Sprintf("%s%s", certHash, key)
	h := sha256.New()
	h.Write([]byte(clearText))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func getCertificateHash() (string, error) {
	certBase64 := "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"
	certBin, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		return "", errors.New("unable to decode certificate from B64 string")
	}

	h := sha1.New()
	h.Write(certBin)
	hash := h.Sum(nil)

	return base64.StdEncoding.EncodeToString(hash), nil
}
