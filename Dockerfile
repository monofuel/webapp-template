FROM debian:jessie

RUN useradd -ms /bin/bash webapp
ENV MONO_USER='webapp'
ENV GO_PACKAGES="github.com/codegangsta/gin"
ENV NODE_TOOLS="webpack typescript nodemon"

ADD mono-scripts /home/webapp/go/src/github.com/monofuel/webapp-template/mono-scripts
ADD Makefile /home/webapp/go/src/github.com/monofuel/webapp-template/
RUN chown -R webapp.webapp /home/webapp/go
WORKDIR /home/webapp/go/src/github.com/monofuel/webapp-template/mono-scripts/setup

RUN bash golang.bash
RUN bash nodejs.bash

USER webapp

# use a login shell with local profile
SHELL ["/bin/bash", "-c", "-l"]

WORKDIR /home/webapp/go/src/github.com/monofuel/webapp-template
ADD package.json /home/webapp/go/src/github.com/monofuel/webapp-template/

RUN make setup

ADD . /home/webapp/go/src/github.com/monofuel/webapp-template

RUN make build

CMD [ "./webapp-template" ]