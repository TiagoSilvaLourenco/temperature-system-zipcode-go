FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o temperature_system

FROM scratch
WORKDIR /app
COPY --from=build /app/temperature_system .
ENTRYPOINT [ "./temperature_system" ]