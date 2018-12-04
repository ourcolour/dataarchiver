// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"strconv"
	"xjh.com/dataarchiver/services"
	"xjh.com/dataarchiver/services/impl"

	"github.com/spf13/cobra"
)

// clearoldCmd represents the clearold command
var clearoldCmd = &cobra.Command{
	Use:   "clearold",
	Short: "Clear the old archives.",
	Long: `Clear the old archives which over {OVER_DAYS}, e.g.:
clearold -o "./dmp" -d 31
`,
	Run: func(cmd *cobra.Command, args []string) {
		outputDir := cmd.Flag("outputdir").Value.String()
		overDays, err := strconv.Atoi(cmd.Flag("overdays").Value.String())
		if nil != err {
			log.Panicf("%s\n", err)
			return
		}

		var dumpSvs services.IDumpSvs = impl.NewMySqlDumpSvs()
		dumpSvs.DeleteOldArchiveByOverDayCount(outputDir, overDays)
	},
}

func init() {
	rootCmd.AddCommand(clearoldCmd)

	clearoldCmd.Flags().StringP("outputdir", "o", "./", "Output dir path")
	clearoldCmd.Flags().StringP("overdays", "d", "31", "Over days count")
}
