all: bin/srv
test: lint unit-test

PLATFORM=local


.PHONY: proto
proto:
	cd message &&  protoc --go_out=. message.proto

gen-tests:
	cd internal/net && gotests -all -w http_servant.go \
		&& gotests -all -w protobuf_servant.go \
		&& gotests -all -w socket_server.go
	cd internal/utils && gotests -all -w client_dict.go \
		&& gotests -all -w safe_counter.go

tests:
	go test ./... -v

.PHONY: bin/srv
bin/srv:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker build . --target lint
