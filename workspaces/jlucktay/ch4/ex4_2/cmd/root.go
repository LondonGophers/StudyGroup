package cmd

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ex4_2",
	Short: "Exercise 4.2 from the Go Programming Language textbook (https://www.gopl.io)",
	Long: `Exercise 4.2

Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print
the SHA384 or SHA512 hash instead.

See page 84 of the Go Programming Language textbook (https://www.gopl.io).`,

	Run: func(cmd *cobra.Command, args []string) {
		for index := 0; index < len(args); index++ {
			fmt.Printf("SHA256 of '%s': %x\n", args[index], sha256.Sum256([]byte(args[index])))

			if cmd.Flag("sha384").Value.String() == "true" {
				fmt.Printf("SHA384 of '%s': %x\n", args[index], sha512.Sum384([]byte(args[index])))
			}

			if cmd.Flag("sha512").Value.String() == "true" {
				fmt.Printf("SHA512 of '%s': %x\n", args[index], sha512.Sum512([]byte(args[index])))
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
// It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ex4_2.yaml)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	rootCmd.Flags().BoolP("sha384", "3", false, "print the SHA384 hash")
	rootCmd.Flags().BoolP("sha512", "5", false, "print the SHA512 hash")
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

		// Search config in home directory with name ".ex4_2" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ex4_2")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
