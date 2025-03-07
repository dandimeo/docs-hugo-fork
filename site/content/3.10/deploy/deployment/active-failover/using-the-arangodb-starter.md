---
title: Using the ArangoDB Starter
menuTitle: Using the ArangoDB Starter
weight: 10
description: >-
  This section describes how to start an Active Failover setup the tool Starter (the arangodb binary program)
archetype: default
---
This section describes how to start an Active Failover setup the tool [_Starter_](../../../components/tools/arangodb-starter/_index.md)
(the _arangodb_ binary program).

As a precondition you should create a _secret_ to activate authentication. The _Starter_ provides a handy
functionality to generate such a file:

```bash
arangodb create jwt-secret --secret=arangodb.secret
```

Set appropriate privilege on the generated _secret_ file, e.g. on Linux:

```bash
chmod 400 arangodb.secret
```

## Local Active Failover test deployment

If you want to start a local _Active Failover_ setup quickly, use the `--starter.local`
option of the _Starter_. This will start all servers within the context of a single
starter process:

```bash
arangodb --starter.local --starter.mode=activefailover --starter.data-dir=./localdata --auth.jwt-secret=/etc/arangodb.secret --agents.agency.supervision-grace-period=30
```

Please adapt the path to your _secret_ file accordingly.

Note that to avoid unnecessary failovers, it may make sense to increase the value
for the startup option `--agents.agency.supervision-grace-period` to a value
beyond 30 seconds.

**Note:** When you restart the _Starter_, it remembers the original `--starter.local` flag.

## Multiple Machines

If you want to start an Active Failover setup using the _Starter_, you need to copy the
_secret_ file to every machine and use the `--starter.mode=activefailover` option of the
_Starter_. A 3 "machine" _Agency_ is started as well as 3 single servers,
that perform asynchronous replication and failover:

```bash
arangodb --starter.mode=activefailover --starter.data-dir=./data --auth.jwt-secret=/etc/arangodb.secret --agents.agency.supervision-grace-period=30 --starter.join A,B,C
```

Please adapt the path to your _secret_ file accordingly.

Note that to avoid unnecessary failovers, it may make sense to increase the value
for the startup option `--agents.agency.supervision-grace-period` to a value
beyond 30 seconds.

Run the above command on machine A, B & C.

Once all the processes started by the _Starter_ are up and running, and joined the
Active Failover setup (this may take a while depending on your system), the _Starter_ will inform
you where to connect the Active Failover from a Browser, shell or your program.

For a full list of options of the _Starter_ please refer to [this](../../../components/tools/arangodb-starter/options.md)
section.

## Using the ArangoDB Starter in Docker

The _Starter_ can also be used to launch an Active Failover setup based on _Docker_
containers. To do this, you can use the normal Docker arguments, combined with
`--starter.mode=activefailover`:

```bash
export IP=<IP of docker host>
docker volume create arangodb
docker run -it --name=adb --rm -p 8528:8528 \
    -v arangodb:/data \
    -v /var/run/docker.sock:/var/run/docker.sock \
    arangodb/arangodb-starter \
    --agents.agency.supervision-grace-period=30 \
    --starter.address=$IP \
    --starter.mode=activefailover \
    --starter.join=A,B,C
```

Run the above command on machine A, B & C.

Note that to avoid unnecessary failovers, it may make sense to increase the value
for the startup option `--agents.agency.supervision-grace-period` to a value
beyond 30 seconds.

The _Starter_ will decide on which 2 machines to run a single server instance.
To override this decision (only valid while bootstrapping), add a
`--cluster.start-single=false` to the machine where the single server
instance should _not_ be started.

If you use the Enterprise Edition Docker image, you have to set the license key
in an environment variable by adding this option to the above `docker` command:

```
    -e ARANGO_LICENSE_KEY=<the-key>
```

You can get a free evaluation license key by visiting:

[www.arangodb.com/download-arangodb-enterprise/](https://www.arangodb.com/download-arangodb-enterprise/)

Then replace `<the-key>` above with the actual license key. The start
will then hand on the license key to the Docker containers it launches
for ArangoDB.

### TLS verified Docker services

Oftentimes, one needs to harden Docker services using client certificate 
and TLS verification. The Docker API allows subsequently only certified access.
As the ArangoDB starter starts the ArangoDB cluster instances using this Docker API, 
it is mandatory that the ArangoDB starter is deployed with the proper certificates
handed to it, so that the above command is modified as follows:

```bash
export IP=<IP of docker host>
export DOCKER_TLS_VERIFY=1
export DOCKER_CERT_PATH=/path/to/certificate
docker volume create arangodb
docker run -it --name=adb --rm -p 8528:8528 \
    -v arangodb:/data \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v /path/to/certificate:/path/to/certificate
    arangodb/arangodb-starter \
    --agents.agency.supervision-grace-period=30 \
    --starter.address=$IP \
    --starter.mode=activefailover \
    --starter.join=A,B,C
```

Note that the environment variables `DOCKER_TLS_VERIFY` and `DOCKER_CERT_PATH` 
as well as the additional mountpoint containing the certificate have been added above. 
directory. The assignment of `DOCKER_CERT_PATH` is optional, in which case it 
is mandatory that the certificates are stored in `$HOME/.docker`. So
the command would then be as follows

```bash
export IP=<IP of docker host>
docker volume create arangodb
docker run -it --name=adb --rm -p 8528:8528 \
    -v arangodb:/data \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v /path/to/cert:/root/.docker \
    -e DOCKER_TLS_VERIFY=1 \
    arangodb/arangodb-starter \
    --agents.agency.supervision-grace-period=30 \
    --starter.address=$IP \
    --starter.mode=activefailover \
    --starter.join=A,B,C
```
