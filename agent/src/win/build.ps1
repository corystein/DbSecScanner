$ErrorActionPreference = "Stop"
$ProgressPreference = "SilentlyContinue"

<#
if ([String]::IsNullOrEmpty($PKG)) {
    echo "PKG must be set"
    exit 1
}
if ([String]::IsNullOrEmpty($ARCH)) {
    echo "ARCH must be set"
    exit 1
}
if ([String]::IsNullOrEmpty($VERSION)) {
    echo "VERSION must be set"
    exit 1
}
#>

#export CGO_ENABLED=0
#export GOARCH="$($ARCH)"

#env GOOS=windows GOARCH=amd64 go build github.com/mholt/caddy/caddy
$env:GOPATH = "$Env:USERPROFILE\go"
$env:GOBIN = "$env:GOPATH\bin"
$env:GOOS = "windows"
$env:GOARCH = "amd64"

<#
go install                                                         `
    -installsuffix "static"                                        `
    -ldflags "-X $($PKG)/pkg/version.VERSION=$($VERSION)"          `
    ./...
#>

# Optional 
#git config --global http.sslVerify false

# Get dependencies
try {go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo} catch {}
try {go get -u golang.org/x/sys/windows/svc} catch {}

# Compile executable
Write-Output "Building..."

if ($invocation.MyCommand.Path -ieq "C:\projects\dbsecscanner") {
    Push-Location -Path "C:\projects\dbsecscanner\agent\src\win"
}

Write-Output "Current Directory: [$pwd]"

go generate
go build -o .\dbsecscanner.exe

Write-Output "Successful"