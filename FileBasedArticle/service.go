package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func createAndStore(val ArticleData) error{
	new := strconv.Itoa(val.ID)
	file:="article"+ new + ".json"
	
	
	data,err:= json.MarshalIndent(val,""," ")
	if err!= nil {
	 
	 return fmt.Errorf("unable to marshall",err)
		
	}
	err = os.WriteFile(file,data,0644)
	if err!= nil {
		
		return fmt.Errorf("unable to write to file",err)
	}
	return nil
}
func delete(id int) error{
	file:= fmt.Sprintf("article%d.json",id)
	err:= os.Remove(file)
	if err!= nil {
		if errors.Is(err,Os.ErrNotExist) {
			return fmt.Errorf("cannot delete %s: the file does not exist",file)
		}
	}

}