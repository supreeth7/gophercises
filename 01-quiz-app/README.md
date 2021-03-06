# CLI Quiz App

- This golang program reads a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question is asked immediately afterwards.
- The quiz has a timer which is set by a ```limit``` flag in seconds.
- The quiz stops as soon as the time limit has exceeded.
- A final score is printed to the console at the end of the quiz.

### Syntax

```
./01-quiz-app [OPTIONS]
```

### Flags

- ```csv``` - The name of the CSV file containing the quiz. Defaults to 'quiz.csv'.
- ```limit``` - The time limit of the quiz. Defaults to 30 seconds.


### Example

```
./01-quiz-app -csv=myquiz.csv -limit=10

```

### Instructions

- Use the ```go build``` command to generate the ```01-quiz-app``` binary file in the current directory.
- Run```./01-quiz-app``` in your terminal to run the program with its default flags.

