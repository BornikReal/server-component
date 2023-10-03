.PHONY: generate
generate:
	protoc \
	--proto_path=api/v1/ \
	--go_out=pkg \
	--go-grpc_out=pkg \
	--grpc-gateway_out pkg \
	--openapiv2_out pkg \
	--grpc-gateway_opt generate_unbound_methods=true \
	api/v1/*.proto