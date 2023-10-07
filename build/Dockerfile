FROM golang:1.19

WORKDIR /app

COPY ../. .

RUN mkdir -p bin && cd cmd/highload && go build -o ../../bin/server

CMD ./bin/server