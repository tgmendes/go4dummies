package yelp_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/tgmendes/go4dummies/internal/platform/yelp"
	"gopkg.in/mgo.v2"
)

type MockDB struct {
	executeFn     func(collName string, f func(*mgo.Collection) error) error
	executeCalled int
}

func (db *MockDB) Execute(collName string, f func(*mgo.Collection) error) error {
	db.executeCalled++
	return db.executeFn(collName, f)
}

func TestCrawler(t *testing.T) {
	// given
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := `
			<body>
				<h3 class="heading">
					<a>FooBar</a>
				</h3>
			</body>
		`
		w.Write([]byte(body))
	}))
	defer ts.Close()

	mockDb := &MockDB{
		executeFn: func(collName string, f func(*mgo.Collection) error) error {
			return nil
		},
	}

	c := yelp.Crawler{
		BaseURL:   ts.URL + "?page=%s&start=%d",
		DB:        mockDb,
		Locations: []string{"foo"},
		Pages:     []int{1},
	}

	c.CollectRestaurants()

	if mockDb.executeCalled != 1 {
		t.Logf("expected: 1")
		t.Logf("got: %d", mockDb.executeCalled)
		t.Fatal("number of calls do not match")
	}
}

func TestCrawl(t *testing.T) {
	// given
	restCh := make(chan yelp.Restaurant)
	doneCh := make(chan bool)

	expRest := yelp.Restaurant{
		Name:     "FooBar",
		Location: "foo",
		Page:     1,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := `
			<body>
				<h3 class="heading">
					<a>FooBar</a>
				</h3>
			</body>
		`
		w.Write([]byte(body))
	}))
	defer ts.Close()

	// when
	go yelp.Crawl(ts.URL+"?page=%s&start=%d", "foo", 1, restCh, doneCh)

	// then
	for c := 0; c < 1; {
		select {
		case rest := <-restCh:
			if !reflect.DeepEqual(rest, expRest) {
				t.Logf("expected: %#v", expRest)
				t.Logf("got %#v", rest)
				t.Fatal("restaurants do not match")
			}
		case <-doneCh:
			c++
		}
	}
}
