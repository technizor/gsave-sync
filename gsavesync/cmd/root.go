package cmd

import (
	"fmt"
	"os"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	gss "github.com/technizor/gsavesync"
	"github.com/technizor/gsavesync/config"
)

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "gsavesync",
		Short: "gsavesync is a game save manager",
		Long: `A useless game save manager built
				with Golang.`,
	}
	hashCmd = &cobra.Command{
		Use:   "hash [file]",
		Short: "Get the SHA-256 hash of a file",
		Long:  `Get the SHA-256 hash of a file in hexadecimal`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]

			fhash := gss.Filehash(filePath)
			os.Stdout.WriteString(fhash)
		},
	}
	saveCmd = &cobra.Command{
		Use:   "save [command]",
		Short: "",
		Long:  ``,
	}
	saveReadCmd = &cobra.Command{
		Use:   "read [file]",
		Short: "Read a save config file",
		Long:  `Read a save config file`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fpath := args[0]

			config.ReadConfig(fpath)
		},
	}
	saveInitCmd = &cobra.Command{
		Use:   "init [file]",
		Short: "Read a save config file",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fpath := args[0]
			v := config.Config{
				Games:        []config.Game{},
				SaveSettings: config.DefaultSaveSettings(),
			}
			config.WriteConfig(fpath, v)
		},
	}
)

// Execute cobra CLI
func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	fmt.Println(runtime.GOOS)
	rootCmd.AddCommand(hashCmd)
	rootCmd.AddCommand(saveCmd)
	saveCmd.AddCommand(saveInitCmd)
	saveCmd.AddCommand(saveReadCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".gsavesync" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gsavesync")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
