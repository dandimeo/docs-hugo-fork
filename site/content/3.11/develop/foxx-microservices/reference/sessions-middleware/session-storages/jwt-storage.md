---
title: JWT Session Storage
menuTitle: JWT storage
weight: 10
description: >-
  JWT Session Storage
archetype: default
---
`const jwtStorage = require('@arangodb/foxx/sessions/storages/jwt');`

The JWT session storage converts sessions to and from
[JSON Web Tokens](https://jwt.io/).

**Examples**

```js
// Pass in a secure secret from the Foxx configuration
const secret = module.context.configuration.jwtSecret;
const sessions = sessionsMiddleware({
  storage: jwtStorage(secret),
  transport: 'header'
});
module.context.use(sessions);
```

## Creating a storage

`jwtStorage(options): Storage`

Creates a [Storage](_index.md) that can be used in the sessions middleware.

**Note:** while the "none" algorithm (i.e. no signature) is supported this
dummy algorithm provides no security and allows clients to make arbitrary
modifications to the payload and should not be used unless you are certain
you specifically need it.

**Arguments**

- **options**: `Object`

  An object with the following properties:

  - **algorithm**: `string` (Default: `"HS512"`)

    The algorithm to use for signing the token.

    Supported values:

    - `"HS256"` (HMAC-SHA256)
    - `"HS384"` (HMAC-SHA384)
    - `"HS512"` (HMAC-SHA512)
    - `"none"` (no signature)

  - **secret**: `string`

    The secret to use for signing the token.

    This field is forbidden when using the "none" algorithm but required otherwise.

  - **ttl**: `number` (Default: `3600`)

    The maximum lifetime of the token in seconds. You may want to keep this
    short as a new token is generated on every request allowing clients to
    refresh tokens automatically.

  - **verify**: `boolean` (Default: `true`)

    If set to `false`, the signature is not verified but still generated
    (unless using the "none" algorithm).

  - **maxExp**: `number` (Default: `Infinity`)

    Largest value that is accepted in an incoming JWT `exp` (expiration) field.

If a string is passed instead of an options object, it is interpreted
as the `secret` option.
