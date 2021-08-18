package results

type MyRSS struct {
	Website string
	Entries []Entry
}

type Entry struct {
	Title     string
	Url       string
	Published string
}
