FROM golang:1.21

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o ./comm-relay

EXPOSE 11073

CMD [ "./comm-relay" ]
