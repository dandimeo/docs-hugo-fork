---
title: _arangosh_ Details
menuTitle: Details
weight: 10
description: >-
  You can paste multiple lines into arangosh, given the first line ends with an opening brace
archetype: default
---
## Interaction

You can paste multiple lines into _arangosh_, given the first line ends with an
opening brace:

```js
---
name: shellPaste
description: ''
---
for (var i = 0; i < 10; i ++) {
  require("@arangodb").print("Hello world " + i + "!\n");
}
```

To load your own JavaScript code into the current JavaScript interpreter context,
use the load command:

```js
require("internal").load("/tmp/test.js")     // <- Linux / macOS
require("internal").load("c:\\tmp\\test.js") // <- Windows
```

You can exit arangosh using the key combination `<CTRL> + D` or by
typing `quit<ENTER>`.

## Shell Output

The ArangoDB shell prints the output of the last evaluated expression
by default:

```js
---
name: lastExpressionResult
description: ''
---
42 * 23
```

In order to prevent printing the result of the last evaluated expression,
the expression result can be captured in a variable, e.g.

```js
---
name: lastExpressionResultCaptured
description: ''
---
var calculationResult = 42 * 23
```

There is also the `print` function to explicitly print out values in the
ArangoDB shell:

```js
---
name: printFunction
description: ''
---
print({ a: "123", b: [1,2,3], c: "test" });
```

By default, the ArangoDB shell uses a pretty printer when JSON documents are
printed. This ensures documents are printed in a human-readable way:

```js
---
name: usingToArray
description: ''
---
db._create("five")
for (i = 0; i < 5; i++) {
  db.five.save({value:i});
}
db.five.toArray()
~db._drop("five");
```

While the pretty-printer produces nice looking results, it needs a lot of
screen space for each document. Sometimes a more dense output might be better.
In this case, the pretty printer can be turned off using the command
`stop_pretty_print()`.

To turn on pretty printing again, use the `start_pretty_print()` command.

## Escaping

In AQL, escaping is done traditionally with the backslash character: `\`.
As seen above, this leads to double backslashes when specifying Windows paths.
_arangosh_ requires another level of escaping, also with the backslash character.
It adds up to four backslashes that need to be written in _arangosh_ for a single
literal backslash (`c:\tmp\test.js`):

```js
db._query('RETURN "c:\\\\tmp\\\\test.js"')
```

You can use [bind variables](../../../aql/how-to-invoke-aql/with-arangosh.md) to
mitigate this:

```js
var somepath = "c:\\tmp\\test.js"
db._query(aql`RETURN ${somepath}`)
```

## Database Wrappers

_arangosh_ provides the `db` object by default, and this object can
be used for switching to a different database and managing collections inside the
current database.

For a list of available methods for the `db` object, type

```js
---
name: shellHelp
description: ''
---
db._help(); 
```

The [`db` object](../../../develop/javascript-api/@arangodb/db-object.md) is available in _arangosh_
as well as on _arangod_ i.e. if you're using [Foxx](../../../develop/foxx-microservices/_index.md). While its
interface is persistent between the _arangosh_ and the _arangod_ implementations,
its underpinning is not. The _arangod_ implementation are JavaScript wrappers
around ArangoDB's native C++ implementation, whereas the _arangosh_ implementation
wraps HTTP accesses to ArangoDB's [RESTful API](../../../develop/http/_index.md).

So while this code may produce similar results when executed in _arangosh_ and
_arangod_, the CPU usage and time required differs since the
_arangosh_ version performs around 100k HTTP requests, and the
_arangod_ version directly writes to the database:

```js
for (i = 0; i < 100000; i++) {
    db.test.save({ name: { first: "Jan" }, count: i});
}
```

## Using `arangosh` via Unix shebang mechanisms
In Unix operating systems, you can start scripts by specifying the interpreter in the first line of the script.
This is commonly called `shebang` or `hash bang`. You can also do that with `arangosh`, i.e. create `~/test.js`:

```sh
#!/usr/bin/arangosh --javascript.execute 
require("internal").print("hello world")
db._query("FOR x IN test RETURN x").toArray()
```

Note that the first line has to end with a blank in order to make it work.
Mark it executable to the OS: 

```sh
> chmod a+x ~/test.js
```

and finally try it out:

```sh
> ~/test.js
```

## Shell Configuration

_arangosh_ looks for a user-defined startup script named `.arangosh.rc` in the
user's home directory on startup. The home directory is likely at `/home/<username>/`
on Unix/Linux, and is determined on Windows by peeking into the environment variables
`%HOMEDRIVE%` and `%HOMEPATH%`. 

If the file `.arangosh.rc` is present in the home directory, _arangosh_ executes
the contents of this file inside the global scope.

You can use this to define your own extra variables and functions that you need often.
For example, you could put the following into the `.arangosh.rc` file in your home
directory:

```js
// "var" keyword avoided intentionally...
// otherwise "timed" would not survive the scope of this script
global.timed = function (cb) {
  console.time("callback");
  cb();
  console.timeEnd("callback");
};
```

This makes a function named `timed` available in _arangosh_ in the global scope.

You can now start _arangosh_ and invoke the function like this:

```js
timed(function () { 
  for (var i = 0; i < 1000; ++i) {
    db.test.save({ value: i }); 
  }
});
```

Please keep in mind that, if present, the `.arangosh.rc` file needs to contain valid
JavaScript code. If you want any variables in the global scope to survive you need to
omit the `var` keyword for them. Otherwise, the variables are only visible inside
the script itself, but not outside.
