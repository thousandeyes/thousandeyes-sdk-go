package pagination

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ErrInvalidContinuation indicates that a next link has no single, non-empty cursor.
var ErrInvalidContinuation = errors.New("pagination: next link must contain exactly one non-empty cursor")

// Pager fetches complete response pages on demand.
type Pager[T any] struct {
	cursor   *string
	fetch    func(context.Context, *string) (*T, *http.Response, error)
	nextHref func(*T) (string, bool)
	complete bool
}

// NewPager creates a lazy page iterator.
func NewPager[T any](
	initialCursor *string,
	fetch func(context.Context, *string) (*T, *http.Response, error),
	nextHref func(*T) (string, bool),
) *Pager[T] {
	if initialCursor != nil {
		cursor := *initialCursor
		initialCursor = &cursor
	}
	return &Pager[T]{cursor: initialCursor, fetch: fetch, nextHref: nextHref}
}

// HasMorePages reports whether NextPage may fetch another page.
func (p *Pager[T]) HasMorePages() bool {
	return !p.complete
}

// NextPage fetches one page or returns io.EOF after the final page.
func (p *Pager[T]) NextPage(ctx context.Context) (*T, *http.Response, error) {
	if p.complete {
		return nil, nil, io.EOF
	}

	page, response, err := p.fetch(ctx, p.cursor)
	if err != nil {
		return page, response, err
	}

	href, present := p.nextHref(page)
	if !present {
		p.complete = true
		return page, response, nil
	}

	cursor, err := cursorFromHref(href)
	if err != nil {
		return page, response, err
	}
	p.cursor = &cursor
	return page, response, nil
}

func cursorFromHref(href string) (string, error) {
	parsed, err := url.Parse(href)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrInvalidContinuation, err)
	}
	query, err := url.ParseQuery(parsed.RawQuery)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrInvalidContinuation, err)
	}
	cursors := query["cursor"]
	if len(cursors) != 1 || cursors[0] == "" {
		return "", ErrInvalidContinuation
	}
	return cursors[0], nil
}
