# Video Protobuf

## Dependencies

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

## Go Code Generation

To generate Go and put artifacts into this project, just run `./generate.sh` without arguments or run `go generate`.
You can also change the output directory, for example, to put generated code to the public Go SDK repo, but this will likely require
import prefix change. To change the import prefix, run `./generate.sh go <output_path> <import_prefix>`. Default import prefix
is `github.com/GetStream/video-proto/protobuf`.

## Client Code Generation

We support generation for some client languages here:

- TypeScript
- Dart
- Swift

For detailed option settings, check out `generate.sh`

To generate code for your language, do this:

1. Switch to the protobuf directory: `cd protobuf`
2. Install dependencies: `./install.sh`
3. Generate: `./generate.sh <language> <output_dir>`
   - `profile: language to generate code for
   - `output_dir` is a directory where you want to put generated code in

## Mock SFU generation for Coordinator QA testing

Run `./generate.sh ts-node <output>`. This will generate a TypeScript Node.js server, the generated code will only contain SFU endpoints and models that are required by the coordinator API.
