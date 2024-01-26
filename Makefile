
PROTO_DIR=proto
OUT_DIR=pb

proto:
	rm -f $(OUT_DIR)/*.go
	rm -f doc/swagger/*.swagger.json
	protoc \
	--proto_path=$(PROTO_DIR) \
	--proto_path=third_party \
	--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
	--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(OUT_DIR) --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=book_store \
	$(PROTO_DIR)/*.proto


mysql:
	docker run \
	--name mysql \
	-p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=my-secret-pw \
	-e MYSQL_DATABASE=bookstore \
	-d mysql:8.0

consul:
	docker run \
	--name=dev-consul \
	-p 8500:8500 \
	-d hashicorp/consul:1.10.0

# http://localhost:8500/ui/dc1/services

# $ docker exec -it bash  <container_id>
# $ docker logs <container_id>


start:
	docker-compose up -d

end:
	docker-compose -f ./docker-compose.yml down

login:
	docker exec -it mysql bash
# mysql -u root -p
# create database bookstore;


help:
	@echo "make gen - 生成 pb 及 grpc 代码"

.PHONY: proto store login help


.PHONY: compile
compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
