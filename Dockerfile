FROM golang

WORKDIR /app
#RUN apk --no-cache add ca-certificates gcc git musl-dev
COPY . .
RUN go mod download

RUN cd cmd && CGO_ENABLED=1 go build -a -ldflags '-linkmode external -extldflags "-static"' -o webserver

FROM scratch

COPY --from=0 /app/cmd/webserver /

ENTRYPOINT ["/webserver"]