# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY data/ ./data
COPY handlers/ ./handlers
COPY model/ ./model
COPY pkg/ ./pkg

WORKDIR /app/cmd/quizapi
RUN env GOOS=linux GOARCH=amd64 go install

EXPOSE 9090
ENV PATH="${PATH}:/go/bin/linux_amd64"

CMD [ "quizapi" ]