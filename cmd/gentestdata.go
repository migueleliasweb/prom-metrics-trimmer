package cmd

import (
	"log"
	"os"

	"github.com/migueleliasweb/prom-metrics-trimmer/pkg/prom"
	"github.com/spf13/cobra"
)

var genTestData = &cobra.Command{
	Use:   "gen-test-data",
	Short: "Generates the open metrics test data in 'data/'",
	Long: `Generates the open metrics test data in 'data/'.
	
If any changes are made to 'pkg/prom/backfill.go this cmd MUST be run to update the test data.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		f, fErr := os.OpenFile("./data/backfill.data", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

		if fErr != nil {
			log.Fatal(fErr)
		}
		_, backfilErr := prom.GenerateBackfillData(f)

		return backfilErr
	},
}

func init() {
	rootCmd.AddCommand(genTestData)
}
