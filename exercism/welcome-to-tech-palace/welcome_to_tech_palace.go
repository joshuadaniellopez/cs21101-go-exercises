package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := ""
	for i := 0; i < numStarsPerLine; i++ {
		stars += "*"
	}
	return stars + "\n" + welcomeMsg + "\n" + stars

}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	message := strings.ReplaceAll(oldMsg, "*", "")
	message = strings.TrimSpace(message)
	return message
}
