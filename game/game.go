package game

import (
	"bufio"
	"os"
	"strings"
)

type Game struct {
	allWords  map[rune][]string
	usedWords map[string]struct{}
}

func NewGame() (*Game, error) {
	f, err := os.Open("words.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	allWords := make(map[rune][]string)
	for {
		word, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		word = strings.Trim(word, "\n")
		wordRune := []rune(word)
		allWords[wordRune[0]] = append(allWords[wordRune[0]], word)
	}
	return &Game{
		allWords:  allWords,
		usedWords: make(map[string]struct{}),
	}, nil
}

func (g *Game) Word(word string) string {
	if _, ok := g.usedWords[word]; ok {
		return "Слово уже использовано"
	}

	wordRune := []rune(word)
	allWords := g.allWords[wordRune[0]]
	flag := false
	for _, allWord := range allWords {
		if allWord == word {
			flag = true
			break
		}
	}
	if !flag {
		return "Слово не найдено"
	}
	g.usedWords[word] = struct{}{}

	answers := g.allWords[wordRune[len(wordRune)-1]]
	for _, answer := range answers {
		if _, ok := g.usedWords[answer]; !ok {
			return answer
		}
	}

	return "ты победил!"
}
