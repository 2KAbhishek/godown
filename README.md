<div align = "center">

<h1><a href="https://2kabhishek.github.io/godown">godown</a></h1>

<a href="https://github.com/2KAbhishek/godown/blob/main/LICENSE">
<img alt="License" src="https://img.shields.io/github/license/2kabhishek/godown?style=flat&color=eee&label="> </a>

<a href="https://github.com/2KAbhishek/godown/graphs/contributors">
<img alt="People" src="https://img.shields.io/github/contributors/2kabhishek/godown?style=flat&color=ffaaf2&label=People"> </a>

<a href="https://github.com/2KAbhishek/godown/stargazers">
<img alt="Stars" src="https://img.shields.io/github/stars/2kabhishek/godown?style=flat&color=98c379&label=Stars"></a>

<a href="https://github.com/2KAbhishek/godown/network/members">
<img alt="Forks" src="https://img.shields.io/github/forks/2kabhishek/godown?style=flat&color=66a8e0&label=Forks"> </a>

<a href="https://github.com/2KAbhishek/godown/watchers">
<img alt="Watches" src="https://img.shields.io/github/watchers/2kabhishek/godown?style=flat&color=f5d08b&label=Watches"> </a>

<a href="https://github.com/2KAbhishek/godown/pulse">
<img alt="Last Updated" src="https://img.shields.io/github/last-commit/2kabhishek/godown?style=flat&color=e06c75&label="> </a>

</div>

godown is a concurrent download manager written in go. 🌐⬇️🗃

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of `go`.

## Installing godown

To install godown, follow these steps:

```bash
git clone https://github.com/2kabhishek/godown
cd godown

# Without build
go run ./godown.go $url $target

# With build
go build
./godown.go $url $target

# After Installing
go install
godown $url $target
```

## Using godown

```bash
USAGE:
    godown [FLAGS] [OPTIONS]

FLAGS:
    -h, --help              Prints help information.

OPTIONS:
    -o, --option <value>    Option description. Options are: a, b. [default: a]

```

### To-Do

- [x] Create initial project
- [ ] Add support for command line arguments
- [ ] Better error reporting
- [ ] Make it pretty

### More Info
