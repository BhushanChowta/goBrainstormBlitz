# goBrainstormBlitz

Simple command-line math (or any Q/A) quiz in Go. It reads `questions.csv`, randomizes questions, collects answers, and tracks score.

## CSV Format
File: `questions.csv`
Each line: `question,answer`
Example:
```
5+5,10
7+3,10
1+1,2
```

(If you already have `operations.csv`, rename or copy it to `questions.csv`.)

## Run
```
go run main.go
```
Or build:
```
go build -o quiz.exe
./quiz.exe   (Windows: quiz.exe)
```

## How It Works
- Loads all rows with `encoding/csv`.
- Random index selection without repeats (map of used indices).
- Reads user input line-by-line via buffered reader.
- Case-insensitive answer check.

## Customize
- Add more rows to `questions.csv`.
- Non-math questions work (e.g. `Capital of France?,Paris`).
- Shuffle strategy could be replaced by pre-shuffling slice.

## Possible Improvements (v1+)
- Timer per question.
- Track and show percentage.
- Accept multiple correct answers (split by `|`).
- Write results to a log file.
