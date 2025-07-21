package service

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	Shows   map[string]media `json:"shows"`
	Changes []struct {
		TimeStamp int64 `json:"timestamp"`
	} `json:"changes"`
	HasMore    bool   `json:"hasMore"`
	NextCursor string `json:"nextCursor"`
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
func (f *dataFetcher) fetch(nextCursor *string, lastUnixDate *int64) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://streaming-availability.p.rapidapi.com/changes", nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	q := req.URL.Query()
	q.Add("change_type", "new")
	q.Add("country", "es")
	q.Add("item_type", "show")
	q.Add("output_language", "es")
	q.Add("order_direction", "asc")
	q.Add("catalogs", "prime.subscription")
	if nextCursor != nil {
		q.Add("cursor", *nextCursor)
	}
	if lastUnixDate != nil {
		d := strconv.FormatInt(*lastUnixDate, 10)
		q.Add("from", d)
	}
	req.URL.RawQuery = q.Encode()

	res, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PERF: muchas conversiones de data res.body -> fetchedData -> types.Media -> realoc en `m`
func (f *dataFetcher) GetLastUpdates(lastUnixDate *int64) (*fetchedData, error) {
	var fd fetchedData
	var err error

	res, err := f.fetch(nil, lastUnixDate)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&fd)
	if err != nil {
		return nil, err
	}

	// FIX: descomentar cuando se guarden los datos en la db para evitar
	// hacer tantas peticiones a la API

	// fetchfor:
	// 	for {
	// 		var data fetchedData
	//
	// 		err = func() error {
	// 			res, err := f.fetch(&fd.NextCursor)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			defer res.Body.Close()
	// 			err = json.NewDecoder(res.Body).Decode(&data)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			return nil
	// 		}()
	//
	// 		maps.Copy(fd.Shows, data.Shows)
	// 		fd.NextCursor = data.NextCursor
	// 		fd.Changes = data.Changes
	//
	// 		if data.HasMore == false {
	// 			break fetchfor
	// 		}
	// 	}
	//
	// 	if err != nil {
	// 		return nil, err
	// 	}

	return &fd, nil
}
