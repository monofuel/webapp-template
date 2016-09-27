monofuel's example webapp template

go backend, react frontend

I threw this together to provide a basic start for making a webapp with go and react.
the makefile makes it real easy to get started.
just run make setup once, make watch and just start coding.

## getting started
- install go, and setup your gopath
- install nodejs
- install make
- check out the github project into $GOPATH/src/github.com/monofuel/webapp-template
- `make setup` to prepare the project initially
- `make watch -j2` to start running both the frontend and backend with hot-reload

## make targets
- `run` build and run the server
- `build` build everything
- `buildServer` build only the server
- `buildClient` build only the frontend
- `watch` runs the server with hot reload for backend and frontend (make sure to run with -j2)
- `watchServer` runs the server with hot reload
- `watchClient` builds the frontend with hot reload
- `setup` runs go get and npm install to set things up
- `test` check/test frontend and backend

![alt text](screenshots/index.png "screenshots rock")

## TODO
- dockerfile for building and running
- add database examples
	- logging system
	- alerting system
	- admin dashboard
	- user auth system
