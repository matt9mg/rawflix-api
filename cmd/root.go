package cmd

import (
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}
var userRepo repositories.UserRepository
var passwordHasher services.PasswordHasher

func Execute(uRepo repositories.UserRepository, pwHasher services.PasswordHasher) {
	userRepo = uRepo
	passwordHasher = pwHasher

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
