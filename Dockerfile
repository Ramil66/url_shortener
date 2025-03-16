FROM golang:1.23

RUN go version
ENV GOPATH=/

WORKDIR /


COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go mod download
RUN go build -o url-shortener ./cmd/main.go

CMD ["./url-shortener"]
