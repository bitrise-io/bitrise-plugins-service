package cmd

import (
	"github.com/bitrise-io/bitrise-plugins-service/generators"
	"github.com/spf13/cobra"
)

var (
	configFilePath  string
	serviceType     string
	databaseDialect string
	projectPath     string
	aws             bool
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
	generateCmd.Flags().StringVar(&configFilePath, "config", "", "Path of config file")
	generateCmd.Flags().StringVar(&serviceType, "type", "", "Type of the service")
	generateCmd.Flags().StringVar(&databaseDialect, "db", "", "Generate database with dialect")
	generateCmd.Flags().StringVar(&projectPath, "path", "", "Source controll path of your project (e.g. github.com/my-org/my-project)")
	generateCmd.Flags().BoolVar(&aws, "aws", false, "Generate AWS specific components")
}

func generate(cmd *cobra.Command, args []string) error {
	if serviceType == "api" {
		return generators.GenerateAPI(generators.Config{
			ProjectPath:    projectPath,
			DBDialect:      databaseDialect,
			AWS:            aws,
			ConfigFilePath: configFilePath,
		})
	}
	return nil
}
