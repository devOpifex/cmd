![](https://img.shields.io/badge/state-experimental-orange)

# cmd

Create command line applications from R packages.

<details>
<summary>How it works</summary>
It's a code generator that outputs Go code which produces
a Command Line Application.

It's a bit ugly and unwieldy but it does the job.
</details>

## Install

```bash
go get -u devOpifex/cmd
```

```bash
go install -u devOpifex/cmd@latest
```

## Config

Create a config file called `cmd.json` at the root of
your R package.

```json
{
	"program": "rshiny",
	"package": "shiny",
	"description": "shiny from the command line",
	"commands": [
		{
			"name": "run",
			"short": "r",
			"description": "run the application",
			"function": "runApp",
			"arguments": [
				{
					"name": "dir",
					"short": "d",
					"description": "directory where the app.R is location",
					"type": "string",
					"required": false,
					"default": "."
				}
			]
		}
	]
}
```

__Root__

Specify the name of the `program` this is what users will call
as root command in their shell. Indicate the `package` this
CLI wraps.

In the example above we create a CLI for the shiny package and
call the program `rshiny`, e.g.: `rshiny -h` will bring up
the help.

__Commands__

List the commands you want to expose to the users of the CLIU,
these map to functions from the `package`.
The `short` is the short version of the command so one can use
`rshiny run` or `rshiny r`: this will run `shiny::runApp()`.

__Arguments__

List the arguments the function should take: _these are currently
passed unnamed so order matters_

Type can be one of `string`, `numeric`, `integer` or, `logical`.

Above makes it such that we can call `rshiny run -d="this/dir"` or
`rshiny run --dir=/some/path`

## Generate

Once the config file created and cmd installed run `cmd generate`
from the root of the package where the config file is located.

This creates a directory named after your program.

Move into the created directory 
(e.g.: `cd rshiny` in the example above) 
then run the folowing.

```bash
go mod init github.com/<username>/<repo>/<program>
go mod tidy
```

The first line will (in part) ensure you can work with go outside
of your `GOROOT`, the second downloads dependencies if there are
any (currently only [cobra](https://github.com/spf13/cobra)).

You should only need to run the above once. No need to repeat the 
process if you want to update your config and regenerate the code
with `cmd generate`.

Then to build run the following.

```bash
go build
```

This should produce a program you can call from the command line.
If you only need it on your own system you can just call
`go install`.

By default `go build` builds for your system, you can easily
build for other system with:

```r
GOOS=linux go build
GOOS=windows go build
GOOS=darwin go build
```
