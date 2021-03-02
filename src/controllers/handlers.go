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

func (serLogs *SerLogs) FirstRoute(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")
	if guid == "" {
		serLogs.errorLog.Print(w, http.StatusBadRequest)
		return
	}
	aToken, rToken, err := serLogs.getTokens(guid)
	if err != nil {
		serLogs.errorLog.Print(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprint(w, aToken.Token+"\n")
	_, _ = fmt.Fprint(w, rToken.Token+"\n")
}

func (serLogs *SerLogs) SecondRoute(w http.ResponseWriter, r *http.Request) {
	guidQ := r.URL.Query().Get("guid")
	at := r.URL.Query().Get("at")
	rt := r.URL.Query().Get("rt")
	if rt == "" || at == "" || guidQ == "" {
		serLogs.errorLog.Print(w, http.StatusBadRequest)
		return
	}
	msg, err := utils.ValidTokens(guidQ, at, rt)
	if err != nil {
		serLogs.errorLog.Print(w, err)
		return
	}
	if msg != "" {
		_, _ = fmt.Fprintln(w, msg)
		return
	}
	aToken, rToken, err := serLogs.getTokens(guidQ)
	if err != nil {
		serLogs.errorLog.Print(w, err)
		return
	}
	base64RToken := base64.StdEncoding.EncodeToString([]byte(rToken.Token))
	_, _ = fmt.Fprintln(w, "access-token: "+aToken.Token+"\n")
	_, _ = fmt.Fprintln(w, "refresh-token: "+rToken.Token+"\n")
	_, _ = fmt.Fprintln(w, "base64 refresh-token: "+base64RToken+"\n")
}

func (serLogs *SerLogs) getTokens(guid string) (typesF.AToken, typesF.RToken, error) {
	rToken := typesF.RToken{
		GUID:          guid,
		Token:         utils.GenerateToken(guid, 20),
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
