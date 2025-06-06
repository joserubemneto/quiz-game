# Quiz Game

The goal of this project is to develop a question and answers game to put in practice the concepts I learned studying go lang.

The game contains some questions with multiple alternatives. The user should pick the correct answer. At the end of the game, 
a score will be generated for the user.

The game is based on a csv file that contains all the questions and answers of the game.

## More features to implement

1. Timer: Add a limit time for the user to answer. The game should end when the user reaches the limit.
2. Define a score for when the user to be approved or not. E.g., needs to score at least 30 points to be approved.
3. Have 3 csv files with the questions by theme, e.g., history, English, general knowledge. Build the csv with the questions you want, you can ask AI to generate some questions. When starting the game, list the themes of each file for the user to choose and then display the questions of the chosen theme.

## Things to improve

1. Do the proper error handling instead of using panic everywhere
2. Add tests for the system
3. Split system into multiple modules