package onepassword

import (
	"context"
	"regexp"
	"strings"
	"sync"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/pkg/errors"
)

var (
	onePasswordCache = map[string]map[string]string{}
	onePasswordUUID  = regexp.MustCompile(`^[a-z0-9]{26}$`)
)

var (
	ErrNotFound    = errors.New("not found")
	ErrNotSignedIn = errors.New("not signed in")
)

var onePasswordGetLock sync.Mutex
var onePasswordGetDocumentLock sync.Mutex

func Secret(ctx context.Context, account, vaultUUID, itemUUID, field string) (string, error) {
	cacheKey := strings.Join([]string{account, vaultUUID, itemUUID}, "#")

	if _, ok := onePasswordCache[cacheKey]; !ok {
		switch {
		case IsConnect():
			client, err := connect.NewClientFromEnvironment()
			if err != nil {
				return "", err
			}
			if res, err := ConnectSecret(client, vaultUUID, itemUUID); err != nil {
				return "", err
			} else {
				onePasswordCache[cacheKey] = res
			}
		default:
			if res, err := CLiSecret(ctx, account, vaultUUID, itemUUID); err != nil {
				return "", err
			} else {
				onePasswordCache[cacheKey] = res
			}
		}
	}

	value, ok := onePasswordCache[cacheKey][field]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func Document(ctx context.Context, account, vaultUUID, itemUUID string) (string, error) {
	// create cache key
	cacheKey := strings.Join([]string{account, vaultUUID, itemUUID}, "#")

	if _, ok := onePasswordCache[cacheKey]; !ok {
		switch {
		case IsConnect():
			if client, err := connect.NewClientFromEnvironment(); err != nil {
				return "", err
			} else if res, err := ConnectDocument(client, vaultUUID, itemUUID); err != nil {
				return "", err
			} else {
				onePasswordCache[cacheKey] = map[string]string{"document": res}
			}
		default:
			if res, err := CLIDocument(ctx, account, vaultUUID, itemUUID); err != nil {
				return "", err
			} else {
				onePasswordCache[cacheKey] = map[string]string{"document": res}
			}
		}
	}

	value, ok := onePasswordCache[cacheKey]["document"]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}
