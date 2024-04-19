FROM golang:alpine as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /build
COPY go.mod . 
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /main cmd/app/main.go
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/cmd/app/main /bin/main
COPY --from=builder build/prod.env /.env
COPY --from=builder build/config/prod.yaml /prod.yaml
ENTRYPOINT ["bin/main"]
