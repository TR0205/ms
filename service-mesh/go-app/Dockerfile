FROM golang:1.22-alpine

# バージョンをbuild時に取得
ARG APP_VERSION=Green
ENV APP_VERSION=${APP_VERSION}

WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o app .

EXPOSE 8080
CMD ["./app"]
