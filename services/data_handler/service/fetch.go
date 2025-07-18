package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/condemo/movie-hub/services/common/types"
)

type media struct {
	Type   string `json:"showType"`
	Title  string `json:"title"`
	Desc   string `json:"overview"`
	Year   int32  `json:"releaseYear"`
	Genres []struct {
		Name string `json:"name"`
	} `json:"genres"`
	Rating int32 `json:"rating"`
	Images struct {
		Vertical struct {
			Poster string `json:"w720"`
		} `json:"verticalPoster"`
	} `json:"imageSet"`
	Seasons int32 `json:"seasonCount"`
	Caps    int32 `json:"episodeCount"`
}

type fetchedData struct {
	Shows map[string]media `json:"shows"`
}

func (fd fetchedData) getShowList() []types.Media {
	md := make([]types.Media, len(fd.Shows))
	i := 0
	for _, d := range fd.Shows {
		var m types.Media
		m.Type = d.Type
		m.Title = d.Title
		m.Year = d.Year
		m.Description = d.Desc
		m.Rating = d.Rating
		m.Image = d.Images.Vertical.Poster

		// TODO: genres
		gl := make([]string, len(d.Genres))
		for i, g := range d.Genres {
			gl[i] = g.Name
		}
		gs := strings.Join(gl, ",")
		m.Genres = gs

		if d.Type == "series" {
			m.Seasons = d.Seasons
			m.Caps = d.Caps
		}

		md[i] = m
		i++
	}
	return md
}

type dataFetcher struct {
	httpClient *http.Client
}

func newDataFetcher() *dataFetcher {
	return &dataFetcher{
		httpClient: &http.Client{},
	}
}

// TODO:
func (f *dataFetcher) GetLastUpdates() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, "https://streaming-availability.p.rapidapi.com/changes", nil)
	if err != nil {
		// FIX: gestionar apropiadamente
		log.Fatal(err)
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	q := req.URL.Query()
	q.Add("change_type", "new")
	q.Add("country", "es")
	q.Add("item_type", "show")
	q.Add("output_language", "es")
	q.Add("order_direction", "asc")
	// TODO: añadir un q param que indique la última fecha de la que se pidió info
	// ej: 	q.Add("from", "[unix-time-stamp]")
	q.Add("catalogs", "prime.subscription")
	req.URL.RawQuery = q.Encode()

	res, err := f.httpClient.Do(req)
	if err != nil {
		// FIX: gestionar apropiadamente
		log.Fatal(err)
	}
	defer res.Body.Close()

	var data fetchedData
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		// FIX: gestionar apropiadamente
		log.Fatal(err)
	}

	// TODO: borrar y devolver data parseada
	for _, d := range data.getShowList() {
		fmt.Println(d.Title, "-", d.Genres, "-", d.Seasons, d.Caps)
	}
}

// TODO:
func (f *dataFetcher) GetLastUpdatesWithCursor() {}
