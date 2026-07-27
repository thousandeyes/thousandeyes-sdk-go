package pagination

import (
	"context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"
)

type testPage struct {
	items    []string
	nextHref string
}

func TestPagerTraversesPagesLazily(t *testing.T) {
	pages := []*testPage{
		{items: []string{"first"}, nextHref: "https://other.example/wrong?cursor=second"},
		{items: []string{"second"}},
	}
	responses := []*http.Response{{StatusCode: http.StatusOK}, {StatusCode: http.StatusCreated}}
	var cursors []string

	fetch := func(_ context.Context, cursor *string) (*testPage, *http.Response, error) {
		if cursor == nil {
			cursors = append(cursors, "")
		} else {
			cursors = append(cursors, *cursor)
		}
		index := len(cursors) - 1
		return pages[index], responses[index], nil
	}
	nextHref := func(page *testPage) (string, bool) {
		return page.nextHref, page.nextHref != ""
	}

	pager := NewPager(nil, fetch, nextHref)
	if len(cursors) != 0 {
		t.Fatal("NewPager fetched a page")
	}

	for index, wantPage := range pages {
		gotPage, gotResponse, err := pager.NextPage(context.Background())
		if err != nil {
			t.Fatalf("NextPage(%d) error = %v", index, err)
		}
		if !reflect.DeepEqual(gotPage, wantPage) {
			t.Fatalf("NextPage(%d) page = %#v, want %#v", index, gotPage, wantPage)
		}
		if gotResponse != responses[index] {
			t.Fatalf("NextPage(%d) response = %p, want %p", index, gotResponse, responses[index])
		}
	}

	if !reflect.DeepEqual(cursors, []string{"", "second"}) {
		t.Fatalf("fetch cursors = %q, want only original cursor then next-link cursor", cursors)
	}

	gotPage, gotResponse, err := pager.NextPage(context.Background())
	if !errors.Is(err, io.EOF) || gotPage != nil || gotResponse != nil {
		t.Fatalf("NextPage after completion = (%#v, %#v, %v), want (nil, nil, io.EOF)", gotPage, gotResponse, err)
	}
	if len(cursors) != len(pages) {
		t.Fatalf("fetch called %d times after completion, want %d", len(cursors), len(pages))
	}
}

func TestPagerRejectsInvalidContinuation(t *testing.T) {
	tests := map[string]string{
		"missing":    "https://other.example/wrong?not_cursor=missing",
		"empty":      "https://other.example/wrong?cursor=",
		"duplicate":  "https://other.example/wrong?cursor=one&cursor=two",
		"wrong case": "https://other.example/wrong?Cursor=next",
	}

	for name, href := range tests {
		t.Run(name, func(t *testing.T) {
			page := &testPage{nextHref: href}
			pager := NewPager(
				nil,
				func(context.Context, *string) (*testPage, *http.Response, error) {
					return page, &http.Response{StatusCode: http.StatusOK}, nil
				},
				func(page *testPage) (string, bool) { return page.nextHref, true },
			)

			if _, _, err := pager.NextPage(context.Background()); !errors.Is(err, ErrInvalidContinuation) {
				t.Fatalf("NextPage error = %v, want ErrInvalidContinuation", err)
			}
		})
	}
}

func TestPagerUsesInitialCursor(t *testing.T) {
	initialCursor := "configured"
	var fetchedCursor string
	pager := NewPager(
		&initialCursor,
		func(_ context.Context, cursor *string) (*testPage, *http.Response, error) {
			fetchedCursor = *cursor
			return &testPage{}, &http.Response{StatusCode: http.StatusOK}, nil
		},
		func(*testPage) (string, bool) { return "", false },
	)
	initialCursor = "changed after construction"

	if _, _, err := pager.NextPage(context.Background()); err != nil {
		t.Fatalf("NextPage error = %v", err)
	}
	if fetchedCursor != "configured" {
		t.Fatalf("fetch cursor = %q, want configured", fetchedCursor)
	}
}

func TestPagerPreservesHTTPResponseOnFetchError(t *testing.T) {
	wantErr := errors.New("fetch failed")
	wantResponse := &http.Response{StatusCode: http.StatusBadGateway}
	pager := NewPager(
		nil,
		func(context.Context, *string) (*testPage, *http.Response, error) {
			return nil, wantResponse, wantErr
		},
		func(*testPage) (string, bool) { return "", false },
	)

	page, response, err := pager.NextPage(context.Background())
	if page != nil || response != wantResponse || !errors.Is(err, wantErr) {
		t.Fatalf("NextPage = (%#v, %#v, %v), want (nil, response, fetch error)", page, response, err)
	}
}
