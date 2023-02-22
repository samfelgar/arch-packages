# Arch Linux Packages Updates Checker

This is a simple checker for Arch Linux Packages' updates.
This was developed for learning purposes and could be done with any scripting language, such as PHP or Python.

## Usage

1. Run `go build .` to generate the executable for your system.
2. Then, you can run the executable, informing the packages you want to know about. Example:

```bash
./arch-packages go
./arch-packages php gnome
./arch-packages python node openjdk
```

> TIP: You may move the binary to something like `/usr/bin`, so it will be available system-wide.