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

## Installation

Run the following command in a terminal and you're ready to go!

**Windows**
```powershell
iwr instl.sh/installer/instl/windows | iex
```

**macOS**
```bash
curl -sSL instl.sh/installer/instl/macos | sudo bash
```

**Linux**
```bash
curl -sSL instl.sh/installer/instl/linux | sudo bash
```
