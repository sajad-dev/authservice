#!/bin/bash

AUTHENTICATION="$(pwd)/authentication"
AUTHORIZATION="$(pwd)/authorization"
OUTPUT="./getway/services.pb"

generate_protos() {
    local DIR="$1"
    find "$DIR" -type f -name "*.proto" | while read file; do
        echo "Found: $file"
        protoc -I="$DIR" -I="$GOOGLEAPIS_DIR" \
            --go_out="$DIR" \
            --go-grpc_out="$DIR" \
            "$file"
    done
}

generate_protos "$AUTHENTICATION"
generate_protos "$AUTHORIZATION"



FILES_PROTO=""

generate_descriptor() {
    local DIR="$1"
    while IFS= read -r file; do
        echo "Processing: $file"
        FILES_PROTO="$FILES_PROTO $file"
    done < <(find "$DIR" -type f -name "*.proto")
}

generate_descriptor "$AUTHENTICATION"
generate_descriptor "$AUTHORIZATION"

protoc -I="$AUTHENTICATION" -I="$AUTHORIZATION" -I="$GOOGLEAPIS_DIR" \
    --include_imports --include_source_info \
    --descriptor_set_out="$OUTPUT" \
    $FILES_PROTO

echo "Generated descriptor set: $OUTPUT"

