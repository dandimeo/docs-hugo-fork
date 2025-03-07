---
title: Endpoints
menuTitle: Endpoints
weight: 5
description: >-
  Endpoints
archetype: default
---
Endpoints are returned by the `use`, `all`
and HTTP verb (e.g. `get`, `post`) methods of [routers](_index.md)
as well as the `use` method of the [service context](../service-context.md).
They can be used to attach metadata to mounted routes, middleware and
child routers that affects how requests and responses are processed or
provides API documentation.

Endpoints should only be used to invoke the following methods.
Endpoint methods can be chained together (each method returns the endpoint itself).

## header

`endpoint.header(name, [schema], [description]): this`

Defines a request header recognized by the endpoint.
Any additional non-defined headers are treated as optional string values.
The definitions is also shown in the route details in the API documentation.

If the endpoint is a child router, all routes of that router use this
header definition unless overridden.

**Arguments**

- **name**: `string`

  Name of the header. This should be considered case insensitive as all header
  names will be converted to lowercase.

- **schema**: `Schema` (optional)

  A schema describing the format of the header value. This can be a joi schema
  or anything that has a compatible `validate` method.

  The value of this header is set to the `value` property of the
  validation result. A validation failure results in an automatic 400
  (Bad Request) error response.

- **description**: `string` (optional)

  A human-readable string that is shown in the API documentation.

Returns the endpoint.

**Examples**

```js
router.get(/* ... */)
.header('arangoVersion', joi.number().min(30000).default(30000));
```

## pathParam

`endpoint.pathParam(name, [schema], [description]): this`

Defines a path parameter recognized by the endpoint.
Path parameters are expected to be filled as part of the endpoint's mount path.
Any additional non-defined path parameters are treated as optional
string values. The definitions are also shown in the route details in
the API documentation.

If the endpoint is a child router, all routes of that router use this
parameter definition unless overridden.

**Arguments**

- **name**: `string`

  Name of the parameter.

- **schema**: `Schema` (optional)

  A schema describing the format of the parameter. This can be a joi schema
  or anything that has a compatible `validate` method.

  The value of this parameter is set to the `value` property of the
  validation result. A validation failure results in the route failing to
  match and being ignored (resulting in a 404 (Not Found) error response if no
  other routes match).

- **description**: `string` (optional)

  A human readable string that is shown in the API documentation.

Returns the endpoint.

**Examples**

```js
router.get('/some/:num/here', /* ... */)
.pathParam('num', joi.number().required());
```

## queryParam

`endpoint.queryParam(name, [schema], [description]): this`

Defines a query parameter recognized by the endpoint.
Any additional non-defined query parameters are treated as optional
string values. The definitions are also shown in the route details in
the API documentation.

If the endpoint is a child router, all routes of that router use this
parameter definition unless overridden.

**Arguments**

- **name**: `string`

  Name of the parameter.

- **schema**: `Schema` (optional)

  A schema describing the format of the parameter. This can be a joi schema or
  anything that has a compatible `validate` method.

  The value of this parameter is set to the `value` property of the
  validation result. A validation failure results in an automatic 400
  (Bad Request) error response.

- **description**: `string` (optional)

  A human-readable string that is shown in the API documentation.

Returns the endpoint.

**Examples**

```js
router.get(/* ... */)
.queryParam('num', joi.number().required());
```

## body

`endpoint.body([model], [mimes], [description]): this`

Defines the request body recognized by the endpoint.
There can only be one request body definition per endpoint. The definition is
also shown in the route details in the API documentation.

In the absence of a request body definition, the request object's `body`
property is initialized to the unprocessed `rawBody` buffer.

If the endpoint is a child router, all routes of that router use this body
definition unless overridden. If the endpoint is a middleware, the request body
is only parsed once (i.e. the MIME types of the route matching the same
request is ignored but the body is still validated again).

**Arguments**

- **model**: `Model | Schema | null` (optional)

  A model or joi schema describing the request body. A validation failure
  results in an automatic 400 (Bad Request) error response.

  If the value is a model with a `fromClient` method, that method is
  applied to the parsed request body.

  If the value is a schema or a model with a schema, the schema is used
  to validate the request body, and the `value` property of the validation
  result of the parsed request body is used instead of the parsed request
  body itself.

  If the value is a model or a schema and the MIME type has been omitted,
  the MIME type defaults to JSON instead.

  If the value is explicitly set to `null`, no request body is expected.

  If the value is an array containing exactly one model or schema, the request
  body is treated as an array of items matching that model or schema.

- **mimes**: `Array<string>` (optional)

  An array of MIME types the route supports.

  Common non-mime aliases like "json" or "html" are also supported and are
  expanded to the appropriate MIME type (e.g. "application/json" and "text/html").

  If the MIME type is recognized by Foxx, the request body is parsed into
  the appropriate structure before being validated. Only JSON,
  `application/x-www-form-urlencoded`, and multipart formats are supported in this way.

  If the MIME type indicated in the request headers does not match any of the
  supported MIME types, the first MIME type in the list is used instead.

  Failure to parse the request body results in an automatic 400
  (Bad Request) error response.

- **description**: `string` (optional)

  A human-readable string that is shown in the API documentation.

Returns the endpoint.

**Examples**

```js
router.post('/expects/some/json', /* ... */)
.body(
  joi.object().required(),
  'This implies JSON.'
);

router.post('/expects/nothing', /* ... */)
.body(null); // No body allowed

router.post('/expects/some/plaintext', /* ... */)
.body(['text/plain'], 'This body will be a string.');
```

## response

`endpoint.response([status], [model], [mimes], [description]): this`

Defines a response body for the endpoint. When using the response object's
`send` method in the request handler of this route, the definition with the
matching status code is used to generate the response body.
The definitions are also shown in the route details in the API documentation.

If the endpoint is a child router, all routes of that router use this
response definition unless overridden. If the endpoint is a middleware,
this method has no effect.

**Arguments**

- **status**: `number | string` (Default: `200` or `204`)

  HTTP status code the response applies to. If a string is provided instead of
  a numeric status code, it is used to look up a numeric status code using
  the [statuses](https://github.com/jshttp/statuses) module.

- **model**: `Model | Schema | null` (optional)

  A model or joi schema describing the response body.

  If the value is a model with a `forClient` method, that method is
  applied to the data passed to `response.send` within the route if the
  response status code matches (but also if no status code has been set).

  If the value is a schema or a model with a schema, the actual schema is
  not used to validate the response body and only serves to document the
  response in more detail in the API documentation.

  If the value is a model or a schema and the MIME type has been omitted,
  the MIME type defaults to JSON instead.

  If the value is explicitly set to `null` and the status code has been omitted,
  the status code defaults to `204` ("no content") instead of `200`.

  If the value is an array containing exactly one model or schema, the response
  body is an array of items matching that model or schema.

- **mimes**: `Array<string>` (optional)

  An array of MIME types the route might respond with for this status code.

  Common non-mime aliases like "json" or "html" are also supported and are
  expanded to the appropriate MIME type (e.g. "application/json" and "text/html").

  When using the `response.send()` method, the response body is converted to
  the appropriate MIME type if possible.

- **description**: `string` (optional)

  A human-readable string that briefly describes the response and is shown
  in the endpoint's detailed documentation.

Returns the endpoint.

**Examples**

```js
// This example only provides documentation
// and implies a generic JSON response body.
router.get(/* ... */)
.response(
  joi.array().items(joi.string()),
  'A list of doodad identifiers.'
);

// No response body is expected here.
router.delete(/* ... */)
.response(null, 'The doodad no longer exists.');

// An endpoint can define multiple response types
// for different status codes -- but never more than
// one for each status code.
router.post(/* ... */)
.response('found', 'The doodad is located elsewhere.')
.response(201, ['text/plain'], 'The doodad was created so here is a haiku.');

// Here, the response body is set to
// the querystring-encoded result of
// FormModel.forClient({some: 'data'})
// because the status code defaults to 200.
router.patch(function (req, res) {
  // ...
  res.send({some: 'data'});
})
.response(FormModel, ['application/x-www-form-urlencoded'], 'OMG.');

// In this case, the response body is set to
// SomeModel.forClient({some: 'data'}) because
// the status code has been set to 201 before.
router.put(function (req, res) {
  // ...
  res.status(201);
  res.send({some: 'data'});
})
.response(201, SomeModel, 'Something amazing happened.');
```

## error

`endpoint.error(status, [description]): this`

Documents an error status for the endpoint.

If the endpoint is a child router, all routes of that router use this
error description unless overridden. If the endpoint is a middleware,
this method has no effect.

This method only affects the generated API documentation and has no other
effect within the service itself.

**Arguments**

- **status**: `number | string`

  HTTP status code for the error (e.g. `400` for "bad request"). If a string is
  provided instead of a numeric status code it is used to look up a numeric
  status code using the [statuses](https://github.com/jshttp/statuses) module.

- **description**: `string` (optional)

  A human-readable string that briefly describes the error condition and is
  shown in the endpoint's detailed documentation.

Returns the endpoint.

**Examples**

```js
router.get(function (req, res) {
  // ...
  res.throw(403, 'Validation error at x.y.z');
})
.error(403, 'Indicates that a validation has failed.');
```

## summary

`endpoint.summary(summary): this`

Adds a short description to the endpoint's API documentation.

If the endpoint is a child router, all routes of that router use this
summary unless overridden. If the endpoint is a middleware, this method has no effect.

This method only affects the generated API documentation and has no other
effect within the service itself.

**Arguments**

- **summary**: `string`

  A human-readable string that briefly describes the endpoint and appears
  next to the endpoint's path in the documentation.

Returns the endpoint.

**Examples**

```js
router.get(/* ... */)
.summary('List all discombobulated doodads')
```

## description

`endpoint.description(description): this`

Adds a long description to the endpoint's API documentation.

If the endpoint is a child router, all routes of that router use
this description unless overridden. If the endpoint is a middleware,
this method has no effect.

This method only affects the generated API documentation and has not
other effect within the service itself.

**Arguments**

- **description**: `string`

  A human-readable string that describes the endpoint in detail and
  will be shown in the endpoint's detailed documentation.

Returns the endpoint.

**Examples**

```js
// The "dedent" library helps formatting
// multi-line strings by adjusting indentation
// and removing leading and trailing blank lines
const dd = require('dedent');
router.post(/* ... */)
.description(dd`
  This route discombobulates the doodads by
  frobnicating the moxie of the request body.
`)
```

## deprecated

`endpoint.deprecated([deprecated]): this`

Marks the endpoint as deprecated.

If the endpoint is a child router, all routes of that router are also
marked as deprecated. If the endpoint is a middleware, this method has no effect.

This method only affects the generated API documentation and has no other
effect within the service itself.

**Arguments**

- **deprecated**: `boolean` (Default: `true`)

  Whether the endpoint should be marked as deprecated. If set to `false`, the
  endpoint is explicitly marked as *not* deprecated.

Returns the endpoint.

**Examples**

```js
router.get(/* ... */)
.deprecated();
```

## tag

`endpoint.tag(...tags): this`

Marks the endpoint with the given tags that are used to group related
routes in the generated API documentation.

If the endpoint is a child router, all routes of that router are also
marked with the tags. If the endpoint is a middleware, this method has no effect.

This method only affects the generated API documentation and has no other
effect within the service itself.

**Arguments**

- **tags**: `string`

  One or more strings that are used to group the endpoint's routes.

Returns the endpoint.

**Examples**

```js
router.get(/* ... */)
.tag('auth', 'restricted');
```

## securityScheme

`endpoint.securityScheme(type, [options], [description])`

Defines an OpenAPI security scheme for this endpoint.

This method only affects the generated API documentation and has no other
effect within the service itself.

**Arguments**

- **type**: `string`

  Type of the security scheme.
  Must be one of `"basic"`, `"apiKey"` or `"oauth2"`.

- **options**: `object`

  An object with the following property:

  - **id**: `string` (optional)

    Unique identifier that can be used to opt in to or out of this scheme or
    to specify OAuth 2 scopes required by this endpoint.

  If **type** is set to `"basic"`, this parameter is optional.

  If **type** is set to `"apiKey"`, the following additional properties are
  required:

  - **name**: `string`

    The name of the header or query parameter that contains the API key.

  - **in**: `string`

    The location of the API key.
    Must be one of `"header"` or `"query"`.

  If **type** is set to `"oauth2"`, the following additional properties are
  required:

  - **flow**: `string`

    The OAuth 2 flow used by this security scheme.
    Must be one of `"implicit"`, `"password"`, `"application"` or
    `"accessCode"`.

  - **scopes**: `object`

    The available scopes for this OAuth 2 security scheme as a mapping of
    scope names to descriptions.

  If **flow** is set to `"implicit"` or `"accessCode"`, the following
  additional property is required:

  - **authorizationUrl**: `string`

    The authorization URL to be used for this OAuth 2 flow.

  If **flow** is set to `"password"`, `"application"` or `"accessCode"`, the
  following additional property is required:

  - **tokenUrl**: `string`

    The token URL to be used for this OAuth 2 flow.

- **description**: `string` (optional)

  A human-readable string that describes the security scheme.

Returns the endpoint.

**Examples**

```js
router.get(/* ... */)
.securityScheme('basic', 'Basic authentication with username and password.')
.securityScheme('apiKey', {name: 'x-api-key', in: 'header'},
  'API key as alternative to password-based authentication.'
);
```

## security

`endpoint.security(id, enabled)`

Opts this endpoint in to or out of the security scheme with the given ID.

- **id**: `string`

  Unique identifier of the security scheme. See `endpoint.securityScheme`.

- **enabled**: `boolean`

  Whether the security scheme should be enabled or disabled for this endpoint.
  Security schemes are enabled for all child routes by default.

**Examples**

```js
router.securityScheme('basic', {id: 'basic-auth'},
  'Basic authentication used by most endpoints on this router.'
);

router.get(/* ... */)
.security('basic-auth', false); // Opt this endpoint out
```

## securityScope

`endpoint.securityScope(id, ...scopes)`

Defines OAuth 2 scopes required by this endpoint for security scheme with the
given ID.

- **id**: `string`

  Unique identifier of the security scheme. See `endpoint.securityScheme`.

- **scopes**: `Array<string>`

  Names of OAuth 2 scopes required by this endpoint.

**Examples**

```js
router.get(/* ... */)
.securityScheme('oauth2', {
  id: 'thebookface-oauth2',
  flow: 'implicit',
  authorizationUrl: 'https://thebookface.example/oauth2/authorization',
  scopes: {
    'profile:read': 'Read user profile',
    'profile:write': 'Modify user profile'
  }
}, 'OAuth 2 authentication for The Bookface.')
.securityScope('thebookface-oauth2', 'profile:read');
```
