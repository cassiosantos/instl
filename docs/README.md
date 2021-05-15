<h1 align="center">instl</h1>
<p align="center">Instl can install GitHub projects on your machine. No setup required.</p>

<p align="center">

<a href="https://github.com/instl-sh/instl/releases">
<img src="https://img.shields.io/github/v/release/instl-sh/instl?style=flat-square" alt="Latest Release">
</a>

<a href="https://github.com/instl-sh/instl/stargazers">
<img src="https://img.shields.io/github/stars/instl-sh/instl.svg?style=flat-square" alt="Stars">
</a>

<a href="https://github.com/instl-sh/instl/fork">
<img src="https://img.shields.io/github/forks/instl-sh/instl.svg?style=flat-square" alt="Forks">
</a>

<a href="https://github.com/instl-sh/instl/issues">
<img src="https://img.shields.io/github/issues/instl-sh/instl.svg?style=flat-square" alt="Issues">
</a>

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

Instl is a CLI tool which detects the right release of a GitHub repository for your system.
It will download the detected release and install the asset files to your computer.
The repositories themself, don't need a setup to be installable with instl. They just need a default release with assets for multiple operating systems.'

## Installation

Run the following command in a terminal and you're ready to go!

**Windows**
```powershell
iwr -useb instl.sh/instl-sh/instl/windows | iex
```

**macOS**
```bash
/bin/bash -c "$(curl -fsSL instl.sh/instl-sh/instl/macos)"
```

**Linux**
```bash
curl -s https://instl.sh/instl-sh/instl/linux | sudo bash
```
