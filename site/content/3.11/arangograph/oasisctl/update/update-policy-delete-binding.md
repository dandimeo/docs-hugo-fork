---
layout: default
description: Description of the oasisctl update policy delete binding command
title: Oasisctl Update Policy Delete Binding
menuTitle: Update Policy Delete Binding
weight: 110
---
## Synopsis
Delete a role binding from a policy

```
oasisctl update policy delete binding [flags]
```

## Options
```
      --group-id strings   Identifiers of the groups to delete bindings for
  -h, --help               help for binding
  -r, --role-id string     Identifier of the role to delete bind for
  -u, --url string         URL of the resource to update the policy for
      --user-id strings    Identifiers of the users to delete bindings for
```

## Options Inherited From Parent Commands
```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also
* [oasisctl update policy delete](update-policy-delete.md)	 - Delete from a policy

