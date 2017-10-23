FROM alpine:3.6

ENV HOME "/home/app/"

WORKDIR "${HOME}/src"

RUN apk update \
  && apk add curl bash tar \
  # Yarn
  && apk add yarn \
  # GO
  && curl https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz | tar xzf - -C / && \
     mv /go /goroot \
  && rm -rf /var/cache/apk/*

ENV GOROOT=/goroot \
    GOPATH=${HOOME} \
    GOBIN=${HOME}/bin \
    PATH=${PATH}:/goroot/bin:${HOME}/bin

COPY . .

RUN rm -rf {node_modules,Dockerfile} && yarn install --non-interactive --silent

EXPOSE 3000

CMD ["yarn", "docker-build"]
