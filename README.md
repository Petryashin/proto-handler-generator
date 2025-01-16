# proto-handler-generator
Handler and UseCases generator  by proto

# Usage
1. add proto files to `protos` directory
Warnings: files must be in .proto format and service and messages
must be separated and not in one file
files must be in the same directory
Example: `service_name_service.proto` and `service_name_messages.proto`
2. add .env and fill variables from .env.example
3. run `go run cmd/proto2handler.go`
4. take code from `generated` directory and use it in your project