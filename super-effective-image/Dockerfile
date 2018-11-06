FROM golang:alpine as builder
LABEL maintainer ciaran
RUN apk update && apk add git
RUN adduser -D -g '' appuser
RUN go get github.com/Code-Hex/Neo-cowsay
ENV SOURCES /go/src/github.com/ciaranRoche/myapp/
COPY . ${SOURCES}
WORKDIR ${SOURCES}
RUN cd ${SOURCES} && CGO_ENABLED=0 go build
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/github.com/ciaranRoche/myapp/myapp /go/src/github.com/ciaranRoche/myapp/myapp
USER appuser
ENTRYPOINT [ "/go/src/github.com/ciaranRoche/myapp/myapp" ]
