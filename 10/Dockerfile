FROM golang:alpine AS build
WORKDIR /go/src/github.com/k8s-homework/10
ADD . .
RUN go build main.go

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /go/src/github.com/k8s-homework/10/main /app/main
EXPOSE 80
ENTRYPOINT ["/app/main"]