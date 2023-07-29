package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Function to fetch the HTML content of a URL
func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching %s: %v", url, err)
	}
	defer resp.Body.Close()

	return html.Parse(resp.Body)
}

// Function to extract SEO content from "meta" tags in the HTML node
func extractSEOContent(node *html.Node) []string {
	var seoContent []string

	if node.Type == html.ElementNode && node.Data == "meta" {
		// Check if the "meta" tag contains a "name" attribute with a value of "description" (common for SEO)
		if nameAttr := getAttribute(node, "name"); nameAttr == "description" {
			if contentAttr := getAttribute(node, "content"); contentAttr != "" {
				seoContent = append(seoContent, "Description: "+contentAttr)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		seoContent = append(seoContent, extractSEOContent(child)...)
	}

	return seoContent
}

// Function to extract the title of the website from the HTML node
func extractTitle(node *html.Node) string {
	var title string

	if node.Type == html.ElementNode && node.Data == "title" {
		if node.FirstChild != nil {
			title = node.FirstChild.Data
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if title != "" {
			break
		}
		title = extractTitle(child)
	}

	return title
}

// Function to extract headings (h1, h2, h3, etc.) from the HTML node
func extractHeadings(node *html.Node) []string {
	var headings []string

	if node.Type == html.ElementNode && strings.HasPrefix(node.Data, "h") && len(node.Data) == 2 {
		if node.FirstChild != nil {
			headings = append(headings, fmt.Sprintf("%s: %s", node.Data, node.FirstChild.Data))
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		headings = append(headings, extractHeadings(child)...)
	}

	return headings
}

// Function to extract image sources from the HTML node
func extractImages(node *html.Node) []string {
	var images []string

	if node.Type == html.ElementNode && node.Data == "img" {
		if srcAttr := getAttribute(node, "src"); srcAttr != "" {
			images = append(images, "Image Source: "+srcAttr)
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		images = append(images, extractImages(child)...)
	}

	return images
}

// Function to extract links from the HTML node
func extractLinks(node *html.Node) []string {
	var links []string

	if node.Type == html.ElementNode && node.Data == "a" {
		if hrefAttr := getAttribute(node, "href"); hrefAttr != "" {
			links = append(links, "Link: "+hrefAttr)
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = append(links, extractLinks(child)...)
	}

	return links
}

// Helper function to get the value of an attribute from an HTML node
func getAttribute(node *html.Node, attributeName string) string {
	for _, attr := range node.Attr {
		if attr.Key == attributeName {
			return attr.Val
		}
	}
	return ""
}

func extractHTMLCode(node *html.Node) string {
	var sb strings.Builder
	html.Render(&sb, node)
	return sb.String()
}

// Function to extract the CSS code from the website (if any)
func extractCSSCode(node *html.Node) string {
	var sb strings.Builder
	extractCSSCodeRecursively(node, &sb)
	return sb.String()
}

func extractCSSCodeRecursively(node *html.Node, sb *strings.Builder) {
	if node.Type == html.ElementNode && node.Data == "style" && node.FirstChild != nil {
		sb.WriteString(node.FirstChild.Data)
		sb.WriteString("\n")
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		extractCSSCodeRecursively(child, sb)
	}
}

// Helper function to get the value of an attribute from an HTML node
// ... (unchanged from previous version)

func main() {
	// Prompt user to enter the website URL
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the website URL: ")
	url, _ := reader.ReadString('\n')

	// Clean the URL to remove leading/trailing whitespace and control characters
	url = strings.TrimSpace(url)

	// Fetch the HTML content of the URL
	doc, err := fetchHTML(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract SEO content, title, headings, images, and links
	seoContent := extractSEOContent(doc)
	title := extractTitle(doc)
	headings := extractHeadings(doc)
	images := extractImages(doc)
	links := extractLinks(doc)

	// Extract the entire HTML code
	htmlCode := extractHTMLCode(doc)

	// Extract the CSS code (if any)
	cssCode := extractCSSCode(doc)

	// Save the extracted information to respective files
	saveToFiles(title, seoContent, headings, images, links, htmlCode, cssCode)

	fmt.Println("Scraped website information has been written to files.")
}

// Function to save the extracted information to separate files
func saveToFiles(title string, seoContent, headings, images, links []string, htmlCode, cssCode string) {
	file, err := os.Create("scraped_website_info.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the extracted information to the file
	fmt.Fprintln(file, "Title:", title)
	fmt.Fprintln(file, "\nSEO Content:")
	for _, content := range seoContent {
		fmt.Fprintln(file, content)
	}
	fmt.Fprintln(file, "\nHeadings:")
	for _, heading := range headings {
		fmt.Fprintln(file, heading)
	}
	fmt.Fprintln(file, "\nImages:")
	for _, image := range images {
		fmt.Fprintln(file, image)
	}
	fmt.Fprintln(file, "\nLinks:")
	for _, link := range links {
		fmt.Fprintln(file, link)
	}

	// Save the entire HTML code to a separate file
	htmlFile, err := os.Create("website.html")
	if err != nil {
		fmt.Println("Error creating HTML file:", err)
		return
	}
	defer htmlFile.Close()
	fmt.Fprintln(htmlFile, htmlCode)

	// Save the CSS code to a separate file
	cssFile, err := os.Create("styles.css")
	if err != nil {
		fmt.Println("Error creating CSS file:", err)
		return
	}
	defer cssFile.Close()
	fmt.Fprintln(cssFile, cssCode)
}
