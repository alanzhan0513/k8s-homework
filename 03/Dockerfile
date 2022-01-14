FROM golang:alpine AS build
WORKDIR /src
ADD . /src
RUN cd /src && go build main.go

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /src/main /app/main
EXPOSE 80
ENTRYPOINT ["/app/main"]
