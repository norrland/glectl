package helper

import (
	"os"

	"github.com/glesys/glesys-go/v6"
)

var glectlUserAgent = "glectl/"
var version = "0.0.1"

func UserAgent() string {
	return glectlUserAgent + version
}

// GetCredentials loads credentials from Env and returns them as `userid` `token`
func GetCredentials() (string, string) {
	return os.Getenv("GLESYS_USERID"), os.Getenv("GLESYS_TOKEN")
}

func NewClient() *glesys.Client {
	u, t := GetCredentials()
	a := UserAgent()
	return glesys.NewClient(u, t, a)
}
