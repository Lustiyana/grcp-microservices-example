#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"

echo "Generating protobuf files..."

# Generate User proto
echo "Generating user.proto..."
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user/user.proto

if [ $? -eq 0 ]; then
    echo "✓ user.proto generated successfully"
else
    echo "✗ Failed to generate user.proto"
    exit 1
fi

# Generate Order proto
echo "Generating order.proto..."
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/order/order.proto

if [ $? -eq 0 ]; then
    echo "✓ order.proto generated successfully"
else
    echo "✗ Failed to generate order.proto"
    exit 1
fi

echo ""
echo "All proto files generated successfully! 🚀"