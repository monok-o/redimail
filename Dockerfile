# Building the binary of the App
FROM golang:1.16.2 AS build

WORKDIR /go/src/redimail

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod tidy

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/redimail/app .

ENV DOMAIN=http://localhost:8080
ENV SERVER_HOST=127.0.0.1
ENV SERVER_PORT=8080
ENV DB_PATH=/tmp/redimail
ENV DB_FILE=database.db

# Exposes port 3000 because our program listens on that port
EXPOSE 8080

CMD ["./app"]