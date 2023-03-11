package templates

import (
	"context"
	"database/sql"

	"github.com/gocopper/copper/csql"
)

var ErrRecordNotFound = sql.ErrNoRows

func NewQueries(querier csql.Querier) *Queries {
	return &Queries{
		querier: querier,
	}
}

type Queries struct {
	querier csql.Querier
}

// Here are some example queries that use Querier to unmarshal results into Go strcuts
func (q *Queries) ListTemplates(ctx context.Context) ([]Template, error) {
	const query = "SELECT * FROM templates ORDER BY created_at DESC"

	var (
		templates []Template
		err       = q.querier.Select(ctx, &templates, query)
	)

	return templates, err
}

func (q *Queries) GetTemplateByID(ctx context.Context, id string) (*Template, error) {
	const query = "SELECT * from templates where id=?"

	var (
		template Template
		err      = q.querier.Get(ctx, &template, query, id)
	)

	return &template, err
}

func (q *Queries) SaveTemplate(ctx context.Context, template *Template) error {
	const query = `
	INSERT INTO templates (id, subject, description, assessment, recommendation)
	VALUES (?, ?, ?, ?, ?)
	ON CONFLICT (id) DO UPDATE SET subject=?, description=? assessment=?, recommendation=?`

	_, err := q.querier.Exec(ctx, query,
		template.ID,
		template.Subject,
		template.Description,
		template.Assessment,
		template.Recommendation,
	)

	return err
}