package model

import "strconv"

type Stats struct {
	UID       int64   `json:"uid"`
	Range     string  `json:"range"`
	StatsName string  `json:"stats_name"`
	Version   string  `json:"version"`
	Score     float64 `json:"data"`
	Rank      int64   `json:"rank"`
}

func (s *Stats) Key() string {
	return s.Range + "_" + s.StatsName + "_" + s.Version
}

func (s *Stats) IsValid() bool {
	return s.UID > 0
}

func (s *Stats) Member() string {
	return strconv.FormatInt(s.UID, 10)
}
