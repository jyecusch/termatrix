package matrix

import (
	"math/rand"
	"time"
)

type RainLine struct {
	// length is the number of characters in the line, so the bottom of the line is at pos + length
	length int
	// speed is the time it takes for the line to move one character
	speed    time.Duration
	previous time.Time
	// pos is the position of the top of the line
	pos int
}

func NewRandomLine(viewHeight int) *RainLine {
	return &RainLine{
		length:   randLineLength(viewHeight),
		speed:    randLineSpeed(),
		pos:      rand.Intn(viewHeight) - 4,
		previous: time.Now(),
	}
}

func (l *RainLine) update(viewHeight int, now time.Time) {
	pos := l.pos

	timeSincePrevious := now.Sub(l.previous)
	l.pos = (l.pos + int(timeSincePrevious/l.speed))
	if l.pos > viewHeight {
		l.length = randLineLength(viewHeight)
		l.pos = -l.length
	}

	if pos != l.pos {
		l.previous = now
	}
}

func (l *RainLine) halfPoint() int {
	return l.headIndex() - l.length/2
}

func (l *RainLine) headIndex() int {
	return l.pos + l.length - 1
}

func (l *RainLine) isHead(charPos int) bool {
	return charPos == l.headIndex()
}

func (l *RainLine) isInside(charPos int) bool {
	return charPos >= l.pos && charPos <= l.headIndex()
}

func (l *RainLine) isTail(charPos int) bool {
	return charPos >= l.pos && charPos <= l.halfPoint()
}

func randLineLength(height int) int {
	if height < 4 {
		return height - 1
	}

	minSize := 3
	maxSize := (height / 5) * 4

	return rand.Intn(maxSize-minSize+1) + minSize
}

func randLineSpeed() time.Duration {
	return time.Second / time.Duration(rand.Intn(20)+5)
}
