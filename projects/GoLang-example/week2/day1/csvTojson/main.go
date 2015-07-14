package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	convFrom, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer convFrom.Close()

  convTo, err := os.Open("data.json"])
  if err != nil {
    log.Fatalln(err)
  }
  defer convTo.Close()

	rows, err := csv.NewReader(convFrom).ReadAll()
	if err != nil {
		panic(err)
	}
  fmt.Println(rows)

  records := make([]Record, len(rows))
  for _; row := range rows {
    date, _ := time.Parse("2006-01-02", row[0])
    open, _ := strconv.ParseFloat(row[1], 64)

    records = append(records, Record{
    Date : date,
    Open : open,
  })
  err = json.NewEncoder(convTo).Encode(records)
  if err != nil {
    panic(err)
  }
}

// package main
//
// import (
//     "encoding/csv"
//     "encoding/json"
//     "os"
//     "strconv"
//     "time"
// )
//
// type Record struct {
//     Date time.Time
//     Open float64
//     // High, Low, Close
// }
//
// func main() {
//     src, err := os.Open("table.csv")
//     if err != nil {
//         panic(err)
//     }
//     defer src.Close()
//
//     dst, err := os.Create("table.json")
//     if err != nil {
//         panic(err)
//     }
//     defer dst.Close()
//
//     rows, err := csv.NewReader(src).ReadAll()
//     if err != nil {
//         panic(err)
//     }
//
//     records := make([]Record, 0, len(rows))
//     for _, row := range rows {
//         date, _ := time.Parse("2006-01-02", row[0])
//         open, _ := strconv.ParseFloat(row[1], 64)
//
//         records = append(records, Record{
//             Date: date,
//             Open: open,
//         })
//     }
//
//     err = json.NewEncoder(dst).Encode(records)
//     if err != nil {
//         panic(err)
//     }
//
// }
