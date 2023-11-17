# Go binaries are standalone, so use a multi-stage build to produce smaller images.
FROM golang:1.21-alpine as builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app

FROM rxmllc/alpine-aws-cli
COPY --from=builder /app /usr/local/bin/go-merkle
ENTRYPOINT ["/usr/local/bin/go-merkle"]
