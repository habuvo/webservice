FROM golang:alpine AS build-env

WORKDIR /app
RUN apk --no-cache add ca-certificates gcc git musl-dev
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .

RUN cd cmd && go build -o webserver

FROM scratch

COPY --from=build-env /app/cmd/webserver /
ENTRYPOINT ["/webserver"]