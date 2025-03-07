---
title: Datacenter-to-Datacenter Replication deployment
menuTitle: Deployment
weight: 5
description: ''
archetype: chapter
---
{{< tag "ArangoDB Enterprise Edition" >}}

## Deployment steps

To deploy all the components needed for _Datacenter-to-Datacenter Replication (DC2DC)_,
you need to set up the following components:

1. An ArangoDB cluster in each data center
2. Multiple ArangoSync Masters
3. Multiple ArangoSync Workers
4. Optional: Prometheus and Grafana for monitoring

## Cluster

Datacenter-to-Datacenter Replication requires an ArangoDB cluster in both data centers.

Since the _Agents_ are so critical to the availability of both the ArangoDB and
the ArangoSync cluster, it is recommended to run _Agents_ on dedicated machines.
They run a real-time system for the elections and bad performance can negatively
affect the availability of the whole cluster.

_DB-Servers_ are also important and you do not want to lose them, but
depending on your replication factor, the system can tolerate some
loss and bad performance slows things down but not stop things from
working.

_Coordinators_ can be deployed on other machines, since they do not hold
persistent state. They might have some in-memory state about running
transactions or queries, but losing a Coordinator does not lose any
persisted data. Furthermore, new Coordinators can be added to a cluster
without much effort.

Please refer to the [Cluster](arangodb-cluster.md) section for
more information.

## ArangoSync Master

The Sync Master is responsible for managing all synchronization, creating tasks and assigning
those to workers.

At least two instances must be deployed in each datacenter.
One instance is the leader cluster, the other is an inactive follower cluster.
When the leader is gone for a short while, one of the other instances takes over.

With clusters of a significant size, the sync master requires a significant set of resources.
Therefore it is recommended to deploy sync masters on their own servers, equipped with sufficient
CPU power and memory capacity.

The sync master must be reachable on a TCP port 8629 (default).
This port must be reachable from inside the datacenter (by sync workers and operations)
and from inside of the other datacenter (by sync masters in the other datacenter).

Since the sync masters can be CPU intensive when running lots of databases & collections,
it is recommended to run them on dedicated machines with a lot of CPU power.

Consider these machines to be crucial for your DC2DC setup.

Please refer to the [ArangoSync Master](arangosync-master.md)
section for more information.

## ArangoSync Workers

The Sync Worker is responsible for executing synchronization tasks.

For optimal performance at least 1 worker instance must be placed on
every machine that has an ArangoDB DB-Server running. This ensures that tasks
can be executed with minimal network traffic outside of the machine.

Since sync workers automatically stop once their TLS server certificate expires
(which is set to 2 years by default),
it is recommended to run at least two instances of a worker on every machine in the datacenter.
That way, tasks can still be assigned in the most optimal way, even when a worker in temporarily
down for a restart.

The sync worker must be reachable on a TCP port 8729 (default).
This port must be reachable from inside the datacenter (by sync masters).

The sync workers should be run on all machines that also contain an ArangoDB DB-Server.
The sync worker can be memory intensive when running lots of databases & collections.

Please refer to the [ArangoSync Workers](arangosync-workers.md)
for more information.

## Prometheus & Grafana (optional)

ArangoSync provides metrics in a format supported by [Prometheus](https://prometheus.io).
We also provide a standard set of dashboards for viewing those metrics in [Grafana](https://grafana.org).

If you want to use these tools, go to their websites for instructions on how to deploy them.

After deployment, you must configure prometheus using a configuration file that instructs
it about which targets to scrape. For ArangoSync you should configure scrape targets for
all sync masters and all sync workers.

Prometheus can be a memory & CPU intensive process. It is recommended to keep them
on other machines than used to run the ArangoDB cluster or ArangoSync components.

Consider these machines to be easily replaceable, unless you configure
alerting on _prometheus_, in which case it is recommended to keep a
close eye on them, such that you do not lose any alerts due to failures
of Prometheus.

Please refer to the [Prometheus & Grafana](prometheus-and-grafana.md)
section for more information.
