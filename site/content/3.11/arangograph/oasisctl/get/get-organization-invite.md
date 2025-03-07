---
layout: default
description: Description of the oasisctl get organization invite command
title: Oasisctl Get Organization Invite
menuTitle: Get Organization Invite
weight: 110
---
## Synopsis
Get an organization invite the authenticated user has access to

```
oasisctl get organization invite [flags]
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
* [oasisctl get organization](get-organization.md)	 - Get an organization the authenticated user is a member of

