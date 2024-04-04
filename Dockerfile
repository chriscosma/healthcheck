FROM golang:1.22.1 as BUILD

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /healthcheck

###############################################################

FROM alpine:latest as RUNTIME

WORKDIR /

COPY --from=BUILD /healthcheck /healthcheck

ENTRYPOINT [ "/healthcheck" ]