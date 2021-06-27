package game

import "sync"

type score struct {
	player int
	bot    int
	mu     sync.Mutex
}

func (s *score) getPlayerScore() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.player
}

func (s *score) addToPlayerScore(x int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.player += x
}

func (s *score) getBotScore() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.bot
}

func (s *score) addToBotScore(x int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bot += x
}
