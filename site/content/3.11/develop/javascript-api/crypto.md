---
title: Crypto Module
menuTitle: '@arangodb/crypto'
weight: 25
description: >-
  Crypto Module
archetype: default
---
`const crypto = require('@arangodb/crypto')`

The crypto module provides implementations of various hashing algorithms as well as cryptography related functions.

## Nonces

These functions deal with
[cryptographic nonces](https://en.wikipedia.org/wiki/Cryptographic_nonce).

For single server use only.

### createNonce

`crypto.createNonce(): string`

Creates a cryptographic nonce consisting of the first 32 bits of a timestamp
and 64 bit of randomness.

The nonce is held in memory for approximately one hour by the server.

Returns the created nonce as base64-encoded string.

### checkAndMarkNonce

`crypto.checkAndMarkNonce(nonce): void`

Checks if the nonce is valid and marks it as used.

**Arguments**

- **nonce**: `string`

  The nonce to check and mark.

Returns `true` if the supplied nonce was issued by the server and not marked
before, otherwise `false`.

## Random values

The following functions deal with generating random values.

### rand

`crypto.rand(): number`

Generates a random integer that may be positive, negative or even zero.

Returns the generated number.

### genRandomAlphaNumbers

`crypto.genRandomAlphaNumbers(length): string`

Generates a string of random alphabetical characters and digits.

**Arguments**

- **length**: `number`

  The length of the string to generate.

Returns the generated string.

### genRandomNumbers

`crypto.genRandomNumbers(length): string`

Generates a string of random digits.

**Arguments**

- **length**: `number`

  The length of the string to generate.

Returns the generated string.

### genRandomSalt

`crypto.genRandomSalt(length): string`

Generates a string of random (printable) ASCII characters.

**Arguments**

- **length**: `number`

  The length of the string to generate.

Returns the generated string.

### genRandomBytes

`crypto.genRandomBytes(length): Buffer`

Generates a buffer of random bytes.

**Arguments**

- **length**: `number`

  The length of the buffer to generate.

Returns the generated buffer.

### uuidv4

`crypto.uuidv4(): string`

Generates a random UUID v4 string.

Returns the generated UUID string.

## JSON Web Tokens (JWT)

These methods implement the JSON Web Token standard.

### jwtEncode

`crypto.jwtEncode(key, message, algorithm): string`

Generates a JSON Web Token for the given message.

**Arguments**

- **key**: `string | null`

  The secret cryptographic key to be used to sign the message using the given algorithm.
  Note that this function will raise an error if the key is omitted but the algorithm expects a key,
  and also if the algorithm does not expect a key but a key is provided (e.g. when using `"none"`).

- **message**: `string`

  Message to be encoded as JWT. Note that the message will only be base64-encoded and signed, not encrypted.
  Do not store sensitive information in tokens unless they will only be handled by trusted parties.

- **algorithm**: `string`

  Name of the algorithm to use for signing the message, e.g. `"HS512"`.

Returns the JSON Web Token.

### jwtDecode

`crypto.jwtDecode(key, token, noVerify): string | null`

**Arguments**

- **key**: `string | null`

  The secret cryptographic key that was used to sign the message using the algorithm indicated by the token.
  Note that this function will raise an error if the key is omitted but the algorithm expects a key.

  If the algorithm does not expect a key but a key is provided, the token will fail to verify.

- **token**: `string`

  The token to decode.

  Note that the function will raise an error if the token is malformed (e.g. does not have exactly three segments).

- **noVerify**: `boolean` (Default: `false`)

  Whether verification should be skipped. If this is set to `true` the signature of the token will not be verified.
  Otherwise the function will raise an error if the signature cannot be verified using the given key.

Returns the decoded JSON message or `null` if no token is provided.

### jwtAlgorithms

A helper object containing the supported JWT algorithms. Each attribute name corresponds to a JWT `alg` and the value is an object with `sign` and `verify` methods.

### jwtCanonicalAlgorithmName

`crypto.jwtCanonicalAlgorithmName(name): string`

A helper function that translates a JWT `alg` value found in a JWT header into the canonical name of the algorithm in `jwtAlgorithms`. Raises an error if no algorithm with a matching name is found.

**Arguments**

- **name**: `string`

  Algorithm name to look up.

Returns the canonical name for the algorithm.

## Hashing algorithms

### md5

`crypto.md5(message): string`

Hashes the given message using the MD5 algorithm.

**Arguments**

- **message**: `string`

  The message to hash.

Returns the cryptographic hash.

### sha1

`crypto.sha1(message): string`

Hashes the given message using the SHA-1 algorithm.

**Arguments**

- **message**: `string`

  The message to hash.

Returns the cryptographic hash.

### sha224

`crypto.sha224(message): string`

Hashes the given message using the SHA-224 algorithm.

**Arguments**

- **message**: `string`

  The message to hash.

Returns the cryptographic hash.

### sha256

`crypto.sha256(message): string`

Hashes the given message using the SHA-256 algorithm.

**Arguments**

- **message**: `string`

  The message to hash.

Returns the cryptographic hash.

### sha384

`crypto.sha384(message): string`

Hashes the given message using the SHA-384 algorithm.

**Arguments**

- **message**: `string`

  The message to hash.

Returns the cryptographic hash.

### sha512

`crypto.sha512(message): string`

Hashes the given message using the SHA-512 algorithm.

**Arguments**

- **message**: `string`

  The message to hash.

Returns the cryptographic hash.

## Miscellaneous

### constantEquals

`crypto.constantEquals(str1, str2): boolean`

Compares two strings.
This function iterates over the entire length of both strings
and can help making certain timing attacks harder.

**Arguments**

- **str1**: `string`

  The first string to compare.

- **str2**: `string`

  The second string to compare.

Returns `true` if the strings are equal, `false` otherwise.

### pbkdf2

`crypto.pbkdf2(salt, password, iterations, keyLength): string`

Generates a PBKDF2-HMAC-SHA1 hash of the given password.

**Arguments**

- **salt**: `string`

  The cryptographic salt to hash the password with.

- **password**: `string`

  The message or password to hash.

- **iterations**: `number`

  The number of iterations.
  This should be a very high number.
  OWASP recommended 64000 iterations in 2012 and recommends doubling that number every two years.

  When using PBKDF2 for password hashes it is also recommended to add a random value
  (typically between 0 and 32000) to that number that is different for each user.

- **keyLength**: `number`

  The key length.

Returns the cryptographic hash.

### hmac

`crypto.hmac(key, message, algorithm): string`

Generates an HMAC hash of the given message.

**Arguments**

- **key**: `string`

  The cryptographic key to use to hash the message.

- **message**: `string`

  The message to hash.

- **algorithm**: `string`

  The name of the algorithm to use.

Returns the cryptographic hash.
