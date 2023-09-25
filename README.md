# slugifier

[![go][go-version-src]][go-version-href]

A small tool, written in Go, to slugify files and directories, recursively.

Files and directories are renamed to their slugified version. If the slugified version already exists, skip the file or directory.

## About

This CLI use [github.com/mozillazg/go-unidecode](https://github.com/mozillazg/go-unidecode) to transform unicode characters to their closest ASCII representation. After that, some treatments are applied to the string to make it more readable.

- Remove all special characters
- Replace all spaces with a dot
- Remove all dots at the beginning and the end of the string
- Replace all dots that are repeated more than once with a single dot
- Keep `-` and `_` characters (and remove spaces before and after them)
- Full lowercase with option `-l`

Examples

- `La Quête d'Ewilan vol.1 : D'un monde à l'autre-·/_,:; (1), [Bottero, Pierre]`Author` @{1} <book> ?!//&` to `la.quete.d.ewilan.vol.1.d.un.monde.a.l.autre-._.1.bottero.pierre.author.1.book` with lowercase
- `00 - Préface` to `00-Preface`
- `Góðan daginn` to `Godan.Daginn`

## Install

```bash
go install github.com/ewilan-riviere/slugifier@v0.0.12
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

Lowercase mode to transform all characters to lowercase.

```bash
slugifier -l path/to/dir
```

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
go run main.go "path/to/dir"
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat-square&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
