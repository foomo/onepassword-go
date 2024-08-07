package onepassword

import (
	"os"
)

func IsServiceAccount() bool {
	return os.Getenv("OP_SERVICE_ACCOUNT_TOKEN") != ""
}
