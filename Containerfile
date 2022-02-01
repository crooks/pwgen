# Build the pwgen binary
FROM golang:1.13 as builder

WORKDIR /workspace

# Copy the go source
COPY pwgenweb/pwgen.go pwgen.go

# Introduce the build arg check in the end of the build stage
# to avoid messing with cached layers
ARG VERSION

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags "-X main.buildVersion=${VERSION} -X main.buildDate=`date -u +%Y-%m-%d`" -o pwgen pwgen.go

RUN test -n "$VERSION" || (echo "VERSION not set" && false)

# Use Red Hat UBI base image
FROM ubi-minimal:latest

ARG VERSION

COPY words.txt /var/local/pwgen/words.txt

LABEL name=pwgen \
      vendor='Western Power Distribution' \
      version=$VERSION \
      release=$VERSION \
      description='pwgen image' \
      summary='pwgen generates secure random passwords'

ENV USER_ID=1001

WORKDIR /
COPY --from=builder /workspace/pwgen .
USER ${USER_ID}

EXPOSE 8080

CMD ["/pwgen"]
