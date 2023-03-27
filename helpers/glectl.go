package helper

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/glesys/glesys-go/v6"
)

var version = "0.0.3"

var Commit string

func Version() string {
	if len(Commit) > 0 {
		commit := strings.TrimSuffix(Commit, "\n")
		version += "-" + commit
	}
	return version
}

func UserAgent() string {
	return fmt.Sprintf("glectl/%s", Version())
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
