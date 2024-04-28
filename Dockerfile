FROM golang:latest as build-env

WORKDIR /app
COPY . .
RUN go build -o ./bin/app ./internal/app/app.go

FROM gcr.io/distroless/base
WORKDIR app/
COPY --from=build-env /app/bin/app .
EXPOSE 80
CMD ["./app"]