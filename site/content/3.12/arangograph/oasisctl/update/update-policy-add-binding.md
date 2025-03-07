---
title: Oasisctl Update Policy Add Binding
menuTitle: Update Policy Add Binding
weight: 95
description: >-
  Description of the oasisctl update policy add binding command
archetype: default
---
Add a role binding to a policy

## Synopsis

Add a role binding to a policy

```
oasisctl update policy add binding [flags]
```

## Options

```
      --group-id strings   Identifiers of the groups to add bindings for
  -h, --help               help for binding
  -r, --role-id string     Identifier of the role to bind to
  -u, --url string         URL of the resource to update the policy for
      --user-id strings    Identifiers of the users to add bindings for
```

## Options inherited from parent commands

```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also

* [oasisctl update policy add](update-policy-add.md)	 - Add to a policy

