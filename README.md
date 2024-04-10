# Abstract Factory Pattern Demonstration in Go

## Overview
This repository is a demonstration of the Abstract Factory design pattern implemented in Go. The project illustrates how to manage families of related products without specifying their concrete classes. The primary focus is on toy products, divided into two categories: Classic and Modern.

## Pattern Description
The Abstract Factory Pattern provides an interface for creating families of related or dependent objects without specifying their concrete classes. This project includes two factories: `ClassicToyFactory` and `ModernToyFactory`, each producing two types of products: `Doll` and `Car`. Each factory and product combination reflects a unique style and behavior, showcasing the pattern's ability to facilitate scalability and testability in software design.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the usage of different factories.
- **pkg/**
    - **factory/**: Includes the abstract factory interface (`factory.go`) and concrete factory implementations (`classic_toy_factory.go` and `modern_toy_factory.go`).
    - **product/**: Contains interfaces and implementations for products (`dolls.go` and `cars.go`).

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Abstract-Factory_Pattern.git
cd Go_Abstract-Factory_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp