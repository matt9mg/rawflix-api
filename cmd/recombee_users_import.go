package cmd

import (
	"github.com/matt9mg/rawflix-api/types"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var recombeeUsersImportCmd = &cobra.Command{
	Use:   "recombee:users:import",
	Short: "Imports the users to recombee",
	Long:  "Imports all the users within the database to recombee that have not already been marked as synced",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := userRepo.FindWhereNotSyncedWithRecombee()

		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			userId := strconv.Itoa(int(user.ID))

			if err := recombee.User.AddUser(userId); err != nil {
				log.Fatal(err)
			}

			err = recombee.User.SetUserValues(userId, &types.RecombeeUserItem{
				RecombeeCascadeCreate: &types.RecombeeCascadeCreate{
					CascadeCreate: true,
				},
				Country: user.Country,
				Gender:  user.Gender,
			})

			if err != nil {
				log.Fatal(err)
			}

			if err = userRepo.MarkAsSyncedWithRecombee(user); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(recombeeUsersImportCmd)
}
