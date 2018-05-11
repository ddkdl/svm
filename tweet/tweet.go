package tweet

// TO DO: Add JSON encoding to each field
type Tweet struct {
	id        int
	Username  string
	Text      string
	Date      string
	Retweets  int
	Favorites int
	Mentions  string
	Hashtags  string
	Geo       string
	Permalink string
}
