[![GoDoc](https://godoc.org/github.com/kreativka/dot-ui?status.svg)](https://godoc.org/github.com/kreativka/dot-ui)
[![Go Report Card](https://goreportcard.com/badge/github.com/kreativka/dot-ui)](https://goreportcard.com/report/github.com/kreativka/dot-ui)
# dot-ui

It's bemenu/dmenu/wldash/wofi like launcher. Best used with sway.

It shows .desktop files from XDG directories, and flatpaks. Fuzzy searching within pretty names from ".desktop and executables. Defaults to localized names.

Navigate through list using arrows and Ctrl+P/N.

## Installation

For Linux you need Wayland and the wayland, x11, xkbcommon, GLES, EGL development packages. On Fedora 28 and newer, install the dependencies with the command

```
$ sudo dnf install wayland-devel libX11-devel libxkbcommon-x11-devel mesa-libGLES-devel mesa-libEGL-devel
```

Go get dot-ui

```
$ go get github.com/kreativka/dot-ui/cmd/dot-ui
```

## Running

Add this to sway's config

```
    for_window [title="^dot-ui$"] floating enable
    bindsym $mod+d exec dot-ui
```

## License

The MIT License
