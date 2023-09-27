package main

import (
	"fmt"
	"time"
)

type Subscrable interface {
	update(post Post, blogName string)
}

type User struct {
	name string
}

func (u User) update(post Post, blogName string) {
	fmt.Printf("Hello %s. In blog '%s' new post: '%s'\n", u.name, blogName, post.Title)
}

func NewUser(name string) User {
	return User{name: name}
}

type Post struct {
	Title, Description string
	UpdateTime         time.Time
}

type Blog struct {
	name        string
	IsActive    bool
	Posts       []Post
	subscribers []Subscrable
}

func NewBlog(name string) Blog {
	return Blog{name: name, IsActive: true}
}

func (b *Blog) AddPost(title, description string) {
	post := Post{title, description, time.Now()}
	b.Posts = append(b.Posts, post)

	b.notifyAll(post, b.name)
}

func (b *Blog) AddSubscriber(subscriber Subscrable) {
	b.subscribers = append(b.subscribers, subscriber)
}

func (b *Blog) DeleteSubscriber(subscriber Subscrable) bool {
	for i, sub := range b.subscribers {
		if sub == subscriber {
			b.subscribers = append(b.subscribers[:i], b.subscribers[i+1:]...)
			return true
		}
	}
	return false
}

func (b *Blog) notifyAll(post Post, blogName string) {
	for _, subscrable := range b.subscribers {
		subscrable.update(post, blogName)
	}
}

func main() {
	sultan := NewUser("Sultan")
	Bob := NewUser("Bob")

	blog := NewBlog("GoLang Blog")
	blog.AddSubscriber(sultan)
	blog.AddSubscriber(Bob)

	blog.AddPost("New Project", "Lets start new project on GoLang")
}
