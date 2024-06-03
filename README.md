
# Simple Statistical Calculator

A simple command-line tool written in Go that calculates the average, median, variance, and standard deviation of a set of numbers from a file.


> ## Features


> - Reads data from a file specified as an argument.
> - Calculates the average of the population data.
> - Calculates the median of the population data.
> - Calculates the variance of the population data.
> - Calculates the standard deviation of the population data.
> - Prints the results of each statistic in a formatted manner.

## Run Locally

Clone the project

```bash
  git clone https://learn.zone01kisumu.ke/git/ferdyodhiambo/math-skills
```

Go to the project directory

```bash
  cd math-skills
```

Run the project

```bash
  go run main.go data.txt
```
 **Note**: You need to have data.txt file in math-skills directory before running the command: go run main.go data.txt

## Usage/Examples
Data.txt content:

124

234

111

664

```bash
go run main.go data.txt
```
**Output**

Average:283

Median:179

Variance:50607

Standard Deviation:225

## Running Tests

To run tests, run the following command

```bash
  cd statmath
```
```bash
go test -v
```

## Contributing

### Contributing to the Statistical Calculator
Thank you for your interest in contributing to the Statistical Calculator project.

 Your contributions are valuable and help make the project better for everyone.
Getting Started

**Fork the Repository**: Fork the Statistical Calculator repository on GitHub by clicking the "Fork" button at the top right corner of the repository page.


**Clone the Repository**: Clone the forked repository to your local machine using the command:

   ``` bash
    git clone https://github.com/your-username/math-skills.git
   ```
**Create a New Branch**: Create a new branch for your feature or bug fix using the command:

```bash
git branch <branch-name>
```
**Make Changes**: Make the necessary changes to the code and commit them using the command:

```bash
git add .
git commit -m "<commit-message>"
```
### Code Style

***Formatting***: Use the Go standard formatting rules for code formatting.

***Naming Conventions***: Follow the Go standard naming conventions for variables, functions, and packages.

### Bug Reports

***Report Issues***: Report any issues you encounter while using the Statistical Calculator on the GitHub issues page.

***Provide Details***: Provide detailed information about the issue, including steps to reproduce it and any relevant code snippets.

### Feature Requests

***Submit a Pull Request***: Submit a pull request to add a new feature to the Statistical Calculator.

***Describe the Feature***: Describe the feature you want to add and how it will improve the project.

### Code Reviews

> - ***Code Review***: Review the code changes made by others to ensure they meet the project's coding standards.
> - ***Provide Feedback***: Provide constructive feedback on the code changes, including suggestions for improvement.

### Communication

***GitHub Issues***: Use the GitHub issues page for discussing issues and feature requests.

## Lessons Learned

What did you learn while building this project? What challenges did you face and how did you overcome them?


**Building the Statistical Calculator project taught me several valuable lessons**:

***Code Organization***: I learned the importance of organizing code into logical packages and files. This helps maintain a clean and structured codebase, making it easier to understand and modify.

***Error Handling***: I realized the importance of proper error handling in Go. This includes using err variables to handle errors and providing informative error messages.

***Testing***: I learned the importance of writing unit tests for the code. This ensures that the code works correctly and helps catch bugs early in the development process.

***Documentation***: I learned the importance of documenting code and projects. This includes writing clear and concise comments and creating a README.md file that provides instructions on how to use the project.

#### Challenges and Solutions

> - ****Calculating Median****: Calculating the median of a set of numbers was a challenge. I overcame this by implementing a simple sorting algorithm to find the middle value.

> - ****Handling Large Data Sets****: Handling large data sets was another challenge. I overcame this by implementing a function to read the data from a file and process it in chunks.

> - ****Error Handling: Handling errors was a challenge****. I overcame this by using err variables to handle errors and providing informative error messages.

> - ****Code Readability: Maintaining code readability was a challenge****. I overcame this by following Go standard formatting rules and using clear and concise variable names.

> - ****Testing****: Writing unit tests was a challenge. I overcame this by using Go's built-in testing library and writing tests for each function.

## Study Reference

[Introduction to Statistics](https://www.khanacademy.org/math/statistics-probability)

[Golang](https://go.dev/)
