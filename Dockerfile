FROM golang:1.23.5 As development
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 4000
CMD reflex -g '*.go' go run main.go --start-service

FROM golang:1.23.5 As builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM alpine:latest As production
RUN apk add --no-cache ca-certificates
COPY --from=builder app/app .
EXPOSE 4000
CMD ./app