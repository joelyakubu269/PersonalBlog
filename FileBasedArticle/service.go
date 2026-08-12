package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func createAndStore(val ArticleData) {
	id:= val.ID
	new:= string(id)
	file,err:= os.Create("article"+ new + ".json")
	if err!= nil {
		fmt.Errorf("unable to read file",err)
		return
	}
	defer file.Close()
	data,err= json.MarshalIndent(val,""," ")
	if err!= nil {
		fmt.Errorf("unable to marshall",err)
		return
	}
	
}