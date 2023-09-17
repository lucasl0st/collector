FROM golang:1.21-alpine AS builder

RUN apk add git

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN mkdir /new_tmp

RUN go mod download

COPY . ./

RUN go build -o /collector

FROM scratch

LABEL org.opencontainers.image.source="https://github.com/lucasl0st/collector"
LABEL org.opencontainers.image.licenses=MIT

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ="Europe/Berlin"
ENV ZONEINFO=/zoneinfo.zip

ENV PORT=8000

COPY --from=builder /collector /usr/bin/collector

COPY --from=builder /new_tmp /tmp

EXPOSE $PORT

COPY --from=tarampampam/curl:7.88.1 /bin/curl /bin/curl
HEALTHCHECK --interval=5s --start-period=60s CMD [ "curl", "--fail", "http://localhost:8000/health"]

CMD [ "/usr/bin/collector" ]