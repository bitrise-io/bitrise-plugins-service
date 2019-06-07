package cmd

import (
	"github.com/bitrise-io/bitrise-plugins-service/generators"
	"github.com/spf13/cobra"
)

var (
	serviceType     string
	databaseDialect string
	projectPath     string
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
	generateCmd.Flags().StringVar(&projectPath, "path", "", "Source controll path of your project (e.g. github.com/my-org/my-project)")
}

func generate(cmd *cobra.Command, args []string) error {
	if serviceType == "api" {
		return generators.GenerateAPI(projectPath, databaseDialect == "postgres")
	}
	return nil
}
