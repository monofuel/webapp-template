
run: build
	go run main.go

build: buildServer buildClient
buildServer:
	go build
buildClient:
	BABEL_ENV=production gulp

watch: watchServer watchClient

watchServer:
	gin --port 8085

watchClient:
	watchify frontend/init.jsx -t babelify -p livereactload -o public/js/index.js

setup:
	go get github.com/codegangsta/gin
	go get
	rm -r ./node_modules
	npm install

test:
	#go vet .
	go test
	flow check
