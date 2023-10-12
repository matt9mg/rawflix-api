package cmd

import (
	"embed"
	"encoding/csv"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/utils"
	"github.com/spf13/cobra"
	"log"
	"time"
)

//go:embed data
var data embed.FS

const movieBatch = 1000

var seedMoviesCmd = &cobra.Command{
	Use:   "seed:movies",
	Short: "Seed movies in the database with data",
	Long:  "Seed movies in the database with data from 3 csvs",
	Run: func(cmd *cobra.Command, args []string) {
		var movies []*entities.Movie

		movies1, err := getMovieData("data/movies.csv")

		if err != nil {
			log.Fatal(err)
		}

		movies2, err := getMovieData("data/movies2.csv")

		if err != nil {
			log.Fatal(err)
		}

		movies3, err := getMovieData("data/movies3.csv")

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movies1...)
		movies = append(movies, movies2...)
		movies = append(movies, movies3...)

		if err := movieRepo.CreateInBatches(movieBatch, movies...); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedMoviesCmd)
}

func toJsonB(data string) []byte {
	datum, err := utils.ToJsonB(data, ",")

	if err != nil {
		log.Fatal(err)
	}

	return datum
}

func getMovieData(filename string) ([]*entities.Movie, error) {
	movieCsv, err := data.Open(filename)

	if err != nil {
		return nil, err
	}

	rows, err := csv.NewReader(movieCsv).ReadAll()

	if err != nil {
		return nil, err
	}

	var movies []*entities.Movie

	for _, row := range rows {
		// if we don't have a valid poster image then I don't want it
		if row[13] == "" {
			continue
		}

		dt, err := time.Parse("02 Jan 2006", row[3])

		if err != nil {
			dt = time.Now()
		}

		if dt.Unix() < 0 {
			dt = time.Now()
		}

		movies = append(movies, &entities.Movie{
			Title:     row[0],
			Year:      row[1],
			Rated:     row[2],
			Released:  row[3],
			Runtime:   row[4],
			Genre:     toJsonB(row[5]),
			Directors: toJsonB(row[6]),
			Writers:   toJsonB(row[7]),
			Actors:    toJsonB(row[8]),
			Plot:      row[9],
			Language:  toJsonB(row[10]),
			Country:   toJsonB(row[11]),
			Awards:    row[12],
			Poster:    row[13],
			Added:     dt,
		})
	}

	return movies, nil
}
