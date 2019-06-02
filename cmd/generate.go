package cmd

import (
	"github.com/spf13/cobra"
)

var (
	serviceType     string
	databaseDialect string
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate service app",
	Long:    `Generate service app`,
	Example: `generate --type=api`,
	RunE:    generate,
	Aliases: []string{"g"},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVar(&serviceType, "type", "", "Type of the service")
	generateCmd.Flags().StringVar(&databaseDialect, "db", "", "Generate database with dialect")
}

func generate(cmd *cobra.Command, args []string) error {

	return nil
}
