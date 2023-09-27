package main

import "fmt"

func (b *Blog) ShareLink(strategy ShareStrategy) {
	strategy.share(b.Link)
}

type ShareStrategy interface {
	share(link string)
}

type TelegramStrategy struct {
}

func (ts TelegramStrategy) share(link string) {
	fmt.Printf("Telegram: %s\n", link)
}

type WhatsappStrategy struct {
}

func (ws WhatsappStrategy) share(link string) {
	fmt.Printf("Whatsapp: %s\n", link)
}

type GmailStrategy struct {
}

func (gs GmailStrategy) share(link string) {
	fmt.Printf("Gmail: %s\n", link)
}

func strategy() {
	telegram := TelegramStrategy{}
	whatsapp := WhatsappStrategy{}
	gmail := GmailStrategy{}

	blog := NewBlog("Learn GoLang")

	blog.ShareLink(telegram)
	blog.ShareLink(whatsapp)
	blog.ShareLink(gmail)

}
