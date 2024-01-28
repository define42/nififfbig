FROM golang:1.21.6 as builder

WORKDIR /app/
COPY go.sum .
COPY go.mod .
RUN go mod download
COPY main.go main.go
RUN CGO_ENABLED=0 go build -o /main
RUN chmod 777 /main

FROM scratch
COPY --from=builder /main /main
ENTRYPOINT ["/main"]

