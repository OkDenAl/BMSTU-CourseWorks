package domain

import "encoding/json"

type StoryStat struct {
	StoryID string
	Count   int
}

func (ss StoryStat) MarshalBinary() ([]byte, error) {
	return json.Marshal(ss)
}

func (ss StoryStat) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, &ss)
}
