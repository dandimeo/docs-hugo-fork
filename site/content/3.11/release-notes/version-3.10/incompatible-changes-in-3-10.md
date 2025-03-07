---
title: Incompatible changes in ArangoDB 3.10
menuTitle: Incompatible changes in 3.10
weight: 15
description: >-
  It is recommended to check the following list of incompatible changes before upgrading to ArangoDB 3.10
archetype: default
---
It is recommended to check the following list of incompatible changes **before**
upgrading to ArangoDB 3.10, and adjust any client programs if necessary.

The following incompatible changes have been made in ArangoDB 3.10:

## Declaration of start vertex collections

In cluster deployments, you need to declare collections that an AQL query
implicitly reads from using the [`WITH` operation](../../aql/high-level-operations/with.md).

From version 3.10.0 onward, it is necessary to also declare the collections of
start vertices that are used for [graph traversals](../../aql/graphs/traversals.md)
if you specify start vertices using strings.

In previous versions, the following query would work:

```aql
WITH managers
FOR v, e, p IN 1..2 OUTBOUND 'users/1' usersHaveManagers
  RETURN { v, e, p }
```

Now, you need to declare the `users` collection as well because the start vertex
`users/1` is specified as a string and the `users` collection is not otherwise
specified explicitly in a language construct:

```aql
WITH users, managers
FOR v, e, p IN 1..2 OUTBOUND 'users/1' usersHaveManagers
  RETURN { v, e, p }
```

In the following case, the `users` collection is accessed explicitly to retrieve
the start vertex. Therefore, the collection does not need to be declared using
`WITH`:

```aql
WITH managers
LET startVertex = (FOR u IN users FILTER u._id == "users/1" RETURN u)[0]
FOR v, e, p IN 1..2 OUTBOUND startVertex usersHaveManagers
  RETURN { v, e, p }
```

## Empty Document Updates

ArangoDB 3.10 adds back the optimization for empty document update operations 
(i.e. updates in which no attributes were specified to be updated).
Such updates were handled in a special way in ArangoDB until including 3.7.12,
so that no actual writes were performed and replicated. This
also included the `_rev` value of documents not being changed on an empty
update.

ArangoDB 3.7.13 and 3.8.0 changed this behavior so that empty document updates
actually performed a write operation, leading to the replication of changes and modification
of the `_rev` value of affected.

ArangoDB 3.10 adds back the optimal behavior for empty document
updates, which no longer perform a write operation, do not need to be
replicated and do not change the documents' `_rev` values.

## Foxx / Server Console

Previously a call to `db._version(true)` inside a Foxx app or the server console
would return a different structure than the same call from arangosh.
Foxx/server console would return `{ <details> }` while arangosh would return
`{ server: ..., license: ..., version: ..., details: { <details> }}`.

This is now unified so that the result structure is always consistent with the
one in arangosh. Any Foxx app or script that ran in the server console which
used `db._version(true)` must now be changed to use `db._version(true).details`
instead.

## Indexes

The fulltext index type is now deprecated in favor of [ArangoSearch](../../index-and-search/arangosearch/_index.md).
Fulltext indexes are still usable in this version of ArangoDB, although their usage is
now discouraged.

### Geo Indexes

After an upgrade to 3.10 or higher, consider to drop and recreate geo
indexes. GeoJSON polygons are interpreted slightly differently (and more
correctly) in the newer versions.

Legacy geo indexes will continue to work and continue to produce the
same results as in earlier versions, since they will have the option
`legacyPolygons` implicitly set to `true`.

Newly created indexes will have `legacyPolygons` set to `false` by default
and thus enable the new polygon parsing.

Note that linear rings are not normalized automatically from version 3.10 onward,
following the [GeoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946).
The "interior" of a polygon strictly conforms to the GeoJSON standard:
it lies to the left of the boundary line (in the direction of travel along the
boundary line on the surface of the Earth). This can be the "larger" connected
component of the surface, or the smaller one. Note that this differs from the
interpretation of GeoJSON polygons in version 3.9 and older:

| `legacyPolygons` enabled | `legacyPolygons` disabled |
|:-------------------------|:--------------------------|
| The smaller of the two regions defined by a linear ring is interpreted as the interior of the ring. | The area to the left of the boundary ring's path is considered to be the interior. |
| A ring can at most enclose half the Earth's surface | A ring can enclose the entire surface of the Earth |

This can mean that old polygon GeoJSON data in the database is
suddenly interpreted in a different way. See
[Legacy Polygons](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#legacy-polygons) for details.
Also see the definition of [Polygons](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#polygon)
and [GeoJSON interpretation](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson-interpretation).

## `geojson` Analyzers

If you use `geojson` Analyzers and upgrade from a version below 3.10 to a
version of 3.10 or higher, the interpretation of GeoJSON Polygons changes.

If you have polygons in your data that mean to refer to a relatively small
region but have the boundary running clockwise around the intended interior,
they are interpreted as intended prior to 3.10, but from 3.10 onward, they are
interpreted as "the other side" of the boundary.

Whether a clockwise boundary specifies the complement of the small region
intentionally or not cannot be determined automatically. Please test the new
behavior manually.

See [Legacy Polygons](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#legacy-polygons) for details.

---

<small>Introduced in: v3.10.5</small>

Analyzers of the `geojson` type have a new `legacy` property.
The default is `false`.

This option controls how GeoJSON Polygons are interpreted.

If you require the old behavior, upgrade to at least 3.10.5, drop your
`geojson` Analyzers, and create new ones with `legacy` set to `true`.
If these Analyzers are used in `arangosearch` Views, then they need to be
dropped as well before dropping the Analyzers, and recreated after creating
the new Analyzers.

| `legacy` enabled | `legacy` disabled |
|:-----------------|:------------------|
| The smaller of the two regions defined by a linear ring is interpreted as the interior of the ring. | The area to the left of the boundary ring's path is considered to be the interior. |
| A ring can at most enclose half the Earth's surface | A ring can enclose the entire surface of the Earth |

Also see the definition of [Polygons](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#polygon) and the
[`geojson` Analyzer](../../index-and-search/analyzers.md#geojson) documentation.

## Maximum Array / Object Nesting

When reading any data from JSON or VelocyPack input or when serializing any data to JSON or 
VelocyPack, there is a maximum recursion depth for nested arrays and objects, which is slightly 
below 200. Arrays or objects with higher nesting than this will cause `Too deep nesting in Array/Object`
exceptions. 
The limit is also enforced when converting any server data to JavaScript in Foxx, or
when sending JavaScript input data from Foxx to a server API.
This maximum recursion depth is hard-coded in arangod and all client tools.

## Validation of traversal collection restrictions

<small>Introduced in: v3.9.11, v3.10.7</small>

In AQL graph traversals, you can restrict the vertex and edge collections in the
traversal options like so:

```aql
FOR v, e, p IN 1..3 OUTBOUND 'products/123' components
  OPTIONS {
    vertexCollections: [ "bolts", "screws" ],
    edgeCollections: [ "productsToBolts", "productsToScrews" ]
  }
  RETURN v 
```

If you specify collections that don't exist, queries now fail with
a "collection or view not found" error (code `1203` and HTTP status
`404 Not Found`). In previous versions, unknown vertex collections were ignored,
and the behavior for unknown edge collections was undefined.

Additionally, the collection types are now validated. If a document collection
or View is specified in `edgeCollections`, an error is raised
(code `1218` and HTTP status `400 Bad Request`).

Furthermore, it is now an error if you specify a vertex collection that is not
part of the specified named graph (code `1926` and HTTP status `404 Not Found`).
It is also an error if you specify an edge collection that is not part of the
named graph's definition or of the list of edge collections (code `1939` and
HTTP status `400 Bad Request`).

## Startup Options

### Handling of Invalid Startup Options

Starting with ArangoDB 3.10, the _arangod_ executable and all other client
tools use more specific process exit codes in the following situations:
  
- An unknown startup option name is used: Previously, the exit code was `1`.
  Now, the exit code when using an invalid option is `3` (symbolic exit code
  name `EXIT_INVALID_OPTION_NAME`).
- An invalid value is used for a startup option (e.g. a number that is
  outside the allowed range for the option's underlying value type, or a
  string value is used for a numeric option): Previously, the exit code was
  `1`. Now, the exit code for these case is `4` (symbolic exit code name
  `EXIT_INVALID_OPTION_VALUE`).
- A config file is specified that does not exist: Previously, the exit code
  was either `1` or `6` (symbolic exit code name `EXIT_CONFIG_NOT_FOUND`).
  Now, the exit code in this case is always `6` (`EXIT_CONFIG_NOT_FOUND`).
- A structurally invalid config file is used, e.g. the config file contains
  a line that cannot be parsed: Previously, the exit code in this situation
  was `1`. Now, it is always `6` (symbolic exit code name `EXIT_CONFIG_NOT_FOUND`).

Note that this change can affect any custom scripts that check for startup
failures using the specific exit code `1`. These scripts should be adjusted so
that they check for a non-zero exit code. They can opt into more specific
error handling using the additional exit codes mentioned above, in order to
distinguish between different kinds of startup errors.

### Web Interface Options

The `--frontend.*` startup options were renamed to `--web-interface.*`:

- `--frontend.proxy-request.check` is now `--web-interface.proxy-request.check`
- `--frontend.trusted-proxy` is now `--web-interface.trusted-proxy`
- `--frontend.version-check` is now `--web-interface.version-check`

The former startup options are still supported for backward compatibility.

### RocksDB Options

The default value of the  `--rocksdb.cache-index-and-filter-blocks` startup option was changed
from `false` to `true`. This makes RocksDB track all loaded index and filter blocks in the 
block cache, so they are accounted against the RocksDB's block cache memory limit. 
The default value for the `--rocksdb.enforce-block-cache-size-limit` startup option was also
changed from `false` to `true` to make the RocksDB block cache not temporarily exceed the 
configured memory limit.

These default value changes will make RocksDB adhere much better to the configured memory limit
(configurable via `--rocksdb.block-cache-size`). 
The changes may have a small negative impact on performance because, if the block cache is 
not large enough to hold the data plus the index and filter blocks, additional disk I/O may 
need to be performed compared to the previous versions. 
This is a trade-off between memory usage predictability and performance, and ArangoDB 3.10
will default to more stable and predictable memory usage. If there is still unused RAM 
capacity available, it may be sensible to increase the total size of the RocksDB block cache,
by increasing `--rocksdb.block-cache-size`. Due to the changed configuration, the block 
cache size limit will not be exceeded anymore.

It is possible to opt out of these changes and get back the memory and performance characteristics
of the previous versions by setting the `--rocksdb.cache-index-and-filter-blocks` 
and `--rocksdb.enforce-block-cache-size-limit` startup options to `false` on startup.

### RocksDB File Format

ArangoDB 3.10 internally switches to RocksDB's `format_version` 5, which can still be 
read by older versions of ArangoDB. 
However, ArangoDB 3.10 uses the LZ4 compression scheme to reduce the size of RocksDB .sst
files from LSM tree level 2 onwards. This compression scheme is not supported in ArangoDB 
versions before 3.10, so any database files created with ArangoDB 3.10 or higher cannot be
opened with versions before 3.10.
The internal checksum type of RocksDB .sst files has been changed to xxHash64 in ArangoDB
3.10 for a slight performance improvement.

### Pregel Options

Pregel jobs now have configurable minimum, maximum and default parallelism values. You can set them
by the following startup options:

- `--pregel.min-parallelism`: minimum parallelism usable in Pregel jobs. Defaults to `1`.
- `--pregel.max-parallelism`: maximum parallelism usable in Pregel jobs. Defaults to the
  number of available cores.
- `--pregel.parallelism`: default parallelism to use in Pregel jobs. Defaults to the number
  of available cores divided by 4. The result will be clamped to a value between 1 and 16.

{{< info >}}
The default values of these options may differ from parallelism values effectively
used by previous versions, so it is advised to explicitly set the desired parallelism
values in ArangoDB 3.10.
{{< /info >}}

Pregel now also stores its temporary data in memory-mapped files on disk by default, whereas 
in previous versions the default behavior was to buffer it to RAM.
Storing temporary data in memory-mapped files rather than in RAM has the advantage of lowering
the RAM usage, which reduces the likelihood of out-of-memory situations.
However, storing the files on disk requires disk capacity, so that instead of running out
of RAM it is now possible to run out of disk space.

{{< info >}}
It is advised to set the storage location for Pregel's memory-mapped files explicitly
in ArangoDB 3.10. The following startup options are available for the configuration of
memory-mapped files: `--pregel.memory-mapped-files` and `--pregel.memory-mapped-files-location-type`.
{{< /info >}}

For more information on the new options, please refer to [ArangoDB Server Pregel Options](../../components/arangodb-server/options.md#pregel).

## HTTP RESTful API

### Validation of collections in named graphs

The `/_api/gharial` endpoints for named graphs have changed:

- If you reference a vertex collection in the `_from` or `_to` attribute of an
  edge that doesn't belong to the graph, an error with the number `1947` is
  returned. The HTTP status code of such an `ERROR_GRAPH_REFERENCED_VERTEX_COLLECTION_NOT_USED`
  error has been changed from `400` to `404`. This change aligns the behavior to
  the similar `ERROR_GRAPH_EDGE_COLLECTION_NOT_USED` error (number `1930`).

- Write operations now check if the specified vertex or edge collection is part
  of the graph definition. If you try to create a vertex via
  `POST /_api/gharial/{graph}/vertex/{collection}` but the `collection` doesn't
  belong to the `graph`, then the `ERROR_GRAPH_REFERENCED_VERTEX_COLLECTION_NOT_USED`
  error is returned. If you try to create an edge via
  `POST /_api/gharial/{graph}/edge/{collection}` but the `collection` doesn't
  belong to the `graph`, then the error is `ERROR_GRAPH_EDGE_COLLECTION_NOT_USED`.

### Encoding of revision IDs

<small>Introduced in: v3.8.8, v3.9.4, v3.10.1</small>

- `GET /_api/collection/<collection-name>/revision`: The revision ID was
  previously returned as numeric value, and now it is returned as
  a string value with either numeric encoding or HLC-encoding inside.
- `GET /_api/collection/<collection-name>/checksum`: The revision ID in
  the `revision` attribute was previously encoded as a numeric value
  in single server, and as a string in cluster. This is now unified so
  that the `revision` attribute always contains a string value with
  either numeric encoding or HLC-encoding inside.

## Client Tools

### arangobench

Renamed the `--concurrency` startup option to `--threads`.

The following deprecated arangobench testcases have been removed from _arangobench_:
- `aqltrx`
- `aqlv8`
- `counttrx`
- `deadlocktrx`
- `multi-collection`
- `multitrx`
- `random-shapes`
- `shapes`
- `shapes-append`
- `skiplist`
- `stream-cursor`

These test cases had been deprecated since ArangoDB 3.9.

The testcase `hash` was renamed to `persistent-index` to better reflect its
scope.

### arangoexport

To improve naming consistency across different client tools, the existing
_arangoexport_ startup options `--query` and `--query-max-runtime` were renamed
to `--custom-query` and `--custom-query-max-runtime`.

Using the old option names (`--query` and `--query-max-runtime`) is still
supported and will implicitly use the `--custom-query` and
`--custom-query-max-runtime` options under the hood. Client scripts should
eventually be updated to use the new option name, however.

### ArangoDB Starter

The _ArangoDB Starter_ comes with the following usability improvements:
- Headers are now added to generated command files, indicating the purpose of
  the file.
- The process output is now shown when errors occur during process startup.
- When passing through other database options, explicit hints are now displayed
  to indicate how to pass those options.
- The Starter now returns exit code `1` if it encounters any errors while
  starting. Previously, the exit code was `0`. Note that this change can affect
  any custom scripts that check for startup errors or invalid command line options.
  These scripts can be adjusted so that they check for a non-zero exit code.  
