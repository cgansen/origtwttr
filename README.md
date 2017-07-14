# origtwttr

A command-line application that shows only original tweets from your timeline. This excludes retweets, quote tweets, and tweets that people you follow have liked.

## Usage

```
$ go get github.com/cgansen/origtwttr
```

Note: this assumes you have [Go](https://golang.org/dl/) installed and working on your local computer.

```
$ origtwttr -h
Usage of origtwttr:
  -access-token string
    	Your access token
  -access-token-secret string
    	Your access token secret
  -consumer-key string
    	Your application's consumer key
  -consumer-secret string
    	Your application's consumer secret
  -num int
    	number of tweets to return (default 50)
```

You must create a Twitter application in order to read your timeline. To do this:

1. Sign in to Twitter
2. Create a new Twitter application at https://apps.twitter.com/app/new
3. Create a new access token in the "Keys and Access Tokens" tab.