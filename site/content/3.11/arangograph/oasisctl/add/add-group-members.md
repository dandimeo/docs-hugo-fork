---
layout: default
description: Description of the oasisctl add group members command
title: Oasisctl Add Group Members
menuTitle: Add Group Members
weight: 25
---
## Synopsis
Add members to group

```
oasisctl add group members [flags]
```

## Options
```
  -g, --group-id string          Identifier of the group to add members to
  -h, --help                     help for members
  -o, --organization-id string   Identifier of the organization
  -u, --user-emails strings      A comma separated list of user email addresses
```

## Options Inherited From Parent Commands
```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also
* [oasisctl add group](add-group.md)	 - Add group resources

