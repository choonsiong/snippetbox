package mock

import (
	"github.com/choonsiong/snippetbox/pkg/models"
	"time"
)

// Create a simple struct which implements the same methods as our production
// mysql.SnippetModel, but have the methods return some fixed dummy data
// instead.

var mockSnippet = &models.Snippet{
	ID: 1,
	Title: "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct {}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

// Get returns a models.ErrNoRecord unless the snippet ID is 1.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}