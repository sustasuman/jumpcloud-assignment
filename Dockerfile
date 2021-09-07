FROM golang:alpine
FROM golang:alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /jumpcloud
EXPOSE 8080
CMD ["/jumpcloud"]