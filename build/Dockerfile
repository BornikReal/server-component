FROM golang:1.19

WORKDIR /app

COPY ../. .

RUN mkdir -p bin && cd cmd/server && go build -o ../../bin/storage_service

CMD ./bin/server