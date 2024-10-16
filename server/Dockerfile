# Install dependencies and compile to a binary
FROM golang:1.20 as builder

LABEL maintainer="Toby Scott <hi@tobyscott.dev>"

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

# Required for sqlite3
ENV CGO_ENABLED=1

ENV GIN_MODE=release
ENV GOOS=linux GOARCH=arm64
RUN go build -o /bin/localflix-server .

# Step 2: build a small image that runs the binary
FROM alpine:3.18.2

# ####
# ## To ensure the application doesn't run as root, we
# ## create a new user and group that it will run as.
# ####
# RUN addgroup localflix && adduser -D -G localflix localflix

# # copy static executable and change the owner to be the user that will run the application
# COPY --chown=localflix:localflix --from=builder /bin/localflix-server /bin/localflix-server

# USER localflix:localflix

EXPOSE 8080

CMD ["/bin/localflix-server"]
