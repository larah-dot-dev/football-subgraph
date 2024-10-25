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

	c := gocal.NewParser(f)
	start, end := time.Now(), time.Now().Add(10*24*time.Hour)
	c.Start, c.End = &start, &end
	c.Parse()

	data := make([]*model.Match, len(c.Events))

	for i, event := range c.Events {
		split := strings.Split(event.Summary, " - ")
		teams, tournament := split[0], split[1]
		home, away := strings.Split(teams, " vs ")[0], strings.Split(teams, " vs ")[1]

		url := fmt.Sprintf("https://www.manutd.com/en/matches/matchcenter?matchId=%s", event.Uid)

		data[i] = &model.Match{
			ID: fmt.Sprintf("fixture%d", i),
			RawTime:	*event.Start,
			Home: &model.Team{
				ID:        strings.ReplaceAll(strings.ToLower(home), " ", "-"),
				Name:      home,
			},
			Away: &model.Team{
				ID:        strings.ReplaceAll(strings.ToLower(away), " ", "-"),
				Name:      away,
			},
			Kickoff:    event.Start.Format(time.RFC3339),
			Tournament: tournament,
			URL: 	  	url,
		}
	}

	slices.SortFunc(data, func(a, b *model.Match) int {
		return a.RawTime.Compare(b.RawTime)
	})

	return data, nil
}

var GetMatches = memoize(getMatches)
