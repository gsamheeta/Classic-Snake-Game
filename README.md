How the code works:
Snake entity: This controls the snake’s movement, direction, and growth.
Food entity: This represents the food that the snake needs to eat to grow.
Level: A container for the snake and food, managing the interaction between the two.
Game loop: The Tick method runs repeatedly and checks user input and game state (e.g., snake’s position, collision detection).


Running the game:
Install the termloop package using go get github.com/JoelOtter/termloop.
Save the code in a file, e.g., snake.go.
Run the game with:
go run snake.go
