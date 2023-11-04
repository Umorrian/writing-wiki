FROM golang:1.21 as build
WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/writingwiki ./main.go

FROM cgr.dev/chainguard/glibc-dynamic:latest

COPY --from=build /bin/writingwiki /bin/writingwiki
COPY ./static /static

VOLUME ["/dbdata"]
ENV VOLUME_PATH=/dbdata
ENV STATIC_PATH=/static
USER root

ENTRYPOINT ["/bin/writingwiki"]