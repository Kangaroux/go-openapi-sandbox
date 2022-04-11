#!/bin/bash

set -e
cd ${0%/*}

function swagger() {
    local cwd=$(pwd)

    docker run \
        --rm \
        -it \
        -e GOPATH=$(go env GOPATH):/go \
        -e API_VERSION \
        -v $cwd:$cwd \
        -w $cwd \
        quay.io/goswagger/swagger \
        "$@"
}

case "$1" in
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
    swagger generate spec -o $out_file

    # Restore the backup
    mv $docs_backup docs/docs.go
    ;;
esac
