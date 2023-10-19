# comm-relay

<img src=https://i.imgur.com/4csVefQ.png width=300 />

A server that handles notifying user using preferred notification channels.

## Running

You can run using the published Docker image.

```bash
docker run --name comm-relay \
    -e MYSQL_DSN='root:password@tcp(localhost:3306)/comm_relay?charset=utf8mb4&parseTime=True&loc=Local' \
    -e COOKIE_SECRET='cookie' \
    -e LISTEN_ADDRESS='0.0.0.0:11073' \
    -e PASSWORD='password' \
    -e SMTP_HOST='smtp.gmail.com' \
    -e SMTP_PORT=587 \
    -e SMTP_USER='hello@gmail.com' \
    -e SMTP_PASSWORD='your-app-password' \
    -e SMTP_FROM='hello@gmail.com' \
    -p 11073:11073 \
    -d \
    ghcr.io/antin0de/comm-relay:latest
```

Then visit `localhost:11073`. Make sure you replace `MYSQL_DSN`, `COOKIE_SECRET` and `PASSWORD` with your own values.

## Developing

### Running Local Server

First copy-paste the `.env.example` file to `.env`, and modify `MYSQL_DSN` to point to your own MySQL server.

Then run

```bash
go run .
```
