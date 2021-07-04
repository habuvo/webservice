FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git gcc ca-certificates
COPY . .
RUN cd /cmd && go build -o webserver

FROM scratch

COPY --from=build-env /cmd/webserver /
ENTRYPOINT ["/main"]