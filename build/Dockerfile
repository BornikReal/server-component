FROM golang:1.19

WORKDIR /app

COPY ../. .

RUN mkdir -p bin && cd cmd/highload && go build -o ../../bin/client

CMD ./bin/client