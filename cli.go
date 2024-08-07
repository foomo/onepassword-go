package onepassword

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type (
	cliItem struct {
		Vault  cliVault   `json:"vault"`
		Fields []cliField `json:"fields"`
	}
	cliVault struct {
		ID string `json:"id"`
	}
	cliField struct {
		ID    string `json:"id"`
		Type  string `json:"type"` // CONCEALED, STRING
		Label string `json:"label"`
		Value any    `json:"value"`
	}
)

func IsCLI(ctx context.Context) bool {
	if _, err := exec.LookPath("op"); err == nil {
		return true
	}
	return false
}

func CLiSecret(ctx context.Context, account, vaultUUID, itemUUID string) (map[string]string, error) {
	onePasswordGetLock.Lock()
	defer onePasswordGetLock.Unlock()

	var v cliItem
	res, err := exec.CommandContext(ctx, "op", "item", "get", itemUUID, "--vault", vaultUUID, "--account", account, "--format", "json").CombinedOutput()
	if err != nil && strings.Contains(string(res), "You are not currently signed in") {
		return nil, ErrNotSignedIn
	} else if err != nil {
		return nil, errors.Wrap(err, string(res))
	}

	if err := json.Unmarshal(res, &v); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal secret")
	}

	if v.Vault.ID != vaultUUID {
		return nil, errors.Errorf("wrong vault UUID %s for item %s", vaultUUID, itemUUID)
	}

	ret := map[string]string{}
	aliases := map[string]string{
		"notesPlain": "notes",
	}
	for _, field := range v.Fields {
		if alias, ok := aliases[field.Label]; ok {
			ret[alias] = fmt.Sprintf("%v", field.Value)
		} else {
			ret[field.Label] = fmt.Sprintf("%v", field.Value)
		}
	}
	return ret, nil
}

func CLIDocument(ctx context.Context, account, vaultUUID, itemUUID string) (string, error) {
	onePasswordGetDocumentLock.Lock()
	defer onePasswordGetDocumentLock.Unlock()

	res, err := exec.CommandContext(ctx, "op", "document", "get", itemUUID, "--vault", vaultUUID, "--account", account).CombinedOutput()
	if err != nil && strings.Contains(string(res), "You are not currently signed in") {
		return "", ErrNotSignedIn
	} else if err != nil {
		return "", errors.Wrap(err, string(res))
	}

	return strings.Trim(string(res), "\n"), nil
}
