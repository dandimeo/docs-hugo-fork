---
layout: default
description: Description of the oasisctl lock deployment command
title: Oasisctl Lock Deployment
menuTitle: Lock Deployment
weight: 15
---
## Synopsis
Lock a deployment, so it cannot be deleted

```
oasisctl lock deployment [flags]
```

## Options
```
  -d, --deployment-id string     Identifier of the deployment
  -h, --help                     help for deployment
  -o, --organization-id string   Identifier of the organization
  -p, --project-id string        Identifier of the project
```

## Options Inherited From Parent Commands
```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also
* [oasisctl lock](_index.md)	 - Lock resources

