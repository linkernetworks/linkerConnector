GO=$(shell which go)
compile-static-amd64:
	$(GO) get -d && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -o docker/linkerConnector .
docker-build:
	docker build -t linkerrepository/linker_connector docker/
docker-push:
	docker push linkerrepository/linker_connector
container-run:
	docker run -d --name linker_connector --net host -v /proc:/linker/porc linkerrepository/linker_connector /linkerConnector -r /linker/proc -i 2000 -f -c http://localhost:8080
container-stop:
	docker stop linker_connector
container-clean:
	docker rm linker_connector
clean:
	rm docker/linkerConnector
