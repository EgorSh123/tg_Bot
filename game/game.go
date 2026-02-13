package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	allWords  map[rune][]string
	usedWords map[string]struct{}
	lastWord  string
	attempts  int
}

func NewGame() (*Game, error) {
	f, err := os.Open("fruits.txt")
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
		attempts:  5,
	}, nil
}

func (g *Game) Word(word string) string {
	if len(g.lastWord) > 0 && lastLetter(g.lastWord) != firstLetter(word) {
		g.attempts--
		if g.attempts == 0 {
			return "ты проиграл!"
		}
		attemptsWord := getAttemptsWord(g.attempts)
		return fmt.Sprintf("Неверное начало слова. У тебя осталось %d %s", g.attempts, attemptsWord)
	}

	if _, ok := g.usedWords[word]; ok {
		g.attempts--
		if g.attempts == 0 {
			return "ты проиграл!"
		}
		attemptsWord := getAttemptsWord(g.attempts)
		return fmt.Sprintf("Слово уже использовано. У тебя осталось %d %s", g.attempts, attemptsWord)
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
		g.attempts--
		if g.attempts == 0 {
			return "ты проиграл!"
		}
		attemptsWord := getAttemptsWord(g.attempts)
		return fmt.Sprintf("Слово не найдено. У тебя осталось %d %s", g.attempts, attemptsWord)
	}
	g.attempts = 5
	g.usedWords[word] = struct{}{}

	answers := g.allWords[lastLetter(word)]
	for _, answer := range answers {
		if _, ok := g.usedWords[answer]; !ok {
			g.lastWord = answer
			g.usedWords[answer] = struct{}{}
			return answer
		}
	}

	return "ты победил!"
}

func (g *Game) RestartGame() {
	g.usedWords = make(map[string]struct{})
	g.attempts = 5
	g.lastWord = ""
}

func lastLetter(word string) rune {
	runes := []rune(strings.ToLower(word))
	last := runes[len(runes)-1]
	if last == 'ь' || last == 'ъ' || last == 'ы' {
		last = runes[len(runes)-2]
	}
	return last
}

func firstLetter(word string) rune {
	runes := []rune(strings.ToLower(word))
	return runes[0]
}

func getAttemptsWord(attempts int) string {
	switch attempts {
	case 1:
		return "попытка"
	case 2, 3, 4:
		return "попытки"
	default:
		return "попыток"
	}
}
