FROM golang:1.22-alpine
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o app .
CMD ["/app/app"]