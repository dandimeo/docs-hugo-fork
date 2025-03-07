---
layout: default
description: Description of the oasisctl delete organization invite command
title: Oasisctl Delete Organization Invite
menuTitle: Delete Organization Invite
weight: 100
---
## Synopsis
Delete an organization invite the authenticated user has access to

```
oasisctl delete organization invite [flags]
```

## Options
```
  -h, --help                     help for invite
  -i, --invite-id string         Identifier of the organization invite
  -o, --organization-id string   Identifier of the organization
```

## Options Inherited From Parent Commands
```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also
* [oasisctl delete organization](delete-organization.md)	 - Delete an organization the authenticated user has access to

