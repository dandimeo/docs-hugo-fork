---
title: Configuration & secrets
menuTitle: Configuration & Secrets
weight: 60
description: >-
  An ArangoDB cluster has lots of configuration options
archetype: default
---
An ArangoDB cluster has lots of configuration options.
Some will be supported directly in the ArangoDB Operator,
others will have to specified separately.

## Built-in options

All built-in options are passed to ArangoDB servers via command-line
arguments configured in the Pod-spec.

## Other configuration options

All command-line options of `arangod` (and `arangosync`) are available
by adding options to the `spec.<group>.args` list of a group
of servers.

These arguments are added to th command-line created for these servers.

## Secrets

The ArangoDB cluster needs several secrets such as JWT tokens
TLS certificates and so on.

All these secrets are stored as Kubernetes Secrets and passed to
the applicable Pods as files, mapped into the Pods filesystem.

The name of the secret is specified in the custom resource.
For example:

```yaml
apiVersion: "cluster.arangodb.com/v1"
kind: "Cluster"
metadata:
  name: "example-arangodb-cluster"
spec:
  mode: Cluster
  auth:
    jwtSecretName: <name-of-JWT-token-secret>
```
