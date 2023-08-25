package env

import "os"

var JWT_KEY string = os.Getenv("JWT_KEY")

/*
	TASK @jose.orozco: create documentation for the database
	QD = qfsd_dashboard
	H = host
	U = user
	P = password
	D = dbname
	PP = port
*/

/* database - qfsd_dashboard_postgresql - QD = qfsd_dashboard */
var DBQDH string = os.Getenv("DBQDH")   /* host */
var DBQDPP string = os.Getenv("DBQDPP") /* port */
var DBQDU string = os.Getenv("DBQDU")   /* user */
var DBQDP string = os.Getenv("DBQDP")   /* password */
var DBQDD string = os.Getenv("DBQDD")   /* dbname */
