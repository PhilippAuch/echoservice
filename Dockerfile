##
## Build
##
FROM golang:1.17.4-alpine3.15 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o goapp-compiled .

##
## Deploy
##
FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /
COPY --from=builder /app/goapp-compiled .

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./goapp-compiled"]