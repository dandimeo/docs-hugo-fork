---
title: Views
menuTitle: Views
weight: 20
description: >-
  Views can index documents of multiple collections and enable sophisticated
  information retrieval possibilities, like full-text search with ranking by
  relevance
archetype: default
---
{{< description >}}

Views allows you to perform complex searches at high performance. They are
accelerated by inverted indexes that are updated near real-time.

A View is conceptually a transformation function over documents from zero or
more collections. The transformation depends on the View type and the View
configuration.

Views are powered by ArangoDB's built-in search engine.
See [ArangoSearch](../../index-and-search/arangosearch/_index.md) for details.

## View types

Available View types:

- The traditional [`arangosearch` Views](../../index-and-search/arangosearch/arangosearch-views-reference.md) to which
  you link collections to.
- The modern [`search-alias` Views](../../index-and-search/arangosearch/search-alias-views-reference.md)
  that can reference inverted indexes that are defined on the collection-level.

You need to specify the type when you create the View.
The type cannot be changed later.

## View names

You can give each View a name to identify and access it. The name needs to
be unique within a [database](databases.md), but not globally
for the entire ArangoDB instance.

The namespace for Views is shared with [collections](collections.md).
There cannot exist a View and a collection with the same name in the same database.

The View name needs to be a string that adheres to either the **traditional**
or the **extended** naming constraints. Whether the former or the latter is
active depends on the `--database.extended-names` startup option.
The extended naming constraints are used if enabled, allowing many special and
UTF-8 characters in database, collection, View, and index names. If set to
`false` (default), the traditional naming constraints are enforced.

{{< info >}}
The extended naming constraints are an **experimental** feature but they will
become the norm in a future version. Check if your drivers and client applications
are prepared for this feature before enabling it.
{{< /info >}}

The restrictions for View names are as follows:

- For the **traditional** naming constraints:
  - The names must only consist of the letters `A` to `Z` (both in lower 
    and upper case), the digits `0` to `9`, and underscore (`_`) and dash   (`-`)
    characters. This also means that any non-ASCII names are not allowed.
  - View names must start with a letter.
  - The maximum allowed length of a name is 256 bytes.
  - View names are case-sensitive.

- For the **extended** naming constraints:
  - Names can consist of most UTF-8 characters, such as Japanese or Arabic
    letters, emojis, letters with accentuation. Some ASCII characters are
    disallowed, but less compared to the  _traditional_ naming constraints.
  - Names cannot contain the characters `/` or `:` at any position, nor any
    control characters (below ASCII code 32), such as `\n`, `\t`, `\r`, and `\0`.
  - Spaces are accepted, but only in between characters of the name. Leading
    or trailing spaces are not allowed.
  - `.` (dot), `_` (underscore) and the numeric digits `0`-`9` are not allowed
    as first character, but at later positions.
  - View names are case-sensitive.
  - View names containing UTF-8 characters must be 
    [NFC-normalized](https://en.wikipedia.org/wiki/Unicode_equivalence#Normal_forms).
    Non-normalized names are rejected by the server.
  - The maximum length for a View name is 256 bytes after normalization. 
    As a UTF-8 character may consist of multiple bytes, this does not necessarily 
    equate to 256 characters.

  Example View names that can be used with the _extended_ naming constraints:
  `España`, `😀`, `犬`, `كلب`, `@abc123`, `København`, `München`, `Бишкек`, `abc? <> 123!`

{{< warning >}}
While it is possible to change the value of the
`--database.extended-names` option from `false` to `true` to enable
extended names, the reverse is not true. Once the extended names have been
enabled, they remain permanently enabled so that existing databases,
collections, Views, and indexes with extended names remain accessible.

Please be aware that dumps containing extended names cannot be restored
into older versions that only support the traditional naming constraints. In a
cluster setup, it is required to use the same naming constraints for all
Coordinators and DB-Servers of the cluster. Otherwise, the startup is
refused. In DC2DC setups, it is also required to use the same naming constraints
for both datacenters to avoid incompatibilities.
{{< /warning >}}

You can rename Views (except in cluster deployments). This changes the
View name, but not the View identifier.

## View identifiers

A View identifier lets you refer to a View in a database, similar to
the name. It is a string value and is unique within a database. Unlike
View names, ArangoDB assigns View IDs automatically and you have no
control over them.

ArangoDB internally uses View IDs to look up Views. However, you
should use the View name to access a View instead of its identifier.

ArangoDB uses 64-bit unsigned integer values to maintain View IDs
internally. When returning View IDs to clients, ArangoDB returns them as
strings to ensure the identifiers are not clipped or rounded by clients that do
not support big integers. Clients should treat the View IDs returned by
ArangoDB as opaque strings when they store or use them locally.

## Views API

The following descriptions cover the JavaScript interface for Views that
you can use to handle Views from the _arangosh_ command-line tool, as
well as in server-side JavaScript code like Foxx microservices.
For other languages see the corresponding language API.

The following examples show the basic usage of the View API.
For more details, see:

- [`arangosearch` Views](../../index-and-search/arangosearch/arangosearch-views-reference.md)
- [`search-alias` Views](../../index-and-search/arangosearch/search-alias-views-reference.md)
- [JavaScript API for Views](../../develop/javascript-api/@arangodb/db-object.md#views)
- [The _view_ object](../../develop/javascript-api/@arangodb/view-object.md)

### Create a View

Create a View with default properties:

```js
---
name: viewUsage_01
description: ''
---
~db._create("colA");
~db._create("colB");
view = db._createView("myView", "arangosearch", {});
~addIgnoreCollection("colA");
~addIgnoreCollection("colB");
~addIgnoreView("myView");
```

### Get a View

Get the View called `myView` by its name:

```js
---
name: viewUsage_02
description: ''
---
view = db._view("myView");
```

### Get the View properties

```js
---
name: viewUsage_03
description: ''
---
view.properties();
```

### Set a View property

```js
---
name: viewUsage_04
description: ''
---
view.properties({cleanupIntervalStep: 12});
```

### Add and remove links from a View

Link a collection to a View:

```js
---
name: viewUsage_05
description: ''
---
view.properties({links: {colA: {includeAllFields: true}}});
```

Add another link to the View:

```js
---
name: viewUsage_06
description: ''
---
view.properties({links: {colB: {fields: {text: {}}}}});
```

Remove the first link again:

```js
---
name: viewUsage_07
description: ''
---
view.properties({links: {colA: null}});
```

### Drop a View

```js
---
name: viewUsage_08
description: ''
---
~removeIgnoreCollection("colA");
~removeIgnoreCollection("colB");
~removeIgnoreView("myView");
db._dropView("myView");
~db._drop("colA");
~db._drop("colB");
```
