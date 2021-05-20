# instl

## Usage
> Instl is an installer that installs GitHub projects on your system with a single command.

instl [username/repo] [global options] command [options] [arguments...]

## Description

```
Instl is an installer that installs GitHub projects on your system with a single command.  
Additionally, Instl provides a server that generates dynamic scripts that install a GitHub project.  

Official docs: https://docs.instl.sh

To use the server you can use the following commands:
  
**Windows**  

    iwr -useb instl.sh/username/reponame/windows | iex  
  
**macOS**  

    sudo /bin/bash -c "$(curl -fsSL instl.sh/username/reponame/macos)"  
  
**Linux**  

    curl -fsSL instl.sh/username/reponame/linux | sudo bash  
  
(Replace username and reponame with the GitHub project you want to install)  

Read more about the web installer here: https://docs.instl.sh/#/web-installer
  
These commands can be executed from any system and install the respective GitHub project.  
You can also provide these commands to your users to make your GitHub project easily installable.
```
## Examples

```bash
instl instl-sh/instl
```

## Flags
|Flag|Usage|
|----|-----|
|`-d, --debug`|enable debug messages|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`instl help`|Help about any command|
# ... help
`instl help`

## Usage
> Help about any command

instl [username/repo] [global options] command [options] [arguments...]

## Description

```
Help provides help for any command in the application.
Simply type instl help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 20 May 2021**
