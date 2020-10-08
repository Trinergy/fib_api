FROM golang:1.14-alpine as build-env


WORKDIR /go/server

# COPY go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/server
FROM scratch 
COPY --from=build-env /go/bin/server /go/bin/server
ENTRYPOINT ["/go/bin/server"]
EXPOSE 8080