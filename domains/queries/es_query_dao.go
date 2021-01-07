package queries

import (
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"io"
)

func (q EsQuery) Build() io.Reader {
	var shouldQuery []map[string]interface{}
	for _, item := range q.Equals {
		terms := map[string]interface{}{
			"term": map[string]interface{}{
				item.Field: item.Value,
			},
		}
		shouldQuery = append(shouldQuery, terms)
	}

	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"bool": map[string]interface{}{
				"should": shouldQuery,
			},
		},
	}

	return esutil.NewJSONReader(query)
}
