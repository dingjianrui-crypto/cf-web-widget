FROM golang:1.25-alpine AS build
WORKDIR /app
COPY server.go .
RUN CGO_ENABLED=0 go build -o /server server.go

FROM alpine:3.20
COPY --from=build /server /server
EXPOSE 8080
CMD ["/server"]