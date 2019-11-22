package adventure

import (
	"encoding/json"
	"io"
)

type Story map[string]Storyarc

func Jsonparser(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	return story, nil
}

type Option struct {
	//USED TO STORE OPTION
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Storyarc struct {
	//USED TO STORE STORY
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
