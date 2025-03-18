# KubeArmor Observability Demo

This demo shows how to use Prometheus to provide observability features for KubeArmor. The current implementation tracks the number of applied policies.

## Usage

1. Make sure KubeArmor is running

2. build and run demo:

```bash
go build
sudo ./Observability_demo
```

3. Access Prometheus metrics:

```bash
curl http://localhost:9090/metrics | grep kubearmor
```

or 

```bash
prometheus --config.file=./conf.yml # host
```

## Implemented Metrics

- `kubearmor_policies_applied_total` - Total number of currently applied policies
- `kubearmor_policies_by_type` - Number of policies categorized by type (container/host)
- `kubearmor_policies_by_namespace` - Number of policies categorized by namespace