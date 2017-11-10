FROM golang:alpine AS build
MAINTAINER Kazumichi Yamamoto <yamamoto.febc@gmail.com>
LABEL MAINTAINER 'Kazumichi Yamamoto <yamamoto.febc@gmail.com>'
ADD . /go/src/github.com/sacloud/tf-release-info
WORKDIR /go/src/github.com/sacloud/tf-release-info
RUN go build -o tf-release-info

FROM alpine:3.6
MAINTAINER Kazumichi Yamamoto <yamamoto.febc@gmail.com>
LABEL MAINTAINER 'Kazumichi Yamamoto <yamamoto.febc@gmail.com>'
COPY --from=build /go/src/github.com/sacloud/tf-release-info/tf-release-info /usr/local/bin/tf-release-info
WORKDIR /workdir
ENTRYPOINT ["/usr/local/bin/tf-release-info"]