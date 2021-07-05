package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	if uid,ok := vars["uid"];ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{Uid: uid}, nil
	}
	return nil, errors.New("param error")
}

func EncodeUserResponse(c context.Context,w http.ResponseWriter, response interface{})  error {
	w.Header().Set("Content-type","application/json")
	return json.NewEncoder(w).Encode(response)
}