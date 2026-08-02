package main

type ArticleData struct {
    ID      int  `json:"id"`
    Title   string `json:"title"`
    Author  string  `json:"author"`
    Summary string  `json:"summary"`
    Content string `json:"content"`
}
articles := []Article{
    {
        ID:      1,
        Title:   "Learning Go",
        Author:  "Joel",
        Summary: "My experience learning web development.",
    },
    {
        ID:      2,
        Title:   "Understanding HTTP",
        Author:  "Joel",
        Summary: "What I learned about HTTP servers.",
    },
}