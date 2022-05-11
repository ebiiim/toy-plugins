# Run the custom scheduler instead of the default scheduler

## 1. Update `/etc/kubernetes/manifests/kube-scheduler.yaml`

Example: `./kube-scheduler.yaml`

```yaml
spec:
  containers:
    - command:
        - kube-scheduler
        # specify --config arg
        - --config=/etc/kubernetes/kube-scheduler-configuration.yaml
        ...
      # specify scheduler image
      image: ghcr.io/ebiiim/toy-scheduler:latest
      # mount config file
      volumeMounts:
        - mountPath: /etc/kubernetes/kube-scheduler-configuration.yaml
        name: kubeschedulerconfig
        readOnly: true
  # specify config file path
  volumes:
    - hostPath:
        path: /etc/kubernetes/kube-scheduler-configuration.yaml
        type: FileOrCreate
      name: kubeschedulerconfig
```

## 2. Create `/etc/kubernetes/kube-scheduler-configuration.yaml`

```sh
sudo cp ./kube-scheduler-configuration.yaml /etc/kubernetes/kube-scheduler-configuration.yaml
```

## 3. Run Pods
