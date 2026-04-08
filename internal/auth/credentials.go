// Package auth handles authentication for the MigrAI Code CLI,
// including OAuth 2.0 Device Authorization Grant (RFC 8628) and
// credential storage.
package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Constants for the MigrAI authentication system.
const (
	DefaultServerURL = "https://app.migrai.com.br"
	DefaultModel     = "migrai/minimax-m2.7"

	credentialsDir  = ".migrai-code"
	credentialsFile = "credentials.json"
)

// Credentials holds the authentication state persisted to disk.
type Credentials struct {
	APIKey      string `json:"api_key"`
	ServerURL   string `json:"server_url"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at,omitempty"`
	Model       string `json:"default_model,omitempty"`
	PlatformKey string `json:"platform_key,omitempty"`
}

// credentialsPath returns the full path to the credentials file.
func credentialsPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, credentialsDir, credentialsFile), nil
}

// SaveCredentials writes credentials to ~/.migrai-code/credentials.json with
// mode 0600 so only the current user can read them.
func SaveCredentials(creds *Credentials) error {
	path, err := credentialsPath()
	if err != nil {
		return err
	}

	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("failed to create credentials directory: %w", err)
	}

	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("failed to write credentials file: %w", err)
	}

	return nil
}

// LoadCredentials reads credentials from ~/.migrai-code/credentials.json.
// Returns nil, nil when the file does not exist (no credentials stored).
func LoadCredentials() (*Credentials, error) {
	path, err := credentialsPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read credentials file: %w", err)
	}

	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("failed to parse credentials file: %w", err)
	}

	return &creds, nil
}

// DeleteCredentials removes the credentials file (logout).
func DeleteCredentials() error {
	path, err := credentialsPath()
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete credentials file: %w", err)
	}

	return nil
}
