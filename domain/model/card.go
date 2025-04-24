package model

import "time"

type Card struct {
	ID          uint       `gorm:"primaryKey"`
	Alphabet    string     `json:"alphabet"`
	Comment     string     `json:"comment"`
	Initial     string     `json:"initial"`
	PictureCard string     `json:"picture_card"`
	Text        string     `json:"text"`
	TextCard    string     `json:"text_card"`
	Yomi        string     `json:"yomi"`
	Score       int        `json:"score"`
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Syllabary string

const (
	SyllabaryA  Syllabary = "あ"
	SyllabaryKa Syllabary = "か"
	SyllabarySa Syllabary = "さ"
	SyllabaryTa Syllabary = "た"
	SyllabaryNa Syllabary = "な"
	SyllabaryHa Syllabary = "は"
	SyllabaryMa Syllabary = "ま"
	SyllabaryYa Syllabary = "や"
	SyllabaryRa Syllabary = "ら"
	SyllabaryWa Syllabary = "わ"
)

func GetInitialsBySyllabary(syllabary Syllabary) []string {
	switch syllabary {
	case SyllabaryA:
		return []string{"あ", "い", "う", "え", "お"}
	case SyllabaryKa:
		return []string{"か", "き", "く", "け", "こ"}
	case SyllabarySa:
		return []string{"さ", "し", "す", "せ", "そ"}
	case SyllabaryTa:
		return []string{"た", "ち", "つ", "て", "と"}
	case SyllabaryNa:
		return []string{"な", "に", "ぬ", "ね", "の"}
	case SyllabaryHa:
		return []string{"は", "ひ", "ふ", "へ", "ほ"}
	case SyllabaryMa:
		return []string{"ま", "み", "む", "め", "も"}
	case SyllabaryYa:
		return []string{"や", "ゆ", "よ"}
	case SyllabaryRa:
		return []string{"ら", "り", "る", "れ", "ろ"}
	case SyllabaryWa:
		return []string{"わ"}
	default:
		return []string{}
	}
}
