package dexigo

import (
	"errors"
)

type Confirm string

const (
	Uncompleted Confirm = "0"
	Completed   Confirm = "1"
)

type CandlesTick struct {
	Timestamp string
	Open      string
	High      string
	Low       string
	Close     string
	Volume    string
	VolumeUSD string
	Confirm   Confirm
}

func ParseCandlestick(data []string) (CandlesTick, error) {
	if len(data) != 8 {
		return CandlesTick{}, errors.New("invalid data")
	}
	return CandlesTick{
		Timestamp: data[0],
		Open:      data[1],
		High:      data[2],
		Low:       data[3],
		Close:     data[4],
		Volume:    data[5],
		VolumeUSD: data[6],
		Confirm:   Confirm(data[7]),
	}, nil
}
