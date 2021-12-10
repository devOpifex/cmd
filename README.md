# cmd

:warning: Experimental

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
