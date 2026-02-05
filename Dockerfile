FROM golang:1.25-alpine AS builder

WORKDIR /campaign-metadata

COPY go.* .

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -o ./app cmd/main.go

FROM scratch

COPY --from=builder /campaign-metadata/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/app" ]
