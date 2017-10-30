Database Security Scanner Agent
=============
This repository contains artifacts to aid in creating a Windows service using Go

[![Build Status](https://ci.appveyor.com/api/projects/status/78ok880rhrc32tmw/branch/master?svg=true)](https://ci.appveyor.com/api/projects/status/78ok880rhrc32tmw) 
[![codecov](https://codecov.io/gh/corystein/DbSecScanner/branch/master/graph/badge.svg)](https://codecov.io/gh/corystein/DbSecScanner)


## Getting Started
The following step will get you started to executing the build.  Some configuration may be required.

## Requirements
- Operating System : Windows or Mac OS
- Go [https://www.golang.org/]

### On Windows
1. Install Chocolately (Requires a session restart)
```powershell
Set-ExecutionPolicy Bypass; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```
2. Install GoLang using Chocolately
```powershell
choco install golang -y
```
3. Install GoLang dependency
```powershell
go get -u golang.org/x/sys/windows/svc
```
4. Compile executable
```powershell
./build.ps1
```

## Links
- [Service Example](https://github.com/golang/sys/tree/master/windows/svc/example)
- [Go Cron](https://github.com/jasonlvhit/gocron)