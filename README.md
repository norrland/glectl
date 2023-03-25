# glectl

CLI tool for managing GleSYS resources.

Heavily influenced by [gleshys](https://github.com/brother/gleshys).

# Use

Obtain a user account and API key [here](https://cloud.glesys.com).

## Environment variables

`GLESYS_USERID` - Project name "CL12345"

`GLESYS_TOKEN` - API Key "abcx98765"

## Usage from `glectl -h`

```
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  glectl [command]

Available Commands:
  completion  Generate completion script
  help        Help about any command
  ip          Manage IPs in your current GleSYS project.
  server      Manage GleSYS servers in your project.

Flags:
  -h, --help     help for glectl
  -t, --toggle   Help message for toggle

Use "glectl [command] --help" for more information about a command.
```
