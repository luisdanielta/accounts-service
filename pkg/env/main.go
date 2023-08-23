package env

import "os"

var JWT_KEY string = os.Getenv("JWT_KEY")
