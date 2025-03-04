package domain

import (
	"regexp"
)

type Meat struct {
	content       string
	counting_meat map[string]int
}

func NewMeat() *Meat {
	m := &Meat{
		content:       "",
		counting_meat: make(map[string]int),
	}
	return m
}

func (m *Meat) SetContent(content string) {
	m.content = content
}

func (m *Meat) GetContent() string {
	return m.content
}

func (m *Meat) GetMeats() map[string]int {
	return m.counting_meat
}

func (m *Meat) CountingMeat() {
	meat_list := m.extractMeat()
	for _, meat := range meat_list {
		m.counting_meat[meat]++
	}
}

func (m *Meat) extractMeat() []string {
	re := regexp.MustCompile(`\b[a-zA-Z]+\b`)

	// Find all words in the text
	words := re.FindAllString(m.content, -1)

	// Print extracted words
	return words
}
