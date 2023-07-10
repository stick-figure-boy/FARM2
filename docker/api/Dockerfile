FROM golang:1.19-buster

ENV GO111MODULE=auto

WORKDIR /go/api

COPY . .

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/gin-gonic/gin && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go install github.com/swaggo/swag/cmd/swag@latest && \
  go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./go.mod ./go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
