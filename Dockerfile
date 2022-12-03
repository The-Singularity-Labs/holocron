# ------------------------------------------------------------------------------
# Base image
# From golang:bullseye
# --

FROM golang:bullseye as horcrux_base
RUN apt-get -y update && apt-get -y  install wget curl

RUN set -eux && \
    apt-get -y install \
        build-essential \
        automake \
        ca-certificates \
        g++ \
        git \
        gcc \
        libc6-dev \
        make \
        pkg-config \
        minify \
        libtool && \
    apt-get -y autoremove --purge && apt-get -y clean && rm -rf /var/lib/apt/lists/*


RUN go get -u github.com/stretchr/testify/assert
RUN go get -u github.com/go-playground/overalls
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.23.1
RUN golangci-lint --version
WORKDIR $GOPATH/src/github.com/the-singularity-lab/holocron
ENV GO111MODULE=on
ENV GOPROXY=direct

COPY go.mod .
COPY go.sum .
RUN go mod download -x


COPY ./templates /templates
COPY ./horcrux.go ./horcrux.go
COPY ./cmd/serve/main.go ./cmd/serve/main.go
COPY ./cmd/generate/main.go ./cmd/generate/main.go

# ------------------------------------------------------------------------------
# Horcrux Builds
# ------------------------------------------------------------------------------
FROM horcrux_base as horcrux_build
RUN mkdir -p /build
CMD ["run", "./cmd/generate/main.go", "/templates/app.gohtml"]
ENTRYPOINT ["go"]

# ------------------------------------------------------------------------------
# Horcrux Server
# ------------------------------------------------------------------------------
FROM horcrux_base as horcrux_server
CMD ["run", "./cmd/serve/main.go", "/templates/app.gohtml"]
ENTRYPOINT ["go"]