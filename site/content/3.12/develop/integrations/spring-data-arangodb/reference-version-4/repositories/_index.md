---
title: Repositories
menuTitle: Repositories
weight: 15
description: >-
  Interact with ArangoDB using the high-level and consistent abstraction
  provided by Spring Data repositories
archetype: chapter
---
{{< description >}}

Spring Data Commons provides a composable repository infrastructure which
Spring Data ArangoDB is built on. These allow for interface-based composition of
repositories consisting of provided default implementations for certain
interfaces (like `CrudRepository`) and custom implementations for other methods.

The base interface of Spring Data ArangoDB is `ArangoRepository`. It extends the
Spring Data interfaces `PagingAndSortingRepository` and `QueryByExampleExecutor`.
To get access to all Spring Data ArangoDB repository functionality simply create
your own interface extending `ArangoRepository<T, ID>`.

The type `T` represents your domain class and type `ID` the type of your field
annotated with `@Id` in your domain class. This field is persisted in ArangoDB
as document field `_key`.

**Examples**

```java
@Document
public class MyDomainClass {
  @Id
  private String id;

}

public interface MyRepository extends ArangoRepository<MyDomainClass, String> {

}
```

Instances of a Repository are created in Spring beans through the auto-wired mechanism of Spring.

```java
public class MySpringBean {

  @Autowired
  private MyRepository rep;

}
```
