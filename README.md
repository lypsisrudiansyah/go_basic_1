# Go Basic 1

This is a basic Go project that demonstrates how to structure code and use modules.

## Project Structure

The project consists of two main files:

- `main.go`: This is the entry point of the application. It imports and uses the `myModule` package.
- `myModule/user.go`: This file contains the `PrintUserDetails` function which prints user details.
- `myModule/arithmetic.go`: This file contains the `PrintArithmetic` function which performs basic arithmetic operations and prints the results.
- `myModule/logic.go`: This file is used for practicing logic statements. It contains the `CheckUserAge`, `CheckStudentAge`, and `CheckWeekend` functions which check the age of a user, whether it's a weekend or a holiday, and print appropriate messages.
- `myModule/struct.go`: This file demonstrates the use of structs in Go. It defines a `Student` struct and includes the `ShowStudentStruct` and `ShowStudentStructWithNamedArgument` functions which create instances of the `Student` struct and print their details.
- `myModule/array.go`: This file demonstrates how to declare and manipulate a slice (dynamic array) in Go.
- `myModule/function.go`: This file demonstrates the use of functions in Go. It includes the `addition` function which performs addition of two integers, and two other functions `DemoTheFunction` and `InputAdditionForOutsideFile` which demonstrate the usage of the `addition` function.
- `myModule/looping.go`: This file demonstrates the use of loops in Go. It includes the `DoTheLooping` function which performs simple looping, looping with an array, and a loop that behaves like a while loop.

## Running the Project

1. **Standard Run**: To run the project, navigate to the project's root directory and use the `go run` command:

```bash
go run main.go
```

This will print the user details, the results of the arithmetic operations, and the user age checks to the console.

2. **Hot Reload with Air**: If you want to use hot reloading, which automatically restarts your application whenever you save a file, you can use Air. Here's how you can set it up:

- Install Air by running the following command in your terminal:

```bash
go get -u github.com/cosmtrek/air
```

- Once Air is installed, you can start your application with hot reload by running:

```bash
air
```

Now, your application will automatically restart whenever you save a file.