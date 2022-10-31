.PHONY: protos

lint:
	golangci-lint run

migrateup: 
	export POSTGRES_URL='postgres://postgres:postgres@localhost:5432/cnord?sslmode=disable' && \
	migrate -database ${POSTGRESQL_URL} -path migrations up
	
migratedown: 
	export POSTGRES_URL='postgres://postgres:postgres@localhost:5432/cnord?sslmode=disable' && \
	migrate -database ${POSTGRESQL_URL} -path migrations up
		
protos:
	protoc --go_out=proto/ -I proto/ proto/user.proto
	protoc --go-grpc_out=proto/ -I proto/ proto/user.proto

run:
	export POSTGRES_URL='postgres://postgres:postgres@localhost:5432/cnord?sslmode=disable' && \
	go build -o bin/main ./cmd/main.go && \
	./bin/main

testapi:
	./test/saveuser.sh && \
	./test/getuserbyid.sh

