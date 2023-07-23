# Install dependencies and compile to a binary
FROM golang:1.20-alpine as builder

LABEL maintainer="Toby Scott <hi@tobyscott.dev>"

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

ENV GIN_MODE=release
ENV GOOS=linux GOARCH=amd64
RUN go build -o /bin/localflix-api .

# Step 2: build a small image that runs the binary
FROM alpine:3.14

####
## To ensure the application doesn't run as root, we
## create a new user and group that it will run as.
####
RUN addgroup localflixapi && adduser -D -G localflixapi localflixapi

# copy static executable and change the owner to be the user that will run the application
COPY --chown=localflixapi:localflixapi --from=builder /bin/localflix-api /bin/localflix-api

USER localflixapi:localflixapi

EXPOSE 8080

CMD ["/bin/localflix-api"]
