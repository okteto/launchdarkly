package cmd

type environmentSource struct {
	Key string
}

type environment struct {
	Key    string
	Name   string
	Color  string
	Source environmentSource
}
