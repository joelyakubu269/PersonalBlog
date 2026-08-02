
package main

type ArticleData struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

var articles = []ArticleData{
	{
		ID:      1,
		Title:   "Learning Go",
		Author:  "Joel",
		Summary: "My experience learning web development.",
		Content: "Learning Go has been an interesting journey for me. When I started learning the language, I was curious about how Go could be used to build real-world applications and web servers. At first, some concepts were difficult to understand, especially working with HTTP handlers, templates, and routing. However, as I continued practicing, I began to understand how the different parts of a Go web application fit together. One of the things I enjoy about Go is its simplicity and straightforward approach to programming. Building small projects has helped me improve my understanding of programming concepts and given me more confidence as I continue learning web development.",
	},
	{
		ID:      2,
		Title:   "Understanding HTTP",
		Author:  "Joel",
		Summary: "What I learned about HTTP servers.",
		Content: "Understanding HTTP is an important part of learning web development. HTTP allows clients, such as web browsers, to communicate with servers by sending requests and receiving responses. While building my Go web application, I learned how an HTTP server listens for requests and how handlers are used to respond to different URLs. I also learned that HTTP methods such as GET and POST describe the type of action a client wants to perform. Understanding concepts such as status codes, query parameters, routes, and request methods has helped me understand what happens behind the scenes when I visit a website. Learning how these pieces work together is an important step in becoming a better web developer.",
	},
}