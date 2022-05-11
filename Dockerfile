FROM golang:1.18
WORKDIR /go/src/github.com/ebiiim/toy-plugins
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build-bin

FROM alpine:3.15
COPY --from=0 /go/src/github.com/ebiiim/toy-plugins/bin/kube-scheduler /bin/kube-scheduler
WORKDIR /bin
CMD ["kube-scheduler"]
