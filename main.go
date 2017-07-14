package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var ck, cs, at, ats string
var n int

var tootTmpl = template.Must(template.New("toot").Funcs(template.FuncMap{"time2local": time2local}).Parse(`-=-=-=-=-=-=-=-

> {{ .Text }}

{{ .Name }} (@{{ .ScreenName }}) @ {{ time2local .Time }}
https://twitter.com/{{ .ScreenName }}/status/{{ .ID }}

`))

func main() {
	flag.StringVar(&at, "access-token", "", "Your access token")
	flag.StringVar(&ats, "access-token-secret", "", "Your access token secret")
	flag.StringVar(&ck, "consumer-key", "", "Your application's consumer key")
	flag.StringVar(&cs, "consumer-secret", "", "Your application's consumer secret")
	flag.IntVar(&n, "num", 50, "number of tweets to return")

	flag.Parse()

	if ck == "" || cs == "" || at == "" || ats == "" {
		fmt.Printf("error: incomplete oAuth credentials")
		os.Exit(-1)
	}

	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, ats)
	client := twitter.NewClient(config.Client(oauth1.NoContext, token))

	// Home Timeline
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: n,
	})
	if err != nil {
		fmt.Println("fatal error:", err, resp)
		os.Exit(-1)
	}

	for _, t := range tweets {
		if t.RetweetedStatus != nil || t.QuotedStatus != nil {
			continue
		}

		ca, _ := time.Parse(time.RubyDate, t.CreatedAt)

		if err := tootTmpl.Execute(os.Stdout, struct {
			Name, ScreenName, Text string
			ID                     int64
			Time                   time.Time
		}{
			Name:       t.User.Name,
			ScreenName: t.User.ScreenName,
			Text:       t.Text,
			Time:       ca,
			ID:         t.ID,
		}); err != nil {
			panic(err)
		}
	}

	os.Exit(0)
}

func time2local(t time.Time) string {
	return t.In(time.Local).Format(time.Kitchen)
}
