package main

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func generatePassword(input string, length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?/|"
	hash := sha256.Sum256([]byte(input))
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])

	var password strings.Builder
	for i := 0; len(password.String()) < length && i < len(hashBase64); i++ {
		index := int(hashBase64[i]) % len(charset)
		password.WriteByte(charset[index])
	}
	return password.String()
}

// Validate if the PIN is a 6-digit number
func isValidPin(pin string) bool {
	if len(pin) != 6 {
		return false
	}
	for _, ch := range pin {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Password Generator")

	pinInput := widget.NewEntry()
	pinInput.SetPlaceHolder("Enter 6-digit primary password (PIN)")

	domainInput := widget.NewEntry()
	domainInput.SetPlaceHolder("Enter full domain or uppercase alias")

	result := widget.NewLabel("Generated password will appear here")
	hintLabel := widget.NewLabel("Hint: Use a commonly remembered 6-digit number for better consistency")

	var copyBtn *widget.Button

	updatePassword := func() {
		pin := pinInput.Text
		domain := domainInput.Text

		if domain == "" {
			result.SetText("Please enter a domain or alias")
			if copyBtn != nil {
				copyBtn.SetText("Copy Password")
			}
			return
		}
		if !isValidPin(pin) {
			result.SetText("Please enter a valid 6-digit PIN")
			if copyBtn != nil {
				copyBtn.SetText("Copy Password")
			}
			return
		}

		combined := domain + ":" + pin
		password := generatePassword(combined, 16)
		result.SetText("Generated Password: " + password)
		if copyBtn != nil {
			copyBtn.SetText("Copy Password")
		}
	}

	copyBtn = widget.NewButton("Copy Password", func() {
		text := result.Text
		if strings.HasPrefix(text, "Generated Password: ") {
			pass := strings.TrimPrefix(text, "Generated Password: ")
			myWindow.Clipboard().SetContent(pass)
			copyBtn.SetText("Copied ✔")
		} else {
			copyBtn.SetText("No password to copy")
		}
	})

	domainInput.OnChanged = func(string) {
		updatePassword()
	}
	pinInput.OnChanged = func(string) {
		updatePassword()
	}

	content := container.NewVBox(
		hintLabel,
		pinInput,
		domainInput,
		result,
		copyBtn,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(450, 280))
	myWindow.ShowAndRun()
}
