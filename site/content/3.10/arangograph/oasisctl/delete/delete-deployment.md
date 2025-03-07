---
title: Oasisctl Delete Deployment
menuTitle: Delete Deployment
weight: 45
description: >-
  Description of the oasisctl delete deployment command
archetype: default
---
Delete a deployment the authenticated user has access to

## Synopsis

Delete a deployment the authenticated user has access to

```
oasisctl delete deployment [flags]
```

## Options

```
  -d, --deployment-id string     Identifier of the deployment
  -h, --help                     help for deployment
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

