---
title: Oasisctl Get Server Status
menuTitle: Get Server Status
weight: 160
description: >-
  Description of the oasisctl get server status command
archetype: default
---
Get the status of servers for a deployment

## Synopsis

Get the status of servers for a deployment

```
oasisctl get server status [flags]
```

## Options

```
  -d, --deployment-id string     Identifier of the deployment
  -h, --help                     help for status
  -o, --organization-id string   Identifier of the organization
  -p, --project-id string        Identifier of the project
```

## Options inherited from parent commands

```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also

* [oasisctl get server](get-server.md)	 - Get server information

