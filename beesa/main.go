package main

import (
	"beesa/models"
	"beesa/redis"
	"beesa/urlgenerator"
	"fmt"
)

var baseURL string = "http://localhost:80/"

func shortLink(link string, userID string) string {
	generatedPath := urlgenerator.GenerateShortLink(link, userID)
	result := baseURL + generatedPath
	setData := models.Link{UserID: userID, OriginalLink: link, ShortLink: result}
	redis.HSetComposite(userID, generatedPath, setData)
	return result
}

func main() {
	// Redis healthcheck
	// redis.Ping()
	// rSetPipe()
	// rGet()
	// values := rGetAllKeys()
	// fmt.Print(values)
	// link := "https://tutorialedge.net/golang/go-redis-tutorial/a"
	link := ""
	id := ""
	fmt.Println("Please paste your link to be shorten below:")
	fmt.Scanf("%s", &link)
	fmt.Println("Specify your ID:")
	fmt.Scanf("%s", &id)
	resultLink := shortLink(link, id)
	fmt.Print(resultLink)

}
