---
title: ArangoDB Starter Architecture
menuTitle: Architecture
weight: 15
description: >-
  The ArangoDB Starter is a program used to create ArangoDB database deployments on bare-metal (or virtual machines) with ease
archetype: default
---
## What does the Starter do

The ArangoDB Starter is a program used to create ArangoDB database deployments
on bare-metal (or virtual machines) with ease.
It enables you to create everything from a simple Single server instance
to a full blown Cluster with Datacenter-to-Datacenter Replication in under 5 minutes.

The Starter is intended to be used in environments where there is no higher
level orchestration system (e.g. Kubernetes) available.

## Starter versions

The Starter is a separate process in a binary called `arangodb` (or `arangodb.exe` on Windows).
This binary has its own version number that is independent of a ArangoDB (database)
version.

This means that Starter version `a.b.c` can be used to run deployments
of ArangoDB databases with different version.
For example, the Starter with version `0.15.5` can be used to create
ArangoDB deployments with ArangoDB version `3.10.<something>` as well
as deployments with ArangoDB version `3.9.<something>`.

It also means that you can update the Starter independently from the ArangoDB
database.

Note that the Starter is also included in all binary ArangoDB packages.

To find the versions of you Starters & ArangoDB database, run the following commands:

```bash
# To get the Starter version
arangodb --version
# To get the ArangoDB database version
arangod --version
```

## Starter deployment modes

The Starter supports 3 different modes of ArangoDB deployments:

1. Single server
1. Active failover
1. Cluster

Note: Datacenter replication is an option for the `cluster` deployment mode.

You select one of these modes using the `--starter.mode` command line option.

Depending on the mode you've selected, the Starter launches one or more
(`arangod` / `arangosync`) server processes.

No matter which mode you select, the Starter always provides you
a common directory structure for storing the servers data, configuration & log files.

## Starter operating modes

The Starter can run as normal processes directly on the host operating system,
or as containers in a docker runtime.

When running as normal process directly on the host operating system,
the Starter launches the servers as child processes and monitors those.
If one of the server processes terminates, a new one is started automatically.

When running in a docker container, the Starter launches the servers
as separate docker containers, that share the volume namespace with
the container that runs the Starter. It monitors those containers
and if one terminates, a new container is launched automatically.

## Starter data-directory

The Starter uses a single directory with a well known structure to store
all data for its own configuration & logs, as well as the configuration,
data & logs of all servers it starts.

This data directory is set using the `--starter.data-dir` command line option.
It contains the following files & sub-directories.

- `setup.json` The configuration of the "cluster of Starters".
  For details see below. DO NOT edit this file.
- `arangodb.log` The log file of the Starter
- `single<port>`, `agent<port>`, `coordinator<port>`, `dbserver<port>`: directories for
  launched servers. These directories contain among others the following files:
  - `apps`: A directory with Foxx applications
  - `data`: A directory with database data
  - `arangod.conf`: The configuration file for the server. Editing this file is possible, but not recommended.
  - `arangod.log`: The log file of the server
  - `arangod_command.txt`: File containing the exact command line of the started server (for debugging purposes only)

## Starter configuration file

The Starter can be configured using a configuration file. The format of the
configuration file is the same as the `arangod` configuration file format.
For more details, refer to the [configuration file format](../../../operations/administration/configuration.md#configuration-file-format)
and [how to use configuration files](../../../operations/administration/configuration.md#using-configuration-files).

The default configuration file of the Starter is `arangodb-starter.conf`.
It can be changed using the `--configuration` option. 
For more information about other
configuration options, see [ArangoDB Starter options](options.md). 

{{< info >}}
The Starter has a different set of supported command line options than `arangod` binary.
Using the `arangod` configuration file as input for `arangodb` binary is not supported.
{{< /info >}}

### Passing through `arangod` options

The configuration file also supports setting pass-through options. Options with
same prefixes can be split into sections.

```conf
# passthrough-example.conf

args.all.log.level = startup=trace
args.all.log.level = warning

[starter]
  mode = single

[args]
all.log.level = queries=debug
all.default-language = de_DE

[args.all.rocksdb]
enable-statistics = true
```

```bash
./arangodb --configuration=passthrough-example.conf
```

### Configuration precedence

When adding a command line option next to a modified configuration
file, the last occurrence of the option becomes the final value.

Running the Starter with the configuration example above and adding the
`default-language=es_419` command line option results in having the
`default-language` set to `es_419` and not the value from the configuration file:

```bash
./arangodb --args.all.default-language=es_419 --configuration=passthrough-example.conf
```

## Running on multiple machines

For the `activefailover` and `cluster` mode, it is required to run multiple
Starters, as every Starter only launches a subset of all servers needed
to form the entire deployment.
In the `cluster` mode, for example, a single Starter launches at most one Agent,
one DB-Server, and one Coordinator.

It is the responsibility of the user to run the Starter on multiple machines such
that enough servers are started to form the entire deployment.
The minimum number of Starters needed is 3.

The Starters running on those machines need to know about each other's existence.
In order to do so, the Starters form a "cluster" of their own (not to be confused
with the ArangoDB database cluster).
This cluster of Starters is formed from the values given to the `--starter.join`
command line option. You should pass the addresses (`<host>:<port>`) of all Starters.

For example, a typical command line for a cluster deployment looks like this:

```bash
arangodb --starter.mode=cluster --starter.join=hostA:8528,hostB:8528,hostC:8528
# this command is run on hostA, hostB and hostC.
```

The state of the cluster (of Starters) is stored in a configuration file called
`setup.json` in the data directory of every Starter and the ArangoDB
The Agency is used to elect a leader (also called _master_) among all Starters.

The leader Starter is responsible for maintaining the list of all Starters
involved in the cluster and their addresses. The follower Starters (all Starters
except the elected leader) fetch this list from the leader Starter on regular
basis and store it to its own `setup.json` config file.

{{< warning >}}
The `setup.json` config file must not be edited manually.
{{< /warning >}}

## Running on multiple machines (under the hood)

As mentioned above, when the Starter is used to create an `activefailover`
or `cluster` deployment, it first creates a "cluster" of Starters.

These are the steps taken by the Starters to bootstrap such a deployment
from scratch.

1. All Starters are started (either manually or by some supervisor)
1. All Starters try to read their config from `setup.json`.
   If that file exists and is valid, this bootstrap-from-scratch process
   is aborted and all Starters go directly to the `running` phase described below.
1. All Starters create a unique ID
1. The list of `--starter.join` arguments is sorted
1. All Starters request the unique ID from the first server in the sorted `--starter.join` list,
   and compares the result with its own unique ID.
1. The Starter that finds its own unique ID, is continuing as _bootstrap leader_
   the other Starters are continuing as _bootstrap followers_.
1. The _bootstrap leader_ waits for at least 2 _bootstrap followers_ to join it.
1. The _bootstrap followers_ contact the _bootstrap leader_ to join its cluster of Starters.
1. Once the _bootstrap leader_ has received enough (at least 2) requests
   to join its cluster of Starters, it continues with the `running` phase.
1. The _bootstrap followers_ keep asking the _bootstrap leader_ about its state.
   As soon as they receive confirmation to do so, they also continue with the `running` phase.

In the _running_ phase, all Starters launch the desired servers and keeps monitoring those
servers. Once a functional Agency is detected, all Starters try to be
_running leader_ by trying to write their ID in a well known location in the Agency.
The first Starter to succeed in doing so wins this leader election.

The _running leader_ keeps writing its ID in the Agency in order to remain
the _running leader_. Since this ID is written with a short time-to-live,
other Starters are able to detect when the current _running leader_ has been stopped
or is no longer responsible. In that case, the remaining Starters perform
another leader election to decide who will be the next _running leader_.

API requests that involve the state of the cluster of Starters are always answered
by the current _running leader_. All other Starters refer the request to
the current _running leader_.
