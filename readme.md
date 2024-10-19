# dollar-skipper

A simple command-line tool to skip the `$` in your shell.

## Motivation

When pasting a command from the internet, some websites include a `$` at the beginning of the command. This is fine, but can be annoying after a while.
This app, which was written in Go, will skip the `$` for you. See the example below for more details.

## Example

When pasting something like this directly to your shell from a website:
```bash
$ echo hi
```

You will get this result:

```bash
$: $ echo hi
```

Leading to the command not being executed, as `$` is not a valid command.
Dollar-skipper acts as a command named `$` which will just execute the command you gave it.

So the result will be:

```bash
$: $ echo hi
hi
```

## Installation

You will need to install it from source, meaning you need to have Go installed on your machine.

So start by cloning the repo, and then run `go build -o $ cmd/dollar-skipper/main.go` inside of the folder.

If you use windows, change `-o $` to `-o $.exe`.

After that, you will most likely want to move the binary to somewhere in your PATH, so you can use it from anywhere. That's it!

## License

This project is licensed under the `Unlicense` license. See the [LICENSE](LICENSE) file for more details.
