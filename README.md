# Go Basic 1

This is a basic Go project that demonstrates how to structure code and use modules.

## Project Structure

The project consists of two main files:

- `main.go`: This is the entry point of the application. It imports and uses the `myModule` package.
- `myModule/user.go`: This file contains the `PrintUserDetails` function which prints user details.
- `myModule/arithmetic.go`: This file contains the `PrintArithmetic` function which performs basic arithmetic operations and prints the results.
- `myModule/logic.go`: This file is used for practicing logic statements. It contains the `CheckUserAge`, `CheckStudentAge`, and `CheckWeekend` functions which check the age of a user, whether it's a weekend or a holiday, and print appropriate messages.


## Usage

To run the project, navigate to the project's root directory and use the `go run` command:

```bash
go run main.go