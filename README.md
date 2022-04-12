<h1 align="center">instl</h1>
<p align="center">Instl is an installer that can install most GitHub projects on your system with a single command.</p>

<p align="center">

<a style="text-decoration: none" href="https://github.com/installer/instl/releases">
<img src="https://img.shields.io/github/v/release/installer/instl?style=flat-square" alt="Latest Release">
</a>

<a style="text-decoration: none" href="https://github.com/installer/instl/releases">
<img src="https://img.shields.io/github/downloads/installer/instl/total.svg?style=flat-square" alt="Downloads">
</a>

<a style="text-decoration: none" href="https://github.com/installer/instl/stargazers">
<img src="https://img.shields.io/github/stars/installer/instl.svg?style=flat-square" alt="Stars">
</a>

<a style="text-decoration: none" href="https://github.com/installer/instl/fork">
<img src="https://img.shields.io/github/forks/installer/instl.svg?style=flat-square" alt="Forks">
</a>

<a style="text-decoration: none" href="https://github.com/installer/instl/issues">
<img src="https://img.shields.io/github/issues/installer/instl.svg?style=flat-square" alt="Issues">
</a>

<a style="text-decoration: none" href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<br/>

<a style="text-decoration: none" href="https://github.com/installer/instl/releases">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

<br/>

</p>

<p align="center">
<img src="https://user-images.githubusercontent.com/31022056/119270816-de43ba80-bbfe-11eb-92af-4b4eaf859399.gif" alt="Instl Demo Animation">
</p>

----

<p align="center">
<strong><a href="https://installer.github.io/instl/#/installation">Installation</a></strong>
|
<strong><a href="https://installer.github.io/instl/#/docs">Documentation</a></strong>
|
<strong><a href="https://installer.github.io/instl/#/CONTRIBUTING">Contributing</a></strong>
</p>

----

<p align="center">
<a href="https://discord.gg/vE2dNkfAmF">
<img width="300" src="https://user-images.githubusercontent.com/31022056/158916278-4504b838-7ecb-4ab9-a900-7dc002aade78.png" alt="Join us on Discord!" />
<br/>
<b>Join us on Discord for support, discussions, updates and general talk!</b>
</a>
</p>

Instl is an installer that can install most GitHub projects on your system with a single command.  
Additionally, Instl provides a server that generates dynamic scripts that install a GitHub project.  

Official docs: https://docs.instl.sh

## Web Installer

> The web install command can be used by anyone and does not require anything to be installed.
> Running the web install command will download instl and install the given GitHub project.
> After that, instl will be removed from the system again.

The instl web installer is a single command, which everyone can run, to install a GitHub project.
This is the basic syntax, which will return an install script from our API server:

                         ┌ The GitHub username of the project
                         |        ┌ The GitHub repository name of the project
                         |        |       ┌ The platform, see "Valid Platforms"
                         |        |       |
	https://instl.sh/username/reponame/platform

### Valid Platforms

| Valid Platforms | Parameter |
|-----------------|-----------|
|     Windows     |  windows  |
|      macOS      |  macos    |
|      Linux      |  linux    |

### Running the web installer command

> Different operating systems need different commands to download and run the web installer script.
> You can include those commands in your GitHub project, to provide a user-friendly installer for your CLI without any setup!

#### Windows

This command will download and execute the web installer script for windows.
You have to execute it in a powershell terminal.

	iwr -useb instl.sh/username/reponame/windows | iex

#### macOS

This command will download and execute the web installer script for macOS.

	curl -fsSL instl.sh/username/reponame/macos | bash

#### Linux

This command will download and execute the web installer script for linux.

	curl -fsSL instl.sh/username/reponame/linux | bash


## Installation

If you want to install instl directly to your system, to be able to install most GitHub projects with ease, you can use the following command:

**Windows**
```powershell
iwr instl.sh/installer/instl/windows | iex
```

**macOS**
```bash
curl -sSL instl.sh/installer/instl/macos | bash
```

**Linux**
```bash
curl -sSL instl.sh/installer/instl/linux | bash
```
