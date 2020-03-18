# Build Stage
FROM golang:1.14.0-alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache ca-certificates make git curl mercurial bzr

LABEL app="build-halodoc-tlv"
LABEL REPO="https://github.com/prasetyowira/halodoc-tlv"

ENV PROJPATH=/go/src/github.com/prasetyowira/halodoc-tlv

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/prasetyowira/halodoc-tlv
WORKDIR /go/src/github.com/prasetyowira/halodoc-tlv

RUN go mod download

RUN make build-alpine

# Final Stage
FROM alpine:latest

RUN apk add --no-cache --update \
    dumb-init \
    bash \
    ca-certificates \
    openssl && \
    rm -rf /var/cache/apk/*

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/prasetyowira/halodoc-tlv"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/halodoc-tlv/bin

WORKDIR /opt/halodoc-tlv/bin

COPY --from=build-stage /go/src/github.com/prasetyowira/halodoc-tlv/bin/halodoc-tlv /opt/halodoc-tlv/bin/
COPY --from=build-stage /go/src/github.com/prasetyowira/halodoc-tlv/input.txt /opt/halodoc-tlv/
RUN ls -la /opt/halodoc-tlv
RUN ls -la /opt/halodoc-tlv/bin
RUN chmod +x /opt/halodoc-tlv/bin/halodoc-tlv

# Create appuser
RUN adduser -D -g '' halodoc-tlv
USER halodoc-tlv

CMD ["/opt/halodoc-tlv/bin/halodoc-tlv", "< input.txt"]
