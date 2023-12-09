FROM golang:alpine as builder
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o application ./cmd/

FROM alpine:latest
WORKDIR /usr/bin
COPY --from=builder /app/application /usr/bin/application
ENV PORT=8000
EXPOSE ${PORT}
CMD [ "application" ]