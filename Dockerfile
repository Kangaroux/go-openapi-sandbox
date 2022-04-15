FROM golang:1.18

# Install the `migrate` tool to `/usr/local/bin`
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz \
    | tar xz -C /usr/local/bin migrate

# Install `entr` for auto restart on code changes
RUN apt-get update && apt-get install -y entr

RUN useradd -m -s /bin/bash apiuser

USER apiuser
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Start the server and watch for changes
CMD \
while sleep 0.5s; do \
    find . -name '*.go' | entr -r go run cmd/server/main.go; \
done
