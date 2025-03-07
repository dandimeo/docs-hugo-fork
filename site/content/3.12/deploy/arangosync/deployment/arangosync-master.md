---
title: ArangoSync Master
menuTitle: ArangoSync Master
weight: 10
description: >-
  The ArangoSync Master is responsible for managing all synchronization, creating tasks and assigning those to the ArangoSync Workers
archetype: default
---
The _ArangoSync Master_ is responsible for managing all synchronization, creating
tasks and assigning those to the _ArangoSync Workers_.

At least 2 instances must be deployed in each datacenter.
One instance will be the "leader", the other will be a "follower". When the
leader is gone for a short while, one of the other instances will take over.

With clusters of a significant size, the _sync master_ will require a
significant set of resources. Therefore it is recommended to deploy the _sync masters_
on their own servers, equipped with sufficient CPU power and memory capacity.

To start an _ArangoSync Master_ using a `systemd` service, use a unit like this:

```conf
[Unit]
Description=Run ArangoSync in master mode
After=network.target

[Service]
Restart=on-failure
EnvironmentFile=/etc/arangodb.env
EnvironmentFile=/etc/arangodb.env.local
LimitNOFILE=8192
ExecStart=/usr/sbin/arangosync run master \
    --log.level=debug \
    --cluster.endpoint=${CLUSTERENDPOINTS} \
    --cluster.jwtSecret=${CLUSTERSECRET} \
    --server.keyfile=${CERTIFICATEDIR}/tls.keyfile \
    --server.client-cafile=${CERTIFICATEDIR}/client-auth-ca.crt \
    --server.endpoint=https://${PRIVATEIP}:${MASTERPORT} \
    --server.port=${MASTERPORT} \
    --master.endpoint=${PUBLICMASTERENDPOINTS} \
    --master.jwtSecret=${MASTERSECRET} \
    --mq.type=direct
TimeoutStopSec=60

[Install]
WantedBy=multi-user.target
```

The _sync master_ needs a TLS server certificate and a
If you want the service to create a TLS certificate & client authentication
certificate, for authenticating with _ArangoSync Masters_ in another datacenter,
for every start, add this to the `Service` section.

```conf
ExecStartPre=/usr/bin/sh -c "mkdir -p ${CERTIFICATEDIR}"
ExecStartPre=/usr/sbin/arangosync create tls keyfile \
    --cacert=${CERTIFICATEDIR}/tls-ca.crt \
    --cakey=${CERTIFICATEDIR}/tls-ca.key \
    --keyfile=${CERTIFICATEDIR}/tls.keyfile \
    --host=${PUBLICIP} \
    --host=${PRIVATEIP} \
    --host=${HOST} \
    --host=${CLUSTERDNSNAME}
```

The _ArangoSync Master_ must be reachable on a TCP port `${MASTERPORT}` (used with `--server.port` option).
This port must be reachable from inside the datacenter (by sync workers and operations)
and from inside of the other datacenter (by sync masters in the other datacenter).

Note that other sync masters in the same datacenter will contact this sync master
through the endpoint specified in `--server.endpoint`.
Sync masters (and sync workers) from the other datacenter will contact this sync master
through the endpoint specified in `--master.endpoint`.

## Recommended deployment environment

Since the _sync masters_ can be CPU intensive when running lots of databases & collections,
it is recommended to run them on dedicated machines with a lot of CPU power.

Consider these machines crucial for you DC2DC replication setup.
