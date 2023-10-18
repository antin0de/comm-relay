FROM golang:1.21

WORKDIR /app

COPY . /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

RUN go build -o ./comm-relay

EXPOSE 11073

CMD [ "./comm-relay" ]
