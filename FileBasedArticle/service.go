package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"errors"
)

func createAndStore(val ArticleData) error{
	
	file:= fmt.Sprintf("article%d.json",val.ID)
	err:= os.Mkdir("articles",0755)
	if err!= nil {
		return fmt.Errorf("unable to create directory :%w" + err.Error(),err)
	}
	
	data,err:= json.MarshalIndent(val,"","  ")
	if err!= nil {
	 
	 return fmt.Errorf("unable to marshall :%w",err)
		
	}
	filepath:= filepath.Join(fi)
	err = os.WriteFile(file,data,0644)
	if err!= nil {
		
		return fmt.Errorf("unable to write to file : %w",err)
	}
	return nil
}
func deleteArticle(id int) error{
	file:= fmt.Sprintf("article%d.json",id)
	err:= os.Remove(file)
	if err!= nil {
		if errors.Is(err,os.ErrNotExist) {
			return fmt.Errorf("cannot delete %s: the file does not exist",file)
		}
		return fmt.Errorf("failed to delete file %s : %w",file,err)
	}
	return nil
}
