package helper

import "os"

var glectlUserAgent = "glectl/"
var version = "0.0.1"

func UserAgent() string {
	return glectlUserAgent + version
}

// GetCredentials loads credentials from Env and returns them as `userid` `token`
func GetCredentials() (string, string) {
	return os.Getenv("GLESYS_USERID"), os.Getenv("GLESYS_TOKEN")
}
