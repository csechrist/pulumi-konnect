# Foo Resource Provider

The Foo Resource Provider lets you manage [Foo](http://example.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @camsechrist/pulumi-konnect
```

or `yarn`:

```bash
yarn add @camsechrist/pulumi-konnect
```


### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/csechrist123/pulumi-foo/sdk/go/...
```


## Configuration

The following configuration points are available for the `konnect` provider:

- `konnect:pat` (environment: `KONNECT_PAT`) - the API key for `konnect`
- `konnect:region` (environment: `KONNECT_REGION`) - the region in which to deploy resources
