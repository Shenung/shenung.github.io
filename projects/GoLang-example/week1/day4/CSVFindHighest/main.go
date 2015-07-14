package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type monument struct {
	station   string
	location  string
	elevation float64
}

func main() {
	f, err := os.Open("City_of_Champaign_GPS_Control_Points.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		if rowCount == 0 {
			info.setColumns(record)
		} else {
			station, err := info.parseState(record)
			if err != nil {
				log.Fatalln(err)
			}
			stateLookup[state.abbreviation] = state
		}
	}
}
