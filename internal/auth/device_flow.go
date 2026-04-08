package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// deviceAuthResponse is the JSON returned by POST /oauth/device/authorize.
type deviceAuthResponse struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

// tokenResponse is the JSON returned by POST /oauth/device/token on success.
type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Error       string `json:"error,omitempty"`
}

// llmKeyResponse is the JSON returned by GET /api/keys/llm/byor.
type llmKeyResponse struct {
	Key   string `json:"key"`
	Model string `json:"model"`
}

// DeviceFlowLogin performs the full OAuth 2.0 Device Authorization Grant flow.
//
//  1. POST to /oauth/device/authorize to start the flow.
//  2. Display the user code and verification URI.
//  3. Attempt to open the browser automatically.
//  4. Poll /oauth/device/token until approval, denial, or expiry.
//  5. On success, fetch the LLM key via /api/keys/llm/byor.
//  6. Return stored Credentials.
func DeviceFlowLogin(serverURL string) (*Credentials, error) {
	if serverURL == "" {
		serverURL = DefaultServerURL
	}
	serverURL = strings.TrimRight(serverURL, "/")

	// --- Step 1: Start device authorization ---
	authResp, err := startDeviceAuthorization(serverURL)
	if err != nil {
		return nil, fmt.Errorf("failed to start device authorization: %w", err)
	}

	// --- Step 2: Display instructions ---
	fmt.Println()
	fmt.Printf("  To authenticate, visit:\n")
	fmt.Printf("  %s\n\n", authResp.VerificationURIComplete)
	fmt.Printf("  And confirm the code: %s\n\n", authResp.UserCode)

	// --- Step 3: Open browser (best-effort) ---
	openBrowser(authResp.VerificationURIComplete)

	// --- Step 4: Poll for token ---
	interval := authResp.Interval
	if interval < 1 {
		interval = 5
	}

	fmt.Printf("  Waiting for authorization...")

	token, err := pollForToken(serverURL, authResp.DeviceCode, interval, authResp.ExpiresIn)
	if err != nil {
		fmt.Println(" failed.")
		return nil, err
	}
	fmt.Println(" done!")

	// --- Step 5: Fetch LLM key ---
	model := DefaultModel
	apiKey := token.AccessToken

	llmKey, err := fetchLLMKey(serverURL, token.AccessToken)
	if err != nil {
		// Non-fatal: we fall back to using the access token directly
		fmt.Printf("  Note: could not fetch LLM key (%v), using access token.\n", err)
	} else if llmKey != nil {
		if llmKey.Key != "" {
			apiKey = llmKey.Key
		}
		if llmKey.Model != "" {
			model = llmKey.Model
		}
	}

	// --- Step 6: Build and return credentials ---
	creds := &Credentials{
		APIKey:      apiKey,
		ServerURL:   serverURL,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		Model:       model,
		PlatformKey: token.AccessToken,
	}

	return creds, nil
}

// startDeviceAuthorization calls POST /oauth/device/authorize.
func startDeviceAuthorization(serverURL string) (*deviceAuthResponse, error) {
	url := serverURL + "/oauth/device/authorize"

	resp, err := http.Post(url, "application/json", bytes.NewReader([]byte("{}")))
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned %d: %s", resp.StatusCode, string(body))
	}

	var authResp deviceAuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if authResp.DeviceCode == "" {
		return nil, fmt.Errorf("server returned empty device_code")
	}

	return &authResp, nil
}

// pollForToken polls POST /oauth/device/token until success, denial, or expiry.
func pollForToken(serverURL, deviceCode string, interval, expiresIn int) (*tokenResponse, error) {
	url := serverURL + "/oauth/device/token"
	deadline := time.Now().Add(time.Duration(expiresIn) * time.Second)

	payload := map[string]string{
		"device_code": deviceCode,
		"grant_type":  "urn:ietf:params:oauth:grant-type:device_code",
	}
	payloadBytes, _ := json.Marshal(payload)

	for {
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("device authorization expired (waited %d seconds)", expiresIn)
		}

		time.Sleep(time.Duration(interval) * time.Second)

		resp, err := http.Post(url, "application/json", bytes.NewReader(payloadBytes))
		if err != nil {
			// Transient network error, keep polling
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			continue
		}

		var tokenResp tokenResponse
		if err := json.Unmarshal(body, &tokenResp); err != nil {
			continue
		}

		// Check for error states per RFC 8628
		switch tokenResp.Error {
		case "authorization_pending":
			// User has not yet approved; keep polling
			continue
		case "slow_down":
			// Back off by adding 5 seconds to the interval
			interval += 5
			continue
		case "access_denied":
			return nil, fmt.Errorf("authorization was denied by the user")
		case "expired_token":
			return nil, fmt.Errorf("device code has expired, please try again")
		case "":
			// No error — check if we got a token
			if tokenResp.AccessToken != "" {
				return &tokenResp, nil
			}
			// Empty error and empty token: keep polling
			continue
		default:
			return nil, fmt.Errorf("authorization failed: %s", tokenResp.Error)
		}
	}
}

// fetchLLMKey retrieves the user's LLM key from the platform.
func fetchLLMKey(serverURL, accessToken string) (*llmKeyResponse, error) {
	url := serverURL + "/api/keys/llm/byor"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned %d: %s", resp.StatusCode, string(body))
	}

	var llmResp llmKeyResponse
	if err := json.Unmarshal(body, &llmResp); err != nil {
		return nil, fmt.Errorf("failed to parse LLM key response: %w", err)
	}

	return &llmResp, nil
}

// openBrowser tries to open a URL in the user's default browser.
// Failures are silently ignored — the user can always copy-paste the URL.
func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return
	}

	_ = cmd.Start()
}
