# Go env library

`goenv` is a  simple go library permitt you to handle env var easier.

## Introduction

Base on driver configuration completion problem (not really one), this library permitt to auto-complete a string with given text and interpolation syntax

## Requirement

- `go 1.13.+ or more recent`

## Usage

### `Sprintf`

This feature is the core and the reason of this library.
Pass an text with interpolation syntax (e.g. "text here with %replacement% then") and this function automatically replace `%replacement%` to the env value that have as key `replacement` in the output text.
In not found case, the function let the syntax interpolation.

*exemple*:

```env
 FUNC_NAME=Sprintf
```

`testing the function with %func_name%` became `testing the function Sprintf`

### GetEnv

The classic function that get value according to the given key

*exemple*:

```env
FUNC_NAME=RemainEnvVar
```

```go
GetEnv("FUNC_NAME") // output => RemainEnvVar
```

### ImportFromOS

This feature is to handle with precaution. It import all environment variable verify given pattern
And then, you can exploit them with `GetEnv` function
*exemple*

```env
FUNC_NAME=ImportFromOS
FUNC_DIR=here
AS_FUNC_NAME=alias
```

```go
ImportFromOS("FUNC_*")
GetEnv("FUNC_NAME") // output => ImportFromOS
GetEnv("FUNC_DIR") // output => here
GetEnv("AS_FUNC_NAME") // empty output
```

## Test

Just run `go test` in current workdir, than go execute each testing function in corresponding test file that respect pattern name `<name>_test.go`

## Warning

One of difficulty you can accounter is the rewriting of env var. The last value of a similar key will be saved. So be carefull about the order of importation.
