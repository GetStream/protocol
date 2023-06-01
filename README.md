# Protocol

This repository contains manifests that describe Stream public APIs, documentation and SDK clients can be generated using this manifests.

## OpenAPI

Stream REST API is described with OpenAPI specs, each product has its own openapi spec file.

At the moment client generation from OpenAPI specs is done inside the specific SDK repositories, generation scripts are not yet part of this repository and should be done using our dockerized fork of [openapi-generator](https://github.com/GetStream/openapi-generator).

```sh
docker pull ghcr.io/getstream/openapi-generator:master

docker run --rm -v "${PWD}:/local" ghcr.io/getstream/openapi-generator:master generate \
   -i https://raw.githubusercontent.com/GetStream/protocol/main/openapi/video-openapi.yaml \
   -g typescript-fetch \
   --additional-properties=supportsES6=true \
   --additional-properties=modelPropertyNaming=original \
   --additional-properties=enumPropertyNaming=original \
   --additional-properties=withoutRuntimeChecks=true \
   -o /local/src/gen/openapi-temp
```

## Protobuf and Twirp

Stream video servers use Twirp to expose an HTTP RPC layer. Websocket events sent by the video server are encoded using protobuf.

### Dependencies

Before generating, you have to install necessary dependencies using `./install.sh` script.

All dependencies are installed into `.protoc` directory. Some of them may be skipped if there's a missing dependency.
Swift plugins are not installed if there's no `swift` executable in the PATH and so on.

#### MacOS

You need to have `realpath`. It can be installed via:

```sh
brew install coreutils
```

If you have an issue with `realpath` in generation, try:

```sh
REALPATH=grealpath ./generate.sh
```

### Go Code Generation

To generate Go and put artifacts into this project, just run `./generate.sh` without arguments or run `go generate`.
You can also change the output directory, for example, to put generated code to the public Go SDK repo, but this will likely require
import prefix change. To change the import prefix, run `./generate.sh go <output_path> <import_prefix>`. Default import prefix
is `github.com/GetStream/protocol/protobuf`.

### Client Code Generation

We support generation for some client languages here:

- TypeScript
- Dart
- Swift

For detailed option settings, check out `generate.sh`

To generate code for your language, do this:

1. Install dependencies: `./install.sh`
2. Generate: `./generate.sh <language> <output_dir>`
   - `language`: language to generate code for can be `ts`, `dart` or `swift`
   - `output_dir` is a directory where you want to put generated code in


## Generate sdk with docker

```zsh
docker run -v $(pwd):/local ghcr.io/getstream/protocol dart /local/dart-sdk
```

Where `/local` is folder mounted to container from `PWD`.
In this way you will find the `dart-sdk` folder in your current directory.
