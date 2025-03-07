---
title: Managing Users in the ArangoDB Shell
menuTitle: In arangosh
weight: 5
description: >-
  The module @arangodb/users exposes a JavaScript API to manage user accounts.
archetype: default
---
Connect with `arangosh` to the server or a Coordinator respectively.
The module `@arangodb/users` exposes a JavaScript API to manage user accounts.

Please note, that for backward compatibility the server access levels
follow from the database access level on the database *_system*.

Also note that the server and database access levels are represented as

- `rw`: for *Administrate*
- `ro`: for *Access*
- `none`: for *No access*

This is again for backward compatibility.

**Example**

Fire up *arangosh* and require the users module. Use it to create a new user:

```js
arangosh --server.endpoint tcp://127.0.0.1:8529 ...
...
> const users = require('@arangodb/users');
> users.save('JohnSmith', 'mypassword');
```

It creates a user called *JohnSmith* with *mypassword* as password. This user
will have no access at all.

Note that running the command like this may store the password literally in
ArangoShell's history. To avoid that, either disable the history
(`--console.history false`) or use a dynamically created password, e.g.:

```js
> passwd = require('internal').genRandomAlphaNumbers(20);
> users.save('JohnSmith', passwd);
```

The above will print the password on screen (so you can memorize it) but will
not store it in the command history.

While there, you probably want to change the password of the default `root`
user too. Otherwise one will be able to connect with the default `root` user
and its empty password. The following commands change the `root` user's password:

```js
> passwd = require('internal').genRandomAlphaNumbers(20);
> require('@arangodb/users').update('root', passwd);
```

Back to our user account *JohnSmith*. Let us create a new database
and grant him access to it with `grantDatabase()`:

```js
> db._createDatabase('testdb');
> users.grantDatabase('JohnSmith', 'testdb', 'rw');
```

This grants the user *Administrate* access to the database
*testdb*. `revokeDatabase()` will revoke this access level setting.

**Note**: Be aware that from 3.2 onwards the `grantDatabase()` will not
automatically grant users the access level to write or read collections in a
database. If you grant access to a database `testdb` you will
additionally need to explicitly grant access levels to individual
collections via `grantCollection()`.

The upgrade procedure from 3.1 to 3.2 sets the wildcard database access
level for all users to *Administrate* and sets the wildcard collection
access level for all user/database pairs to *Read/Write*.

Before we can grant *JohnSmith* access to a collection, we first have to
connect to the new database and create a collection. Disconnect `arangosh`
by pressing Ctrl+C twice. Then reconnect, but to the database we created:

```js
arangosh --server.endpoint tcp://127.0.0.1:8529 --server.database testdb ...
...
> db._create('testcoll');
> const users = require('@arangodb/users');
> users.grantCollection('JohnSmith', 'testdb', 'testcoll', 'rw');
```

It is not necessary to reconnect to the `_system` database in order to grant
access to the collection.

To confirm that the authentication works as expected, try to connect to
different databases as *JohnSmith*:

```
arangosh --server.endpoint tcp://127.0.0.1:8529 --server.username JohnSmith --server.database testdb ...
...
> Connected to ArangoDB 'http+tcp://127.0.0.1:8529, version: 3.5.2 [SINGLE, server], database: 'testdb', username: 'JohnSmith'
```

```
arangosh --server.endpoint tcp://127.0.0.1:8529 --server.username JohnSmith --server.database _system ...
...
> Could not connect to endpoint 'tcp://127.0.0.1:8529', database: '_system', username: 'JohnSmith'
> Error message: 'not authorized to execute this request'
```

You can also use curl to check that you are actually getting HTTP 401
(Unauthorized) server responses for requests that require authentication:

```
curl --dump - http://127.0.0.1:8529/_api/version
```

## Save

`users.save(user, passwd, active, extra)`

This will create a new ArangoDB user. The user name must be specified in *user*
and must not be empty. Note that usernames *must* not start with `:role:`
(reserved for LDAP authentication).

The password must be given as a string, too, but can be left empty if
required. If you pass the special value *ARANGODB_DEFAULT_ROOT_PASSWORD*, the
password will be set the value stored in the environment variable
`ARANGODB_DEFAULT_ROOT_PASSWORD`. This can be used to pass an instance
variable into ArangoDB. For example, the instance identifier from Amazon.

If the *active* attribute is not specified, it defaults to *true*. The *extra*
attribute can be used to save custom data with the user.

This method will fail if either the user name or the passwords are not
specified or given in a wrong format, or there already exists a user with the
specified name.

**Note**: The user will not have permission to access any database. You need
to grant the access rights for one or more databases using
[grantDatabase](#grant-database).

*Examples*

```js
---
name: USER_02_saveUser
description: ''
---
require('@arangodb/users').save('my-user', 'my-secret-password');
```

## Grant Database

`users.grantDatabase(user, database, type)`

This grants *type* ('rw', 'ro' or 'none') access to the *database* for
the *user*. If *database* is `"*"`, this sets the wildcard database access
level for the user *user*.

The server access level follows from the access level for the database
`_system`.

## Revoke Database

`users.revokeDatabase(user, database)`

This clears the access level setting to the *database* for the *user* and
the wildcard database access setting for this user kicks in. In case no wildcard
access was defined the default is *No Access*. This will also
clear the access levels for all the collections in this database.

## Grant Collection

`users.grantCollection(user, database, collection, type)`

This grants *type* ('rw', 'ro' or 'none') access level to the *collection*
in *database* for the *user*. If *collection* is `"*"` this sets the
wildcard collection access level for the user *user* in database
*database*.

## Revoke Collection

`users.revokeCollection(user, database)`

This clears the access level setting to the collection *collection* for the
user *user*. The system will either fallback to the wildcard collection access
level or default to *No Access*

## Replace

`users.replace(user, passwd, active, extra)`

This will look up an existing ArangoDB user and replace its user data.

The username must be specified in *user*, and a user with the specified name
must already exist in the database.

The password must be given as a string, too, but can be left empty if required.

If the *active* attribute is not specified, it defaults to *true*.  The
*extra* attribute can be used to save custom data with the user.

This method will fail if either the user name or the passwords are not specified
or given in a wrong format, or if the specified user cannot be found in the
database.

**Note**: this function will not work from within the web interface

*Examples*

```js
---
name: USER_03_replaceUser
description: ''
---
require("@arangodb/users").replace("my-user", "my-changed-password");
```

## Update

`users.update(user, passwd, active, extra)`

This will update an existing ArangoDB user with a new password and other data.

The user name must be specified in *user* and the user must already exist in
the database.

The password must be given as a string, too, but can be left empty if required.

If the *active* attribute is not specified, the current value saved for the
user will not be changed. The same is true for the *extra* attribute.

This method will fail if either the user name or the passwords are not specified
or given in a wrong format, or if the specified user cannot be found in the
database.

*Examples*

```js
---
name: USER_04_updateUser
description: ''
---
require("@arangodb/users").update("my-user", "my-secret-password");
```

## isValid

`users.isValid(user, password)`

Checks whether the given combination of user name and password is valid. The
function will return a boolean value if the combination of user name and password
is valid.

Each call to this function is penalized by the server sleeping a random
amount of time.

*Examples*

```js
---
name: USER_05_isValidUser
description: ''
---
require("@arangodb/users").isValid("my-user", "my-secret-password");
```

## Remove

`users.remove(user)`

Removes an existing ArangoDB user from the database.

The user name must be specified in *User* and the specified user must exist in
the database.

This method will fail if the user cannot be found in the database.

*Examples*

```js
---
name: USER_07_removeUser
description: ''
---
require("@arangodb/users").remove("my-user");
~require('@arangodb/users').save('my-user', 'my-secret-password');
```

## Document

`users.document(user)`

Fetches an existing ArangoDB user from the database.

The user name must be specified in *user*.

This method will fail if the user cannot be found in the database.

*Examples*

```js
---
name: USER_04_documentUser
description: ''
---
require("@arangodb/users").document("my-user");
```

## All

`users.all()`

Fetches all existing ArangoDB users from the database.

*Examples*

```js
---
name: USER_06_AllUsers
description: ''
---
require("@arangodb/users").all();
```

## Reload

`users.reload()`

Reloads the user authentication data on the server

All user authentication data is loaded by the server once on startup only and is
cached after that. When users get added or deleted, a cache flush is done
automatically, and this can be performed by a call to this method.

*Examples*

```js
---
name: USER_03_reloadUser
description: ''
---
require("@arangodb/users").reload();
```

## Permission

`users.permission(user, database[, collection])`

Fetches the access level to the database or a collection.

The user and database name must be specified, optionally you can specify
the collection name.

This method will fail if the user cannot be found in the database.

*Examples*

```js
---
name: USER_05_permission
description: ''
---
~require("@arangodb/users").grantDatabase("my-user", "testdb");
require("@arangodb/users").permission("my-user", "testdb");
```
