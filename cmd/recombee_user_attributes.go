package cmd

import (
	"github.com/matt9mg/rawflix-api/services"
	"github.com/spf13/cobra"
	"log"
)

var recombeeUserAttributesCreateCmd = &cobra.Command{
	Use:   "recombee:user:attributes-create",
	Short: "Creates the attributes in recombee",
	Long:  "Creates the attributes via the api in recombee",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("creating user attributes")

		if err := recombee.UserProperties.AddUserProperty("country", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		if err := recombee.UserProperties.AddUserProperty("gender", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		log.Println("user attributes created")
	},
}

func init() {
	rootCmd.AddCommand(recombeeUserAttributesCreateCmd)
}
