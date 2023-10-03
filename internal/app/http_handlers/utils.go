package http_handlers

import (
	"fmt"
	"google.golang.org/grpc/status"
	"net/http"

	"google.golang.org/grpc/codes"
)

var GrpcCodeToHttp = map[codes.Code]int{
	codes.NotFound: http.StatusNotFound,
	codes.Internal: http.StatusInternalServerError,
}

func takeSingleFromQuery(w http.ResponseWriter, req *http.Request, key string) *string {
	value, ok := req.URL.Query()[key]
	if !ok {
		http.Error(w, fmt.Sprintf("argument %s not found", key), http.StatusBadRequest)
		return nil
	}
	if len(value) != 1 {
		http.Error(w, fmt.Sprintf("too many values(%v) in argument %s", value, key), http.StatusBadRequest)
		return nil
	}
	return &value[0]
}

func processGrpcError(w http.ResponseWriter, err error) {
	httpCode := GrpcCodeToHttp[status.Code(err)]
	http.Error(w, err.Error(), httpCode)
	return
}
