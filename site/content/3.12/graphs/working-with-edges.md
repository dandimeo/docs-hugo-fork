---
title: Edges
menuTitle: Working with Edges
weight: 105
description: >-
  Edges are documents with special `_from` and `_to` attributes to reference
  other documents by their ID to form a graph
archetype: default
---
{{< description >}}

Edges are used in [graphs](first-steps.md) to link vertices together.
They are documents with two additional system attributes that you set, a
`_from` and a `_to` attribute. They need to be stored in **edge collections**.

Together with vertex documents stored in _document collections_ (also called
_vertex collections_ in this context), edges form graphs.

**Example**

- An edge collection stores the information that a company's reception is a
  sub-unit to the services unit and the services unit is sub-unit to the
  CEO. You would express this relationship with the `_from` and `_to` attributes.
- A vertex collection stores all the properties about the reception, for example,
  that 20 people are working there, the room number, and so on.
- The `_from` attribute of an edge is the document ID of the linked vertex
  (incoming relation).
- The `_to` attribute of an edge is the document ID of the linked vertex
  (outgoing relation).

## Edges API

Edges are normal documents but additionally have a `_from` and a `_to` attribute
that you need to set.
You can update the `_from` and `_to` attributes like any other document attribute
using the [Documents API](../develop/javascript-api/@arangodb/collection-object.md#documents).

For methods you can call on edge collections, see the edge document methods of
the [_collection_ object](../develop/javascript-api/@arangodb/collection-object.md#edge-documents).
