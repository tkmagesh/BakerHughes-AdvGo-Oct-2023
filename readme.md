# Magesh Kuppan

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Docker

## Schedule:
    Commence    : 9:30 AM
    Tea Break   : 11:00 AM (15 mins)
    Lunch Break : 1:00 PM (1 hour)
    Tea Break   : 3:30 PM (15 mins)
    Wind up     : 5:30 PM

## Methodology
- No powerpoint
- Discussion & Code

## Repository
- https://github.com/tkmagesh/bakerhughes-advgo-oct-2023 (https://bit.ly/bh-ago)

## Go
- Functional Programming
- Interfaces
- Struct Composition
- Concurrency

## Functional Programming
- Assign a function as a value to a variable
- Higher Order Functions
    - Pass functions as arguments to other functions
    - Return functions as return values

## Concurrency in Go
### Concurrency
- The ability to have more than one execution path in the application
### Race Detection
    - go run --race <filename.go>
    - go build --race <filename.go> 
        - DO NOT use a build with race detector in production
### Channel
- data type to enable communication between goroutines
- declare
```
var <var_name> chan <data_type>
ex : var ch chan int
```
- initialize
```
<var_name> = make(chan <data_type>)
ex: ch = make(chan int)
```
- declare & initalize
```
ch := make(chan int)
```