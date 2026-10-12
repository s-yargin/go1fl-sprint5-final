package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		log.Println("Dataset is empty")
		return
	}
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("Error parsing data: %v", err)
			continue
		}
		actionInfo, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Error getting action info: %v", err)
			continue
		}
		fmt.Println(actionInfo)
	}
}
