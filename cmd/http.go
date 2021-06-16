/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"go_cli/pkg/app"
	"go_cli/pkg/infra/repo"

	"github.com/spf13/cobra"
)

var userName string

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Github informations.",
	Long:  `Go CLI To Get Github User Informations.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo := repo.NewGetGithubUserDetails()
		app := app.NewGithubStatusApp(repo)

		app.ShowGithubUserDetails(userName)
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&userName, "user", "u", "ralvescosta", "Github user")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
