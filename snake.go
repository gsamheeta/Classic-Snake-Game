package main

import (
	"github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

// Snake represents the snake in the game
type Snake struct {
	*termloop.Entity
	body        []termloop.Cell
	direction   string
	gameOver    bool
	grow        bool
	width, height int
}

// NewSnake creates a new snake at the starting position
func NewSnake(width, height int) *Snake {
	snake := &Snake{
		Entity:    termloop.NewEntity(1, 1, 1, 1),
		direction: "right",
		width:     width,
		height:    height,
	}
	snake.body = append(snake.body, termloop.Cell{Ch: 'O'})
	snake.SetPosition(2, 2)
	return snake
}

// Tick updates the snake's position and handles input
func (s *Snake) Tick(ev termloop.Event) {
	if s.gameOver {
		return
	}

	switch ev.Type {
	case termloop.EventKey:
		switch ev.Key {
		case termloop.KeyArrowRight:
			if s.direction != "left" {
				s.direction = "right"
			}
		case termloop.KeyArrowLeft:
			if s.direction != "right" {
				s.direction = "left"
			}
		case termloop.KeyArrowUp:
			if s.direction != "down" {
				s.direction = "up"
			}
		case termloop.KeyArrowDown:
			if s.direction != "up" {
				s.direction = "down"
			}
		}
	}

	s.move()
}

// Move the snake in the current direction
func (s *Snake) move() {
	x, y := s.Position()

	switch s.direction {
	case "right":
		x++
	case "left":
		x--
	case "up":
		y--
	case "down":
		y++
	}

	// Check if the snake hits the wall
	if x < 0 || y < 0 || x >= s.width || y >= s.height {
		s.gameOver = true
		return
	}

	s.SetPosition(x, y)

	// Check if snake needs to grow
	if s.grow {
		s.body = append(s.body, termloop.Cell{Ch: 'O'})
		s.grow = false
	} else {
		s.body = s.body[1:]
	}
}

// Level is the game level where the snake and food exist
type Level struct {
	*termloop.BaseLevel
	snake *Snake
	food  *Food
}

// NewLevel creates a new game level with a snake and food
func NewLevel(width, height int) *Level {
	level := &Level{
		BaseLevel: termloop.NewBaseLevel(termloop.Cell{Ch: ' '}),
		snake:     NewSnake(width, height),
		food:      NewFood(width, height),
	}
	level.AddEntity(level.snake)
	level.AddEntity(level.food)
	return level
}

// Food represents the food that the snake eats
type Food struct {
	*termloop.Entity
	width, height int
}

// NewFood creates new food at a random position
func NewFood(width, height int) *Food {
	food := &Food{
		Entity: termloop.NewEntity(1, 1, 1, 1),
		width:  width,
		height: height,
	}
	food.randomizePosition()
	return food
}

// Tick checks if the snake has eaten the food
func (f *Food) Tick(ev termloop.Event) {
	x, y := f.Position()
	sx, sy := f.Level().(*Level).snake.Position()

	if x == sx && y == sy {
		f.randomizePosition()
		f.Level().(*Level).snake.grow = true
	}
}

// RandomizePosition changes the food's location to a random position
func (f *Food) randomizePosition() {
	rand.Seed(time.Now().UnixNano())
	f.SetPosition(rand.Intn(f.width), rand.Intn(f.height))
}

func main() {
	game := termloop.NewGame()
	game.SetDebugOn(true)

	level := NewLevel(40, 20)
	game.Screen().SetLevel(level)

	game.Start()
}
