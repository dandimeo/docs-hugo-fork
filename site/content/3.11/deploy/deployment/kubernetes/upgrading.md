---
title: Upgrading
menuTitle: Upgrading
weight: 55
description: >-
  The ArangoDB Kubernetes Operator supports upgrading an ArangoDB from one version to the next
archetype: default
---
The ArangoDB Kubernetes Operator supports upgrading an ArangoDB from
one version to the next.

{{< warning >}}
It is highly recommended to take a backup of your data before upgrading ArangoDB
using [_arangodump_](../../../components/tools/arangodump/_index.md).
{{< /warning >}}

## Upgrade an ArangoDB deployment

To upgrade a cluster, change the version by changing
the `spec.image` setting and the apply the updated
custom resource using:

```bash
kubectl apply -f yourCustomResourceFile.yaml
```

The ArangoDB operator will perform an sequential upgrade
of all servers in your deployment. Only one server is upgraded
at a time.

For patch level upgrades (e.g. 3.9.2 to 3.9.3) each server
is stopped and restarted with the new version.

For minor level upgrades (e.g. 3.9.2 to 3.10.0) each server
is stopped, then the new version is started with `--database.auto-upgrade`
and once that is finish the new version is started with the normal arguments.

The process for major level upgrades depends on the specific version.

## Upgrade the operator itself

To update the ArangoDB Kubernetes Operator itself to a new version,
update the image version of the deployment resource
and apply it using:

```bash
kubectl apply -f examples/yourUpdatedDeployment.yaml
```

## See also

- [Scaling](scaling.md)
