---
title: Oasisctl Get Cacertificate
menuTitle: Get CA Certificate
weight: 30
description: >-
  Description of the oasisctl get cacertificate command
archetype: default
---
Get a CA certificate the authenticated user has access to

## Synopsis

Get a CA certificate the authenticated user has access to

```
oasisctl get cacertificate [flags]
```

## Options

```
  -c, --cacertificate-id string   Identifier of the CA certificate
  -h, --help                      help for cacertificate
  -o, --organization-id string    Identifier of the organization
  -p, --project-id string         Identifier of the project
```

## Options inherited from parent commands

```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also

* [oasisctl get](_index.md)	 - Get information

