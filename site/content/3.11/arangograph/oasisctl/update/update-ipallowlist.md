---
layout: default
description: Description of the oasisctl update ipallowlist command
title: Oasisctl Update Ipallowlist
menuTitle: Update Ipallowlist
weight: 40
---
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

## Options Inherited From Parent Commands
```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also
* [oasisctl update](_index.md)	 - Update resources

