# pik
![image](https://github.com/immanelg/pik/assets/119798691/a7156404-c752-4899-b837-217dc1acc471)

Color picker for terminal

# Installation

## From source
```
CGO_ENABLED=0 go install -trimpath -ldflags="-s -w" github.com/immanelg/pik@latest
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
