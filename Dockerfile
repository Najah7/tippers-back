FROM golang:1.20

WORKDIR /app
# webpに変換する用
RUN  apt-get update -y
RUN apt-get install libwebp-dev -y
RUN go install github.com/cosmtrek/air@latest

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

CMD ["air", "-c", ".air.toml"]