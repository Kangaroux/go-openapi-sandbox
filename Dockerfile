FROM golang:1.18

# Install the `migrate` tool to `/usr/local/bin`
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz \
    | tar xz -C /usr/local/bin migrate

RUN useradd -m -s /bin/bash apiuser

USER apiuser
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD go run cmd/server/main.go
