# Quick Start - Install instl

> [!TIP]
> instl is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
iwr instl.sh/installer/instl/windows | iex
```

#### ** Linux **

### Linux Command

```bash
curl -sSL instl.sh/installer/instl/linux | bash
```

#### ** macOS **

### macOS Command

```bash
curl -sSL instl.sh/installer/instl/macos | bash
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile instl from source, you have to have [Go](https://golang.org/) installed.

Compiling instl from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of instl.

```command
go install github.com/installer/instl@latest
```

<!-- tabs:end -->
