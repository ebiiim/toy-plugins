# Run the custom scheduler locally without container images (for development)

## 1. Run the custom scheduler on one of Master Nodes.

```sh
sudo ../../bin/kube-scheduler --authentication-kubeconfig=/etc/kubernetes/scheduler.conf --authorization-kubeconfig=/etc/kubernetes/scheduler.conf --config=./kube-scheduler-configuration.yaml --secure-port=10260
```

## 2. Run Pods with scheduler specified.

```yaml
spec:
  schedulerName: toy-scheduler
```
