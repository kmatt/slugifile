# slugifile

Forked from https://github.com/ewilan-riviere/slugifile

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

- `La Quête d'Ewilan vol.1 : D'un monde à l'autre-·/_,:; (1), [Bottero, Pierre]`Author` @{1} <book> ?!//&` to `la.quete.dewilan.vol.1.dun.monde.a.lautre-._.1.bottero.pierre.author.{1}.book` with lowercase
- `00 - Préface` to `00-Preface`
- `Góðan daginn` to `Godan.Daginn`

## Install

```bash
go install github.com/kmatt/slugifile@latest
```

## Usage

For a directory

```bash
slugifile path/to/dir
```

Or for a file

```bash
slugifile path/to/file
```

### Options

Lowercase mode to transform all characters to lowercase.

```bash
slugifile -l path/to/dir
```

Verbose mode to enable preview and confirmation before renaming.

```bash
slugifile -v path/to/dir
```

## Build

Build the script.

```bash
go build -o slugifile
```

You can use `./slugifile` to run the script.

```bash
./slugifile path/to/dir
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
