---
title: ArangoDB Starter Removal Procedure
menuTitle: Removal Procedure
weight: 5
description: >-
  This procedure describes how to remove a machine from a cluster that was started with ArangoDB Starter
archetype: default
---
{{< danger >}}
**Do not** apply this procedure to machines that have an **Agent** on it.
{{< /danger >}}

This procedure is intended to remove a machine from a cluster
that was started with the ArangoDB _Starter_.

It is possible to run this procedure while the machine is still running
or when it has already been removed.

It is not possible to remove machines that have an Agent on it!
The _Agency_ needs to remain functional for the cluster to operate.
Use the [recovery procedure](recovery-procedure.md) if you have
a failed machine with an Agent on it.

Note that it is highly recommended to remove a machine while it is still running.

To remove a machine from a cluster, run the following command:

```bash
arangodb remove starter --starter.endpoint=<endpoint> [--starter.id=<id>] [--force]
```

Where `<endpoint>` is the endpoint of the starter that you want to remove,
or the endpoint of one of the remaining starters. E.g. `http://localhost:8528`.

If you want to remove a machine that is no longer running, use the `--starter.id`
option. Set it to the ID of the ArangoDB _Starter_ on the machine that you want to remove.

You can find this ID in a `setup.json` file in the data directory of one of
the remaining ArangoDB _Starters_. Example:

```json
{
  ...
  "peers": {
    "Peers": [
      {
        "ID": "21e42415",
        "Address": "10.21.56.123",
        "Port": 8528,
        "PortOffset": 0,
        "DataDir": "/mydata/server1",
        "HasAgent": true,
        "IsSecure": false
      },
      ...
    ]
  }
}
```

If the machine you want to remove has address `10.21.56.123` and was listening
on port `8528`, use ID `21e42415`.

The `remove starter` command will attempt the cleanout all data from the servers
of the machine that you want to remove. This can take a long of time.
If the cleanout fails, the `remove starter` command will fail.

If you want to remove the machine even when the cleanout has failed, use
the `--force` option. Note that this may lead to data loss!
