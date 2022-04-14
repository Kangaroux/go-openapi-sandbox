#!/bin/bash

set -e
cd ${0%/*}

conn_str='postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST/$POSTGRES_DB?sslmode=disable'

function migrate() {
    args="$@"
    docker-compose run --rm api \
        bash -c "migrate \
            -path migrations \
            -database `echo $conn_str` \
            $args"
}

function swagger() {
    local cwd=$(pwd)

    docker run \
        --rm \
        -it \
        -e GOPATH=$(go env GOPATH):/go \
        -e API_VERSION \
        -e SWAGGER_GENERATE_EXTENSION \
        -v $cwd:$cwd \
        -w $cwd \
        quay.io/goswagger/swagger \
        "$@"
}

case "$1" in
"migrate")
    shift
    migrate "$@"
    ;;

"migrate:new")
    shift
    migrate \
        create \
        -dir migrations \
        -ext sql \
        "$@"
    ;;

"psql")
    shift
    args="$@"
    docker-compose run --rm db bash -c "psql $conn_str $args"
    ;;

"swagger")
    shift
    swagger "$@"
    ;;

"spec")
    docs_backup=docs/.docs.go
    out_file=swagger.yml

    # Make a backup of docs.go and substitute the API_VERSION env
    mv docs/docs.go $docs_backup
    API_VERSION=`head -n 1 VERSION` envsubst < $docs_backup > docs/docs.go

    # Touch the file first so it has the correct permissions
    touch $out_file
    SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o $out_file

    # Restore the backup
    mv $docs_backup docs/docs.go
    ;;
esac
