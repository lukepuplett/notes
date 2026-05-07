# Google Docs to Markdown Converter

A Go tool that converts Google Docs documents (via the Google Docs API) into clean, self-contained Markdown format with downloaded images.

## Overview

This tool works in conjunction with the **GWS (Google Workspace CLI)** to:
1. Fetch Google Docs documents as JSON using GWS
2. Parse the JSON using the official Google Docs API Go SDK
3. **Download all images** locally into an `images/` folder
4. Convert the document structure to properly formatted Markdown
5. **Update image links to local references** (no external dependencies)
6. Preserve text styling, links, and document structure

## Prerequisites

- [Go](https://golang.org/) 1.20 or later
- [GWS CLI](https://github.com/googleworkspace/cli) installed and authenticated
- Access to your Google Docs (via GWS authentication)

## Installation

```bash
# Dependencies are already in go.mod
go mod download
```

## Usage

### Quick Start (Recommended: Piping)

Convert directly from GWS to Markdown in one command:

```bash
GOOGLE_WORKSPACE_CLI_LOG="" \
  gws docs documents get --params '{"documentId": "YOUR_DOC_ID", "includeTabsContent": true}' 2>/dev/null \
  | go run main.go
```

**Why piping?**
- No intermediate files
- Works with shell scripts and pipelines
- Cleaner, more Unix-idiomatic

### Alternative: File-Based Workflow

**Step 1: Export Document with GWS**

```bash
gws docs documents get --params '{"documentId": "YOUR_DOC_ID", "includeTabsContent": true}' > journal.json
```

**Step 2: Convert to Markdown**

```bash
go run main.go
```

This reads `journal.json` from the current directory.

### Find Your Document ID

```bash
# From Google Drive URL
https://docs.google.com/document/d/YOUR_DOC_ID/edit

# Or list your documents
gws drive files list --params '{"pageSize": 10, "q": "mimeType='\''application/vnd.google-apps.document'\''"}'
```

### What Gets Generated

The tool creates:
- **`<doc-name>/`** - Folder named after document
- **`tab-1.md`, `tab-2.md`, etc.** - One file per tab (if multi-tab)
- **`README.md`** - Single file for non-tabbed documents
- **`images/`** - Folder with all downloaded images (MD5-hashed filenames)

### View the Result

```bash
code 2026-journal/        # Open folder in VS Code
code 2026-journal/README.md  # Open markdown file
cat 2026-journal/README.md   # View in terminal
```

### Output Structure

```
2026-journal/
├── README.md
└── images/
    ├── image-1.png
    ├── image-2.jpg
    ├── image-3.jpg
    └── ...
```

## What Gets Converted

| Google Docs Feature | Markdown Output |
|---|---|
| Document Title | `# Title` |
| Heading Levels | `##`, `###`, `####`, etc. |
| Bold Text | `**text**` |
| Italic Text | `*text*` |
| Strikethrough | `~~text~~` |
| Underline | `__text__` |
| Hyperlinks | `[text](url)` |
| **Images** | **`![image](images/image-N.png)`** ← **Downloaded locally** ✨ |
| Section Breaks | `---` |
| Tables | Markdown table format |

## Example Workflow

```bash
# 1. List your Google Docs
gws drive files list --params '{"pageSize": 20, "q": "mimeType='\''application/vnd.google-apps.document'\''"}'

# 2. Export your journal
gws docs documents get --params '{"documentId": "1sAWs4ikiX0kix7BBYa3gWVWezWngNwQeSqubsjZQFtE"}' > journal.json

# 3. Convert to Markdown (with images)
go run main.go

# Output:
# ✓ Downloaded image: image-1.png
# ✓ Downloaded image: image-2.jpg
# ✓ Converted to Markdown
# Document: 2026 - Journal
# Output folder: 2026-journal/
# Images: 5 downloaded to 2026-journal/images/

# 4. View the result
code 2026-journal/
```

## Output Structure

### Single Tab / No Tabs

Creates a **self-contained folder** with:

```
2026-journal/
├── README.md           (Markdown file with content)
└── images/
    ├── bb25171b371f.png     (MD5 hash named)
    ├── fcf05aa53e86.jpg
    └── ...
```

### Multiple Tabs

Each tab gets its **own markdown file**, with a **shared images folder**:

```
my-document/
├── overview.md         (Tab 1: "Overview")
├── detailed-notes.md   (Tab 2: "Detailed Notes")
├── appendix.md         (Tab 3: "Appendix")
└── images/
    ├── bb25171b371f.png     (Image from any tab)
    ├── fcf05aa53e86.jpg     (De-duplicated)
    └── ...
```

**Filenames** are derived from tab titles:
- Tab title "Project Overview" → `project-overview.md`
- Tab title "Q1 Results" → `q1-results.md`
- Unnamed tabs → `tab-1.md`, `tab-2.md`, etc.

**Images** use MD5-based filenames (first 12 chars of hash):
- De-duplicated: Same image in multiple tabs = same file
- Deterministic: Same URL always produces same filename
- Example: Long Google URL → `bb25171b371f.png`

## Implementation Details

### Architecture

1. **JSON Unmarshaling**: Uses `google.golang.org/api/docs/v1` to deserialize GWS output
2. **Folder Creation**: Creates output folder based on document title (sanitized)
3. **Image Download**: Fetches images from Google CDN and saves locally
   - Detects file format from HTTP Content-Type header
   - Names files `image-1.png`, `image-2.jpg`, etc.
   - Skips failed downloads but continues processing
4. **Image Mapping**: Tracks URL → local filename for markdown linking
5. **Paragraph Processing**: Iterates through document body elements
6. **Text Styling**: Checks `TextStyle` properties for formatting (bold, italic, etc.)
7. **Link Handling**: Extracts URL from `TextStyle.Link` structure
8. **Image References**: Uses local filenames in markdown instead of URLs

### Image Handling: MD5 Hash-Based Filenames

Images are saved with MD5 hash-based filenames (first 12 characters of URL hash):

**Benefits:**
- **De-duplication**: If the same image appears in multiple documents, it's only saved once
- **Deterministic**: Same URL always produces same filename
- **Collision-resistant**: Extremely unlikely two different images get same name

**Example:**
```
https://lh7-rt.googleusercontent.com/...very-long-url... → bb25171b371f.png
```

### Tabs Structure (Google Docs API)

The Google Docs API represents documents with an optional `Document.Tabs` field:

```go
type Document struct {
    Tabs []*Tab              // Array of tabs (can be empty)
    Body *Body               // Root body (legacy, used if no tabs)
    // ... other fields
}

type Tab struct {
    TabProperties *TabProperties   // Title, ID, index
    DocumentTab   *DocumentTab      // Actual content (Body, InlineObjects, etc.)
    ChildTabs     []*Tab            // Nested tabs (up to 3 levels)
}

type TabProperties struct {
    TabId  string  // Unique tab identifier
    Title  string  // User-visible tab name
    Index  int64   // Tab position
    // ... other fields
}

type DocumentTab struct {
    Body          *Body
    InlineObjects map[string]InlineObject  // Images in this tab
    // ... other content fields
}
```

### Input Handling: stdin vs File

The tool supports both piping and file input:

```go
func readJSONInput() []byte {
    // Check if stdin has data (not a terminal)
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        // stdin is not a terminal, read from pipe
        return ioutil.ReadAll(os.Stdin)
    }
    
    // Fall back to file for backwards compatibility
    return ioutil.ReadFile("./journal.json")
}
```

**Priority:**
1. If data is piped to stdin → use stdin (Unix philosophy)
2. Else if `journal.json` exists → use file (backwards compatible)
3. Else → error

### Key Code Functions

- `readJSONInput()` - Reads JSON from stdin or file (stdin preferred)
- `main()` - Orchestrates tab detection, image download, markdown generation
- `downloadImagesFromMap()` - Downloads images from any InlineObjects map
- `generateMarkdownFromTab()` - Generates markdown from a single Tab
- `sanitizeFolderName()` - Converts document/tab title to safe filename
- `downloadImageWithHash()` - Fetches image, uses MD5 hash for filename, detects format
- `hashURL()` - Computes 12-char MD5 hash of image URL
- `generateMarkdown()` - Generates markdown from root-level document
- `paragraphToMarkdown()` - Converts paragraphs to MD with styling, local image links
- `tableToMarkdown()` - Converts tables to MD table syntax

## Features & Limitations

### ✅ Supported
- **Text styling**: Bold, italic, strikethrough, underline
- **Headings**: All 6 heading levels + title style
- **Links**: Hyperlinks preserved with URLs
- **Images**: Automatically downloaded and linked locally ✨
  - **De-duplication**: Same image in multiple documents = same filename
  - **Content-based naming**: MD5 hash of image URL ensures unique, stable filenames
- **Tables**: Basic markdown table format
- **Section breaks**: Converted to horizontal rules (`---`)
- **Document structure**: Paragraphs, spacing preserved

### ✨ Advanced Features
- **Multi-tab documents**: Each tab gets its own markdown file ✨ NEW
- **Advanced formatting**: Colors, custom fonts, font sizes
- **Complex tables**: Merged cells, nested tables
- **Comments/Suggestions**: Document must be finalized
- **Text boxes/Shapes**: Floating elements
- **Embedded objects**: Charts, embedded spreadsheets
- **Page breaks**: Converted to section breaks

## Extending the Tool

To add new features:

1. **Custom styling**: Modify `paragraphToMarkdown()` to handle additional `TextStyle` properties
2. **Different output format**: Create a new converter function (e.g., `paragraphToHTML()`)
3. **Batch processing**: Add a loop to process multiple document IDs
4. **File output options**: Modify `main()` to support custom output paths

## Troubleshooting

### "document not found" error
- Verify the Document ID is correct (check the Google Docs URL)
- Ensure you have read access to the document
- Run `gws docs documents get --params '{"documentId": "..."}' --dry-run` to validate

### JSON file not found
- Ensure `journal.json` exists in the working directory
- Run from the same directory as `main.go`, or provide a full path in the code

### Go compilation errors
- Update dependencies: `go mod tidy`
- Check Go version: `go version` (requires 1.20+)

## Related Tools

- **GWS CLI**: https://github.com/googleworkspace/cli
- **Google Docs API**: https://developers.google.com/workspace/docs/api
- **Google Docs Go SDK**: https://pkg.go.dev/google.golang.org/api/docs/v1
