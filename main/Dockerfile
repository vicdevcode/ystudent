FROM golang:alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main cmd/app/main.go
FROM scratch
ARG CONFIG_FILE=prod.yaml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main
COPY --from=builder build/prod.env /.env
COPY --from=builder build/config/${CONFIG_FILE} /config.yaml
ENTRYPOINT ["bin/main"]
