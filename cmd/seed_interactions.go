package cmd

import (
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"time"
)

const batchSize = 1000

var seedInteractionCmd = &cobra.Command{
	Use:   "seed:interactions",
	Short: "Creates the interactions data",
	Long:  "Creates the interactions data for every user with random movies",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := userRepo.FindAll()

		if err != nil {
			log.Fatal(err)
		}

		comedyMovies, err := movieRepo.FindMoviesWithGenre("Comedy")

		if err != nil {
			log.Fatal(err)
		}

		dramaMovies, err := movieRepo.FindMoviesWithGenre("Drama")

		if err != nil {
			log.Fatal(err)
		}

		actionMovies, err := movieRepo.FindMoviesWithGenre("Action")

		if err != nil {
			log.Fatal(err)
		}

		crimeMovies, err := movieRepo.FindMoviesWithGenre("Crime")

		if err != nil {
			log.Fatal(err)
		}

		thrillerMovies, err := movieRepo.FindMoviesWithGenre("Thriller")

		if err != nil {
			log.Fatal(err)
		}

		familyMovies, err := movieRepo.FindMoviesWithGenre("Family")

		if err != nil {
			log.Fatal(err)
		}

		scifiMovies, err := movieRepo.FindMoviesWithGenre("Sci-Fi")

		if err != nil {
			log.Fatal(err)
		}

		var interactions []*entities.Interaction

		for k, user := range users {
			if k < 20 {
				interactions = append(interactions, buildInteractions(comedyMovies, user)...)
				continue
			}

			if k > 19 && k < 25 {
				interactions = append(interactions, buildInteractions(dramaMovies, user)...)
				continue
			}

			if k > 24 && k < 50 {
				interactions = append(interactions, buildInteractions(actionMovies, user)...)
				continue
			}

			if k > 49 && k < 55 {
				interactions = append(interactions, buildInteractions(crimeMovies, user)...)
				continue
			}

			if k > 54 && k < 60 {
				interactions = append(interactions, buildInteractions(thrillerMovies, user)...)
				continue
			}

			if k > 59 && k < 80 {
				interactions = append(interactions, buildInteractions(scifiMovies, user)...)
				continue
			}

			interactions = append(interactions, buildInteractions(familyMovies, user)...)
		}

		if err = interactionRepo.CreateInBatches(batchSize, interactions...); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedInteractionCmd)
}

func buildInteractions(movies []*entities.Movie, user *entities.User) []*entities.Interaction {
	var interactions []*entities.Interaction

	for _, movie := range rando10mMovies(movies) {
		interactions = append(interactions, &entities.Interaction{
			UserID:  user.ID,
			MovieID: movie.ID,
		})
	}

	return interactions
}

func rando10mMovies(movies []*entities.Movie) []*entities.Movie {
	var interactionsMovies []*entities.Movie

	for i := 0; i < 10; i++ {
		rand.NewSource(time.Now().UnixNano())
		interactionsMovies = append(interactionsMovies, movies[rand.Intn((len(movies)-1)+1)])
	}

	return interactionsMovies
}
