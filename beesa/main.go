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

// TODO LIST
// 1. Create HTTP Router using fast http or mux (https://deepsource.io/blog/go-web-frameworks/)
// 2. Create Dockerfile
// 3. Install redis on kubernetes or using the existing redis on digital ocean
// 4. Deploy by using jenkins to Kubernetes (using secret)
// 5.

func main() {
	link := ""
	id := ""
	fmt.Println("Please paste your link to be shorten below:")
	fmt.Scanf("%s", &link)
	fmt.Println("Specify your ID:")
	fmt.Scanf("%s", &id)
	resultLink := shortLink(link, id)
	fmt.Print(resultLink)

}
