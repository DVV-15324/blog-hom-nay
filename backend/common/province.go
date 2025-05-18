package common

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Checkprovince(province string, district string, ward string) error {
	dir, _ := os.Getwd()
	data, err := os.Open(fmt.Sprintf("%s%s", dir, "/common/province/province.csv"))
	if err != nil {
		return err
	}
	defer data.Close()
	reader := csv.NewReader(data)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, d := range records {
		if d[2] == ward && d[1] == district && d[0] == province {
			return nil
		}
	}
	return nil
}
