FROM golang:1.17.2 as build-image

ENV GOPATH=/go \
    HOME=/go \
    REPO=demo-backend

COPY . $HOME/src/$REPO
RUN cd /$HOME/src/$REPO && make get && make build
WORKDIR /$HOME/bin

EXPOSE 8080
CMD ./sample