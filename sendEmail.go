package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

var View *template.Template

func init() {
	// Parse email template(s) during initialization
	// Example: View, _ = template.ParseFiles("template.html")
	// This assumes you have a template file named template.html
	// Feel free to adjust this as per your actual template setup.
	View = template.Must(template.New("emailTemplate").Parse(`
	<html>
	<head></head>
	<body>
		<h1>{{ .title }}</h1>
		<p>{{ .content }}</p>
		<p>Regards,<br>{{ .app_name }}</p>
	</body>
	</html>`))
}

/* SMTP send email */
func SendEmailSMTP(to []string, subject string, body string) bool {
	// Sender data
	from := os.Getenv("FROM_EMAIL")

	// Set up email headers
	header := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	msg := []byte("From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\r\n" +
		header + "\r\n" +
		body)

	// Sending email using SMTP
	err := smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		smtp.PlainAuth("", os.Getenv("FROM_APIKEY"), os.Getenv("EMAILSECRATE"), os.Getenv("SMTP_HOST")),
		from, to, msg,
	)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return false
	}
	return true
}

/*
* Main function to call for sending email
# Inputs:
- to: string array (email addresses to send to)
- template: templatePath (path to the email template)
- data: map[string]interface{} (data to be used in the email template)
*/
func SendEmail(to []string, template string, data map[string]interface{}) bool {
	// Prepare buffer to render email content
	buf := new(bytes.Buffer)

	// Add default data to be used in the template
	data["app_url"] = os.Getenv("APPURL")
	data["app_name"] = os.Getenv("APPNAME")

	// Render email template with provided data
	err := View.ExecuteTemplate(buf, template, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return false
	}

	// Send email using SMTP
	return SendEmailSMTP(to, fmt.Sprintf("%v", data["subject"]), buf.String())
}

func main() {
	// Example usage:
	to := []string{"recipient@example.com"}
	template := "emailTemplate"
	data := map[string]interface{}{
		"title":   "Hello!",
		"content": "This is a test email.",
		"subject": "Test Email",
	}
	success := SendEmail(to, template, data)
	fmt.Println("Email sent successfully:", success)
}
