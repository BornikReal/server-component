package http_handlers

import (
	"context"
	"net/http"
	desc "service-component/pkg/service-component/pb"
)

func (h *HttpService) Get(w http.ResponseWriter, req *http.Request) {
	key := takeSingleFromQuery(w, req, "key")
	if key == nil {
		return
	}
	resp, err := h.grpcService.Get(context.Background(), &desc.GetRequest{
		Key: *key,
	})
	if err != nil {
		processGrpcError(w, err)
		return
	}

	_, err = w.Write([]byte(resp.Value))
	if err != nil {
		http.Error(w, "error while writing", http.StatusInternalServerError)
		return
	}
}
