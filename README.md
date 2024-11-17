# AutoMate: Playwright Automation Framework for Golang

<h1 align="center">autoMate</h1>
<p align="center">
  <img src="https://img.shields.io/github/languages/top/0x30c4/autoMate?style=flat-square" alt="Test">
</p>
<p align="justify">
AutoMate is a lightweight and extensible automation tool built with Playwright for Golang. It provides a simple, configuration-driven approach to automate web interactions such as clicking elements, selecting options, filling inputs, and waiting for elements. The tool is ideal for scenarios involving repetitive UI interactions and testing.
</p>
<p align="center">
  <a href="mailto:automate@0x30c4.dev"> automate@0x30c4.dev </a>
</p>

<br>
<img src="./assets/demo.gif">
<br>

---

## Features
- **Config-Driven Automation:** Define automation steps in a simple YAML configuration file.
- **Supports Core UI Actions:**
  - Clicking buttons and links.
  - Selecting dropdown options by value.
  - Filling input fields.
- **Dynamic Waits:** Automatically wait for elements to become visible or the page to load.
- **Extensibility:** Easy to add more actions or customize existing behavior.

---

## Installation

1. **Install Playwright for Go:**
   ```bash
   go install github.com/playwright-community/playwright-go@latest
   ```

2. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/autoMate.git
   cd autoMate
   ```

3. **Install Dependencies:**
   Ensure all required dependencies are installed:
   ```bash
   go mod tidy
   ```

---

## Usage

### Configuration File (Example)
Create a YAML file to define the automation steps:

```yaml
steps:
  - name: Close popup
    xpath: //*[@id="popup"]/span
    wait_until: true
    click: true

  - name: Select your option
    xpath: "//select[@id='select']"
    value: "2"
    select: true
    wait_until: true
    wait_for: 1

  - name: Enter your web file ID
    xpath: "#application > div:nth-child(1) > div.form-group.inner-addon.right-addon > input"
    value: "hello"
    fill: true
    wait_for: 1
    wait_until: true
```

### Main Code
Initialize and run the automation tool in your main Go application:

```go
package main

import (
	"log/slog"

	"github.com/playwright-community/playwright-go"
	"your-project-path/automator"
	"your-project-path/configparser"
)

func main() {
	// Load YAML configuration
	config := configparser.LoadSteps("config.yaml") // Replace with your YAML path

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	// Launch browser and navigate
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // Set to false for debugging
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Navigate to the target page
	err = page.Goto("https://example.com") // Replace with your URL
	if err != nil {
		log.Fatalf("could not navigate: %v", err)
	}

	// Initialize logger
	logger := slog.New(slog.NewTextHandler())

	// Initialize automation bot
	bot := automator.NewPlayWriteBot(page, config, logger)

	// Start automation
	if bot.AutoMate() {
		logger.Info("Automation completed successfully!")
	} else {
		logger.Error("Automation encountered an error.")
	}
}
```

### Running the Tool
Run the tool with:
```bash
go run main.go
```

---

## Key Components

### 1. **ConfigParser**
Parses the YAML configuration file containing the steps for automation. Each step defines:
- **name**: Description of the step.
- **xpath**: Locator for the element.
- **value**: Value to fill or select.
- **actions**: `click`, `fill`, or `select`.
- **wait_for**: Seconds to wait after this step.
- **wait_until**: Wait for the element to be visible before executing.

### 2. **PlayWriteBot**
The core automation engine that processes each step, waits for elements, and performs actions based on the configuration.

### 3. **Logging**
Uses `slog` for detailed logging of actions, errors, and debugging information.

---

## Debugging

- **Screenshots:** Capture screenshots to debug unexpected behavior.
- **Logs:** Check logs for detailed information about each step.

---

## Contribution

Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request.

---

## License
This project is licensed under the MIT License.

---

For any questions or feedback, feel free to contact automate@0x30c4.dev.
