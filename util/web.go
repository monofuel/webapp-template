package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

type AppError struct {
	Error   error
	Message string
	Code    int
}

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errorResponse struct {
		Err     error  `json:"err"`
		Message string `json:"message"`
		Code    int    `json:"status"`
	}
	if e := fn(w, r); e != nil {
		fmt.Printf("error: %v, %v", e.Message, e.Error)
		//http.Error(w, e.Message, e.Code)
		errorResponse.Err = e.Error
		errorResponse.Message = e.Message
		errorResponse.Code = e.Code
		w.WriteHeader(e.Code)
		json.NewEncoder(w).Encode(errorResponse)
	}
}

func GetPageLimitParameters(r *http.Request) (int, int, error) {
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return 0, 0, err
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return 0, 0, err
	}
	return page, limit, nil
}
