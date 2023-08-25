$ENV = @{
    "JWT_KEY" = "88U)3F!!SlirpLZP(a^v-l0pj(ivN+f+~&J9m-6t~TZ51@l2F2"

    # database - qfsd_dashboard_postgresql
    "DBQDH" = "192.168.1.201"
    "DBQDPP" = "5432"
    "DBQDU" = "admin@qfsd-linux"
    "DBQDP" = "CAL3158focas?"
    "DBQDD" = "qfsd"
}

foreach ($key in $ENV.Keys)
{
    $value = $ENV[$key]
    [Environment]::SetEnvironmentVariable($key, $value)
}
Write-Host "All Environment variables loaded successfully."