---
layout: default
description: Description of the oasisctl delete auditlog archive events command
title: Oasisctl Delete Auditlog Archive Events
menuTitle: Delete Audit Log Archive Events
weight: 25
---
## Synopsis
Delete auditlog archive events

```
oasisctl delete auditlog archive events [flags]
```

## Options
```
  -i, --auditlog-archive-id string   Identifier of the auditlog archive to delete events from.
  -h, --help                         help for events
      --to string                    Remove events created before this timestamp.
```

## Options Inherited From Parent Commands
```
      --endpoint string   API endpoint of the ArangoDB Oasis (default "api.cloud.arangodb.com")
      --format string     Output format (table|json) (default "table")
      --token string      Token used to authenticate at ArangoDB Oasis
```

## See also
* [oasisctl delete auditlog archive](delete-auditlog-archive.md)	 - Delete an auditlog archive

