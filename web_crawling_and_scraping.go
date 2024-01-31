package main

// Import the appropriate packages
import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Define the webpages that we would like to scrape
	webpages := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	for _, pageURL := range webpages {
		// Define the appropriate input variables to store scraped information
		c := colly.NewCollector()

		var pageTitle string
		var pageContent strings.Builder

		// Extract the webpage title from the webpage
		c.OnHTML("html", func(e *colly.HTMLElement) {
			pageTitle = e.ChildText("title")
		})

		// Extract the text from the body of the webpage
		c.OnHTML(".mw-parser-output p", func(e *colly.HTMLElement) {
			pageContent.WriteString(e.Text + " ")
		})

		// Compile the JSON using all the webpage data that we've scraped
		c.OnScraped(func(r *colly.Response) {
			data := map[string]interface{}{
				"page_url":   pageURL,
				"page_title": pageTitle,
				"text":       strings.TrimSpace(pageContent.String()),
			}

			// Convert data to JSON format
			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}

			// Print the JSON data
			//			println(string(jsonData))

			// Save the JSON data to an output .JL file
			saveToFile(jsonData, "scrapedOutputGo.jl")
		})

		// Start the scraping for the current Wikipedia page
		err := c.Visit(pageURL)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func saveToFile(data []byte, filename string) {
	// Export the scraped data to an output file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(append(data, []byte("\n")...))
	if err != nil {
		log.Fatal(err)
	}
}
