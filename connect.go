package onepassword

import (
	"os"
	"strings"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
	"github.com/pkg/errors"
)

func IsConnect() bool {
	return os.Getenv("OP_CONNECT_HOST") != "" && os.Getenv("OP_CONNECT_TOKEN") != ""
}

func ConnectSecret(client connect.Client, vaultUUID, itemUUID string) (map[string]string, error) {
	var err error
	var item *onepassword.Item
	if onePasswordUUID.MatchString(itemUUID) {
		item, err = client.GetItem(itemUUID, vaultUUID)
		if err != nil {
			return nil, err
		}
	} else {
		item, err = client.GetItemByTitle(itemUUID, vaultUUID)
		if err != nil {
			return nil, err
		}
	}

	ret := map[string]string{}
	for _, f := range item.Fields {
		ret[f.Label] = f.Value
	}

	return ret, nil
}

func ConnectDocument(client connect.Client, vaultUUID, itemUUID string) (string, error) {
	var err error
	var item *onepassword.Item
	if onePasswordUUID.MatchString(itemUUID) {
		item, err = client.GetItem(itemUUID, vaultUUID)
		if err != nil {
			return "", err
		}
	} else {
		item, err = client.GetItemByTitle(itemUUID, vaultUUID)
		if err != nil {
			return "", err
		}
	}

	if item.Category != onepassword.Document {
		return "", errors.Errorf("unexpected document type: %s", item.Category)
	} else if len(item.Files) != 0 {
		return "", errors.Errorf("unexpected document files length: %d", len(item.Files))
	}

	res, err := client.GetFileContent(item.Files[0])
	if err != nil {
		return "", err
	}

	return strings.Trim(string(res), "\n"), nil
}
