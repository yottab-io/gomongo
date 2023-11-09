# GoLang Simplified MongoDB Library

This library provides a simplified interface for common operations in MongoDB, such as retrieving data, deleting items, and searching within collections. It aims to streamline and simplify the usage of these core functions, reducing the complexity and boilerplate code typically associated with such operations.

## Installation

To include this library in your project, you can simply use Go modules:


```bash
go get -u github.com/yottab-io/gomongo
```

## Features

- **Simplified API:** The library offers a simplified API for performing common data operations, making it easier and more intuitive to work with.
- **Connection Management:** The library provides two key functions for connection management:
 + Connect: This function establishes a connection with the data source and returns a handle to the connection.
 + CloseConnection: This function safely closes the connection when exiting the program, ensuring that resources are released properly.

## Usage

To use the library, simply import it into your Go project and leverage the simplified functions for efficient data manipulation. Below is an example of how to get started:

```go
import "github.com/yottab-io/gomongo"

// Example usage
result, err := gomongo.Get("exampleKey")
if err != nil {
    // Handle error
}
// Process result