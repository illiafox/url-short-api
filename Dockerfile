# build stage
FROM golang:1.18.2-alpine AS build-env
RUN apk --no-cache add build-base git curl
ADD . /temp
RUN cd /temp/app/cmd/app && go build -o server

# final stage
FROM alpine
COPY --from=build-env /temp/app/cmd/app/ /app/

WORKDIR /app

ENTRYPOINT "./server" $ARGS