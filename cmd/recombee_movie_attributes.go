package cmd

import (
	"github.com/matt9mg/rawflix-api/services"
	"github.com/spf13/cobra"
	"log"
)

var recombeeMovieAttributesCreateCmd = &cobra.Command{
	Use:   "recombee:movie:attributes-create",
	Short: "Creates the attributes in recombee",
	Long:  "Creates the attributes via the api in recombee",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("creating product attributes")
		if err := recombee.ItemProperties.AddItemProperty("title", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("year", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("rated", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("released", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("runtime", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("genres", services.RecombeePropertyTypeSet); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("writers", services.RecombeePropertyTypeSet); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("actors", services.RecombeePropertyTypeSet); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("plot", services.RecombeePropertyTypeString); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("language", services.RecombeePropertyTypeSet); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("country", services.RecombeePropertyTypeSet); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("poster", services.RecombeePropertyTypeImage); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("directors", services.RecombeePropertyTypeSet); err != nil {
			log.Fatal(err)
		}

		if err := recombee.ItemProperties.AddItemProperty("added", services.RecombeePropertyTypeTimestamp); err != nil {
			log.Fatal(err)
		}

		log.Println("product attributes created")
	},
}

func init() {
	rootCmd.AddCommand(recombeeMovieAttributesCreateCmd)
}
