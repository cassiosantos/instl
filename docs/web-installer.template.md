# Web Installer

> [!NOTE]
> The web install command can be used by anyone and does not require anything to be installed.
> Running the web install command will download instl and install the given GitHub project.
> After that, instl will be removed from the system again.

The instl web installer is a single command, which everyone can run, to install a GitHub project.
This is the basic syntax, which will return an install script from our API server:

```
                     ┌ The GitHub username of the project
                     |          ┌ The GitHub repository name of the project
                     |          |         ┌ The platform, see "Valid Platforms"
                     |          |         |
https://instl.sh/:username/:reponame/:platform
```

## Valid Platforms

| Valid Platforms | Parameter |
|-----------------|-----------|
|     Windows     |  `windows |
|      macOS      |  `macos`  |
|      Linux      |  `linux`  |

## Running the web installer command

> [!NOTE]
> Different operating systems need different commands to download and run the web installer script.

### Windows

This command will download and execute the web installer script for windows.
You have to execute it in a `powershell` terminal.

```powershell
iwr -useb instl.sh/:username/:reponame/windows | iex
```

### macOS

This command will download and execute the web installer script for macOS.

```bash
curl -fsSL instl.sh/:username/:reponame/macos | bash
```

### Linux

This command will download and execute the web installer script for linux.

```bash
curl -fsSL instl.sh/:username/:reponame/linux | bash
```
