package controllers

import (
	"encoding/base64"
	"fmt"
	dao "medods/src/dao"
	typesF "medods/src/types"
	utils "medods/src/utils"
	"net/http"
	"time"
)

func (serLogs *SerLogs) getTokens(guid string) (typesF.AToken, typesF.RToken, error) {
	rToken := typesF.RToken{
		GUID:          guid,
		Token:         utils.GenToken(guid, 20),
		Created:       time.Now().Unix(),
		AccessCreated: time.Now().Unix(),
		Expires:       time.Now().AddDate(1, 0, 0).Unix(),
		AccessExpires: time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	jwtTokenString, err := utils.GetToken(guid, rToken.Token, rToken.AccessCreated, rToken.AccessExpires)
	if err != nil {
		return typesF.AToken{}, typesF.RToken{}, err
	}
	aToken := typesF.AToken{
		GUID:    guid,
		Token:   jwtTokenString,
		Created: rToken.AccessCreated,
		Expires: rToken.AccessExpires,
	}
	err = dao.SaveToken(rToken)
	if err != nil {
		return typesF.AToken{}, typesF.RToken{}, err
	}
	return aToken, rToken, nil
}

func (serLogs *SerLogs) FirstRoute(writer http.ResponseWriter, request *http.Request) {
	guid := request.URL.Query().Get("guid")
	if guid == "" {
		serLogs.errorLog.Print(writer, http.StatusBadRequest)
		return
	}
	TokenA, TokenR, err := serLogs.getTokens(guid)
	if err != nil {
		serLogs.errorLog.Print(writer, err)
		return
	}
	base64t := base64.StdEncoding.EncodeToString([]byte(TokenR.Token))
	_, _ = fmt.Fprint(writer, "<a href='/second?guid="+guid+
		"&at="+TokenA.Token+"&rt="+base64t+"'>Second route</a>")
}

func (serLogs *SerLogs) SecondRoute(writer http.ResponseWriter, request *http.Request) {
	guidQ := request.URL.Query().Get("guid")
	at := request.URL.Query().Get("at")
	rt := request.URL.Query().Get("rt")
	if rt == "" || at == "" || guidQ == "" {
		serLogs.errorLog.Print(writer, http.StatusBadRequest)
		return
	}
	msg, err := utils.ValToken(guidQ, at, rt)
	if err != nil {
		serLogs.errorLog.Println(writer, err)
		return
	}
	if msg != "" {
		_, _ = fmt.Fprintln(writer, msg)
		return
	}
	aToken, rToken, err := serLogs.getTokens(guidQ)
	if err != nil {
		serLogs.errorLog.Print(writer, err)
		return
	}
	base64RToken := base64.StdEncoding.EncodeToString([]byte(rToken.Token))
	_, _ = fmt.Fprintln(writer, aToken.Token+"\n")
	_, _ = fmt.Fprintln(writer, rToken.Token+"\n")
	_, _ = fmt.Fprintln(writer, base64RToken+"\n")
}
