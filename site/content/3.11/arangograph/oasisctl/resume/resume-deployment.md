---
layout: default
description: Description of the oasisctl resume deployment command
title: Oasisctl Resume Deployment
menuTitle: Resume Deployment
weight: 10
---
## Synopsis
Resume a paused deployment the authenticated user has access to

```
oasisctl resume deployment [flags]
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
* [oasisctl resume](_index.md)	 - Resume resources

