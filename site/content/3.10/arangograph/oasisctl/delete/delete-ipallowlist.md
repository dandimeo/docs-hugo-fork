---
title: Oasisctl Delete Ipallowlist
menuTitle: Delete IP Allowlist
weight: 70
description: >-
  Description of the oasisctl delete ipallowlist command
archetype: default
---
Delete an IP allowlist the authenticated user has access to

## Synopsis

Delete an IP allowlist the authenticated user has access to

```
oasisctl delete ipallowlist [flags]
```

## Options

```
  -h, --help                     help for ipallowlist
  -i, --ipallowlist-id string    Identifier of the IP allowlist
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

* [oasisctl delete](_index.md)	 - Delete resources

