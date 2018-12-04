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
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"xjh.com/dataarchiver/constants/errs"
	"xjh.com/dataarchiver/services"
	"xjh.com/dataarchiver/services/impl"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup database",
	Long: `Backup database by mysqldump, e.g.:
backup -H "127.0.0.1" -P 3306 -u root -p 123456 -d northworld -o "./dmp"
`,
	Run: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()
		port, err := strconv.Atoi(cmd.Flag("port").Value.String())
		if nil != err {
			log.Panicf("%s\n", errs.ERR_INVALID_PARAMETERS)
			return
		}
		user := cmd.Flag("user").Value.String()
		pass := cmd.Flag("pass").Value.String()
		dbName := cmd.Flag("dbname").Value.String()
		tableName := cmd.Flag("tablename").Value.String()
		outputDirPath := cmd.Flag("outputdir").Value.String()

		var dumpSvs services.IDumpSvs = impl.NewMySqlDumpSvs()
		dumpSvs.Backup(host, port, user, pass, dbName, tableName, outputDirPath)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringP("host", "H", "localhost", "Database host")
	backupCmd.Flags().StringP("port", "P", "3306", "Database port")
	backupCmd.Flags().StringP("user", "u", "root", "Username")
	backupCmd.Flags().StringP("pass", "p", "", "Password")
	backupCmd.Flags().StringP("dbname", "d", "", "Database name")
	backupCmd.Flags().StringP("tablename", "t", "", "Table name")
	backupCmd.Flags().StringP("outputdir", "o", "./", "Output dir path")
}
