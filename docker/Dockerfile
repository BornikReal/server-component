FROM golang:1.19
WORKDIR /app
RUN git clone https://github.com/BornikReal/server-component.git
WORKDIR server-component
RUN git pull && make build
WORKDIR bin
CMD ./client