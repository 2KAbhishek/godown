# godown

![Size](https://img.shields.io/github/repo-size/2kabhishek/godown?style=plastic&color=0f0&label=Size)
![Updated](https://img.shields.io/github/last-commit/2kabhishek/godown?style=plastic&color=f00&label=Updated)
![Stars](https://img.shields.io/github/stars/2kabhishek/godown?style=plastic&color=ffc801&label=Stars)
![Forks](https://img.shields.io/github/forks/2kabhishek/godown?style=plastic&color=003cff&label=Forks)
![Watchers](https://img.shields.io/github/watchers/2kabhishek/godown?style=plastic&color=ff5500&label=Watchers)
![Contributors](https://img.shields.io/github/contributors/2kabhishek/godown?style=plastic&color=f0f&label=Contributors)
![License](https://img.shields.io/github/license/2kabhishek/godown?style=plastic&color=555&label=License)

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
