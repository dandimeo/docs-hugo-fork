---
title: AQL Documentation
menuTitle: AQL
weight: 70
description: >-
  The ArangoDB Query Language (AQL) can be used to retrieve and modify data that are stored in ArangoDB.
archetype: chapter
---
The ArangoDB Query Language (AQL) can be used to retrieve and modify data that 
are stored in ArangoDB.

The general workflow when executing a query is as follows:

- A client application ships an AQL query to the ArangoDB server. The query text
  contains everything ArangoDB needs to compile the result set
- ArangoDB will parse the query, execute it and compile the results. If the
  query is invalid or cannot be executed, the server will return an error that
  the client can process and react to. If the query can be executed
  successfully, the server will return the query results (if any) to the client

AQL is mainly a declarative language, meaning that a query expresses what result
should be achieved but not how it should be achieved. AQL aims to be
human-readable and therefore uses keywords from the English language. Another
design goal of AQL was client independency, meaning that the language and syntax
are the same for all clients, no matter what programming language the clients
may use.  Further design goals of AQL were the support of complex query patterns
and the different data models ArangoDB offers.

In its purpose, AQL is similar to the Structured Query Language (SQL). AQL supports 
reading and modifying collection data, but it doesn't support data-definition
operations such as creating and dropping databases, collections and indexes.
It is a pure data manipulation language (DML), not a data definition language
(DDL) or a data control language (DCL).

The syntax of AQL queries is different to SQL, even if some keywords overlap.
Nevertheless, AQL should be easy to understand for anyone with an SQL background.

For example queries, see the [Data Queries](data-queries.md) and
[Examples & Query Patterns](examples-and-query-patterns/_index.md) chapters.
