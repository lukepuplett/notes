package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"google.golang.org/api/docs/v1"
)

func main() {
	// Read JSON from stdin or file
	jsonData := readJSONInput()

	// Unmarshal into a Document struct
	var doc *docs.Document
	err := json.Unmarshal(jsonData, &doc)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Create output folder based on document title
	outputFolder := sanitizeFolderName(doc.Title)
	imagesFolder := filepath.Join(outputFolder, "images")

	// Create directories
	err = os.MkdirAll(imagesFolder, 0755)
	if err != nil {
		log.Fatalf("Failed to create output folder: %v", err)
	}

	// Track image URLs to local filenames (shared across all tabs)
	imageMap := make(map[string]string)

	// Download images from root-level document (for backward compatibility)
	if doc.InlineObjects != nil {
		downloadImagesFromMap(doc.InlineObjects, imagesFolder, imageMap)
	}

	// Download images from all tabs
	if len(doc.Tabs) > 0 {
		for _, tab := range doc.Tabs {
			if tab.DocumentTab != nil && tab.DocumentTab.InlineObjects != nil {
				downloadImagesFromMap(tab.DocumentTab.InlineObjects, imagesFolder, imageMap)
			}
		}
	}

	fmt.Printf("\n✓ Converted to Markdown\n")
	fmt.Printf("Document: %s\n", doc.Title)
	fmt.Printf("Output folder: %s/\n", outputFolder)

	// Check if document has tabs
	hasMultipleTabs := len(doc.Tabs) > 1

	if hasMultipleTabs {
		// Process each tab as a separate markdown file
		for i, tab := range doc.Tabs {
			var filename string
			var title string

			if tab.TabProperties != nil && tab.TabProperties.Title != "" {
				title = tab.TabProperties.Title
				filename = sanitizeFolderName(title) + ".md"
			} else {
				title = fmt.Sprintf("Tab %d", i+1)
				filename = fmt.Sprintf("tab-%d.md", i+1)
			}

			outputFile := filepath.Join(outputFolder, filename)
			markdown := generateMarkdownFromTab(tab, imageMap)

			err = ioutil.WriteFile(outputFile, []byte(markdown), 0644)
			if err != nil {
				log.Fatalf("Failed to write output file %s: %v", outputFile, err)
			}

			fmt.Printf("Markdown file: %s\n", outputFile)
		}
	} else if len(doc.Tabs) == 1 {
		// Single tab - use it as README
		tab := doc.Tabs[0]
		outputFile := filepath.Join(outputFolder, "README.md")
		markdown := generateMarkdownFromTab(tab, imageMap)

		err = ioutil.WriteFile(outputFile, []byte(markdown), 0644)
		if err != nil {
			log.Fatalf("Failed to write output file: %v", err)
		}

		fmt.Printf("Markdown file: %s\n", outputFile)
		fmt.Printf("\n=== MARKDOWN PREVIEW ===\n\n")
		fmt.Println(markdown)
	} else {
		// No tabs - use root-level document (legacy behavior)
		markdown := generateMarkdown(doc, imageMap)
		outputFile := filepath.Join(outputFolder, "README.md")
		err = ioutil.WriteFile(outputFile, []byte(markdown), 0644)
		if err != nil {
			log.Fatalf("Failed to write output file: %v", err)
		}

		fmt.Printf("Markdown file: %s\n", outputFile)
		fmt.Printf("\n=== MARKDOWN PREVIEW ===\n\n")
		fmt.Println(markdown)
	}

	if len(imageMap) > 0 {
		fmt.Printf("Images: %d downloaded (with MD5 hashing) to %s/\n", len(imageMap), imagesFolder)
	}
}

// readJSONInput reads JSON from stdin or from journal.json file
func readJSONInput() []byte {
	// Check if stdin has data (not a terminal)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// stdin is not a terminal, read from pipe
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Failed to read from stdin: %v", err)
		}
		if len(data) > 0 {
			return data
		}
	}

	// Fall back to reading from file (backwards compatibility)
	data, err := ioutil.ReadFile("./journal.json")
	if err != nil {
		log.Fatalf("Failed to read from stdin or file: provide piped JSON or create journal.json")
	}
	return data
}

// downloadImagesFromMap downloads images from an InlineObjects map
func downloadImagesFromMap(objects map[string]docs.InlineObject, imagesFolder string, imageMap map[string]string) {
	for objID, inlineObj := range objects {
		if inlineObj.InlineObjectProperties != nil &&
			inlineObj.InlineObjectProperties.EmbeddedObject != nil &&
			inlineObj.InlineObjectProperties.EmbeddedObject.ImageProperties != nil {
			url := inlineObj.InlineObjectProperties.EmbeddedObject.ImageProperties.ContentUri

			// Skip if already downloaded
			if _, exists := imageMap[url]; exists {
				continue
			}

			filename, err := downloadImageWithHash(url, imagesFolder)
			if err != nil {
				fmt.Printf("⚠ Failed to download image %s: %v\n", objID, err)
			} else {
				imageMap[url] = filename
				fmt.Printf("✓ Downloaded image: %s\n", filename)
			}
		}
	}
}

// generateMarkdownFromTab generates markdown from a DocumentTab
func generateMarkdownFromTab(tab *docs.Tab, imageMap map[string]string) string {
	var result strings.Builder

	// Add tab title if available
	if tab.TabProperties != nil && tab.TabProperties.Title != "" {
		result.WriteString(fmt.Sprintf("# %s\n\n", tab.TabProperties.Title))
	}

	// Process tab content
	if tab.DocumentTab != nil && tab.DocumentTab.Body != nil {
		for _, elem := range tab.DocumentTab.Body.Content {
			if elem.Paragraph != nil {
				// Create a temporary document-like structure for tab content
				doc := &docs.Document{
					InlineObjects: tab.DocumentTab.InlineObjects,
				}
				result.WriteString(paragraphToMarkdown(elem.Paragraph, doc, imageMap))
				result.WriteString("\n")
			} else if elem.Table != nil {
				result.WriteString(tableToMarkdown(elem.Table))
				result.WriteString("\n")
			} else if elem.SectionBreak != nil {
				result.WriteString("---\n\n")
			}
		}
	}

	return result.String()
}

// sanitizeFolderName converts document title to a safe folder name
func sanitizeFolderName(title string) string {
	// Remove special characters, keep only alphanumeric, spaces, and hyphens
	reg := regexp.MustCompile("[^a-zA-Z0-9\\s-]")
	sanitized := reg.ReplaceAllString(title, "")
	// Replace spaces with hyphens
	sanitized = regexp.MustCompile("\\s+").ReplaceAllString(sanitized, "-")
	// Remove consecutive hyphens
	sanitized = regexp.MustCompile("-+").ReplaceAllString(sanitized, "-")
	// Convert to lowercase
	sanitized = strings.ToLower(strings.TrimSpace(sanitized))
	return sanitized
}

// downloadImageWithHash downloads an image from URL and saves it with MD5 hash filename
func downloadImageWithHash(url string, imagesFolder string) (string, error) {
	// Check if already downloaded
	hash := hashURL(url)

	// Fetch the image
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Determine file extension from Content-Type header
	contentType := resp.Header.Get("Content-Type")
	extension := ".jpg"
	if strings.Contains(contentType, "png") {
		extension = ".png"
	} else if strings.Contains(contentType, "gif") {
		extension = ".gif"
	} else if strings.Contains(contentType, "webp") {
		extension = ".webp"
	}

	// Create filename using MD5 hash of URL
	filename := fmt.Sprintf("%s%s", hash, extension)
	filePath := filepath.Join(imagesFolder, filename)

	// Skip if already exists (de-duplication)
	if _, err := os.Stat(filePath); err == nil {
		return filename, nil
	}

	// Save to file
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// hashURL returns the first 12 characters of the MD5 hash of the URL
func hashURL(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])[:12]
}

func generateMarkdown(doc *docs.Document, imageMap map[string]string) string {
	var result strings.Builder

	// Add title
	result.WriteString(fmt.Sprintf("# %s\n\n", doc.Title))

	// Extract content
	if doc.Body != nil {
		for _, elem := range doc.Body.Content {
			if elem.Paragraph != nil {
				result.WriteString(paragraphToMarkdown(elem.Paragraph, doc, imageMap))
				result.WriteString("\n")
			} else if elem.Table != nil {
				result.WriteString(tableToMarkdown(elem.Table))
				result.WriteString("\n")
			} else if elem.SectionBreak != nil {
				result.WriteString("---\n\n")
			}
		}
	}

	return result.String()
}


func paragraphToMarkdown(para *docs.Paragraph, doc *docs.Document, imageMap map[string]string) string {
	var content strings.Builder

	// Determine heading level
	headingPrefix := ""
	if para.ParagraphStyle != nil {
		switch para.ParagraphStyle.NamedStyleType {
		case "TITLE":
			headingPrefix = "# "
		case "HEADING_1":
			headingPrefix = "## "
		case "HEADING_2":
			headingPrefix = "### "
		case "HEADING_3":
			headingPrefix = "#### "
		case "HEADING_4":
			headingPrefix = "##### "
		case "HEADING_5":
			headingPrefix = "###### "
		}
	}

	content.WriteString(headingPrefix)

	// Process elements
	for _, elem := range para.Elements {
		if elem.TextRun != nil {
			text := elem.TextRun.Content

			// Apply text styling
			if elem.TextRun.TextStyle != nil {
				if elem.TextRun.TextStyle.Bold {
					text = fmt.Sprintf("**%s**", text)
				}
				if elem.TextRun.TextStyle.Italic {
					text = fmt.Sprintf("*%s*", text)
				}
				if elem.TextRun.TextStyle.Strikethrough {
					text = fmt.Sprintf("~~%s~~", text)
				}
				if elem.TextRun.TextStyle.Underline {
					text = fmt.Sprintf("__%s__", text)
				}
				// Handle links
				if elem.TextRun.TextStyle.Link != nil && elem.TextRun.TextStyle.Link.Url != "" {
					text = fmt.Sprintf("[%s](%s)", text, elem.TextRun.TextStyle.Link.Url)
				}
			}

			content.WriteString(text)
		} else if elem.InlineObjectElement != nil {
			// Get image info
			objID := elem.InlineObjectElement.InlineObjectId
			if inlineObj, ok := doc.InlineObjects[objID]; ok {
				if inlineObj.InlineObjectProperties != nil &&
					inlineObj.InlineObjectProperties.EmbeddedObject != nil &&
					inlineObj.InlineObjectProperties.EmbeddedObject.ImageProperties != nil {
					url := inlineObj.InlineObjectProperties.EmbeddedObject.ImageProperties.ContentUri
					// Use local image filename if available
					var imgPath string
					if localFilename, ok := imageMap[url]; ok {
						imgPath = fmt.Sprintf("images/%s", localFilename)
					} else {
						// Fallback to URL if download failed
						imgPath = url
					}
					// Markdown image syntax: ![alt](path)
					content.WriteString(fmt.Sprintf("![image](%s)", imgPath))
				}
			}
		}
	}

	return content.String()
}

func tableToMarkdown(table *docs.Table) string {
	var result strings.Builder

	if len(table.TableRows) == 0 {
		return ""
	}

	// Process each row
	for i, row := range table.TableRows {
		result.WriteString("|")

		for _, cell := range row.TableCells {
			// Extract text from cell content
			cellText := ""
			for _, elem := range cell.Content {
				if elem.Paragraph != nil {
					for _, paraElem := range elem.Paragraph.Elements {
						if paraElem.TextRun != nil {
							cellText += paraElem.TextRun.Content
						}
					}
				}
			}

			result.WriteString(" ")
			result.WriteString(strings.TrimSpace(cellText))
			result.WriteString(" |")
		}

		result.WriteString("\n")

		// Add separator after header row
		if i == 0 {
			result.WriteString("|")
			for range row.TableCells {
				result.WriteString(" --- |")
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
