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
var movieRepo repositories.MovieRepository
var recombee *services.Recoombe
var interactionRepo repositories.InteractionRepository

func Execute(uRepo repositories.UserRepository, pwHasher services.PasswordHasher, mvieRepo repositories.MovieRepository, recombeeApi *services.Recoombe, iRepo repositories.InteractionRepository) {
	userRepo = uRepo
	passwordHasher = pwHasher
	movieRepo = mvieRepo
	recombee = recombeeApi
	interactionRepo = iRepo

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
