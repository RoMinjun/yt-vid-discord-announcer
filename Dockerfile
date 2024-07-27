FROM golang:1.22.5-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /yt-vid-discord-announcer ./cmd/server
EXPOSE 8080
CMD ["/yt-vid-discord-announcer"]
