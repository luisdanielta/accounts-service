$ENV = @{
    "JWT_KEY" = "88U)3F!!SlirpLZP(a^v-l0pj(ivN+f+~&J9m-6t~TZ51@l2F2"

    # database - qfsd_dashboard_postgresql
    "DBQDH" = "localhost"
    "DBQDPP" = "5432"
    "DBQDU" = "postgres"
    "DBQDP" = "postgres"
    "DBQDD" = "postgres"
}

foreach ($key in $ENV.Keys)
{
    $value = $ENV[$key]
    [Environment]::SetEnvironmentVariable($key, $value)
}
Write-Host "All Environment variables loaded successfully."