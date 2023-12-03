PWD=$(shell pwd)
SERVICE=eater-svc
MIGRATION_PATH=${PWD}/src/infrastructure/migrations
PROTOS_PATH=$(PWD)/src/infrastructure/protos

migration-file:
	Project setup 2
	docker run -v ${MIGRATION_PATH}:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(NAME)
clear:
	rm -rf ${PWD}/bin/${SERVICE}
bin: clear
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
	chmod +x ${PWD}/bin/${SERVICE}
bin-linux: clear
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
	chmod +x ${PWD}/bin/${SERVICE}
bin-windows: clear
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
	chmod +x ${PWD}/bin/${SERVICE}

add-protos-submodule:
	git submodule add https://github.com/Azamjon99/services-proto.git ./src/infrastructure/protos

pull-protos-submodule:
	git submodule update --recursive --remote

gen-restaurant-support:
	protoc \
	--go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PWD)/src/infrastructure/protos/restaurant_support \
	$(PWD)/src/infrastructure/protos/restaurant_support/*.proto

docker-build:
	docker build --rm -t restaurant-support-svc -f ${PWD}/deploy/Dockerfile .

compose-up:
	docker-compose -f ./deploy/docker-compose.yml up

compose-down:
	docker-compose -f ./deploy/docker-compose.yml down