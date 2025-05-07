package model

import "time"

type EventScore struct {
	ID         uint64     `json:"id"`
	Rank       int        `json:"rank"`
	TotalUsers int        `json:"total_users"`
	Name       string     `json:"name"`
	Feeling    string     `json:"feeling"`
	Score      int        `json:"score"`
	CardsTaken int        `json:"cards_taken"`
	FaultCount int        `json:"fault_count"`
	CreatedAt  *time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type EventScoreForCreate struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	Feeling    string     `json:"feeling"`
	Score      int        `json:"score"`
	CardsTaken int        `json:"cards_taken"`
	FaultCount int        `json:"fault_count"`
	CreatedAt  *time.Time `gorm:"autoCreateTime" json:"created_at"`
}

const (
	// カード取得枚数あたりの乗数 (cards_taken * CardsTakenMultiplier)
	CardsTakenMultiplier = 1200000
	// お手付き回数あたりの減算量 (fault_count * FaultCountMultiplier)
	FaultCountMultiplier = 25000
	// スコアを割る除数
	ScoreDivisor = 10000
)
