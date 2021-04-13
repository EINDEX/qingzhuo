FROM golang:1.16 AS BUILDER
WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server


FROM scratch AS RUNNER

COPY --from=BUILDER /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=BUILDER /app/server /app/
CMD ["/app/server"]