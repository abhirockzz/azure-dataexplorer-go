package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{Use: "azdatax", Short: "CLI to test sample program for Azure Data Explorer"}

const dbnameFlag = "dbname"

// Init serves as entrypoint for CLI
func Init() {
	rootCmd.Execute()
}
