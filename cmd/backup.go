// Copyright © 2019 dbArchiver<CC.Yao>
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
	"github.com/ourcolour/dataarchiver/configs"
	"github.com/ourcolour/dataarchiver/services"
	"github.com/ourcolour/dataarchiver/services/impl"
	"github.com/spf13/cobra"
	"log"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup database",
	Long: `Backup database by mysqldump, e.g.:
backup -H "127.0.0.1" -P 3306 -u root -p 123456 -d northworld -o "./dmp --clearBinLogs"`,
	Run: func(cmd *cobra.Command, args []string) {
		// Arguments
		host, err := cmd.Flags().GetString("host")
		port, err := cmd.Flags().GetInt("port")
		user, err := cmd.Flags().GetString("user")
		pass, err := cmd.Flags().GetString("pass")
		dbName, err := cmd.Flags().GetString("dbname")
		tblName, err := cmd.Flags().GetString("tblname")
		outputDir, err := cmd.Flags().GetString("outputdir")
		compress, err := cmd.Flags().GetBool("compress")
		clearBinLogs, err := cmd.Flags().GetBool("clearbinlogs")
		debug, err := cmd.Flags().GetBool("debug")

		if nil != err {
			log.Panicf("%s\n", err.Error())
			return
		}

		var cmdOptions []string = nil
		if clearBinLogs {
			cmdOptions = []string{"--delete-master-logs", "--flush-logs"}
		}

		// 如果是 debug 模式，打印所有参数
		if debug {
			log.Printf("[DEBUG] arguments: host -> %s\n", host)
			log.Printf("[DEBUG] arguments: port -> %d\n", port)
			log.Printf("[DEBUG] arguments: user -> %s\n", user)
			log.Printf("[DEBUG] arguments: pass -> %s\n", pass)
			log.Printf("[DEBUG] arguments: dbName -> %s\n", dbName)
			log.Printf("[DEBUG] arguments: tblName -> %s\n", tblName)
			log.Printf("[DEBUG] arguments: outputDir -> %s\n", outputDir)
			log.Printf("[DEBUG] arguments: compress -> %v\n", compress)

			for idx, cmdOpt := range cmdOptions {
				log.Printf("[DEBUG] cmdOptions[%d] -> %s\n", idx, cmdOpt)
			}
		}

		var dumpSvs services.IDumpSvs = impl.NewMySqlDumpSvs()
		dumpSvs.Backup(host, port, user, pass, dbName, tblName, outputDir, compress, cmdOptions...)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringP("host", "H", "localhost", "Database host")
	backupCmd.Flags().StringP("port", "P", "3306", "Database port")
	backupCmd.Flags().StringP("user", "u", "root", "Username")
	backupCmd.Flags().StringP("pass", "p", "root", "Password")
	backupCmd.Flags().StringP("outputdir", "o", configs.DEFAULT_OUTPUT_DIR, "Output dir path")

	backupCmd.Flags().StringP("dbname", "d", "", "Database name")
	backupCmd.Flags().StringP("tblname", "t", "", "Table name")
	backupCmd.Flags().BoolP("compress", "c", false, "Compress dump file")

	backupCmd.Flags().Bool("clearbinlogs", true, "Delete master logs and rebuild bin-logs.")

	backupCmd.Flags().Bool("debug", false, "Enable debug mode.")
}
