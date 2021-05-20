<h1 align="center">instl</h1>
<p align="center">Instl is an installer that installs GitHub projects on your system with a single command.</p>

<p align="center">

<a href="https://github.com/instl-sh/instl/releases">
<img src="https://img.shields.io/github/v/release/instl-sh/instl?style=flat-square" alt="Latest Release">
</a>
&nbsp;
<a href="https://github.com/instl-sh/instl/releases">
<img src="https://img.shields.io/github/downloads/instl-sh/instl/total.svg?style=flat-square" alt="Downloads">
</a>
&nbsp;
<a href="https://github.com/instl-sh/instl/stargazers">
<img src="https://img.shields.io/github/stars/instl-sh/instl.svg?style=flat-square" alt="Stars">
</a>
&nbsp;
<a href="https://github.com/instl-sh/instl/fork">
<img src="https://img.shields.io/github/forks/instl-sh/instl.svg?style=flat-square" alt="Forks">
</a>
&nbsp;
<a href="https://github.com/instl-sh/instl/issues">
<img src="https://img.shields.io/github/issues/instl-sh/instl.svg?style=flat-square" alt="Issues">
</a>
&nbsp;
<a href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<br/>

<a href="https://github.com/instl-sh/instl/releases">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

<br/>

</p>

----

<p align="center">
<strong><a href="https://instl-sh.github.io/instl/#/installation">Installation</a></strong>
|
<strong><a href="https://instl-sh.github.io/instl/#/docs">Documentation</a></strong>
|
<strong><a href="https://instl-sh.github.io/instl/#/CONTRIBUTING">Contributing</a></strong>
</p>

----

Instl is an installer that installs GitHub projects on your system with a single command.  
Additionally, Instl provides a server that generates dynamic scripts that install a GitHub project.  

Official docs: https://docs.instl.sh

To use the server you can use the following commands:
  
**Windows**  

    iwr -useb instl.sh/username/reponame/windows | iex  
  
**macOS**  

    curl -fsSL instl.sh/username/reponame/macos | bash   
  
**Linux**  

    curl -fsSL instl.sh/username/reponame/linux | bash  
  
(Replace username and reponame with the GitHub project you want to install)  

Read more about the web installer here: https://docs.instl.sh/#/web-installer
  
These commands can be executed from any system and install the respective GitHub project.  
You can also provide these commands to your users to make your GitHub project easily installable.

## Installation

Run the following command in a terminal and you're ready to go!

**Windows**
```powershell
iwr -useb instl.sh/instl-sh/instl/windows | iex
```

**macOS**
```bash
curl -fsSL instl.sh/instl-sh/instl/macos | bash "
```

**Linux**
```bash
curl -fsSL instl.sh/instl-sh/instl/linux | bash
```
