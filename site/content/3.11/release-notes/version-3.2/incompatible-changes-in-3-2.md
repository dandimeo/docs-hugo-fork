---
title: Incompatible changes in ArangoDB 3.2
menuTitle: Incompatible changes in 3.2
weight: 15
description: >-
  It is recommended to check the following list of incompatible changes before upgrading to ArangoDB 3
archetype: default
---
It is recommended to check the following list of incompatible changes **before**
upgrading to ArangoDB 3.2, and adjust any client programs if necessary.

## AQL

- AQL breaking change in cluster:
  The SHORTEST_PATH statement using edge-collection names instead
  of a graph name now requires to explicitly name the vertex-collection names
  within the AQL query in the cluster. It can be done by adding `WITH <name>`
  at the beginning of the query.

  Example:
  ```
  FOR v,e IN OUTBOUND SHORTEST_PATH @start TO @target edges [...]
  ```

  Now has to be:

  ```
  WITH vertices
  FOR v,e IN OUTBOUND SHORTEST_PATH @start TO @target edges [...]
  ```

  This change is due to avoid dead-lock situations in clustered case.
  An error stating the above is included.

## REST API

- Removed undocumented internal HTTP API:
  - PUT /_api/edges

  The documented GET /_api/edges and the undocumented POST /_api/edges remains unmodified.

- change undocumented behavior in case of invalid revision ids in
  `If-Match` and `If-None-Match` headers from returning HTTP status code 400 (bad request)
  to returning HTTP status code 412 (precondition failed).

- the REST API for fetching the list of currently running AQL queries and the REST API
  for fetching the list of slow AQL queries now return an extra *bindVars* attribute which
  contains the bind parameters used by the queries.

  This affects the return values of the following API endpoints:
  - GET /_api/query/current
  - GET /_api/query/slow

- The REST API for retrieving indexes (GET /_api/index) now returns the *deduplicate*
  attribute for each index

- The REST API for creating indexes (POST /_api/index) now accepts the optional *deduplicate*
  attribute

- The REST API for executing a server-side transaction (POST /_api/transaction) now accepts the optional attributes: `maxTransactionSize`, `intermediateCommitCount`, `intermediateCommitSize`

- The REST API for creating a cursor (POST /_api/cursor) now accepts the optional attributes: `failOnWarning`, `maxTransactionSize`, `maxWarningCount`, `intermediateCommitCount`, `satelliteSyncWait`, `intermediateCommitSize`. `skipInaccessibleCollections`

## JavaScript API

- change undocumented behavior in case of invalid revision ids in
  JavaScript document operations from returning error code 1239 ("illegal document revision")
  to returning error code 1200 ("conflict").

- the `collection.getIndexes()` function now returns the *deduplicate* attribute for each index

- the `collection.ensureIndex()` function now accepts the optional *deduplicate* attribute

## Foxx

- JWT tokens issued by the built-in [JWT session storage](../../develop/foxx-microservices/reference/sessions-middleware/session-storages/jwt-storage.md) now correctly specify the `iat` and `exp` values in seconds rather than milliseconds as specified in the JSON Web Token standard.

  This may result in previously expired tokens using milliseconds being incorrectly accepted. For this reason it is recommended to replace the signing `secret` or set the new `maxExp` option to a reasonable value that is smaller than the oldest issued expiration timestamp.

  For example setting `maxExp` to `10**12` would invalidate all incorrectly issued tokens before 9 September 2001 without impairing new tokens until the year 33658 (at which point these tokens are hopefully no longer relevant).

- ArangoDB running in standalone mode will commit all services in the `javascript.app-path` to the database on startup. This may result in uninstalled services showing up in ArangoDB if they were not properly removed from the filesystem.

- ArangoDB Coordinators in a cluster now perform a self-healing step during startup to ensure installed services are consistent across all Coordinators. We recommend backing up your services and configuration before upgrading to ArangoDB 3.2, especially if you have made use of the development mode.

- Services installed before upgrading to 3.2 (including services installed on alpha releases of ArangoDB 3.2) are **NOT** picked up by the Coordinator self-healing watchdog. This can be solved by either upgrading/replacing these services or by using the ["commit" route of the Foxx service management HTTP API](../../develop/http/foxx.md#miscellaneous), which commits the exact services installed on a given Coordinator to the cluster. New services will be picked up automatically.

- The format used by Foxx to store internal service metadata in the database has been simplified and existing documents will be updated to the new format. If you have made any changes to the data stored in the `_apps` system collection, you may wish to export these changes as they will be overwritten.

- There is now an [official HTTP API for managing services](../../develop/http/foxx.md). If you were previously using any of the undocumented APIs or the routes used by the administrative web interface we highly recommend migrating to the new API. The old undocumented HTTP API for managing services is deprecated and will be removed in a future version of ArangoDB.

- Although changes to the filesystem outside of development mode were already strongly discouraged, this is a reminder that they are no longer supported. All files generated by services (whether by a setup script or during normal operation such as uploads) should either be stored outside the service directory or be considered extremely volatile.

- Introduced distinction between `arangoUser` and `authorized` in Foxx requests. Cluster internal requests will never have an `arangoUser` but are authorized. In earlier versions of ArangoDB parts of the statistics were not accessible by the Coordinators because the underlying Foxx service couldn't authorize the requests. It now correctly checks the new `req.authorized` property. `req.arangoUser` still works as before. End-users may use this new property as well to easily check if a request is authorized or not regardless of a specific user.

## Command-line options changed

- --server.maximal-queue-size is now an absolute maximum. If the queue is
  full, then 503 is returned. Setting it to 0 means "no limit". The default
  value for this option is now `0`.

- the default value for `--ssl.protocol` has been changed from `4` (TLSv1) to `5` (TLSv1.2).

- the startup options `--database.revision-cache-chunk-size` and
  `--database.revision-cache-target-size` are now obsolete and do nothing

- the startup option `--database.index-threads` option is now obsolete

- the option `--javascript.v8-contexts` is now an absolute maximum. The server
  may start less V8 contexts for JavaScript execution at startup. If at some
  point the server needs more V8 contexts it may start them dynamically, until
  the number of V8 contexts reaches the value of `--javascript.v8-contexts`.

  the minimum number of V8 contexts to create at startup can be configured via
  the new startup option `--javascript.v8-contexts-minimum`.

- added command-line option `--javascript.allow-admin-execute`

  This option can be used to control whether user-defined JavaScript code
  is allowed to be executed on server by sending via HTTP to the API endpoint
  `/_admin/execute`  with an authenticated user account.
  The default value is `false`, which disables the execution of user-defined
  code. This is also the recommended setting for production. In test environments,
  it may be convenient to turn the option on in order to send arbitrary setup
  or teardown commands for execution on the server.

  The introduction of this option changes the default behavior of ArangoDB 3.2:
  3.2 now by default disables the execution of JavaScript code via this API,
  whereas earlier versions allowed it. To restore the old behavior, it is
  necessary to set the option to `true`.

## Users Management

- It is no longer supported to access the `_users` collection in any way directly, except through the official `@arangodb/users` module or the `/_api/users` REST API.

- The access to the `_users` collection from outside of the arangod server process is now forbidden (Through drivers, arangosh or the REST API). Foxx services are still be able to access the `_users` collection for now, but this might change in future minor releases.

- The internal format of the documents in the `_users` collection has changed from previous versions

- The `_queues` collection only allows read-only access from outside of the arangod server process.

- Accessing `_queues` is only supported through the official `@arangodb/queues` module for Foxx apps.
