# cfget
A terminal-accessible database of configuration file snippets.

Usage:
```
cfget -s [tags]
```
Where `[tags]` is the category of configuration file you're looking for, like `hyprland/decoration/blur` or `alacritty/window`.

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
To contribute a snippet, make a pull request with the snippet in its proper format in the proper directory.
