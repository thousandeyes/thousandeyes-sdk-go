package pagination

import (
	"context"
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/url"
)

var (
	errInvalidContinuation  = errors.New("pagination: next link must contain exactly one non-empty cursor")
	errRepeatedContinuation = errors.New("pagination: next link repeated a cursor")
)

// Pager fetches response pages lazily.
type Pager[T any] struct {
	cursor      *string
	seenCursors map[string]struct{}
	fetch       func(context.Context, *string) (*T, *http.Response, error)
	nextHref    func(*T) (string, bool)
	complete    bool
}

// NewPager creates a Pager.
func NewPager[T any](
	initialCursor *string,
	fetch func(context.Context, *string) (*T, *http.Response, error),
	nextHref func(*T) (string, bool),
) *Pager[T] {
	if initialCursor != nil {
		cursor := *initialCursor
		initialCursor = &cursor
	}
	seenCursors := make(map[string]struct{})
	if initialCursor != nil {
		seenCursors[*initialCursor] = struct{}{}
	}
	return &Pager[T]{
		cursor:      initialCursor,
		seenCursors: seenCursors,
		fetch:       fetch,
		nextHref:    nextHref,
	}
}

// All returns a lazy, single-use item sequence.
func All[Page, Item any](
	ctx context.Context,
	pager *Pager[Page],
	items func(*Page) []Item,
) iter.Seq2[Item, error] {
	return func(yield func(Item, error) bool) {
		for pager.HasMorePages() {
			page, _, err := pager.NextPage(ctx)
			if err != nil {
				var zero Item
				yield(zero, err)
				return
			}
			for _, item := range items(page) {
				if !yield(item, nil) {
					return
				}
			}
		}
	}
}

// HasMorePages reports whether another page may be fetched.
func (p *Pager[T]) HasMorePages() bool {
	return !p.complete
}

// NextPage fetches the next response page.
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
	if _, repeated := p.seenCursors[cursor]; repeated {
		return page, response, errRepeatedContinuation
	}
	p.seenCursors[cursor] = struct{}{}
	p.cursor = &cursor
	return page, response, nil
}

func cursorFromHref(href string) (string, error) {
	parsed, err := url.Parse(href)
	if err != nil {
		return "", fmt.Errorf("%w: %v", errInvalidContinuation, err)
	}
	query, err := url.ParseQuery(parsed.RawQuery)
	if err != nil {
		return "", fmt.Errorf("%w: %v", errInvalidContinuation, err)
	}
	cursors := query["cursor"]
	if len(cursors) != 1 || cursors[0] == "" {
		return "", errInvalidContinuation
	}
	return cursors[0], nil
}
