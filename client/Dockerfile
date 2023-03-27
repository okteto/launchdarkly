FROM golang:1.19-buster as builder

LABEL org.opencontainers.image.source=https://github.com/okteto/launchdarkly
LABEL org.opencontainers.image.description="Integrate Launchdarkly and Okteto"
LABEL org.opencontainers.image.licenses=Apache-2.0

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY Makefile .
COPY cmd /app/cmd
COPY main.go .
ARG COMMIT_SHA
RUN make

FROM busybox

COPY --from=builder /app/launchdarkly /usr/local/bin/launchdarkly
RUN chmod +x /usr/local/bin/launchdarkly