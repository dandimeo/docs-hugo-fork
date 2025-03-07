---
title: Oasisctl Update Ipallowlist
menuTitle: Update IP Allowlist
weight: 35
description: >-
  Description of the oasisctl update ipallowlist command
archetype: default
---
Update an IP allowlist the authenticated user has access to

## Synopsis

Update an IP allowlist the authenticated user has access to

```
oasisctl update ipallowlist [flags]
```

## Options

```
      --add-cidr-range strings      List of CIDR ranges to add to the IP allowlist
      --description string          Description of the CA certificate
  -h, --help                        help for ipallowlist
  -i, --ipallowlist-id string       Identifier of the IP allowlist
      --name string                 Name of the CA certificate
  -o, --organization-id string      Identifier of the organization
  -p, --project-id string           Identifier of the project
      --remote-inspection-allowed   If set, remote connectivity checks by the Oasis platform are allowed
      --remove-cidr-range strings   List of CIDR ranges to remove from the IP allowlist
```

## Options inherited from parent commands

```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also

* [oasisctl update](_index.md)	 - Update resources

