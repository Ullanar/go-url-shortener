FROM golang:alpine3.19

WORKDIR /app
COPY . .
RUN go build -o server ./cmd/main.go

EXPOSE 8000
CMD ["./server"]
