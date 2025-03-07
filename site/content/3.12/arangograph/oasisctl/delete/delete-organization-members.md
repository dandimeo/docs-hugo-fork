---
title: Oasisctl Delete Organization Members
menuTitle: Delete Organization Members
weight: 100
description: >-
  Description of the oasisctl delete organization members command
archetype: default
---
Delete members from organization

## Synopsis

Delete members from organization

```
oasisctl delete organization members [flags]
```

## Options

```
  -h, --help                     help for members
  -o, --organization-id string   Identifier of the organization
  -u, --user-emails strings      A comma separated list of user email addresses
```

## Options inherited from parent commands

```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also

* [oasisctl delete organization](delete-organization.md)	 - Delete an organization the authenticated user has access to

