package football_data

import (
	"fmt"
	"strings"
	"os"
	"time"
	"slices"

	"football-subgraph/graph/model"
	"github.com/apognu/gocal"

)

type memo[T any] struct {
	written_at time.Time
	data       T
}

func memoize[Out any](cb func() (Out, error)) func() (Out, error) {
	cache := new(memo[Out])

	return func() (Out, error) {
		one_hour_ago := time.Now().Add(-time.Hour)
		var err error
	
		if cache.written_at.IsZero() || cache.written_at.Before(one_hour_ago) {
			cache.data, err = cb()
			if err == nil {
				cache.written_at = time.Now()
			}
		}

		return cache.data, err
	}
}

func getMatches() (matches []*model.Match, err error) {
	ics_file, err := DownloadIcsFile()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(ics_file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	start, end := time.Now(), time.Now().Add(10*24*time.Hour)
	c := gocal.NewParser(f)
	c.Start, c.End = &start, &end
	c.Parse()

	data := make([]*model.Match, len(c.Events))

	for i, event := range c.Events {
		split := strings.Split(event.Summary, " at ")
		home, away := split[0], split[1]

		data[i] = &model.Match{
			ID: fmt.Sprintf("fixture%d", i),
			Home: &model.Team{
				ID:        "",
				Name:      home,
				Shortcode: "",
				Crest:     "",
			},
			Away: &model.Team{
				ID:        "",
				Name:      away,
				Shortcode: "",
				Crest:     "",
			},
			Kickoff:    event.Start.Format(time.RFC3339),
			Tournament: "tournament",
			RawTime:	*event.Start,
		}
	}

	slices.SortFunc(data, func(a, b *model.Match) int {
		return a.RawTime.Compare(b.RawTime)
	})

	return data, nil
}

var GetMatches = memoize(getMatches)
