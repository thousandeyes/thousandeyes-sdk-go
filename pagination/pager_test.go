package pagination

import (
	"context"
	"errors"
	"io"
	"iter"
	"net/http"
	"reflect"
	"testing"
)

type testPage struct {
	items    []string
	nextHref string
}

func testPager(pages []*testPage, fetchCalls *int) *Pager[testPage] {
	return NewPager(
		nil,
		func(context.Context, *string) (*testPage, *http.Response, error) {
			index := *fetchCalls
			*fetchCalls = *fetchCalls + 1
			if index >= len(pages) {
				return nil, nil, errors.New("unexpected page fetch")
			}
			return pages[index], &http.Response{StatusCode: http.StatusOK}, nil
		},
		func(page *testPage) (string, bool) {
			return page.nextHref, page.nextHref != ""
		},
	)
}

func testPageItems(page *testPage) []string {
	return page.items
}

func TestAllIsLazy(t *testing.T) {
	fetchCalls := 0
	pager := testPager([]*testPage{{items: []string{"first"}}}, &fetchCalls)

	sequence := All(context.Background(), pager, testPageItems)

	if fetchCalls != 0 {
		t.Fatalf("All fetched %d pages before iteration, want 0", fetchCalls)
	}

	for item, err := range sequence {
		if err != nil {
			t.Fatalf("All yielded error = %v", err)
		}
		if item != "first" {
			t.Fatalf("All yielded item = %q, want first", item)
		}
		break
	}

	if fetchCalls != 1 {
		t.Fatalf("All fetched %d pages after first item, want 1", fetchCalls)
	}
}

func TestAllYieldsItemsAcrossPages(t *testing.T) {
	fetchCalls := 0
	pager := testPager(
		[]*testPage{
			{items: []string{"first", "second"}, nextHref: "?cursor=page-two"},
			{items: []string{"third", "fourth"}},
		},
		&fetchCalls,
	)

	var got []string
	for item, err := range All(context.Background(), pager, testPageItems) {
		if err != nil {
			t.Fatalf("All yielded error = %v", err)
		}
		got = append(got, item)
	}

	want := []string{"first", "second", "third", "fourth"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("All yielded items = %q, want %q", got, want)
	}
	if fetchCalls != 2 {
		t.Fatalf("All fetched %d pages, want 2", fetchCalls)
	}
}

func TestAllSkipsEmptyIntermediatePages(t *testing.T) {
	fetchCalls := 0
	pager := testPager(
		[]*testPage{
			{items: []string{"first"}, nextHref: "?cursor=empty"},
			{nextHref: "?cursor=last"},
			{items: []string{"last"}},
		},
		&fetchCalls,
	)

	var got []string
	for item, err := range All(context.Background(), pager, testPageItems) {
		if err != nil {
			t.Fatalf("All yielded error = %v", err)
		}
		got = append(got, item)
	}

	want := []string{"first", "last"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("All yielded items = %q, want %q", got, want)
	}
	if fetchCalls != 3 {
		t.Fatalf("All fetched %d pages, want 3", fetchCalls)
	}
}

func TestAllYieldsFetchErrorOnceThenStops(t *testing.T) {
	wantErr := errors.New("fetch failed")
	fetchCalls := 0
	pager := NewPager(
		nil,
		func(context.Context, *string) (*testPage, *http.Response, error) {
			fetchCalls++
			if fetchCalls == 1 {
				return &testPage{
					items:    []string{"first"},
					nextHref: "?cursor=failed",
				}, &http.Response{StatusCode: http.StatusOK}, nil
			}
			return nil, nil, wantErr
		},
		func(page *testPage) (string, bool) {
			return page.nextHref, page.nextHref != ""
		},
	)

	next, stop := iter.Pull2(All(context.Background(), pager, testPageItems))
	defer stop()

	item, err, ok := next()
	if !ok || item != "first" || err != nil {
		t.Fatalf("first yield = (%q, %v, %t), want (first, nil, true)", item, err, ok)
	}

	item, err, ok = next()
	if !ok || item != "" || !errors.Is(err, wantErr) {
		t.Fatalf("error yield = (%q, %v, %t), want (zero, fetch error, true)", item, err, ok)
	}

	item, err, ok = next()
	if ok || item != "" || err != nil {
		t.Fatalf("yield after error = (%q, %v, %t), want (zero, nil, false)", item, err, ok)
	}
	if fetchCalls != 2 {
		t.Fatalf("All fetched %d pages, want 2", fetchCalls)
	}
}

func TestAllPropagatesContextCancellation(t *testing.T) {
	type contextKey struct{}

	ctx := context.WithValue(context.Background(), contextKey{}, "marker")
	ctx, cancel := context.WithCancel(ctx)
	cancel()

	fetchCalls := 0
	pager := NewPager(
		nil,
		func(gotCtx context.Context, _ *string) (*testPage, *http.Response, error) {
			fetchCalls++
			if got := gotCtx.Value(contextKey{}); got != "marker" {
				t.Errorf("fetch context value = %v, want marker", got)
			}
			return nil, nil, gotCtx.Err()
		},
		func(*testPage) (string, bool) { return "", false },
	)

	next, stop := iter.Pull2(All(ctx, pager, testPageItems))
	defer stop()

	item, err, ok := next()
	if !ok || item != "" || !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled yield = (%q, %v, %t), want (zero, context.Canceled, true)", item, err, ok)
	}

	item, err, ok = next()
	if ok || item != "" || err != nil {
		t.Fatalf("yield after cancellation = (%q, %v, %t), want (zero, nil, false)", item, err, ok)
	}
	if fetchCalls != 1 {
		t.Fatalf("All fetched %d pages, want 1", fetchCalls)
	}
}

func TestAllStopsFetchingWhenConsumerBreaks(t *testing.T) {
	fetchCalls := 0
	pager := testPager(
		[]*testPage{
			{items: []string{"first"}, nextHref: "?cursor=page-two"},
			{items: []string{"should not be fetched"}},
		},
		&fetchCalls,
	)

	for item, err := range All(context.Background(), pager, testPageItems) {
		if err != nil {
			t.Fatalf("All yielded error = %v", err)
		}
		if item != "first" {
			t.Fatalf("All yielded item = %q, want first", item)
		}
		break
	}

	if fetchCalls != 1 {
		t.Fatalf("All fetched %d pages after consumer break, want 1", fetchCalls)
	}
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
	if !pager.HasMorePages() {
		t.Fatal("new pager reported no pages")
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
		if got, want := pager.HasMorePages(), index < len(pages)-1; got != want {
			t.Fatalf("HasMorePages() after page %d = %t, want %t", index+1, got, want)
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
