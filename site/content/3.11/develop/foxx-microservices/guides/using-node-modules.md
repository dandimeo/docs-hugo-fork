---
title: Bundled Node modules
menuTitle: Using Node modules
weight: 40
description: >-
  Bundled Node modules
archetype: default
---
You can use the `node_modules` folder to bundle Node.js modules with your Foxx
service. Note that many third-party libraries written for Node.js or the
browser rely on async or filesystem logic
[which may not be compatible with Foxx](../_index.md#compatibility-caveats).

{{< info >}}
Bundled node modules are often referred to as _dependencies_. In ArangoDB this
term can often be ambiguous because Foxx also provides a
[dependency mechanism](linking-services-together.md) for linking services together.
{{< /info >}}

Use a tool like [yarn](https://yarnpkg.com) or
[npm](https://npmjs.com) to
create a `package.json` file in your service source directory and add node
dependencies as you would for any other Node.js application or library:

```sh
cd my-foxx-service/
echo '{"private": true}' > package.json
yarn add lodash # or:
npm install --save lodash
```

Make sure to include the actual `node_modules` folder in your Foxx service
bundle as ArangoDB will not automatically install these dependencies for you.
Also keep in mind that bundling extraneous modules like development
dependencies may bloat the file size of your Foxx service bundle.

If you are using the [Foxx CLI](../../../components/tools/foxx-cli/_index.md)
command-line tool, you can exclude individual modules by ignoring them:

```sh
npm install --save prettier
foxx ignore '/node_modules/prettier/'
# the 'prettier' folder will now be excluded
# in service bundles generated by foxx-cli
foxx install /my-foxx-service
```

Keep in mind that both yarn and npm typically also install dependencies of
your dependencies to the `node_modules` folder which you'll need to ignore as
well if you want to exclude these modules from your service bundle.

If you are using the npm package manager, you can use
`npm install --global-style` to force these indirect dependencies
to be nested to make them easier to exclude.