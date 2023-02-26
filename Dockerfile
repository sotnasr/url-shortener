FROM golang:1.20

WORKDIR /usr/src/url-shortener

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/url-shortener /usr/src/url-shortener/cmd

EXPOSE 9000
RUN sleep 15

CMD ["url-shortener"]