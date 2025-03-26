package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

// Load .env file
func LoadEnv() {
	err := godotenv.Load("src/.env") // Load file .env từ thư mục src
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Hàm gửi email reset mật khẩu
func SendResetPasswordEmail(toEmail string, token string) error {
	// Lấy thông tin SMTP từ .env
	LoadEnv()
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASS")
	fromEmail := os.Getenv("SMTP_FROM")

	// Chuyển smtpPort từ string -> int
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Println("Lỗi khi chuyển đổi SMTP_PORT:", err)
		return err
	}

	// Tạo đường link đặt lại mật khẩu
	resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", token)
	// Tạo nội dung email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fromEmail) // Email gửi
	mailer.SetHeader("To", toEmail)     // Email người nhận
	mailer.SetHeader("Subject", "Reset Password")
	mailer.SetBody("text/html", fmt.Sprintf(`
		<h3>Yêu cầu đặt lại mật khẩu</h3>
		<p>Nhấn vào đường link dưới đây để đặt lại mật khẩu:</p>
		<a href="%s">%s</a>
	`, resetLink, resetLink))
	dialer := gomail.NewDialer(smtpHost, port, smtpUser, smtpPassword)
	return dialer.DialAndSend(mailer)
}
