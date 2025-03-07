---
title: Cloud Migration Tool
menuTitle: Migrate to the cloud
weight: 30
description: >-
  Migrating data from bare metal servers to the cloud with minimal downtime
archetype: default
---
{{< description >}}

The `arangosync-migration` tool allows you to easily move from on-premises to
the cloud while ensuring a smooth transition with minimal downtime.
Start the cloud migration, let the tool do the job and, at the same time,
keep your local cluster up and running.

Some of the key benefits of the cloud migration tool include:
- Safety comes first - pre-checks and potential failures are carefully handled.
- Your data is secure and fully encrypted.
- Ease-of-use with a live migration while your local cluster is still in use.
- Get access to what a cloud-based fully managed service has to offer:
  high availability and reliability, elastic scalability, and much more.

## Downloading the tool

The `arangosync-migration` tool is available to download for the following
operating systems:

**Linux**
- [AMD64 (x86_64) architecture](https://download.arangodb.com/arangosync-migration/linux/amd64/arangosync-migration)
- [ARM64 (AArch64) architecture](https://download.arangodb.com/arangosync-migration/linux/arm64/arangosync-migration)

**macOS / Darwin**
- [AMD64 (x86_64) architecture](https://download.arangodb.com/arangosync-migration/darwin/amd64/arangosync-migration)
- [ARM64 (AArch64) architecture](https://download.arangodb.com/arangosync-migration/darwin/arm64/arangosync-migration)

**Windows**
- [AMD64 (x86_64) architecture](https://download.arangodb.com/arangosync-migration/windows/amd64/arangosync-migration.exe)
- [ARM64 (AArch64) architecture](https://download.arangodb.com/arangosync-migration/windows/arm64/arangosync-migration.exe)

For macOS as well as other Unix-based operating systems, run the following
command to make sure you can execute the binary:

```bash
chmod 755 ./arangosync-migration
```

## Prerequisites

Before getting started, make sure the following prerequisites are in place:

- Go to the [ArangoGraph Insights Platform](https://cloud.arangodb.com/home)
  and sign in. If you don’t have an account yet, sign-up to create one.

- Generate an ArangoGraph API key and API secret. See a detailed guide on
  [how to create an API key](arangograph-api/getting-started-with-the-api.md#creating-an-api-key).

{{< info >}}
The cloud migration tool is only available for clusters.
{{< /info >}}

### Setting up the target deployment in ArangoGraph

Continue by [creating a new ArangoGraph deployment](deployments/_index.md#how-to-create-a-new-deployment)
or choose an existing one.

The target deployment in ArangoGraph requires specific configuration rules to be
set up before the migration can start:

- **Configuration settings**: The target deployment must be compatible with the
  source data cluster. This includes the ArangoDB version that is being used,
  the DB-Servers count, and disk space.
- **Deployment region and cloud provider**: Choose the closest region to your
  data cluster. This factor can speed up your migration to the cloud.

After setting up your ArangoGraph deployment, wait for a few minutes for it to become
fully operational.

{{< info >}}
Note that Developer mode deployments are not supported.
{{< /info >}}

## Running the migration tool

The `arangosync-migration` tool provides a set of commands that allow you to:
- start the migration process
- check whether your source and target clusters are fully compatible
- get the current status of the migration process
- stop or abort the migration process
- switch the local cluster to read-only mode

### Starting the migration process

To start the migration process, run the following command:

```bash
arangosync-migration start
```
The `start` command runs some pre-checks. Among other things, it measures
the disk space which is occupied by your ArangoDB cluster. If you are using the
same data volume for ArangoDB servers and other data as well, the measurements
can be incorrect. Provide the `--source.ignore-metrics` option to overcome this.

You also have the option of doing a `--check-only` without starting the actual
migration. If specified, this checks if your local cluster and target deployment
are compatible without sending any data to ArangoGraph.

Once the migration starts, the local cluster enters into monitoring mode and the
synchronization status is displayed in real-time. If you don't want to see the
status you can terminate this process, as the underlying agent process
continues to work. If something goes wrong, restarting the same command restores
the replication state.

To restart the migration, first `stop` or `stop --abort` the migration. Then,
start it again using the `start` command.

{{< warning >}}
Starting the migration creates a full copy of all data from the source cluster
to the target deployment in ArangoGraph. All data that has previously existed in the
target deployment will be lost.
{{< /warning >}}

### During the migration

The following takes place during an active migration:
- The source data cluster remains usable.
- The target deployment in ArangoGraph is switched to read-only mode.
- Your root user password is not copied to the target deployment in ArangoGraph.
  To get your root password, select the target deployment from the ArangoGraph
  Dashboard and go to the **Overview** tab. All other users are fully synchronized.

{{< warning >}}
The migration tool increases the CPU and memory usage of the server you are
running it on. Depending on your ArangoDB usage pattern, it may take a lot of CPU
to handle the replication. You can stop the migration process anytime
if you see any problems.
{{< /warning >}}

```bash
./arangosync-migration start \
  --source.endpoint=$COORDINATOR_ENDPOINT \
  --source.jwt-secret=/path-to/jwt-secret.file \
  --arango-graph.api-key=$ARANGO_GRAPH_API_KEY \
  --arango-graph.api-secret=$ARANGO_GRAPH_API_SECRET \
  --arango-graph.deployment-id=$ARANGO_GRAPH_DEPLOYMENT_ID
```

### How long does it take?

The total time required to complete the migration depends on how much data you
have and how often write operations are executed during the process.

You can also track the progress by checking the **Migration status** section of
your target deployment in ArangoGraph dashboard.

![ArangoGraph Cloud Migration Progress](../../images/arangograph-migration-agent.png)

### Getting the current status

To print the current status of the migration, run the following command:

```bash
./arangosync-migration status \
  --arango-graph.api-key=$ARANGO_GRAPH_API_KEY \
  --arango-graph.api-secret=$ARANGO_GRAPH_API_SECRET \
  --arango-graph.deployment-id=$ARANGO_GRAPH_DEPLOYMENT_ID
```

You can also add the `--watch` option to start monitoring the status in real-time.

### Stopping the migration process

The `arangosync-migration stop` command stops the migration and terminates
the migration agent process.

If replication is running normally, the command waits until all shards are
in sync. The local cluster is then switched into read-only mode.
After all shards are in-sync and the migration stopped, the target deployment
is switched into the mode specified in `--source.server-mode` option. If no
option is specified, it defaults to the read/write mode.

```bash
./arangosync-migration stop \
  --arango-graph.api-key=$ARANGO_GRAPH_API_KEY \
  --arango-graph.api-secret=$ARANGO_GRAPH_API_SECRET \
  --arango-graph.deployment-id=$ARANGO_GRAPH_DEPLOYMENT_ID
```

The additional `--abort` option is supported. If specified, the `stop` command
will not check anymore if both deployments are in-sync and stops all
migration-related processes as soon as possible.

### Switching the local cluster to read-only mode

The `arangosync-migration set-server-mode` command allows switching
[read-only mode](../develop/http/administration.md#set-the-server-mode-to-read-only-or-default)
for your local cluster on and off.

In a read-only mode, all write operations are going to fail with an error code
of `1004` (ERROR_READ_ONLY).
Creating or dropping databases and collections are also going to fail with
error code `11` (ERROR_FORBIDDEN).

```bash
./arangosync-migration set-server-mode \
  --source.endpoint=$COORDINATOR_ENDPOINT \
  --source.jwt-secret=/path-to/jwt-secret.file \
  --source.server-mode=readonly
```  
The `--source.server-mode` option allows you to specify the desired server mode.
Allowed values are `readonly` or `default`.

### Supported environment variables

The `arangosync-migration` tool supports the following environment variables:

- `$ARANGO_GRAPH_API_KEY`
- `$ARANGO_GRAPH_API_SECRET`
- `$ARANGO_GRAPH_DEPLOYMENT_ID`

Using these environment variables is highly recommended to ensure a secure way
of providing sensitive data to the application.

### Restrictions and limitations

When running the migration, ensure that your target deployment has the same (or
bigger) amount of resources (CPU, RAM) than your cluster. Otherwise, the
migration process might get stuck or require manual intervention. This is closely
connected to the type of data you have and how it is distributed between shards
and collections.

In general, the most important parameters are:
- Total number of leader shards
- The amount of data in bytes per collection

Both parameters can be retrieved from the ArangoDB Web Interface.

The `arangosync-migration` tool supports migrating large datasets of up to
5 TB of data and 3800 leader shards, as well as collections as big as 250 GB.

In case you have any questions, please
[reach out to us](https://www.arangodb.com/contact).

## Cloud migration workflow for minimal downtime

1. Download and start the `arangosync-migration` tool. The target deployment
   is switched into read-only mode automatically.
2. Wait until all shards are in sync. You can use the `status` or the `start`
   command with the same parameters to track that.
3. Optionally, when all shards are in-sync, you can switch your applications
   to use the endpoint of the ArangoGraph deployment, but note that it stays in
   read-only mode until the migration process is fully completed.
4. Stop the migration using the `stop` subcommand. The following steps are executed:
    - The source data cluster is switched into read-only mode.
    - It waits until all shards are synchronized.
    - The target deployment is switched into default read/write mode.

   {{< info >}}
   If you switched the source data cluster into read-only mode,
   you can switch it back to default (read/write) mode using the
   `set-server-mode` subcommand.
   {{< /info >}}
