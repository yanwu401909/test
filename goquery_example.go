package main

import(
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
)

type Book struct{
	ImgUrl string
	Title string
	Url string 
	Author string 
	StarsNum float32
	Desc string
}

func main(){
	doc, err := goquery.NewDocument("http://yuedu.163.com/book/category/category/3000")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Print(doc.Text())
	doc.Find(".yd-book-content .yd-book-item").Each(func(i int, s *goquery.Selection){
		title := s.Find("a h2").Text()
		imgurl, ok := s.Find("a img").Attr("src")
		if(ok){
			fmt.Printf("Review %d : %s-%s \n", i, title, imgurl)
		}else{
			fmt.Printf("Review %d : %s-%s \n", i, title, "nil")
		}
	
	})
}