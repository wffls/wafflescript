# wafflescript

A Waffles interpreter

## Introduction

Wafflescript is a compiled Go binary that embeds [Waffles](https://github.com/wffls/waffles). Instead of installing Waffles on a local system, you can install this binary to `/usr/local/bin` (or your location of choice).

Under the covers, `wafflescript` is still running `/bin/bash`, so all valid Bash scripts are valid wafflescripts.

## Example

```shell
#!/usr/local/bin/wafflescript

# Install memcached
apt.pkg --package memcached --version latest

# Set the listen option
file.line --file /etc/memcached.conf --line "-l 0.0.0.0" --match "^-l"

# Determine the amount of memory available and use half of that for memcached
memory_bytes=$(elements System.Memory.Total 2>/dev/null)
memory=$(( $memory_bytes / 1024 / 1024 / 2 ))

# Set the memory available to memcached
file.line --file /etc/memcached.conf --line "-m $memory" --match "^-m"

# Manage the memcached service
service.sysv --name memcached

# If any changes happened, restart memcached
if [[ -n $waffles_total_changes ]]; then
  exec.mute /etc/init.d/memcached restart
fi
```

## Download

See the [Releases](https://github.com/wffls/wafflescript/releases) page.

## Release Information

New Wafflescript releases will be created when:

* New features are added to Wafflescript itself.
* A new version of Waffles is released.

Because of these two conditions, Wafflescript versions will look like the following:

```
x-y
```

Where `x` is the Wafflescript release and `y` is the Waffles relese.

For example:

```
0.1-0.30.1
```

Means the version of Wafflescript is `0.1` and contains Waffles `0.30.1`.

## Development

* [Install](https://github.com/travis-ci/gimme) and [Configure](https://golang.org/doc/code.html) Go.
* Download this repository:

```shell
$ go get github.com/jtopjian/wafflescript
```

* Pull in the dependencies:

```shell
$ make deps
```

* Make any changes
* Pull in [Waffles](https://github.com/wffls/waffles) and build a binary:

```shell
$ make build
$ echo log.info foobar | ~/go/bin/wafflescript
```
