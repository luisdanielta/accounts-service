$ENV = @{
    "JWT_KEY" = "88U)3F!!SlirpLZP(a^v-l0pj(ivN+f+~&J9m-6t~TZ51@l2F2"

    # database - qfsd_dashboard_postgresql
    "DBQDH" = "containers-us-west-34.railway.app"
    "DBQDPP" = "5938"
    "DBQDU" = "postgres"
    "DBQDP" = "ABjpX5SSjHqN29LJivPF"
    "DBQDD" = "railway"
}

foreach ($key in $ENV.Keys)
{
    $value = $ENV[$key]
    [Environment]::SetEnvironmentVariable($key, $value)
}
Write-Host "All Environment variables loaded successfully."