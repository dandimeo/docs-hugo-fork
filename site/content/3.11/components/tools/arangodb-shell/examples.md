---
title: _arangosh_ Examples
menuTitle: Examples
weight: 5
description: >-
  By default, arangosh tries to connect to an ArangoDB server running on server localhost on port 8529
archetype: default
---
## Connecting to a server

By default, _arangosh_ tries to connect to an ArangoDB server running on
server `localhost` on port `8529`. It uses the username `root` and an
empty password by default. Additionally, it connects to the default database
(`_system`). All these defaults can be changed using the following 
command-line options:

- `--server.database <string>`: name of the database to connect to
- `--server.endpoint <string>`: endpoint to connect to
- `--server.username <string>`: database username
- `--server.password <string>`: password to use when connecting 
- `--server.authentication <bool>`: whether or not to use authentication

For example, to connect to an ArangoDB server on IP `192.168.173.13` on port
8530 with the user `foo` and using the database `test`, use:

```
arangosh --server.endpoint tcp://192.168.173.13:8530 --server.username foo --server.database test --server.authentication true
```

_arangosh_ then displays a password prompt and tries to connect to the 
server after the password is entered.

{{< warning >}}
At signs `@` in startup option arguments need to be escaped as `@@`.
ArangoDB programs and tools support a
[special syntax `@envvar@`](../../../operations/administration/configuration.md#environment-variables-as-parameters)
that substitutes text wrapped in at signs with the value of an equally called
environment variable. This is most likely an issue with passwords and the
`--server.password` option.

For example, `password@test@123` needs to be passed as
`--server.password password@@test@@123` to work correctly, unless you want
`@test@` to be replaced by whatever the environment variable `test` is set to.
{{< /warning >}}

The shell prints its own version number, and if successfully connected
to a server, the version number of the ArangoDB server.

{{< tip >}}
If the server endpoint is configured for SSL then clients such as _arangosh_
need to connect to it using an SSL socket as well. For example, use `http+ssl://`
as schema in `--server.endpoint` for an SSL-secured HTTP connection.
{{< /tip >}}

The schema of an endpoint is comprised of a protocol and a socket in the format
`protocol+socket://`. There are alternatives and shorthands for some combinations,
`ssl://` is equivalent to `http+ssl://` and `https://` for instance:

Protocol     | Socket           | Schema
-------------|------------------|-----------
HTTP         | TCP              | `http+tcp`, `http+srv`, `http`, `tcp`
HTTP         | TCP with SSL/TLS | `http+ssl`, `https`, `ssl`
HTTP         | Unix             | `http+unix`, `unix`
VelocyStream | TCP              | `vst+tcp`, `vst+srv`, `vst`
VelocyStream | TCP with SSL/TLS | `vst+ssl`, `vsts`
VelocyStream | Unix             | `vst+unix`

## Using _arangosh_

To change the current database after the connection has been made, you
can use the `db._useDatabase()` command in _arangosh_:

```js
---
name: shellUseDB
description: ''
---
db._createDatabase("myapp");
db._useDatabase("myapp");
db._useDatabase("_system");
db._dropDatabase("myapp");
```

To get a list of available commands, _arangosh_ provides a `help()` function.
Calling it displays helpful information.

_arangosh_ also provides auto-completion. Additional information on available 
commands and methods is thus provided by typing the first few letters of a
variable and then pressing the tab key. It is recommend to try this with entering
`db.` (without pressing return) and then pressing tab.
