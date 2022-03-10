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

    curl -sSL instl.sh/username/reponame/macos | bash   
  
**Linux**  

    curl -sSL instl.sh/username/reponame/linux | bash  
  
(Replace username and reponame with the GitHub project you want to install)  

Read more about the web installer here: https://docs.instl.sh/#/web-installer
  
These commands can be executed from any system and install the respective GitHub project.  

## Installable Projects

Instl can install every public GitHub project, that has releases which contain a single binary.  
It will search the release for a binary and install it. Instl will also search inside archives.
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
|`instl check`|Performs various tests to ensure instl is working correctly on the host machine|
|`instl completion`|Generate the autocompletion script for the specified shell|
|`instl help`|Help about any command|
# ... check
`instl check`

## Usage
> Performs various tests to ensure instl is working correctly on the host machine

instl check

## Description

```
The check command performs various tests to ensure instl is working correctly on the host machine.
Use the -d (debug) flag to see more detailed output.
```
# ... completion
`instl completion`

## Usage
> Generate the autocompletion script for the specified shell

instl completion

## Description

```
Generate the autocompletion script for instl for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`instl completion bash`|Generate the autocompletion script for bash|
|`instl completion fish`|Generate the autocompletion script for fish|
|`instl completion powershell`|Generate the autocompletion script for powershell|
|`instl completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`instl completion bash`

## Usage
> Generate the autocompletion script for bash

instl completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(instl completion bash)

To load completions for every new session, execute once:

#### Linux:

	instl completion bash > /etc/bash_completion.d/instl

#### macOS:

	instl completion bash > /usr/local/etc/bash_completion.d/instl

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`instl completion fish`

## Usage
> Generate the autocompletion script for fish

instl completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	instl completion fish | source

To load completions for every new session, execute once:

	instl completion fish > ~/.config/fish/completions/instl.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`instl completion powershell`

## Usage
> Generate the autocompletion script for powershell

instl completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	instl completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`instl completion zsh`

## Usage
> Generate the autocompletion script for zsh

instl completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	instl completion zsh > "${fpath[1]}/_instl"

#### macOS:

	instl completion zsh > /usr/local/share/zsh/site-functions/_instl

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
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
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 10 March 2022**
