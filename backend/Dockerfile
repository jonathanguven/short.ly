FROM golang:1.23.0-alpine

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /backend/cmd/server

RUN go build -o shortly .

EXPOSE 8080

CMD ["/backend/cmd/server/shortly"]