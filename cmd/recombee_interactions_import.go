package cmd

import (
	"github.com/golang-module/carbon/v2"
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/utils"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"time"
)

var recombeeInteractionsImporteCmd = &cobra.Command{
	Use:   "recombee:interactions:import",
	Short: "Import the interactions",
	Long:  "Imports the interactions that have not been synced to recombee",
	Run: func(cmd *cobra.Command, args []string) {
		interactions, err := interactionRepo.FindAllNotSyncedWithRecombee()

		if err != nil {
			log.Fatal(err)
		}

		totalInteractions := len(interactions)

		log.Printf("We have a total of %d to sync", totalInteractions)

		for key, interaction := range interactions {
			err := recombee.UserItemInteraction.AddDetailView(&types.RecombeeUserItemInteraction{
				RecombeeCascadeCreate: &types.RecombeeCascadeCreate{CascadeCreate: true},
				UserID:                utils.UintToString(interaction.UserID),
				ItemID:                utils.UintToString(interaction.MovieID),
				TimeStamp:             randomDateInLast7DaysToUnix(),
			})

			if err != nil {
				log.Println(err)
			}

			if err = interactionRepo.MarkAsSynced(interaction); err != nil {
				log.Fatal(err)
			}

			if key%100 == 0 {
				log.Printf("We have synced %d/%d to sync", (key + 1), totalInteractions)
			}
		}

		log.Println("all interactions synced")
	},
}

func init() {
	rootCmd.AddCommand(recombeeInteractionsImporteCmd)
}

func randomDateInLast7DaysToUnix() int64 {
	rand.NewSource(time.Now().UnixNano())

	return carbon.Now().SubDays(rand.Intn((7 - 1) + 1)).ToStdTime().Unix()
}
