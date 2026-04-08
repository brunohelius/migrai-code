package cmd

import (
	"fmt"

	"github.com/brunohelius/migrai-code/internal/auth"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with MigrAI using device flow",
	Long: `Authenticate with the MigrAI platform using the OAuth 2.0 Device
Authorization Grant (RFC 8628). This will open your browser to authorize
the CLI and store credentials locally.`,
	Example: `
  # Login with default server
  migrai-code login

  # Login with a custom server URL
  migrai-code login --server https://my-migrai.example.com
  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverURL, _ := cmd.Flags().GetString("server")

		fmt.Println("Authenticating with MigrAI...")

		creds, err := auth.DeviceFlowLogin(serverURL)
		if err != nil {
			return fmt.Errorf("login failed: %w", err)
		}

		if err := auth.SaveCredentials(creds); err != nil {
			return fmt.Errorf("failed to save credentials: %w", err)
		}

		fmt.Println()
		fmt.Println("  Successfully authenticated!")
		fmt.Printf("  Server:  %s\n", creds.ServerURL)
		fmt.Printf("  Model:   %s\n", creds.Model)
		fmt.Println()
		fmt.Println("  Credentials saved. You can now use migrai-code.")
		return nil
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored MigrAI credentials",
	Long:  `Remove locally stored MigrAI credentials, effectively logging out.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := auth.DeleteCredentials(); err != nil {
			return fmt.Errorf("logout failed: %w", err)
		}
		fmt.Println("Successfully logged out. Credentials removed.")
		return nil
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	Long:  `Display the current MigrAI authentication status and stored credentials.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		creds, err := auth.LoadCredentials()
		if err != nil {
			return fmt.Errorf("failed to load credentials: %w", err)
		}

		if creds == nil {
			fmt.Println("Not authenticated. Run 'migrai-code login' to authenticate.")
			return nil
		}

		fmt.Println("  Authentication status: logged in")
		fmt.Printf("  Server:     %s\n", creds.ServerURL)
		fmt.Printf("  Model:      %s\n", creds.Model)
		fmt.Printf("  Created at: %s\n", creds.CreatedAt)
		if creds.ExpiresAt != "" {
			fmt.Printf("  Expires at: %s\n", creds.ExpiresAt)
		}
		if creds.APIKey != "" {
			// Show only last 8 chars of the key for safety
			masked := creds.APIKey
			if len(masked) > 8 {
				masked = "..." + masked[len(masked)-8:]
			}
			fmt.Printf("  API key:    %s\n", masked)
		}
		return nil
	},
}

func init() {
	loginCmd.Flags().String("server", auth.DefaultServerURL,
		"MigrAI server URL")

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(statusCmd)
}
