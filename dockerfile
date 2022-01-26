FROM golang:1.17-alpine

WORKDIR /

COPY go.mod . 

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o program

CMD ./program
