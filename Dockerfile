FROM golang:1.18

WORKDIR /app

# Install the `migrate` tool to `/usr/local/bin`
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz \
    | tar xvz -C /usr/local/bin

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD go run .
