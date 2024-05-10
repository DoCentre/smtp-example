package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/akamensky/argparse"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Sender   string
	Password string
}

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}

	parser := argparse.NewParser("",
		"This is a simple example of how to send an email using the `net/smtp` package in Go. The example reads the SMTP server host, port, sender email address, and password from environment variables. The email recipients are specified as command-line arguments.")
	rcpts := parser.StringList("r", "recipient",
		&argparse.Options{Required: true, Help: "Recipient email address; can be specified multiple times"})
	err = parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	log.Println("Connecting to the remote SMTP server...")
	client, err := smtp.Dial(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
		ServerName:         config.Host,
	}
	log.Println("Sending the STARTTLS command and encrypting all further communication...")
	err = client.StartTLS(&tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	auth := smtp.PlainAuth("", config.Sender, config.Password, config.Host)
	if err := client.Auth(auth); err != nil {
		log.Fatal(err)
	}

	log.Println("Setting the sender and recipient...")
	if err := client.Mail(config.Sender); err != nil {
		log.Fatal(err)
	}
	for _, rcpt := range *rcpts {
		if err := client.Rcpt(rcpt); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Sending the email body...")
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	subject := "Sample email"
	header := fmt.Sprintf(
		"To: %s\r\nFrom: %s\r\nSubject: %s\r\n\r\n",
		*rcpts, config.Sender, subject)
	message := "This is a sample email sent using SMTP."
	_, err = fmt.Fprintf(wc, header+message)
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sending the QUIT command and closing the connection...")
	err = client.Quit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}

func parseConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("error loading .env file")
	}

	var config Config
	config.Host = os.Getenv("HOST")
	config.Port = os.Getenv("PORT")
	config.Sender = os.Getenv("SENDER")
	config.Password = os.Getenv("PASSWORD")
	return config, nil
}
