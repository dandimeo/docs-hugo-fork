---
title: Faceted Search with ArangoSearch
menuTitle: Faceted search
weight: 45
description: >-
  Combine aggregation with search queries to retrieve how often values occur
  overall
archetype: default
---
{{< description >}}

A popular method for filtering items in an online shop is to display product
categories in a list, together with the number of items in each category.
This way, users get an idea of how many items will be left after applying a
certain filter before they actually enable it. This concept can be extended to
any properties, also called facets.

To implement such a feature in ArangoDB, you can use a `COLLECT` operation
to group and count how many documents share an attribute value. This is also
possible with `arangosearch` and `search-alias` Views by simply iterating over a
View instead of a collection.

## Dataset

[IMDB movie dataset](example-datasets.md#imdb-movie-dataset)

## View definition

### `search-alias` View

```js
db.imdb_vertices.ensureIndex({ name: "inv-exact", type: "inverted", fields: [ "title" ] });
db._createView("imdb", "search-alias", { indexes: [ { collection: "imdb_vertices", index: "inv-exact" } ] });
```

### `arangosearch` View

```json
{
  "links": {
    "imdb_vertices": {
      "fields": {
        "genre": {
          "analyzers": [
            "identity"
          ]
        }
      }
    }
  }
}
```

## AQL queries

Find out all genre values by grouping by the `genre` attribute and count the
number of occurrences:

```aql
FOR doc IN imdb
  COLLECT genre = doc.genre WITH COUNT INTO count
  RETURN { genre, count }
```

| genre     | count |
|:----------|------:|
| null      | 51287 |
| Action    |  2449 |
| Adventure |   312 |
| Animation |   426 |
| British   |     1 |
| Comedy    |  3188 |
| …         |     … |

The `COLLECT` operation is applied as a post-operation. In other words, it is
not accelerated by the View index. On the other hand, the `genre` field does
not need to be indexed for this query.

To look up a specific genre, the field needs to be indexed. The lookup itself
utilizes the View index, but the `COLLECT` operation is still a post-operation:

```aql
FOR doc IN imdb
  SEARCH doc.genre == "Action"
  COLLECT WITH COUNT INTO count
  RETURN count
```

```json
[
  2690
]
```

For above query with a single, simple condition, there is an optimization you
can enable that can accurately determine the count from index data faster than
the standard `COLLECT`. Also see
[Count Approximation](performance.md#count-approximation).

```aql
FOR doc IN imdb
  SEARCH doc.genre == "Action"
  OPTIONS { countApproximate: "cost" }
  COLLECT WITH COUNT INTO count
  RETURN count
```

To apply this optimization to the faceted search paradigm over all genres, you
can run and **cache** the following query that determines all unique genre
values:

```aql
FOR doc IN imdb
  RETURN DISTINCT doc.genre
```

You may use the AQL query cache if the data does not change much, or you could
execute above query periodically and store the result in a separate collection.
The numbers will not be fully up to date in that case, but it is often an
acceptable tradeoff for a faceted search.

You can then use the genre list to look up each genre and retrieve the count
while utilizing the count approximation optimization:

```aql
LET genres = [ "Action", "Adventure", "Animation", /* ... */ ]
FOR genre IN genres
  LET count = FIRST(
    FOR doc IN imdb
      SEARCH doc.genre == genre
      OPTIONS { countApproximate: "cost" }
      COLLECT WITH COUNT INTO count
      RETURN count
  )
  RETURN { genre, count }
```
