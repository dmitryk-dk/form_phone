FROM golang:1.9-alpine

WORKDIR "/home/app/"

RUN apk update && \
    apk add git curl yarn && \
    mkdir /opt && \
    rm -rf /var/cache/apk/*

COPY . .

RUN go get -v -t -u -d github.com/dmitryk-dk/form_phone

RUN curl https://glide.sh/get | sh && \
    glide i

RUN rm -rf node_modules && \
    yarn install --non-interactive --silent
    yarn build

EXPOSE 3000

CMD ["go", "run", "main.go"]
