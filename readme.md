## CoS Cli

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/chief-of-state/cos-cli/main)

The Chief Of State Command line Tool(a.k.a cos-cli) is command line tool that will help execute some
commands against a running Chief Of State instance. 

## Features
With the cos-cli one can manage the various read sides that will be running with CoS.
- Resume a read side  across the whole cluster or given a shard number.
- Pause a read side  across the whole cluster or given a shard number.
- Skip a read side offsets across the whole cluster or given a shard number. 
- List a read side offsets across the whole cluster or given a shard number.

## Usage

- One needs at least CoS version [1.0.0-beta.1](https://github.com/chief-of-state/chief-of-state/pkgs/container/chief-of-state/15359929?tag=1.0.0-beta.1) running before using the cos-cli
- Download any of the artifacts that matches your OS at [releases](https://github.com/chief-of-state/cos-cli/releases).
- Run `cos-cli` and you will see the output displayed below:
```
cos-cli is command line tool that helps send commands to a running CoS to manage the various read sides.

Usage:
  cos-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  readside    readside command helps manage the various read-sides

Flags:
      --cosHost string   CoS service host address
      --cosPort int      CoS service port (default 9000)
  -h, --help             help for cos-cli

Use "cos-cli [command] --help" for more information about a command.
```

### Examples

- List offsets across the whole cluster for read side `READSIDE_1`: 
```bash 
cos-cli readside offset --cosHost=localhost --cosPort=9000 --id=READSIDE_1
```
- Get offset for a given shard for read side `READSIDE_1`: 
```bash
cos-cli readside offset --cosHost=localhost --cosPort=9000 --id=READSIDE_1 --shard-number=2
 ```
- Pause read side across the whole cluster `READSIDE_1`: 
```bash 
cos-cli readside pause --cosHost=localhost --cosPort=9000 --id=READSIDE_1
```
- Pause the read side `READSIDE_1` for the shard number 2:
```bash 
cos-cli readside pause --cosHost=localhost --cosPort=9000 --id=READSIDE_1 --shard-number=2
```
- Resume paused read side across the whole cluster `READSIDE_1`: 
```bash 
cos-cli readside resume --cosHost=localhost --cosPort=9000 --id=READSIDE_1
```
- Pause the paused read side `READSIDE_1` for the shard number 2: 
```bash 
cos-cli readside resume --cosHost=localhost --cosPort=9000 --id=READSIDE_1 --shard-number=2
```
