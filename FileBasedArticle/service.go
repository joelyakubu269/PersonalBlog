package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"errors"
)
	const articlesDir= "articles"

func createAndStore(val ArticleData) error{
	
	file:= fmt.Sprintf("article%d.json",val.ID)
	 err:= os.Mkdir( articlesDir,0755) // version control removes empty files, make dir is a check that creates it and prevents it from failing
	if err!= nil {
		return fmt.Errorf("unable to create directory :%w" + err.Error(),err)
	}
	
	data,err:= json.MarshalIndent(val,"","  ")
	if err!= nil {
	 
	 return fmt.Errorf("unable to marshall :%w",err)
		
	}
	filepath:= filepath.Join(articlesDir,file)
	err = os.WriteFile(filepath,data,0644)
	if err!= nil {
		
		return fmt.Errorf("unable to write to file : %w",err)
	}
	return nil
}
func deleteArticle(id int) error{
	file:= fmt.Sprintf("article%d.json",id)
	filepath:= filepath.Join(articlesDir,file)
	err:= os.Remove(filepath)
	if err!= nil {
		if errors.Is(err,os.ErrNotExist) {
			return fmt.Errorf("cannot delete %s: the file does not exist",file)
		}
		return fmt.Errorf("failed to delete file %s : %w",file,err)
	}
	return nil
}

