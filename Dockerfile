# syntax=docker/dockerfile:1.2

# Stage 1: build the client files
# get node image and alias it as client-builder
FROM node:16.15.1-alpine3.16 AS client-builder
# set our docker image directory
WORKDIR /client
# copy our package.json and package-lock.json to our current image directory /client
COPY /client/package.json /client/package-lock.json ./
# similar to npm i but to ran in automated environment
RUN npm ci
# copy the whole content of client folder in our current image directory
COPY /client .
# run the build command to generate our site in dist folder
RUN npm run build

# Stage 2: build the server binary
# get golang image and alias it as server-builder
FROM golang:1.18.3-alpine3.15 AS server-builder
# define a variable for the engine to use (http, echo, fiber)
ARG ENGINE_NAME=http
# update to get the latest from the alpine linux distro and add git
RUN apk update && apk upgrade && apk --update add git
# set our directory to be named builder
WORKDIR /builder
# copy our local go.mod and go.sum file to the image directory folder /builder
COPY go.mod go.sum ./
# pull all the depedencies mentioned in the go.mod file
RUN go mod download
# copy all the folders to the image directory folder /builder
COPY . .
# copy using client builder info the local dist folder to the image dist folder 
COPY --from=client-builder /client/dist ./client/dist
# build the server binary using the server engine variable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o engine ./server/${ENGINE_NAME}/main.go

# Stage 3: run the binary
# get a linux distro that is lighter than alpine
FROM gcr.io/distroless/static
ENV APP_PORT=5050
WORKDIR /app
COPY --from=server-builder --chown=nonroot:nonroot /builder/engine .
EXPOSE ${APP_PORT}
ENTRYPOINT [ "./engine" ]