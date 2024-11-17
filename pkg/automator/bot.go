package automator

import (
	"log/slog"
	"time"

	"github.com/playwright-community/playwright-go"

	"indcdi/pkg/configparser"
)

type PlayWriteBot struct {
	page   playwright.Page
	logger *slog.Logger
	steps  configparser.StepsContainer
}

func NewPlayWriteBot(page playwright.Page, steps configparser.StepsContainer, logger *slog.Logger) *PlayWriteBot {
	return &PlayWriteBot{
		page:   page,
		logger: logger,
		steps:  steps,
	}
}

func (pb *PlayWriteBot) AutoMate() bool {
	for i := range pb.steps.Steps {
		pb.logger.Info("Now doing :",
			slog.String("Step Name", pb.steps.Steps[i].Name),
			slog.String("Step xPath/Selector", pb.steps.Steps[i].XPath),
		)

		pb.WaitForLoadState()

		if pb.steps.Steps[i].WaitFor > 0 {
			time.Sleep(time.Duration(pb.steps.Steps[i].WaitFor) * time.Second)
		}

		if pb.steps.Steps[i].WaitUntil {
			pb.WaitForElement(pb.steps.Steps[i].XPath)
		}

		if pb.steps.Steps[i].Click {
			pb.Click(pb.steps.Steps[i].XPath)
		}

		if pb.steps.Steps[i].Select {
			pb.SelectOptionByValue(pb.steps.Steps[i].XPath, pb.steps.Steps[i].Value)
		}

		if pb.steps.Steps[i].Select {
			pb.SelectOptionByValue(pb.steps.Steps[i].XPath, pb.steps.Steps[i].Value)
		}

		if pb.steps.Steps[i].Fill {
			pb.InputFill(pb.steps.Steps[i].XPath, pb.steps.Steps[i].Value)
		}

		pb.logger.Info("Now doing :",
			slog.String("Step Name", pb.steps.Steps[i].Name),
			slog.String("Step xPath/Selector", pb.steps.Steps[i].XPath),
		)
	}
	return true
}

func (pb *PlayWriteBot) Click(xpath string) bool {
	err := pb.page.Locator(xpath).Click()
	if err != nil {
		pb.logger.Error("while clicking: ",
			slog.String("xpath", xpath),
			slog.String("err", err.Error()),
		)
		return false
	}
	return true
}

func (pb *PlayWriteBot) SelectOptionByValue(xpath string, value string) bool {
	if _, err := pb.page.Locator(xpath).SelectOption(playwright.SelectOptionValues{
		ValuesOrLabels: &[]string{value},
	}); err != nil {
		pb.logger.Error("could not select option: ",
			slog.String("xpath", xpath),
			slog.String("value", value),
			slog.String("err", err.Error()),
		)
		return false
	}
	return true
}

func (pb *PlayWriteBot) InputFill(xpath string, value string) bool {
	err := pb.page.Locator(xpath).Fill(value)
	if err != nil {
		pb.logger.Error("could not fill the input field: ",
			slog.String("xpath", xpath),
			slog.String("value", value),
			slog.String("err", err.Error()),
		)
		return false
	}
	return true
}

func (pb *PlayWriteBot) GetInputValue(xpath string) string {
	selectedValue, err := pb.page.Locator(xpath).InputValue()
	if err != nil {
		pb.logger.Error("could not get selected value: ",
			slog.String("xpath", xpath),
			slog.String("err", err.Error()),
		)
		return ""
	}
	return selectedValue
}

func (pb *PlayWriteBot) WaitForElement(xpath string) bool {
	// Wait for the element to become visible
	err := pb.page.Locator(xpath).WaitFor(playwright.LocatorWaitForOptions{
		State: playwright.WaitForSelectorStateVisible,
	})
	if err != nil {
		pb.logger.Error("element not visible: ",
			slog.String("xpath", xpath),
			slog.String("err", err.Error()),
		)
		return false
	}
	return true
}

func (pb *PlayWriteBot) GetPageContent() string {
	content, err := pb.page.Content()
	if err != nil {
		pb.logger.Error("cant load content: ",
			slog.String("err", err.Error()),
		)
		return ""
	}
	return content
}

func (pb *PlayWriteBot) WaitForLoadState() bool {
	err := pb.page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateLoad,
	})
	if err != nil {
		pb.logger.Error("cant load content: ",
			slog.String("err", err.Error()),
		)
		return false
	}
	return true
}
