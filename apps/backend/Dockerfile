FROM golang:1.24.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# airのインストール
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air"]