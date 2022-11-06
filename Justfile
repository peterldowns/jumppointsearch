# This Justfile contains rules/targets/scripts/commands that are used when
# developing. The most important ones are:
#
# - test: runs the test suite
# - lint: lints the package

# allow passing arguments through to tasks, see the docs here
# https://just.systems/man/en/chapter_24.html#positional-arguments
set positional-arguments

# print all available commands by default
default:
        just --list

# runs `go test`.
test *args='./...':
        go test "$@"

# run `golangci-lint`
lint *args:
  golangci-lint run --fix --config .golangci.yaml "$@"

