# Slugifier

[![go][go-version-src]][go-version-href]

A small tool, written in Go, to slugify files and directories.

## Install

```bash
go install github.com/ewilan-riviere/slugifier-cli@latest
```

## Usage

For a directory

```bash
slugifier path/to/dir
```

Or for a file

```bash
slugifier path/to/file
```

### Options

Verbose mode to enable preview and confirmation before renaming.

```bash
slugifier -v path/to/dir
```

## Build

Build the script.

```bash
go build -o slugifier
```

You can use `./slugifier` to run the script.

```bash
./slugifier path/to/dir
```

Or you can install it globally.

```bash
go install
```

## Test

Check with `curl` if the webhook is working.

```bash
go test
```

```bash
go test ./pkg/... -coverprofile=coverage.out
go test -v ./...
go test -v ./pkg/file
```

Direct usage

```bash
go run main.go "path/to/file"
go run main.go "/Users/ewilan/Workspace/renamer-cli/misc/file/La Quête d'Ewilan vol.1 D'un monde à l'autre.md"
go run main.go "/Users/ewilan/Workspace/renamer-cli/misc/file"
go run main.go /Users/ewilan/Downloads/renamer
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat-square&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
