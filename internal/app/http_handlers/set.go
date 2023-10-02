package http_handlers

import (
	"context"
	"net/http"
	desc "service-component/pkg/service-component/pb"
)

func (h *HttpService) Set(w http.ResponseWriter, req *http.Request) {
	key := takeSingleFromQuery(w, req, "key")
	if key == nil {
		return
	}

	value := takeSingleFromQuery(w, req, "value")
	if value == nil {
		return
	}
	_, err := h.grpcService.Set(context.Background(), &desc.SetRequest{
		Key:   *key,
		Value: *value,
	})
	if err != nil {
		processGrpcError(w, err)
		return
	}
}
