# cfget
Terminal-accessible config file snippets by topic.

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/aarikpokras/cfget/.github%2Fworkflows%2Fgo.yml)
![GitHub License](https://img.shields.io/github/license/aarikpokras/cfget)
![GitHub repo size](https://img.shields.io/github/repo-size/aarikpokras/cfget)
![GitHub Release](https://img.shields.io/github/v/release/aarikpokras/cfget)

Consolidates config files so that you don't have to go searching all over the Internet for them.

Usage:
```
cfget [options] [tags]
```
Where `[tags]` is the category of configuration file you're looking for, like `hyprland/decoration/blur` or `alacritty/window`.

`options`:
|Flag|Meaning|
|----|-------|
|`--ctmax [count]`|Get up to the `[count]`th snippet.|
|`-s`|Get the first snippet.|
|`--gos [num]`|Get the `[num]`th snippet for tags by number.|

The `CFGET_HIDE_ERRORS` environment variable can have the value "0" or "1" to decide whether to hide errors. "0" shows all errors and "1" hides all errors.

## Installation
To build cfmake from the source, you need:
* `go`
* `make`

Clone cfget:
```bash
git clone https://github.com/aarikpokras/cfget.git
```

Build the program:
```bash
make
```

Or you can get the build from the artifacts page and either:
* Add an alias for it in your rc file
* Put it in a directory and add it to your `PATH`

## Contributing
To contribute a snippet, make a pull request with the snippet in its proper format in the proper directory. Additionally, at the top of the snippet, add a link to the full documentation for that section.
