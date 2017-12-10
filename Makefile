
run: build
	go run main.go

build:
	make -j 2  buildServer buildClient
buildServer: webapp-template
webapp-template:
	go build

buildClient: public/js/index.js
public/js/index.js:
	webpack

watch:
	make -j 2  watchServer watchClient

watchServer:
	gin --port 8085

watchClient:
	webpack -d -w

setup: setupGo setupNode

# installing dependencies manually so docker can cache them
# otherwise `go get` would re-install on every code change
setupGo:
	go get github.com/codegangsta/gin
	go get github.com/golang/lint/golint

setupNode:
	npm install

test:
	make -j 2 testGo testNode

testGo:
	golint ./...
	go tool vet ./
	go test
	
testNode:
	tslint frontend/

cover:
	go test -coverprofile=coverage.out
	go tool cover  -html=coverage.out

clean:
	rm public/js/*.js || true
	rm public/js/*.js.map || true
	rm -r ./node_modules || true