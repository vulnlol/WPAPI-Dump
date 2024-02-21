package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type PostInfo struct {
	ID    int    `json:"id"`
	Title TitleInfo `json:"title"`
	Slug  string `json:"slug"`
}

type TitleInfo struct {
	Rendered string `json:"rendered"`
}

type PageInfo struct {
	ID    int    `json:"id"`
	Title TitleInfo `json:"title"`
	Slug  string `json:"slug"`
}

type CategoryInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TagInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type MediaInfo struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

type SettingInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Permalink   string `json:"permalink"`
}

func main() {
	baseURL := getUserInput("Enter the base URL of the WordPress site (e.g., https://example.com): ")

	fmt.Println("Choose the information you want to fetch:")
	fmt.Println("1. Users: Retrieve information about users.")
	fmt.Println("2. Posts: Retrieve information about posts.")
	fmt.Println("3. Pages: Retrieve information about pages.")
	fmt.Println("4. Categories: Retrieve information about categories.")
	fmt.Println("5. Tags: Retrieve information about tags.")
	fmt.Println("6. Media: Retrieve information about media files.")
	fmt.Println("7. Settings: Retrieve configuration settings for the WordPress site.")
	fmt.Println("8. All: Fetch all available information.")

	infoTypes := getUserInput("Enter the numbers corresponding to the information you want to fetch (e.g., '1 2 3', 'all' for all): ")

	var data []interface{}
	logData := make(map[string]interface{}) // Initialize logData map

	for _, choice := range strings.Split(infoTypes, " ") {
		switch choice {
		case "1":
			users := fetchUsers(baseURL + "/wp-json/wp/v2/users")
			data = append(data, users)
			logData["User Data"] = users
		case "2":
			posts := fetchPosts(baseURL + "/wp-json/wp/v2/posts")
			data = append(data, posts)
			logData["Post Data"] = posts
		case "3":
			pages := fetchPages(baseURL + "/wp-json/wp/v2/pages")
			data = append(data, pages)
			logData["Page Data"] = pages
		case "4":
			categories := fetchCategories(baseURL + "/wp-json/wp/v2/categories")
			data = append(data, categories)
			logData["Category Data"] = categories
		case "5":
			tags := fetchTags(baseURL + "/wp-json/wp/v2/tags")
			data = append(data, tags)
			logData["Tag Data"] = tags
		case "6":
			media := fetchMedia(baseURL + "/wp-json/wp/v2/media")
			data = append(data, media)
			logData["Media Data"] = media
		case "7":
			settings, err := fetchSettings(baseURL + "/wp-json/wp/v2/settings")
			if err != nil {
				fmt.Println("Error fetching settings info:", err)
				continue
			}
			data = append(data, settings)
			logData["Settings Data"] = settings
		case "8":
			users := fetchUsers(baseURL + "/wp-json/wp/v2/users")
			data = append(data, users)
			logData["User Data"] = users
			posts := fetchPosts(baseURL + "/wp-json/wp/v2/posts")
			data = append(data, posts)
			logData["Post Data"] = posts
			pages := fetchPages(baseURL + "/wp-json/wp/v2/pages")
			data = append(data, pages)
			logData["Page Data"] = pages
			categories := fetchCategories(baseURL + "/wp-json/wp/v2/categories")
			data = append(data, categories)
			logData["Category Data"] = categories
			tags := fetchTags(baseURL + "/wp-json/wp/v2/tags")
			data = append(data, tags)
			logData["Tag Data"] = tags
			media := fetchMedia(baseURL + "/wp-json/wp/v2/media")
			data = append(data, media)
			logData["Media Data"] = media
			settings, err := fetchSettings(baseURL + "/wp-json/wp/v2/settings")
			if err != nil {
				fmt.Println("Error fetching settings info:", err)
				continue
			}
			data = append(data, settings)
			logData["Settings Data"] = settings
		default:
			fmt.Println("Invalid choice:", choice)
		}
	}

	fmt.Println("Fetching information...")

	// Write fetched information to a log file
	logFileName := "wordpress_data.log"
	writeToLogFile(logFileName, logData)

	fmt.Println("Fetched information. See", logFileName, "for details.")
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + " ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func fetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &target)
}

func fetchUsers(url string) []UserInfo {
	var users []UserInfo
	err := fetchJSON(url, &users)
	if err != nil {
		fmt.Println("Error fetching user info:", err)
	}
	return users
}

func fetchPosts(url string) []PostInfo {
	var posts []PostInfo
	err := fetchJSON(url, &posts)
	if err != nil {
		fmt.Println("Error fetching post info:", err)
	}
	return posts
}

func fetchPages(url string) []PageInfo {
	var pages []PageInfo
	err := fetchJSON(url, &pages)
	if err != nil {
		fmt.Println("Error fetching page info:", err)
	}
	return pages
}

func fetchCategories(url string) []CategoryInfo {
	var categories []CategoryInfo
	err := fetchJSON(url, &categories)
	if err != nil {
		fmt.Println("Error fetching category info:", err)
	}
	return categories
}

func fetchTags(url string) []TagInfo {
	var tags []TagInfo
	err := fetchJSON(url, &tags)
	if err != nil {
		fmt.Println("Error fetching tag info:", err)
	}
	return tags
}

func fetchMedia(url string) []MediaInfo {
	var media []MediaInfo
	err := fetchJSON(url, &media)
	if err != nil {
		fmt.Println("Error fetching media info:", err)
	}
	return media
}

func fetchSettings(url string) (SettingInfo, error) {
	var settings SettingInfo
	err := fetchJSON(url, &settings)
	return settings, err
}

func writeToLogFile(filename string, data map[string]interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for section, item := range data {
		sectionTitle := strings.ToUpper(section)
		writer.WriteString(sectionTitle + ":\n\n")
		jsonData, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			fmt.Println("Error encoding JSON data:", err)
			continue
		}
		_, err = writer.Write(jsonData)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
		_, err = writer.WriteString("\n\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
	}
	writer.Flush()
}
