mock:
	mockgen -source=internal/ports/resource/service.go -destination=internal/core/services/mock/service_mock.go -package=services

mock-repository:
	mockgen -source=internal/ports/resource/repository.go -destination=internal/core/services/mock/repository_mock.go -package=services

test:
	go test -v ./...

run:
	go run .