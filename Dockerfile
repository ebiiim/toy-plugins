FROM golang:1.18
WORKDIR /go/src/github.com/ebiiim/toy-plugins
COPY . .
ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X k8s.io/component-base/version.gitVersion=${VERSION} -w" -o bin/kube-scheduler cmd/scheduler/main.go

FROM alpine:3.15
COPY --from=0 /go/src/github.com/ebiiim/toy-plugins/bin/kube-scheduler /bin/kube-scheduler
WORKDIR /bin
CMD ["kube-scheduler"]
