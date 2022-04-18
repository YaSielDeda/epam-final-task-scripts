package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type magazine struct {
	id   int
	name string
}

type article struct {
	id              int
	magazines_id    int
	article_type_id int
	author_id       int
}

type articleType struct {
	id   int
	Type string
}

type author struct {
	id     int
	author string
}

func main() {
	/*magazines := []magazine{
		{1, "it herald"},
		{2, "it with kids"},
		{3, "it stories"},
	}

	articles := []article{
		{1, 1, 2, 3},
		{2, 3, 3, 2},
		{3, 2, 2, 4},
		{4, 1, 1, 1},
	}

	articleTypes := []articleType{
		{1, "news"},
		{2, "tech"},
		{3, "enertainment"},
	}

	authors := []author{
		{1, "Chappie"},
		{2, "Wall-e"},
		{3, "Atom"},
		{4, "T1000"},
	}*/

	db, err := sql.Open("mysql", "production:prodpass@tcp(db)/magazines_db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	magazines := selectMagazines(db)

	articles := selectArticles(db)

	articleTypes := selectArticleTypes(db)

	authors := selectAuthors(db)

	var processed strings.Builder

	processed.WriteString("<h1>Magazines</h1>")
	processed.WriteString("<table border='1'>")
	processed.WriteString("<tr><th>id</th><th>name</th></tr>")
	for _, item := range magazines {
		processed.WriteString("<tr>")
		processed.WriteString("<td>")
		processed.WriteString(fmt.Sprint(item.id))
		processed.WriteString("</td>")

		processed.WriteString("<td>")
		processed.WriteString(string(item.name))
		processed.WriteString("</td>")
		processed.WriteString("</tr>")
	}
	processed.WriteString("</table>")

	//--

	processed.WriteString("<h1>Article types</h1>")
	processed.WriteString("<table border='1'>")
	processed.WriteString("<tr><th>id</th><th>type</th></tr>")
	for _, item := range articleTypes {
		processed.WriteString("<tr>")
		processed.WriteString("<td>")
		processed.WriteString(fmt.Sprint(item.id))
		processed.WriteString("</td>")

		processed.WriteString("<td>")
		processed.WriteString(string(item.Type))
		processed.WriteString("</td>")
		processed.WriteString("</tr>")
	}
	processed.WriteString("</table>")

	//--

	processed.WriteString("<h1>Articles</h1>")
	processed.WriteString("<table border='1'>")
	processed.WriteString("<tr><th>id</th><th>magazines_id</th><th>article_type_id</th><th>author_id</th></tr>")
	for _, item := range articles {
		processed.WriteString("<tr>")
		processed.WriteString("<td>")
		processed.WriteString(fmt.Sprint(item.id))
		processed.WriteString("</td>")

		processed.WriteString("<td>")
		processed.WriteString(fmt.Sprint(item.magazines_id))
		processed.WriteString("</td>")

		processed.WriteString("<td>")
		processed.WriteString(fmt.Sprint(item.article_type_id))
		processed.WriteString("</td>")

		processed.WriteString("<td>")
		processed.WriteString(fmt.Sprint(item.author_id))
		processed.WriteString("</td>")
		processed.WriteString("</tr>")
	}
	processed.WriteString("</table>")

	//--

	processed.WriteString("<h1>Author</h1>")
	processed.WriteString("<table border='1'>")
	processed.WriteString("<tr><th>id</th><th>author</th></tr>")
	for _, item := range authors {
		processed.WriteString("<tr>")
		processed.WriteString("<td>")
		processed.WriteString(fmt.Sprint(item.id))
		processed.WriteString("</td>")

		processed.WriteString("<td>")
		processed.WriteString(string(item.author))
		processed.WriteString("</td>")
		processed.WriteString("</tr>")
	}
	processed.WriteString("</table>")

	f, _ := os.Create("/var/www/html/test-88.ru/local/scripts/index.html")
	w := bufio.NewWriter(f)
	w.WriteString(processed.String())
	w.Flush()

	/*fmt.Println("MAGAZINES")
	for _, m := range magazines {
		fmt.Println(m.id, m.name)
	}

	fmt.Println("ARTICLES")
	for _, a := range articles {
		fmt.Println(a.id, a.magazines_id, a.article_type_id, a.author_id)
	}

	fmt.Println("ARTICLE TYPES")
	for _, at := range articleTypes {
		fmt.Println(at.id, at.Type)
	}

	fmt.Println("AUTHORS")
	for _, au := range authors {
		fmt.Println(au.id, au.author)
	}*/
}

func selectAuthors(db *sql.DB) []author {
	rows, err := db.Query("select * from Author a")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	authors := []author{}

	for rows.Next() {
		a := author{}

		err := rows.Scan(&a.id, &a.author)
		if err != nil {
			fmt.Println(err)
			continue
		}
		authors = append(authors, a)
	}
	return authors
}

func selectArticleTypes(db *sql.DB) []articleType {
	rows, err := db.Query("select * from Article_types a")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	articleTypes := []articleType{}

	for rows.Next() {
		a := articleType{}

		err := rows.Scan(&a.id, &a.Type)
		if err != nil {
			fmt.Println(err)
			continue
		}
		articleTypes = append(articleTypes, a)
	}
	return articleTypes
}

func selectArticles(db *sql.DB) []article {
	rows, err := db.Query("select * from Articles a")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	articles := []article{}

	for rows.Next() {
		a := article{}

		err := rows.Scan(&a.id, &a.magazines_id, &a.article_type_id, &a.author_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		articles = append(articles, a)
	}
	return articles
}

func selectMagazines(db *sql.DB) []magazine {
	rows, err := db.Query("select * from Magazines m")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	magazines := []magazine{}

	for rows.Next() {
		m := magazine{}

		err := rows.Scan(&m.id, &m.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		magazines = append(magazines, m)
	}
	return magazines
}
