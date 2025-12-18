# cfget
Config file snippets by topic.

Usage:
```
cfget [options] [tags]
```
Where `[tags]` is the category of configuration file you're looking for, like `hyprland/decoration/blur` or `alacritty/window`.

`[options]` can be either `--ctmax` or `-s`; `--ctmax` followed by a count, like `--ctmax X` will get `X` snippets. `-s` will only get the first.

## Installation
Clone cfget:
```bash
git clone https://github.com/aarikpokras/cfget.git
```

Build the program:
```bash
go build main.go
```

Or you can get the build from the artifacts page.

## Contributing
To contribute a snippet, make a pull request with the snippet in its proper format in the proper directory. Additionally, at the top of the snippet, add a link to the full documentation for that section.
