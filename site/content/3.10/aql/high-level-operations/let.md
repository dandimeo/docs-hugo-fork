---
title: LET operation in AQL
menuTitle: LET
weight: 35
description: >-
  You can use the `LET` operation to assign an arbitrary value to a variable
archetype: default
---
{{< description >}}

The variable is introduced in the scope the `LET` statement is placed in.
You cannot change the value once assigned.

## Syntax

<pre><code>LET <em>variableName</em> = <em>expression</em></code></pre>

*expression* can be a simple expression or a subquery.

For allowed variable names [AQL Syntax](../fundamentals/syntax.md#names).

## Usage

Variables are immutable in AQL, which means they cannot be re-assigned:

```aql
LET a = [1, 2, 3]  // initial assignment

a = PUSH(a, 4)     // syntax error, unexpected identifier
LET a = PUSH(a, 4) // parsing error, variable 'a' is assigned multiple times
LET b = PUSH(a, 4) // allowed, result: [1, 2, 3, 4]
```

`LET` statements are mostly used to declare complex computations and to avoid
repeated computations of the same value at multiple parts of a query.

```aql
FOR u IN users
  LET numRecommendations = LENGTH(u.recommendations)
  RETURN {
    "user" : u,
    "numRecommendations" : numRecommendations,
    "isPowerUser" : numRecommendations >= 10
  }
```

In the above example, the computation of the number of recommendations is
factored out using a `LET` statement, thus avoiding computing the value twice in
the `RETURN` statement.

Another use case for `LET` is to declare a complex computation in a subquery,
making the whole query more readable.

```aql
FOR u IN users
  LET friends = (
  FOR f IN friends 
    FILTER u.id == f.userId
    RETURN f
  )
  LET memberships = (
  FOR m IN memberships
    FILTER u.id == m.userId
      RETURN m
  )
  RETURN { 
    "user" : u, 
    "friends" : friends, 
    "numFriends" : LENGTH(friends), 
    "memberShips" : memberships 
  }
```
