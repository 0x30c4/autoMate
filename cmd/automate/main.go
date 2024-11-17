package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/playwright-community/playwright-go"

	"indcdi/pkg/automator"
	"indcdi/pkg/configparser"
)

func main() {
	// start logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Start Playwright
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	// Launch a browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	// Open a new browser context and page
	context, err := browser.NewContext()
	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}

	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Navigate to the target URL
	url := "https://google.com/" // Replace with your target URL
	if _, err := page.Goto(url); err != nil {
		log.Fatalf("could not go to URL: %v", err)
	}

	steps, err := configparser.ParserSteps("./automator.yml")

	pwBot := automator.NewPlayWriteBot(page, steps, logger)

	pwBot.AutoMate()
}
