---
title: Durability
menuTitle: Durability
weight: 20
description: >-
  Transactions are executed until there is either a rollback or a commit
archetype: default
---
Transactions are executed until there is either a rollback
or a commit. On rollback, the operations from the transaction are reversed.

The RocksDB storage engine applies operations of a transaction in main memory
only until they are committed. In case of an a rollback the entire transaction
is just cleared, no extra rollback steps are required.

<!-- TODO: point out data loss (query accepted by server, but will be lost) -->
<!-- TODO: intermediate commits?! -->

In the event of a server crash, the storage engine scans the write-ahead log
to restore certain meta-data like the number of documents in collection 
or the selectivity estimates of secondary indexes.

<!-- TODO: obsolete?
There is thus the potential risk of losing data between the commit of the 
transaction and the actual (delayed) disk synchronization. This is the same as 
writing into collections that have the `waitForSync` property set to `false`
outside of a transaction.
In case of a crash with `waitForSync` set to false, the operations performed in
the transaction are either visible completely or not at all, depending on
whether the delayed synchronization had kicked in or not.

To ensure durability of transactions on a collection that have the `waitForSync`
property set to `false`, you can set the `waitForSync` attribute of the object
that is passed to `executeTransaction`. This forces a synchronization of the
transaction to disk even for collections that have `waitForSync` set to `false`:

```js
db._executeTransaction({
  collections: { 
    write: "users"
  },
  waitForSync: true,
  action: function () { ... }
});
```

An alternative is to perform an operation with an explicit `sync` request in
a transaction, e.g.

```js
db.users.save({ _key: "1234" }, true); 
```

In this case, the `true` value makes the whole transaction be synchronized
to disk at the commit.

In any case, ArangoDB gives users the choice of whether or not they want 
full durability for single collection transactions. Using the delayed synchronization
(i.e. `waitForSync` with a value of `false`) potentially increases throughput 
and performance of transactions, but introduces the risk of losing the last
committed transactions in the case of a crash.

The call to the `db._executeTransaction()` function 
only returns after the data of all modified collections has been synchronized
to disk and the transaction has been made fully durable. This not only reduces the
risk of losing data in case of a crash but also ensures consistency after a
restart.
-->
