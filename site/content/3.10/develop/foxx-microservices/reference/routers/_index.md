---
title: Routers
menuTitle: Routers
weight: 20
description: >-
  Routers
archetype: chapter
---
`const createRouter = require('@arangodb/foxx/router');`

Routers let you define routes that extend ArangoDB's HTTP API with custom endpoints.

Routers need to be mounted using the `use` method of a
[service context](../service-context.md) to expose their HTTP routes at a service's mount path.

You can pass routers between services mounted in the same database
[as dependencies](../../guides/linking-services-together.md). You can even nest routers
within each other.

## Creating a router

`createRouter(): Router`

This returns a new, clean router object that has not yet been mounted in the
service and can be exported like any other object.

## Request handlers

`router.get([path], [...middleware], handler, [name]): Endpoint`

`router.post([path], [...middleware], handler, [name]): Endpoint`

`router.put([path], [...middleware], handler, [name]): Endpoint`

`router.patch([path], [...middleware], handler, [name]): Endpoint`

`router.delete([path], [...middleware], handler, [name]): Endpoint`

`router.all([path], [...middleware], handler, [name]): Endpoint`

These methods let you specify routes on the router.
The `all` method defines a route that matches any supported HTTP verb. The
other methods define routes that only match the HTTP verb with the same name.

**Arguments**

- **path**: `string` (Default: `"/"`)

  The path of the request handler relative to the base path the Router is mounted at.
  If omitted, the request handler handles requests to the base path of the Router.
  For information on defining dynamic routes see the section on
  [path parameters in the chapter on router endpoints](endpoints.md#pathparam).

- **middleware**: `Function` (optional)

  Zero or more middleware functions that take the following arguments:

  - **req**: `Request`

    An incoming server request object.

  - **res**: `Response`

    An outgoing server response object.

  - **next**: `Function`

    A callback that passes control over to the next middleware function
    and returns when that function has completed.

    If a truthy argument is passed, that argument is thrown as an error.

    If there is no next middleware function, the `handler` is
    invoked instead (see below).

- **handler**: `Function`

  A function that takes the following arguments:

  - **req**: `Request`

    An incoming server request object.

  - **res**: `Response`

    An outgoing server response.

- **name**: `string` (optional)

  A name that can be used to generate URLs for the endpoint.
  For more information see the `reverse` method of the [request object](request.md).

Returns an [Endpoint](endpoints.md) for the route.

**Examples**

Simple index route:

```js
router.get(function (req, res) {
  res.set('content-type', 'text/plain');
  res.write('Hello World!');
});
```

Restricting access to authenticated ArangoDB users:

```js
router.get('/secrets', function (req, res, next) {
  if (req.arangoUser) {
    next();
  } else {
    res.throw(404, 'Secrets? What secrets?');
  }
}, function (req, res) {
  res.download('allOurSecrets.zip');
});
```

Multiple middleware functions:

```js
function counting (req, res, next) {
  if (!req.counter) req.counter = 0;
  req.counter++;
  next();
  req.counter--;
}
router.get(counting, counting, counting, function (req, res) {
  res.json({counter: req.counter}); // {"counter": 3}
});
```

## Mounting child routers and middleware

`router.use([path], middleware, [name]): Endpoint`

The `use` method lets you mount a child router or middleware at a given path.

**Arguments**

- **path**: `string` (optional)

  The path of the middleware relative to the base path the Router is mounted at.
  If omitted, the middleware handles requests to the base path of the Router.
  For information on defining dynamic routes see the section on
  [path parameters in the chapter on router endpoints](endpoints.md#pathparam).

- **middleware**: `Router | Middleware`

  An unmounted router object or a [middleware](middleware.md).

- **name**: `string` (optional)

  A name that can be used to generate URLs for endpoints of this router.
  For more information see the `reverse` method of the [request object](request.md).
  Has no effect if `handler` is a Middleware.

Returns an [Endpoint](endpoints.md) for the middleware or child router.

{{< warning >}}
When mounting child routers at multiple paths, effects of methods
invoked on each endpoint only affect routes of that endpoint.
{{< /warning >}}

**Examples**

```js
const child = createRouter();

router.use("/a", child);

router.use("/b", child)
.queryParam("number", joi.number().required(), "Required number parameter.");

child.get(function (req, res) {
  // The query parameter "number" is required if the request was made via "/b"
  // but optional if the request was made via "/a", we can't rely on it.
  res.json({number: req.queryParams.number});
});
```

```js
const child = createRouter()
.queryParam("number", joi.number().required(), "Required number parameter.");

router.use("/a", child);

router.use("/b", child);

child.get(function (req, res) {
  // The query parameter "number" is always required, regardless if the
  // request was made via "/a" or "/b".
  res.json({number: req.queryParams.number});
});
```

## Additional metadata

In addition to the router-specific methods, all methods available on
[endpoints](endpoints.md) are also available on
router objects and can be used to define shared defaults.

**Examples**

```js
router.header(
  'x-common-header',
  joi.string().required(),
  'Common header shared by all routes on this router.'
);

router.get('/', function (req, res) {
  // This route requires the header to be set but we don't need to explicitly
  // specify that.
  const value = req.header('x-common-header');
  // ...
});
```
