package persistant

import (
	"encoding/json"
	"log"
	"os"

	"github.com/condemo/movie-hub/services/common/config"
)

var RequestData = newReqData()

type reqData struct {
	LastMediaDate *int64 `json:"lastMediaDate"`
}

func newReqData() reqData {
	rd := reqData{}
	if _, err := os.Stat(config.DefaultPaths.DataFile); os.IsNotExist(err) {
		f, err := os.Create(config.DefaultPaths.DataFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		rd.LastMediaDate = nil
		err = json.NewEncoder(f).Encode(&rd)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		f, err := os.Open(config.DefaultPaths.DataFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = json.NewDecoder(f).Decode(&rd)
		if err != nil {
			log.Fatal(err)
		}
	}

	return rd
}

func (rd reqData) Save() error {
	f, err := os.Create(config.DefaultPaths.DataFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(rd)
	if err != nil {
		return err
	}

	return nil
}
