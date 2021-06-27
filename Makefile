gen-proto:
	protoc --proto_path=$(PWD)/internal/domains/$(SERVICE)/ --go_out=$(PWD)/internal/shared/infrastructure/grpc/$(SERVICE)_service --go-grpc_out=$(PWD)/internal/shared/infrastructure/grpc/$(SERVICE)_service $(PWD)/internal/domains/$(SERVICE)/$(SERVICE).proto
test:
	go test ./... -coverprofile=coverage.out