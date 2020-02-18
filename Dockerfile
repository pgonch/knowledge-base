FROM golang:1.13.0-alpine3.10 AS builder

ARG GITHUB_TOKEN

RUN apk add --no-cache make git && \
    echo "[url \"https://$GITHUB_TOKEN:x-oauth-basic@github.com/\"]"$'\n\t'"insteadOf = https://github.com/" \
      >> /root/.gitconfig

WORKDIR /go/src/github.com/pgonch/knowledge-base/

COPY . .
RUN make

FROM alpine

COPY --from=builder /go/src/github.com/pgonch/knowledge-base/bin/knowledge-base /knowledge-base

ENTRYPOINT ["/knowledge-base"]
