FROM golang:1.9-alpine

WORKDIR "/home/app/"

RUN apk update && \
    apk add git curl yarn && \
    mkdir /opt && \
    rm -rf /var/cache/apk/*

COPY . .

RUN mkdir -p $GOPATH/src/github.com/dmitryk-dk/form_phone
COPY . $GOPATH/src/github.com/dmitryk-dk/form_phone/

RUN curl https://glide.sh/get | sh && \
    glide i

RUN rm -rf node_modules && \
    yarn install --non-interactive --silent

EXPOSE 3000

CMD ["yarn", "docker-build"]
