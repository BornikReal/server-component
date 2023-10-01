.PHONY: generate
generate:
	protoc \
	--proto_path=api/v1/ \
	--go_out=pkg \
	--go-grpc_out=pkg \
	api/v1/*.proto