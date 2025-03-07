---
title: Fulltext functions in AQL
menuTitle: Fulltext
weight: 30
description: >-
  AQL offers the following functions to filter data based on fulltext indexes
archetype: default
---
AQL offers the following functions to filter data based on
[fulltext indexes](../../index-and-search/indexing/working-with-indexes/fulltext-indexes.md).

{{< warning >}}
The fulltext index type is deprecated from version 3.10 onwards.
It is recommended to use [Inverted indexes](../../index-and-search/indexing/working-with-indexes/inverted-indexes.md) or
[ArangoSearch](../../index-and-search/arangosearch/_index.md) for advanced full-text search capabilities.
{{< /warning >}}

## FULLTEXT()

`FULLTEXT(coll, attribute, query, limit) → docArray`

Return all documents from collection *coll*, for which the attribute *attribute*
matches the fulltext search phrase *query*, optionally capped to *limit* results.

**Note**: the *FULLTEXT()* function requires the collection *coll* to have a
fulltext index on *attribute*. If no fulltext index is available, this function
will fail with an error at runtime. It doesn't fail when explaining the query however.

- **coll** (collection): a collection
- **attribute** (string): the attribute name of the attribute to search in
- **query** (string): a fulltext search expression as described below
- **limit** (number, *optional*): if set to a non-zero value, it will cap the result
  to at most this number of documents
- returns **docArray** (array): an array of documents

*FULLTEXT()* is not meant to be used as an argument to `FILTER`,
but rather to be used as the expression of a `FOR` statement:

```aql
FOR oneMail IN FULLTEXT(emails, "body", "banana,-apple")
    RETURN oneMail._id
```

*query* is a comma-separated list of sought words (or prefixes of sought words). To
distinguish between prefix searches and complete-match searches, each word can optionally be
prefixed with either the `prefix:` or `complete:` qualifier. Different qualifiers can
be mixed in the same query. Not specifying a qualifier for a search word will implicitly
execute a complete-match search for the given word:

- `FULLTEXT(emails, "body", "banana")`\
  Will look for the word *banana* in the
  attribute *body* of the collection *collection*.

- `FULLTEXT(emails, "body", "banana,orange")`\
  Will look for both words
  *banana* and *orange* in the mentioned attribute. Only those documents will be
  returned that contain both words.

- `FULLTEXT(emails, "body", "prefix:head")`\
  Will look for documents that contain any
  words starting with the prefix *head*.

- `FULLTEXT(emails, "body", "prefix:head,complete:aspirin")`\
  Will look for all
  documents that contain a word starting with the prefix *head* and that also contain
  the (complete) word *aspirin*. Note: specifying `complete:` is optional here.

- `FULLTEXT(emails, "body", "prefix:cent,prefix:subst")`\
  Will look for all documents
  that contain a word starting with the prefix *cent* and that also contain a word
  starting with the prefix *subst*.

If multiple search words (or prefixes) are given, then by default the results will be
AND-combined, meaning only the logical intersection of all searches will be returned.
It is also possible to combine partial results with a logical OR, and with a logical NOT:

- `FULLTEXT(emails, "body", "+this,+text,+document")`\
  Will return all documents that
  contain all the mentioned words. Note: specifying the `+` symbols is optional here.

- `FULLTEXT(emails, "body", "banana,|apple")`\
  Will return all documents that contain
  either (or both) words *banana* or *apple*.

- `FULLTEXT(emails, "body", "banana,-apple")`\
  Will return all documents that contain
  the word *banana*, but do not contain the word *apple*.

- `FULLTEXT(emails, "body", "banana,pear,-cranberry")`\
  Will return all documents that
  contain both the words *banana* and *pear*, but do not contain the word
  *cranberry*.

No precedence of logical operators will be honored in a fulltext query. The query will simply
be evaluated from left to right.
