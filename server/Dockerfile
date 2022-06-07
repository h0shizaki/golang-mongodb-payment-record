FROM golang:1.18.1-alpine
RUN apk add git
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

ENV ENVIRONMENT=Production
ENV MONGODB_URI=mongodb://mongo:cosmos@localhost:27017/
ENV PORT=4000

RUN go build -o ./build/server.exe ./cmd/api

CMD ["./build/server.exe"]