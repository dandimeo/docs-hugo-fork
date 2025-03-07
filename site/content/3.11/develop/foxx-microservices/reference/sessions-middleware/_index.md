---
title: Session Middleware
menuTitle: Sessions middleware
weight: 25
description: >-
  Session Middleware
archetype: chapter
---
`const sessionMiddleware = require('@arangodb/foxx/sessions');`

The session middleware adds the `session` and `sessionStorage` properties to
the [request object](../routers/request.md) and deals with serializing and
deserializing the session as well as extracting session identifiers from
incoming requests and injecting them into outgoing responses.

**Examples**

```js
// Create a session middleware
const sessions = sessionsMiddleware({
  storage: module.context.collection('sessions'),
  transport: ['header', 'cookie']
});
// First enable the middleware for this service
module.context.use(sessions);
// Now mount the routers that use the session
const router = createRouter();
module.context.use(router);

router.get('/', function (req, res) {
  res.send(`Hello ${req.session.uid || 'anonymous'}!`);
}, 'hello');

router.post('/login', function (req, res) {
  req.session.uid = req.body;
  req.sessionStorage.save(req.session);
  res.redirect(req.reverse('hello'));
});
.body(['text/plain'], 'Username');
```

## Creating a session middleware

`sessionMiddleware(options): Middleware`

Creates a session middleware.

**Arguments**

- **options**: `Object`

  An object with the following properties:

  - **storage**: `Storage`

    Storage that is used to persist the sessions.

    The storage is also exposed as the `sessionStorage` on all request objects
    and as the `storage` property of the middleware.

    If a string or collection is passed instead of a Storage, it is used
    to create a [Collection Storage](session-storages/collection-storage.md).

  - **transport**: `Transport | Array<Transport>`

    Transport or array of transports that is used to extract the session
    identifiers from incoming requests and inject them into outgoing responses.
    When attempting to extract a session identifier, the transports are
    used in the order specified until a match is found. When injecting
    (or clearing) session identifiers, all transports are invoked.

    The transports are also exposed as the `transport` property of the middleware.

    If the string `"cookie"` is passed instead of a Transport, the
    [Cookie Transport](session-transports/cookie-transport.md) are used with the default
    settings instead.

    If the string `"header"` is passed instead of a Transport, the
    [Header Transport](session-transports/header-transport.md) are used with the default
    settings instead.

  - **autoCreate**: `boolean` (Default: `true`)

    If enabled the session storage's `new` method is invoked to create an
    empty session whenever the transport failed to return a session for the
    incoming request. Otherwise, the session is initialized as `null`.

Returns the session middleware.
