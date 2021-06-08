# instl

## Usage
> Instl is an installer that can install most GitHub projects on your system with a single command.

instl [username/repo]

## Description

```
Instl is an installer that can install most GitHub projects on your system with a single command.  
Additionally, Instl provides a server that generates dynamic scripts that install a GitHub project.  

Official docs: https://docs.instl.sh

## Web Installer

Use these commands, if you don't have instl on your system to install any GitHub project:  
(If you own a GitHub repo, you can put these commands into your readme, to let users install your tool easily)
  
**Windows**  

    iwr instl.sh/username/reponame/windows | iex  
  
**macOS**  

    curl -sSL instl.sh/username/reponame/macos | sudo bash   
  
**Linux**  

    curl -sSL instl.sh/username/reponame/linux | sudo bash  
  
(Replace username and reponame with the GitHub project you want to install)  

Read more about the web installer here: https://docs.instl.sh/#/web-installer
  
These commands can be executed from any system and install the respective GitHub project.  

## Installable Projects

Instl can install every public GitHub project, that has releases which contain a single binary.  
Instl will search the release for a binary and install it. Instl will also search archives.
```
## Examples

```bash
instl installer/instl
```

## Flags
|Flag|Usage|
|----|-----|
|`-d, --debug`|enable debug messages|
|`--raw`|print unstyled raw output (set it if output is written to a file)|
|`-s, --silent`|only outputs errors|

## Commands
|Command|Usage|
|-------|-----|
|`instl help`|Help about any command|
# ... help
`instl help`

## Usage
> Help about any command

instl help [command]

## Description

```
Help provides help for any command in the application.
Simply type instl help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 08 June 2021**
