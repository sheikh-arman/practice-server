FROM golang:latest
RUN mkdir /build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
RUN go build

EXPOSE 8080

CMD [ "./practice-server"]
