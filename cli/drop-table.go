package cli

import (
	"github.com/abhirockzz/datax-go/cli/util"
	"github.com/abhirockzz/datax-go/setup"

	"github.com/spf13/cobra"
)

// azdatax drop-table --dbname testdb
var dropTableCLICommand = cobra.Command{Use: "drop-table", Short: "drops the StormEvents table", Long: "drops the StormEvents table from the database passed using --dbname", Run: dropTable}

func init() {
	dropTableCLICommand.Flags().String(dbnameFlag, "", "Database name in Azure Data Explorer")
	dropTableCLICommand.MarkFlagRequired(dbnameFlag)
	rootCmd.AddCommand(&dropTableCLICommand)
}

func dropTable(c *cobra.Command, args []string) {
	dbname := c.Flag(dbnameFlag).Value.String()
	setup.DropTable(util.Connect(), dbname)
}
