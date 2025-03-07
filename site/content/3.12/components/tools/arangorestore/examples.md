---
title: _arangorestore_ Examples
menuTitle: Examples
weight: 5
description: >-
  To restore data from a dump previously created with arangodump, ArangoDB provides the arangorestore tool
archetype: default
---
To restore data from a dump previously created with [_arangodump_](../arangodump/_index.md),
ArangoDB provides the _arangorestore_ tool.

## Invoking _arangorestore_

_arangorestore_ can be invoked from the command-line as follows:

```
arangorestore --input-directory "dump"
```

This connects to an ArangoDB server (`tcp://127.0.0.1:8529` by default), then restores the
collection structure and the documents from the files found in the input directory `dump`.
Note that the input directory must have been created by running `arangodump` before.

_arangorestore_ connects to the `_system` database by default, using the default
endpoint. To override the endpoint, or specify a different user, use one of the
following startup options:

- `--server.endpoint <string>`: endpoint to connect to
- `--server.username <string>`: username
- `--server.password <string>`: password to use
  (omit this and you'll be prompted for the password)
- `--server.authentication <bool>`: whether or not to use authentication

If you want to connect to a different database or dump all databases you can additionally
use the following startup options:

- `--server.database <string>`: name of the database to connect to.
  Defaults to the `_system` database.
- `--all-databases true`: restore multiple databases from a dump which used the same option.

Note that the specified user must have access to the database(s).
 
The _arangorestore_ tool provides the `--create-database` option. Setting this
option to `true` creates the target database if it does not exist. When creating the
target database, the username and passwords passed to _arangorestore_ (in options
`--server.username` and `--server.password`) are used to create an initial user for the
new database.

The option `--force-same-database` allows restricting arangorestore operations to a
database with the same name as in the source dump's `dump.json` file. It can thus be used
to prevent restoring data into a "wrong" database by accident.

For example, if a dump was taken from database ***A***, and the restore is attempted into 
database ***B***, then with the `--force-same-database` option set to `true`, arangorestore
aborts instantly.

The `--force-same-database` option is set to `false` by default to ensure backwards-compatibility.

Here's an example of reloading data to a non-standard endpoint, using a dedicated
[database name](../../../concepts/data-structure/databases.md#database-names):

```
arangorestore \
  --server.endpoint tcp://192.168.173.13:8531 \
  --server.username backup \
  --server.database mydb \
  --input-directory "dump" \
```

Also, more than one endpoint can be provided, such as:

```
arangorestore \
  --server.endpoint tcp://192.168.173.13:8531 \
  --server.endpoint tcp://192.168.173.13:8532 \
  --server.username backup \
  --server.database mydb \
  --input-directory "dump"
```

To create the target database when restoring, use a command like this:

```
arangorestore --server.username backup --server.database newdb --create-database true --input-directory "dump"
```

In contrast to the above calls, when working with multiple databases using `--all-databases true`
the parameter `--server.database mydb` must not be specified:

```
arangorestore --server.username backup --all-databases true --create-database true --input-directory "dump-multiple"
```

_arangorestore_ prints out its progress while running, and ends with a line
showing some aggregate statistics:

```
Processed 2 collection(s), read 2256 byte(s) from datafiles, sent 2 batch(es)
```

By default, _arangorestore_ re-creates all non-system collections found in the input
directory and loads data into them. If the target database already contains collections
which are also present in the input directory, the existing collections in the database
are dropped and re-created with the properties and data found in the input directory.

The following parameters are available to adjust this behavior:

- `--create-collection <bool>`: set to `true` to create collections in the target
  database if they don't yet exist. If the target database already contains a 
  collection with the same name, then it is dropped and recreated with the
  same properties as in the dump if the `overwrite` option is also set. 
  If the `overwrite` option is not set, an existing collection is used as is,
  and its properties are not updated nor is its data discarded before restoring.
  If `--create-collection` is set to `false`, then _arangorestore_ does not make any
  attempts to create the collection or modify its properties. Data is restored
  into the existing collections without wiping the collection beforehand.
  If set to `false` and _arangorestore_ encounters a collection that is present in the
  input directory but not in the target database, it aborts with a
  "collection not found" error.
  The default value for `--create-collection` is `true`.
- `--overwrite <bool>`: controls whether existing collections are dropped if
  `--create-collection true` is used. The default value is `true`.
- `--import-data <bool>`: set to `true` to load document data into the collections in
  the target database. Set to `false` to not load any document data. The default value 
  is `true`.
- `--include-system-collections <bool>`: whether or not to include system collections
  when re-creating collections or reloading data. The default value is `false`.

For example, to (re-)create all non-system collections and load document data into them, use:

```
arangorestore --create-collection true --import-data true --input-directory "dump"
```

This drops potentially existing collections in the target database that are also present
in the input directory.

To include system collections too, use `--include-system-collections true`:

```
arangorestore --create-collection true --import-data true --include-system-collections true --input-directory "dump"
```

To (re-)create all non-system collections without loading document data, use:

```
arangorestore --create-collection true --import-data false --input-directory "dump"
```

This also drops existing collections in the target database that are also present in the
input directory.

To just load document data into existing non-system collections, use:

```
arangorestore --create-collection false --import-data true --input-directory "dump"
```

To restrict reloading to just specific collections, use the `--collection` option.
It can be specified multiple times if required:

```
arangorestore --collection myusers --collection myvalues --input-directory "dump"
```

Collections are processed in alphabetical order by _arangorestore_, with all document
collections being processed before all [edge collections](../../../concepts/data-models.md#graph-model).
This remains valid also when multiple threads are in use.

Note however that when restoring an edge collection no internal checks are made in order to validate that
the documents that the edges connect exist. As a consequence, when restoring individual collections
which are part of a graph you are not required to restore in a specific order. 

{{< warning >}}
When restoring only a subset of collections of your database, and graphs are in use, you need
to make sure that you restore all the needed collections (the ones that are part of the graph).
Otherwise, you might end up with edges pointing to non-existing documents.
{{< /warning >}}

To restrict reloading to specific Views, there is the `--view` option.
Should you specify the `--collection` parameter, Views are not restored _unless_ you explicitly
specify them via the `--view` option.

```
arangorestore --collection myusers --view myview --input-directory "dump"
```

In the case of an `arangosearch` View, you must make sure that the linked collections are either
also restored or already present on the server.

## Encryption

See [_arangodump_](../arangodump/examples.md#encryption) for details.

## Reloading Data into a different Collection

_arangorestore_ restores documents and edges with the exact same `_key`,
`_from`, and `_to` values as found in the input directory.

With some creativity you can also use _arangodump_ and _arangorestore_ to transfer data from one
collection into another (either on the same server or not). For example, to copy data from
a collection `myvalues` in database `mydb` into a collection `mycopyvalues` in database `mycopy`,
you can start with the following command:

```
arangodump --collection myvalues --server.database mydb --output-directory "dump"
```

This creates two files, `myvalues.structure.json` and `myvalues.data.json`, in the output 
directory. To load data from the datafile into an existing collection `mycopyvalues` in database 
`mycopy`, rename the files to `mycopyvalues.structure.json` and `mycopyvalues.data.json`.

After that, run the following command:

```
arangorestore --collection mycopyvalues --server.database mycopy --input-directory "dump"
```

## Enabling revision trees for older dumps

<small>Introduced in: v3.8.7, v3.9.2</small>

Collections in ArangoDB 3.8 and later can use an internal format that is based
on revision trees for replication. Using this format has advantages over the
previous format, because changes to the collection on the leader can quickly be
detected when trying to get follower shards in sync.

Dumps taken from older versions of ArangoDB, i.e. ArangoDB 3.7 or before, do not
contain any information about revision trees.
The _arangorestore_ behavior for these collections is as follows:

- In ArangoDB versions before 3.8.7 and 3.9.2, the collections are
  restored without revision trees.
- In ArangoDB versions 3.8.7, 3.9.2 or later, the
  collections use revision trees by default, but you can opt out of this by
  invoking arangorestore with the `--enable-revision-trees false` option.

If the `--enable-revision-trees` startup option is `true` (which is the default value),
then _arangorestore_ adds the necessary attributes for using revision trees
when restoring the collections. It's only done for the attributes which are not
contained in the dump. If the option is set to `false`, _arangorestore_ does not
add the attributes when restoring collections.

Regardless of the setting of this option, _arangorestore_ does not add the
attributes, when they are already present in the dump. You may modify the
attributes manually in the dump if you want to change their values.

## Restoring in a Cluster

To restore data into a Cluster, simply point _arangorestore_ to one of the
_Coordinators_ in your Cluster.

If _arangorestore_ is asked to restore a collection, it uses the same
number of shards, replication factor, and shard keys as when the collection
was dumped. The distribution of the shards to the servers are also the
same as at the time of the dump, provided that the number of _DB-Servers_ in
the cluster dumped from is identical to the number of DB-Servers in the
to-be-restored-to cluster.

To modify the number of _shards_ or the _replication factor_ for all or just
some collections, _arangorestore_ provides the options `--number-of-shards`
and `--replication-factor`. These options
can be specified multiple times as well, in order to override the settings
for dedicated collections, e.g.

```
arangorestore --number-of-shards 2 --number-of-shards mycollection=3 --number-of-shards test=4
```

The above command restores all collections except "mycollection" and "test" with
2 shards. "mycollection" has 3 shards when restored, and "test" has 4.
It is possible to omit the default value and only use
collection-specific overrides. In this case, the number of shards for any
collections not overridden is determined by looking into the
`numberOfShards` values contained in the dump.

The `--replication-factor` options works in the same way, e.g.

```
arangorestore --replication-factor 2 --replication-factor mycollection=1
```

sets the replication factor to `2` for all collections but "mycollection",
which gets a replication factor of just `1`.

{{< info >}}
The options `--number-of-shards` and `replication-factor`, as well as the deprecated
options `--default-number-of-shards` and `--default-replication-factor`, are
**not applicable to system collections**. They are managed by the server.
{{< /info >}}

If a collection was dumped from a single instance and is then restored into
a cluster, the sharding is done by the `_key` attribute by default. You can
manually edit the structural description for the shard keys in the dump files if
required (`*.structure.json`).

If you restore a collection that was dumped from a cluster into a single
ArangoDB instance, the number of shards, replication factor and shard keys are
silently ignored.

### Factors affecting speed of arangorestore in a Cluster

The following factors affect speed of _arangorestore_ in a Cluster:

- **Replication Factor**: the higher the _replication factor_, the more
  time the restore takes. To speed up the restore you can restore
  using a _replication factor_ of `1` and then increase it again
  after the restore. This reduces the number of network hops needed
  during the restore.
- **Restore Parallelization**: if the collections are not restored in
  parallel, the restore speed is highly affected. A parallel restore can
  be done by using the `--threads` option of _arangorestore_.
  Depending on your specific case, you might be able to achieve additional
  parallelization by restoring on multiple _Coordinators_ at the same time.

### Restoring collections with sharding prototypes

_arangorestore_ yields an error when you try to restore a collection whose shard
distribution follows a collection (`distributeShardsLike` property) which does not
exist in the cluster and which was not dumped together with the other collection:

```
arangorestore --collection clonedCollection --server.database mydb --input-directory "dump"

WARNING [c6658] {restore} Error while creating document collection 'clonedCollection': got invalid response from server: HTTP 500: 'Collection not found: prototypeCollection in database _system' while executing restoring collection with this requestPayload: ...

ERROR [cb69f] {restore} got invalid response from server: HTTP 500: 'Collection not found: prototypeCollection in database _system' while executing restoring collection with this requestPayload: ...

INFO [a66e1] {restore} Processed 0 collection(s) from 1 database(s) in 0.04 s total time. Read 0 bytes from datafiles, sent 0 data batch(es) of 0 bytes total size.
```

You need to dump and restore collections that follow the sharding of a
prototype collection together with the prototype collection:

```
arangorestore --collection clonedCollection --collection prototypeCollection --server.database mydb --input-directory "dump"

...
INFO [a66e1] {restore} Processed 2 collection(s) from 1 database(s) in 1.12 s total time. Read 0 bytes from datafiles, sent 0 data batch(es) of 0 bytes total size.
```

## Restore into an authentication-enabled ArangoDB

Of course you can restore data into a password-protected ArangoDB as well.
However this requires certain user rights for the user used in the restore process.
The rights are described in detail in the [Managing Users](../../../operations/administration/user-management/_index.md) chapter.
For restore this short overview is sufficient:

- When importing into an existing database, the given user needs `Administrate`
  access on this database.
- When creating a new database during restore, the given user needs `Administrate`
  access on `_system`. The user is promoted to `Administrate` access on the
  newly created database.
