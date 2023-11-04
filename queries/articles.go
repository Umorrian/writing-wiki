package queries

import (
	"arnesteen.de/writing-wiki/queries/gen"
	"context"
	"log"
)

func (db *DB) GetArticleList(ctx context.Context) []gen.Article {
	result := make(chan []gen.Article)
	db.WorkChannel <- func(queries *gen.Queries) {
		articles, err := queries.SelectAllArticles(ctx)
		if err != nil {
			log.Printf("Error fetching data")
			result <- []gen.Article{}
		}
		result <- articles
	}
	return <-result
}

func (db *DB) GetArticleByName(ctx context.Context, name string) *gen.Article {
	result := make(chan *gen.Article)
	db.WorkChannel <- func(queries *gen.Queries) {
		fetchedArticle, err := queries.SelectArticleByName(ctx, name)
		if err != nil {
			log.Printf("Error fetching data")
			return
		}
		result <- &fetchedArticle
	}
	return <-result
}
