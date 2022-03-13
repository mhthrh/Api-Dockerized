# syntax= docker /dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /main



CMD ["/main"]


#docker run --publish 8585:8585 kir-khar