# Parsing Azure's Logs - Pal

----

## Introduction

Pal is a simple tool to parse Terraform Azure provider logs.
It can output the request traces in Markdown format, which can be used to create a GitHub issue or a forum post.

The Markdown output supports expanding/collapsing the request traces. This is useful when the request trace is very long.

## Usage

```bash
$ pal {path to terraform_log_file}
```

## Example

```bash
$ cd ./testdata
$ pal ./input.txt
```

Above command will generate a [markdown file named "output.md"](https://github.com/ms-henglu/pal/tree/main/testdata/output.md) in the same working directory.

## How to install?

```bash
$ go install github.com/ms-henglu/pal@latest
```