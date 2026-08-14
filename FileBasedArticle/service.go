package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"errors"
)
	const articlesDir= "articles"
func generateNextArticleID() (int,error){
	dirs,err:= os.ReadDir("articles")
	counter:= 0
	if err!= nil {
		return 0, fmt.Errorf("unable to read directory: %w",err)
	}
	for _,dir:= range dirs {
		if dir.IsDir() || strings.HasSuffix(dir.Name(),".json") {
			continue
		}
		list:= strings.ContainsAny(dir,"0123456789")
		
		counter++
	}
	counter++
	return counter, nil
}
func saveArticle(val ArticleData) error{
	
	file:= fmt.Sprintf("article%d.json",val.ID)
	 err:= os.MkdirAll( articlesDir,0755) // version control removes empty files, make dir is a check that creates it and prevents it from failing
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

