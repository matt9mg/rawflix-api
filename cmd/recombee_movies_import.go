package cmd

import (
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/utils"
	"github.com/spf13/cobra"
	"gorm.io/datatypes"
	"log"
	"strconv"
)

var recombeeMoviesImportCmd = &cobra.Command{
	Use:   "recombee:movies:import",
	Short: "Imports the movies to recombee",
	Long:  "Imports all the movies within the database to recombee that have not already been marked as synced",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("creating movie items")

		movies, err := movieRepo.FindWhereNotSyncedWithRecombee()

		if err != nil {
			log.Fatal(err)
		}

		for _, movie := range movies {
			itemId := strconv.Itoa(int(movie.ID))

			if err := recombee.Item.AddItem(itemId); err != nil {
				log.Fatal(err)
			}

			err = recombee.Item.SetItemValues(itemId, &types.RecombeeMovieItem{
				RecombeeCascadeCreate: &types.RecombeeCascadeCreate{
					CascadeCreate: true,
				},
				Actors:    fromJsonDataTypeToSliceString(movie.Actors),
				Country:   fromJsonDataTypeToSliceString(movie.Country),
				Directors: fromJsonDataTypeToSliceString(movie.Directors),
				Genres:    fromJsonDataTypeToSliceString(movie.Genre),
				Language:  fromJsonDataTypeToSliceString(movie.Language),
				Plot:      movie.Plot,
				Poster:    movie.Poster,
				Rated:     movie.Rated,
				Released:  movie.Released,
				Runtime:   movie.Runtime,
				Title:     movie.Title,
				Writers:   fromJsonDataTypeToSliceString(movie.Writers),
				Year:      movie.Year,
				Added:     movie.Added.Unix(),
			})

			if err != nil {
				log.Println(movie.ID)
				log.Fatal(err)
			}

			if err = movieRepo.MarkAsSyncedWithRecombee(movie); err != nil {
				log.Fatal(err)
			}
		}

		log.Println("movie items created")
	},
}

func init() {
	rootCmd.AddCommand(recombeeMoviesImportCmd)
}

func fromJsonDataTypeToSliceString(json datatypes.JSON) []string {
	sliceString, err := utils.FromJsonDataTypeToSliceString(json)

	if err != nil {
		log.Fatal(err)
	}

	return sliceString
}
