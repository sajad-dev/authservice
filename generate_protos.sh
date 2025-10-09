#!/bin/bash
set -e

PROTO_DIR="internal/domain"

PROTO_FILES=$(find $PROTO_DIR -name "*.proto")

for f in $PROTO_FILES; do
  echo "Processing $f ..."

  protoc -I=$GOOGLEAPIS_DIR -I=. \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
    $f
done

echo "✅ All proto files processed!"

