package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPage(url string)(int,error){
	resp,err:=http.Get(url)
	if err!=nil{
		return 0,err
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		return 0,err
	}
    return len(body),nil
}


func main() {

   urls:=[]string{"http://www.google.com","http://www.yahoo.com","http://www.bing.com","http://bbc.co.uk"}

   for _,url :=range urls {
	   pageLength, err := GetPage(url)
	   if err != nil {

	   }
	   fmt.Printf("Page length %d \n", pageLength)
   }
}
