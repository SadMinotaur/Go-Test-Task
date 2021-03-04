package utils

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"medods/src/dao"
	typesF "medods/src/types"
	"time"
)

func ValidTokens(guid string, jwtToken string, base64RToken string) (string, error) {
	rtQuery, err := base64.StdEncoding.DecodeString(base64RToken)
	if err != nil {
		return "", err
	}
	rtDB, err := dao.GetToken(guid, rtQuery)
	if err != nil {
		return "", err
	}
	if (rtDB == typesF.RToken{}) {
		return "Different refresh token", nil
	}
	rtDB.Token = string(rtQuery)
	if rtDB.Expires <= time.Now().Unix() {
		return "Refresh token expired", nil
	}
	jwtRtString, err := GetToken(rtDB.GUID, rtDB.Token, rtDB.AccessCreated, rtDB.AccessExpires)
	if err != nil {
		return "", err
	}
	if jwtRtString != jwtToken {
		return "Wrong token", nil
	}
	if rtDB.AccessExpires <= time.Now().Unix() {
		return "Access token expired", nil
	}
	return "", nil
}

func GenerateToken(guid string, n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b) + guid
}

func GetToken(guid string, rToken string, created, expires int64) (string, error) {
	jwtRt := jwt.New(jwt.SigningMethodHS512)
	claims := make(jwt.MapClaims)
	claims["guid"] = guid
	claims["created"] = created
	claims["expires"] = expires
	claims["rt"] = rToken
	jwtRt.Claims = claims
	jwtRtS, err := jwtRt.SignedString([]byte("sign"))
	if err != nil {
		return "", err
	}
	return jwtRtS, nil
}
