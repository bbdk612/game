package items

import (
	"encoding/json"
	"log"
	"os"
)

type ItemData struct {
	Name      string
	ImagePath string
	Health    int
	ID        int
}

func JsonFileDecodeItem(filePath string) *ItemData {
	data, err := os.ReadFile(filePath)
	ItemD := &ItemData{}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &ItemD)
	return ItemD
}
