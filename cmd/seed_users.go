package cmd

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"time"
)

var countries = []string{
	"Brazil",
	"Croatia",
	"Denmark",
	"France",
	"Germany",
	"Moldova",
	"Poland",
	"Turkey",
	"United Kingdom",
	"United States",
}

const userBatches = 1000

var seedUsersCmd = &cobra.Command{
	Use:   "seed:users",
	Short: "Seed users in the database with data",
	Long:  "Seed users in the database with fake data",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("generating 100 users entries with random data")

		var users []*entities.User

		faker.SetGenerateUniqueValues(true)

		for i := 0; i < 100; i++ {
			pw, err := passwordHasher.HashPassword("rawnet100")

			if err != nil {
				log.Fatal(err)
			}

			users = append(users, &entities.User{
				Username: faker.Email(),
				Password: pw,
				Country:  entities.UserCountry(randomCountry()),
				Gender: entities.UserGender(faker.Gender(func(oo *options.Options) {
					oo.GenerateUniqueValues = false
				})),
			})

			log.Printf("user %d created\n", i+1)
		}

		if err := userRepo.CreateInBatches(userBatches, users...); err != nil {
			log.Fatal(err)
		}

		log.Println("users created")
	},
}

func init() {
	rootCmd.AddCommand(seedUsersCmd)
}

func randomCountry() string {
	rand.NewSource(time.Now().UnixNano())

	return countries[rand.Intn((10-1)+1)]
}
