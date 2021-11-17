FROM golang:1.17

RUN mkdir -p /streamer

WORKDIR /streamer
COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o go_streams cmd/main.go

ENTRYPOINT [ "./go_streams"]