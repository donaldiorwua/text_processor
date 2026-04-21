# Basic Calculator

Simple command-line calculator written in Go.

## Overview
This small program provides a basic interactive calculator that supports add, subtract, multiply and divide operations. It demonstrates simple user input handling and function organization in Go.

## Features
- Interactive CLI menu
- Add, subtract, multiply (integers)
- Divide (floating point)
- Exit command to terminate

## Prerequisites
- Go 1.18+ installed

## Build & Run
From the project root:

- Run without building:
  go run .

- Build and run binary:
  go build -o text_processor
  ./text_processor

(You can also run `go run main.go`.)

## Usage
When started the program shows a simple menu. Example session:

- Enter operation: `add`, `subtract`, `multiply`, or `divide`
- Enter the two numbers when prompted
- Enter `exit` to terminate

Example:
> Please, choose an operation:  
> add  
> Enter Two numbers to add!  
> Enter a number:  
> 5  
> Enter another number:  
> 3  
> Result: 8

## Project layout
- main.go — program entry point
- functions — package with calculator logic
  - calc.go — interactive calculator functions
  - `textprocessor.go` — (other text processing utilities, if present)
- go.mod — module definition

## Notes
- Integer operations use `int`. Division uses `float64`.
- The package in the functions folder is `TEXT_PROCESSOR`.

## Contributing
Small fixes and improvements welcome. Open an issue or submit a PR.
