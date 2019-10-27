/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	//"fmt"
	"github.com/spf13/agenda/service"
	"github.com/spf13/cobra"
)

// logInCmd represents the logIn command
var logInCmd = &cobra.Command{
	Use:   "logIn -u [username] -p [password]",
	Short: "User log in",
	Long: `Input like: logIn -u Go -p 123456`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		tmp_n, _ := cmd.Flags().GetString("name")
		tmp_p, _ := cmd.Flags().GetString("password")
		service.LogIn(tmp_n, tmp_p)
	},
}

func init() {
	rootCmd.AddCommand(logInCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logInCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logInCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	logInCmd.Flags().StringP("name", "u", "", "user name")
	logInCmd.Flags().StringP("password", "p", "", "user password")
}
