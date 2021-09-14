package pp

import (
	"encoding/json"
	"fmt"
)

func Print(i interface{}) {
	bs, err := json.MarshalIndent(i, "", "        ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
