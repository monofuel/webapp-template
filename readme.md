monofuel's example webapp template

go backend, react frontend

I threw this together to provide a basic start for making a webapp with go, typescript and react.

## Getting Started (bare metal)
- install go, and setup your gopath
	- on debian, you can run `./mono-scripts/setup/golang.bash`
- install nodejs
	- on debian, you can run `./mono-scripts/setup/nodejs.bash`
- install make
- check out the github project into `$GOPATH/src/github.com/monofuel/webapp-template`
- `make setup` to prepare the project initially
- `make watch -j2` to start running both the frontend and backend with hot-reload

## Getting Started (Docker)
- `docker-compose build`
- `docker-compose up`
- if dependencies are changed, re-run docker-compose build

## Getting Started (Vagrant)
- `vagrant up`
- if dependencies are changed, re-run `vagrant provision` for fast updates

## make targets
- `run` build and run the server
- `build` build everything
- `buildServer` build only the server
- `buildClient` build only the frontend
- `watch` runs the server with hot reload for backend and frontend
- `watchServer` runs the server with hot reload
- `watchClient` builds the frontend with hot reload
- `setup` runs go get and npm install to set things up
- `test` check/test frontend and backend

![alt text](screenshots/index.png "screenshots rock")

## Visual Studio Code
- recommended plugins
  - Go
  - TSLint
- configuration
  - set `"editor.formatOnSave": true`


## TODO
- dockerfile for building and running
- add database examples
	- logging system
	- alerting system
	- admin dashboard
	- user auth system
