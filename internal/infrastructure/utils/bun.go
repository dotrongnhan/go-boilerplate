package utils

import (
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type CommonQueryOptions struct {
	Id             int64
	Ids            []int64
	SinceId        int64
	CreatedAtMin   *time.Time
	CreatedAtMax   *time.Time
	UpdatedAtMin   *time.Time
	UpdatedAtMax   *time.Time
	Limit          uint64
	Page           uint64
	Prefix         string
	NoLimit        bool
	ExcludeIds     []int64
	OrderBy        string
	OrderDirection string
}

func AddUpdateCols(q *bun.UpdateQuery, updatedFields ...string) *bun.UpdateQuery {
	if len(updatedFields) > 0 {
		for _, field := range updatedFields {
			q = q.Column(field)
		}
	}
	return q
}

func AddCommonConditions(qb *bun.SelectQuery, opts CommonQueryOptions) *bun.SelectQuery {
	// Ignore paginate for request get by specific ids
	if !opts.NoLimit && len(opts.Ids) < 1 {
		opts.Limit = GetQueryLimit(opts.Limit, 50, 250)
		qb = qb.Limit(int(opts.Limit))
		qb = qb.Offset(int(GetQueryOffset(opts.Limit, opts.Page)))
	}

	if opts.Id > 0 {
		qb = AddWhereClauseWithPrefix(qb, "%s = ?", opts.Prefix, "id", opts.Id)
	}

	if len(opts.Ids) > 0 {
		qb = AddWhereClauseWithPrefix(qb, "%s IN (?)", opts.Prefix, "id", bun.In(opts.Ids))
	}

	if len(opts.ExcludeIds) > 0 {
		qb = AddWhereClauseWithPrefix(qb, "%s NOT IN (?)", opts.Prefix, "id", bun.In(opts.ExcludeIds))
	}

	if opts.SinceId > 0 {
		qb = AddWhereClauseWithPrefix(qb, "%s > ?", opts.Prefix, "id", opts.SinceId)
	}

	if opts.SinceId < 0 {
		qb = AddWhereClauseWithPrefix(qb, "%s < ?", opts.Prefix, "id", -1*opts.SinceId)
	}

	if opts.CreatedAtMin != nil {
		qb = AddWhereClauseWithPrefix(qb, "%s >= ?", opts.Prefix, "created_at", opts.CreatedAtMin)
	}

	if opts.CreatedAtMax != nil {
		qb = AddWhereClauseWithPrefix(qb, "%s <= ?", opts.Prefix, "created_at", opts.CreatedAtMax)
	}

	if opts.UpdatedAtMin != nil {
		qb = AddWhereClauseWithPrefix(qb, "%s >= ?", opts.Prefix, "updated_at", opts.UpdatedAtMin)
	}

	if opts.UpdatedAtMax != nil {
		qb = AddWhereClauseWithPrefix(qb, "%s <= ?", opts.Prefix, "updated_at", opts.UpdatedAtMax)
	}

	if len(opts.OrderBy) > 0 {
		if len(opts.OrderDirection) == 0 {
			opts.OrderDirection = "asc"
		}
		if opts.OrderDirection != "asc" {
			opts.OrderDirection = "desc"
		}
		qb = qb.Order(fmt.Sprintf("%s %s", AddPrefixToColumn(opts.Prefix, opts.OrderBy), opts.OrderDirection))
	}
	return qb
}

func AddWhereClauseWithPrefix(qb *bun.SelectQuery, format string, prefix string, columnName string, value interface{}) *bun.SelectQuery {
	return qb.Where(fmt.Sprintf(format, bun.Ident(AddPrefixToColumn(prefix, columnName))), value)
}

func AddPrefixToColumn(prefix string, columnName string) string {
	if prefix != "" {
		columnName = prefix + "." + columnName
	}
	return columnName
}

// GetQueryLimit Check and add default limit or maximum limit in cases limit input out of the allowed range
func GetQueryLimit(limitInput, defaultLimit, maximumLimit uint64) uint64 {
	if limitInput < 1 {
		return defaultLimit
	} else if limitInput > maximumLimit {
		return maximumLimit
	}

	return limitInput
}

// GetQueryOffset Give me a number of query limit and the page number, I give you the number of query offset
func GetQueryOffset(limit, pageInput uint64) uint64 {
	if pageInput < 1 {
		pageInput = 1
	}

	return pageInput*limit - limit
}
