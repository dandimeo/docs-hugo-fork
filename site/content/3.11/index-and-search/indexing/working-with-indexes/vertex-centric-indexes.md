---
title: Vertex-Centric Indexes
menuTitle: Vertex-Centric Indexes
weight: 35
description: >-
  In ArangoDB there are special indexes designed to speed up graph operations, especially if the graph contains supernodes (vertices that have an exceptionally high amount of connected edges)
archetype: default
---
## Introduction to Vertex-Centric Indexes

In ArangoDB there are special indexes designed to speed up graph operations,
especially if the graph contains supernodes (vertices that have an exceptionally
high amount of connected edges).
These indexes are called vertex-centric indexes and can be used in addition
to the existing edge index.

## Motivation

The idea of this index is to index a combination of a vertex, the direction and any arbitrary
set of other attributes on the edges.
To take an example, if we have an attribute called `type` on the edges, we can use an outbound
vertex-centric index on this attribute to find all edges attached to a vertex with a given `type`.
The following query example could benefit from such an index:

```aql
FOR v, e, p IN 3..5 OUTBOUND @start GRAPH @graphName
  FILTER p.edges[*].type ALL == "friend"
  RETURN v
```

Using the built-in edge-index ArangoDB can find the list of all edges attached to the vertex fast,
but still it has to walk through this list and check if all of them have the attribute `type == "friend"`.
Using a vertex-centric index would allow ArangoDB to find all edges for the vertex having the attribute `type == "friend"`
in the same time and can save the iteration to verify the condition.

## Index creation

A vertex-centric has to be of the type [Persistent Index](persistent-indexes.md)
and is created using its normal creation operations. However, in the list of
fields used to create the index we have to include either `_from` or `_to`.

Let us again explain this by an example.
Assume we want to create an hash-based outbound vertex-centric index on the attribute `type`.
This can be created with the following way:

```js
---
name: ensureVertexCentricIndex
description: ''
---
~db._createEdgeCollection("collection");
db.collection.ensureIndex({ type: "persistent", fields: [ "_from", "type" ] })
~db._drop("collection");
```

All options that are supported by persistent indexes are supported by the
vertex-centric index as well.

## Index usage

The AQL optimizer can decide to use a vertex-centric whenever suitable, however it is not guaranteed that this
index is used, the optimizer may estimate that an other index is assumed to be better.
The optimizer will consider this type of indexes on explicit filtering of `_from` respectively `_to`:

```aql
FOR edge IN collection
  FILTER edge._from == "vertices/123456" AND edge.type == "friend"
  RETURN edge
```

and during pattern matching queries:

```aql
FOR v, e, p IN 3..5 OUTBOUND @start GRAPH @graphName
  FILTER p.edges[*].type ALL == "friend"
  RETURN v
```
