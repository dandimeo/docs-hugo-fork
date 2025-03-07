---
title: API Changes in ArangoDB 3.9
menuTitle: API changes in 3.9
weight: 20
description: >-
  ArangoDB v3.9 Release Notes API Changes
archetype: default
---
This document summarizes the HTTP API changes and other API changes in ArangoDB 3.9.
The target audience for this document are developers who maintain drivers and
integrations for ArangoDB 3.9.

## HTTP RESTful API

### Behavior changes

#### Graph API (Gharial)

The following changes affect the behavior of the RESTful graph APIs at
endpoints starting with path `/_api/gharial/`:

The options object now supports a new optional field `satellites` in the
Enterprise Edition when creating a graph (POST method). If set, it needs to be
an array of collection names. Each name must be a string and valid as collection
name. The `satellites` option is ignored in the Community Edition.

Using `satellites` during SmartGraph creation will result in a SmartGraph
with SatelliteCollections.
Using `satellites` during Disjoint SmartGraph creation will result in a
Disjoint SmartGraph with SatelliteCollections.

(Disjoint) SmartGraphs using SatelliteCollections are capable of having
SatelliteCollections in their
graph definitions. If a collection is named in `satellites` and also used in the
graph definition itself (e.g. EdgeDefinition), this collection will be created
as a SatelliteCollection. (Disjoint) SmartGraphs using SatelliteCollections are
then capable of executing all types of graph queries between the regular
SmartCollections and SatelliteCollections.

The following changes affect the behavior of the RESTful graph APIs at
endpoints starting with path `/_api/gharial/{graph}/edge` and
`/_api/gharial/{graph}/vertex`:

Added new optional `options` object that can be set when creating a new or
modifying an existing edge definition (POST / PUT method), as well as when
creating a new vertex collection (POST method). This was not available in
previous ArangoDB versions. The `options` object can currently contain a field
called `satellites` only.

The `satellites` field must be an array with one or more collection name strings.
If an EdgeDefinition contains a collection name that is also contained in the
`satellites` option, or if the vertex collection to add is contained in the
`satellites` option, the collection will be created as a SatelliteCollection.
Otherwise, it will be ignored. This option only takes effect using SmartGraphs.

Also see [Graph Management](../../develop/http/graphs/named-graphs.md#management).

#### Extended naming convention for databases

There is a new startup option `--database.extended-names-databases` to allow
database names to contain most UTF-8 characters.

The feature is disabled by default to ensure compatibility with existing client
drivers and applications that only support ASCII names according to the
traditional database naming convention used in previous ArangoDB versions.

If the feature is enabled, then any endpoints that contain database names in the URL 
may contain special characters that were previously not allowed
(percent-encoded). They are also to be expected in payloads that contain
database names. 

For client applications and drivers that assemble URLs containing database names,
it is required that database names are properly URL-encoded in URLs. In addition,
database names containing UTF-8 characters must be 
[NFC-normalized](https://en.wikipedia.org/wiki/Unicode_equivalence#Normal_forms).
Non-NFC-normalized names will be rejected by arangod.
This is true for any REST API endpoint in arangod if the extended database naming
convention is used.

{{< info >}}
The extended naming convention is an **experimental** feature
but will become the norm in a future version. Check if your drivers and
client applications are prepared for this feature before enabling it.
{{< /info >}}

Also see [Database names](../../concepts/data-structure/databases.md#database-names).

#### Overload control

Starting with version 3.9.0, ArangoDB returns an `x-arango-queue-time-seconds`
HTTP header with all responses. This header contains the most recent request
queueing/dequeuing time (in seconds) as tracked by the server's scheduler.
This value can be used by client applications and drivers to detect server
overload and react on it.

The arangod startup option `--http.return-queue-time-header` can be set to
`false` to suppress these headers in responses sent by arangod.

In a cluster, the value returned in the `x-arango-queue-time-seconds` header
is the most recent queueing/dequeuing request time of the Coordinator the
request was sent to, except if the request is forwarded by the Coordinator to
another Coordinator. In that case, the value will indicate the current
queueing/dequeuing time of the forwarded-to Coordinator.

In addition, client applications and drivers can optionally augment the
requests they send to arangod with the header `x-arango-queue-time-seconds`.
If set, the value of the header should contain the maximum server-side
queuing time (in seconds) that the client application is willing to accept.
If the header is set in an incoming request, arangod will compare the current
dequeuing time from its scheduler with the maximum queue time value contained
in the request header. If the current queueing time exceeds the value set
in the header, arangod will reject the request and return HTTP 412
(precondition failed) with the error code 21004 (queue time violated).
In a cluster, the `x-arango-queue-time-seconds` request header will be
checked on the receiving Coordinator, before any request forwarding.

#### Cursor API

<small>Introduced in: v3.9.11, v3.10.7</small>

In AQL graph traversals (`POST /_api/cursor` endpoint), you can restrict the
vertex and edge collections in the traversal options like so:

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

#### Document API

<small>Introduced in: v3.9.12</small>

Using the Document API for reading multiple documents used to return an error
if the request body was an empty array. Example:

```bash
> curl -XPUT -d '[]' 'http://localhost:8529/_api/document/coll?onlyget=true'
{"code":500,"error":true,"errorMessage":"internal error","errorNum":4}
```

Now, a request like this succeeds and returns an empty array as response.

### Endpoint return value changes

- All collections in ArangoDB are now always in the `loaded` state. APIs return
  return a collection's status will now return it as `loaded`, unconditionally.

- The HTTP endpoints for loading and unloading collections (i.e. HTTP PUT
  `/_api/collection/<collection>/load` and HTTP PUT `/_api/collection/<collection>/unload`)
  have been turned into no-ops. They still exist in ArangoDB 3.9, but do not
  serve any purpose and are deprecated.

- Changed the encoding of revision IDs returned by the below listed REST APIs.

  <small>Introduced in: v3.8.8, v3.9.4</small>

  - `GET /_api/collection/<collection-name>/revision`: The revision ID was
    previously returned as numeric value, and now it is returned as
    a string value with either numeric encoding or HLC-encoding inside.
  - `GET /_api/collection/<collection-name>/checksum`: The revision ID in
    the `revision` attribute was previously encoded as a numeric value
    in single server, and as a string in cluster. This is now unified so
    that the `revision` attribute always contains a string value with
    either numeric encoding or HLC-encoding inside.

### Endpoints added

#### Support Info API

The HTTP REST API endpoint `GET /_admin/support-info` was added for retrieving
deployment information for support purposes. The endpoint returns data about the
ArangoDB version used, the host (operating system, server ID, CPU and storage capacity,
current utilization, a few metrics) and the other servers in the deployment
(in case of active failover or cluster deployments).

As this API may reveal sensitive data about the deployment, it can only be 
accessed from inside the `_system` database. In addition, there is a policy control 
startup option `--server.support-info-api` that controls if and to whom the API 
is made available. This option can have the following values:

- `disabled`: support info API is disabled.
- `jwt`: support info API can only be accessed via superuser JWTs.
- `admin` (default): the support info API can only be accessed by admin users and superuser JWTs.
- `public`: everyone with access to the `_system` database can access the support info API.

#### License Management (Enterprise Edition)

Two endpoints were added for the new
[License Management](../../operations/administration/license-management.md). They can be called on
single servers, Coordinators and DB-Servers:

- `GET /_admin/license`: Query license information and status.

  ```json
  {
    "features": {
      "expires": 1640255734
    },
    "license": "JD4EOk5fcx...HgdnWw==",
    "version": 1,
    "status": "good"
  }
  ```

  - `features`:
    - `expires`: Unix timestamp (seconds since January 1st, 1970 UTC)
  - `license`: Encrypted and base64-encoded license key
  - `version`: License version number
  - `status`:
    - `good`: The license is valid for more than 2 weeks.
    - `expiring`: The license is valid for less than 2 weeks.
    - `expired`: The license has expired. In this situation, no new
      Enterprise Edition features can be utilized.
    - `read-only`: The license is expired over 2 weeks. The instance is now
      restricted to read-only mode.

- `PUT /_admin/license`: Set a new license key. Expects the key as string in the
  request body (wrapped in double quotes).

  Server reply on success:

  ```json
  {
    "result": {
      "error": false,
      "code": 201
    }
  }
  ```

  If the new license expires sooner than the current one, an error will be
  returned. The query parameter `?force=true` can be set to update it anyway.

  ```json
  {
    "code": 400,
    "error": true,
    "errorMessage": "This license expires sooner than the existing. You may override this by specifying force=true with invocation.",
    "errorNum": 9007
  }
  ```

### Endpoints augmented

#### Cursor API

The HTTP REST API endpoint `POST /_api/cursor` can now handle an 
additional sub-attribute `fillBlockCache` for its `options` attribute.
`fillBlockCache` controls whether the to-be-executed query should
populate the RocksDB block cache with the data read by the query.
This is an optional attribute, and its default value is `true`, meaning
that the block cache will be populated. This functionality was also backported
to v3.8.1.

The HTTP REST API endpoint `POST /_api/cursor` can also handle the
sub-attribute `maxNodesPerCallstack`, which controls after how many
execution nodes in a query a stack split should be performed. This is
only relevant for very large queries. If this option is not specified,
the default value is 200 on MacOS, and 250 for other platforms.
Please note that this option is only useful for testing and debugging 
and normally does not need any adjustment.

#### Log API

The HTTP REST API endpoint `PUT /_admin/log/level` can now handle the
pseudo log topic `"all"`. Setting the log level for the "all" log topic will
adjust the log level for **all existing log topics**.
For example, sending the JSON object to this API

```json
{"all":"debug"}
```

will set all log topics to log level "debug".

#### Authentication API

The HTTP REST API endpoint `POST /_open/auth` now returns JWTs with a shorter
lifetime of one hour by default. You can adjust the lifetime with the
`--server.session-timeout` startup option.

#### Analyzers API

Analyzers with a `locale` property use a new syntax. The encoding (`.utf-8`)
does not need to be set anymore. The `collation` Analyzer supports
`language[_COUNTRY][_VARIANT][@keywords]` (square bracket denote optional parts).
The `text` and `norm` Analyzers support `language[_COUNTRY][_VARIANT]`, the `stem`
Analyzer only `language`. The former syntax is still supported but automatically
normalized to the new syntax.

#### Views API

Views of the type `arangosearch` support new caching options in the
Enterprise Edition.

<small>Introduced in: v3.9.5</small>

- A `cache` option for individual View links or fields (boolean, default: `false`).
- A `cache` option in the definition of a `storedValues` View property
  (boolean, immutable, default: `false`).

<small>Introduced in: v3.9.6</small>

- A `primarySortCache` View property (boolean, immutable, default: `false`).
- A `primaryKeyCache` View property (boolean, immutable, default: `false`).

The `POST /_api/view` endpoint accepts these new options for `arangosearch`
Views, the `GET /_api/view/<view-name>/properties` endpoint may return these
options, and you can change the `cache` View link/field property with the
`PUT /_api/view/<view-name>/properties` and `PATCH /_api/view/<view-name>/properties`
endpoints.

See the [`arangosearch` Views Reference](../../index-and-search/arangosearch/arangosearch-views-reference.md#link-properties)
for details.

#### Document API

<small>Introduced in: v3.9.6</small>

The following endpoints support a new, experimental `refillIndexCaches` query
parameter to repopulate the edge cache after requests that insert, update,
replace, or remove single or multiple edge documents:

- `POST /_api/document/{collection}`
- `PATCH /_api/document/{collection}/{key}`
- `PUT /_api/document/{collection}/{key}`
- `DELETE /_api/document/{collection}/{key}`

It is a boolean option and the default is `false`.

#### Metrics API

<small>Introduced in: v3.8.7, v3.9.2</small>

I/O heartbeat metrics have been added:

| Label | Description |
|:------|:------------|
| `arangodb_ioheartbeat_delays_total` | Total number of delayed I/O heartbeats. |
| `arangodb_ioheartbeat_duration` | Histogram of execution times in microseconds. |
| `arangodb_ioheartbeat_failures_total` | Total number of failures. |

---

<small>Introduced in: v3.9.5</small>

The `GET /_admin/metrics/v2` and `GET /_admin/metrics` endpoints includes a new
metrics `arangodb_search_columns_cache_size` which reports the ArangoSearch
column cache size.

---

<small>Introduced in: v3.8.9, v3.9.6</small>

The metrics endpoints include the following new traffic accounting metrics:

- `arangodb_client_user_connection_statistics_bytes_received`
- `arangodb_client_user_connection_statistics_bytes_sent`
- `arangodb_http1_connections_total`

---

<small>Introduced in: v3.9.6</small>

The metrics endpoints include the following new edge cache (re-)filling metrics:

- `rocksdb_cache_auto_refill_loaded_total`
- `rocksdb_cache_auto_refill_dropped_total`
- `rocksdb_cache_full_index_refills_total`

---

<small>Introduced in: v3.9.10</small>

The following metrics for write-ahead log (WAL) file tracking have been added:

| Label | Description |
|:------|:------------|
| `rocksdb_live_wal_files` | Number of live RocksDB WAL files. |
| `rocksdb_wal_released_tick_flush` | Lower bound sequence number from which WAL files need to be kept because of external flushing needs. |
| `rocksdb_wal_released_tick_replication` | Lower bound sequence number from which WAL files need to be kept because of replication. |
| `arangodb_flush_subscriptions` | Number of currently active flush subscriptions. |

---

The following metrics for diagnosing delays in cluster-internal network requests
have been added:

<small>Introduced in: v3.9.11</small>

| Label | Description |
|:------|:------------|
| `arangodb_network_dequeue_duration` | Internal request duration for the dequeue in seconds. |
| `arangodb_network_response_duration` | Internal request duration from fully sent till response received in seconds. |
| `arangodb_network_send_duration` | Internal request send duration in seconds. |
| `arangodb_network_unfinished_sends_total` | Number of internal requests for which sending has not finished. |

### Endpoints moved

#### Cluster API redirects

Since ArangoDB 3.7, some cluster APIs were made available under different
paths. The old paths were left in place and simply redirected to the new
address. These redirects have now been removed in ArangoDB 3.9.

The following list shows the old, now dysfunctional paths and their
replacements:

- `/_admin/clusterNodeVersion`: replaced by `/_admin/cluster/nodeVersion`
- `/_admin/clusterNodeEngine`: replaced by `/_admin/cluster/nodeEngine`
- `/_admin/clusterNodeStats`: replaced by `/_admin/cluster/nodeStatistics`
- `/_admin/clusterStatistics`: replaced by `/_admin/cluster/statistics`

Using the replacements will work from ArangoDB 3.7 onwards already, so
any client applications that still call the old addresses can be adjusted
to call the new addresses from 3.7 onwards.

### Endpoints deprecated

The REST API endpoint GET `/_api/replication/logger-follow` is deprecated
since ArangoDB 3.4.0 and will be removed in a future version. Client
applications should use the endpoint `/_api/wal/tail` instead, which is
available since ArangoDB 3.3. This is a reminder to migrate to the other
endpoint.

### Endpoints removed

#### Redirects

The following API redirect endpoints have been removed in ArangoDB 3.9.
These endpoints have been only been redirections since ArangoDB 3.7. Any
caller of these API endpoints should use the updated endpoints:

- `/_admin/clusterNodeVersion`: use `/_admin/cluster/nodeVersion`
- `/_admin/clusterNodeEngine`: use `/_admin/cluster/nodeEngine`
- `/_admin/clusterNodeStats`: use `/_admin/cluster/nodeStatistics`
- `/_admin/clusterStatistics`: use `/_admin/cluster/statistics`

The REST API endpoint `/_msg/please-upgrade-handler` has been removed in 
ArangoDB 3.9 as it is no longer needed. Its purpose was to display a static 
message.

#### Export API

The REST API endpoint `/_api/export` has been removed in ArangoDB 3.9.
This endpoint was previously only present in single server, but never
supported in cluster deployments.

The purpose of the endpoint was to provide the full data of a collection
without holding collection locks for a long time, which was useful for
the MMFile storage engine with its collection-level locks.

The MMFiles engine is gone since ArangoDB 3.7, and the only remaining
storage engine since then is RocksDB. For the RocksDB engine, the
`/_api/export` endpoint internally used a streaming AQL query such as

```aql
FOR doc IN @@collection RETURN doc
```

anyway. To remove API redundancy, the API endpoint has been deprecated
in ArangoDB 3.8 and is now removed. If the functionality is still required
by client applications, running a streaming AQL query can be used as a
substitution.

## JavaScript API

### Loading and unloading of collections

All collections in ArangoDB are now always in the "loaded" state. Any 
JavaScript functions for returning a collection's status will now return 
"loaded", unconditionally.

The JavaScript functions for loading and unloading collections (i.e. 
`db.<collection>.load()` and `db.<collection>.unload()`) have been turned 
into no-ops. They still exist in ArangoDB 3.9, but do not serve any purpose 
and are deprecated.

### AQL queries

<small>Introduced in: v3.9.11, v3.10.7</small>

If you specify collections that don't exist in the options of AQL graph traversals
(`vertexCollections`, `edgeCollections`), queries now fail. In previous versions,
unknown vertex collections were ignored, and the behavior for unknown
edge collections was undefined.

Additionally, queries fail if you specify a document collection or View
in `edgeCollections`.
