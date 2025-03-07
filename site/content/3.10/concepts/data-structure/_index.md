---
title: Data Structure
menuTitle: Data Structure
weight: 60
description: >-
  ArangoDB stores graphs and documents as JSON objects that can be organized in
  collections and databases
archetype: chapter
---
{{< description >}}

## How Data is Structured in ArangoDB

The hierarchy that data is organized in is **documents** (data records) in
**collections**, and collections in **databases**.

### Documents

ArangoDB lets you store documents as JSON objects.

```json
{
  "name": "ArangoDB",
  "tags": ["graph", "database", "NoSQL"],
  "scalable": true,
  "company": {
    "name": "ArangoDB Inc.",
    "founded": 2015
  }
}
```

Each record that you store is a JSON object at the top-level, also referred to
as a [**document**](documents/_index.md).
Each key-value pair is called an **attribute**, comprised
of the attribute name and the attribute value. Attributes can also be called
*properties* or *fields*.

You can freely model your data using the available data types. Each document is
self-contained and can thus have a unique structure. You do not need to define a
schema upfront. However, sets of documents typically have some common
attributes. If you want to enforce a specific structure, then you can do so with a
[schema validation](documents/schema-validation.md).

Documents are internally stored in a binary format called
[VelocyPack](https://github.com/arangodb/velocypack#readme).

### Collections

Documents are stored in [**collections**](collections.md),
similar to how files are stored in folders. A collection can hold an arbitrary
number of documents.

You can group related documents together using collections, such as by
entity type. For example, you can store _book_ documents in a `books`
collections. All book records have some common attributes like a title,
author, and publisher. You can later create indexes for some of the often-used
attributes to speed up queries. This is done at the collection level.

### Databases

Each collection is part of a [**database**](databases.md).
Databases allow you to isolate sets of collections from one another, usually for
multi-tenant applications, where each of your clients has their own database to
work with. You cannot run queries across several databases.

Every server instance has a default database called  `_system`. It is special
because it cannot be removed and it holds a couple of system collections that
are used internally by the server. Other than that, you may create your own
collections in this database like in any other.