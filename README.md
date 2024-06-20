# pik

Color picker for terminal

# Installation

## From source
```
go install github.com/immanelg/pik
```

# Usage

## Flags
- `-init string`: initial color (`#123456`, `rgb(...)`, `hsl(...)`)
- `-logfile /dev/null`: log debug messages to file

## Keybindings

- `h`, `l`, `b`, `w`, `[`, `]`: change current slider value
- `j`, `k`: change focused slider
- `i`: cycle input mode (RGB, HSL)
- `o`: cycle output mode (RGB, HEX, HSL)
- `y`: copy to clipboard
- `p`: paste from clipboard
- `enter`: quit and print current color
- `q`: quit
- `<c-l>`: redraw
