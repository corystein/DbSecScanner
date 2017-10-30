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
$env:GOPATH = "$($PSScriptRoot)"
$env:GOOS = "windows"
$env:GOARCH = "amd64"

<#
go install                                                         `
    -installsuffix "static"                                        `
    -ldflags "-X $($PKG)/pkg/version.VERSION=$($VERSION)"          `
    ./...
#>

go build