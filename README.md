# cfget
Terminal-accessible config file snippets by topic.

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
