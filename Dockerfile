FROM golang:1.21-alpine

RUN apk add --no-cache git wget

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/cyverse-de/road-runner
COPY . .
RUN go install

ENTRYPOINT ["road-runner"]
CMD ["--help"]

ARG git_commit=unknown
ARG version="3.0.1"
ARG descriptive_version=unknown

LABEL org.cyverse.git-ref="$git_commit"
LABEL org.cyverse.version="$version"
LABEL org.cyverse.descriptive-version="$descriptive_version"
LABEL org.label-schema.vcs-ref="$git_commit"
LABEL org.label-schema.vcs-url="https://github.com/cyverse-de/road-runner"
LABEL org.label-schema.version="$descriptive_version"
