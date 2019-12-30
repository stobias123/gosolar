// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stobias123/gosolar"
)

var cfgFile string
var orionUsername string
var orionPassword string
var orionHost string
var orionSSL bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "solarcmd",
	Short: "solarcmd is a CLI to interact with Solarwinds Orion API",
	Long:  `solarcmd is a CLI application that provides several common orion functions.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version is a CLI to interact with Solarwinds Orion API",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { fmt.Println("1.0") },
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("orion")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.solar.yaml)")

	rootCmd.PersistentFlags().StringVar(&orionHost, "server", "", "orion server")
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))

	rootCmd.PersistentFlags().StringVar(&orionUsername, "username", "", "orion username")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))

	rootCmd.PersistentFlags().StringVar(&orionPassword, "password", "", "orion password")
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))

	rootCmd.PersistentFlags().BoolVar(&orionSSL, "ssl", true, "orion use ssl")
	viper.BindPFlag("ssl", rootCmd.PersistentFlags().Lookup("ssl"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".solar" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".solar")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// GetClient is a convenience function to create a client object with provided strings.
func GetClient(cmd *cobra.Command, args []string) *gosolar.Client {
	return gosolar.NewClient(orionHost, orionUsername, orionPassword, orionSSL, true)
}
