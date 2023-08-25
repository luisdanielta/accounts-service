Param (
    [string]$Option = "go"
)

Invoke-Expression -Command ".\env.ps1"

if ($option -eq "go")
{
    go run app.go
    exit 0
}
elseif ($Option -eq "dev")
{
    Write-Host "mode dev"
    air -c .air.toml
    exit 0
}
elseif ($Option -eq "build")
{
    Write-Host "mode build"
    $path = "./build"

    if (!(Test-Path $path))
    {
        mkdir "build"
    }

    go build -o ./build/ app.go

    ./build/app.exe

    exit 0
}