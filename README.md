# Wits - The 🥦 Information Tracking System

Wits aims to help cannabis patients and users to manage and monitor their cannabis consumption and inventory.

## Building the Application

Building the binary requires only a single step:

```shell
$ make build
go build -v -o ./bin/wits ./cmd/wits/main.go
```

## Running the Application

After building, simply invoke:

```shell
🥦 Welcome to Wits!

> 🌿 Strains
 🚀 Devices
 🔧 Settings
 📊 Statistics

Press ctrl+c or q to quit.
```

Or, do it all in one step by invoking:

```shell
$ make
go build -v -o ./bin/wits ./cmd/wits/main.go

🥦 Welcome to Wits!

> 🌿 Strains
 🚀 Devices
 🔧 Settings
 📊 Statistics

Press ctrl+c or q to quit.
```

## Running Tests

- Run the testsuite with coverage enabled:

```shell
$ make test
go test -race -v ./... -coverprofile coverage.out
        github.com/TheDonDope/wits-tui/cmd/tui          coverage: 0.0% of statements
?       github.com/TheDonDope/wits-tui/pkg/cannabis     [no test files]
        github.com/TheDonDope/wits-tui/pkg/service              coverage: 0.0% of statements
        github.com/TheDonDope/wits-tui/pkg/storage              coverage: 0.0% of statements
        github.com/TheDonDope/wits-tui/pkg/tui          coverage: 0.0% of statements
```

- Generate the coverage results as html:

```shell
$ make cover
go test -race -v ./... -coverprofile coverage.out
        github.com/TheDonDope/wits-tui/cmd/tui          coverage: 0.0% of statements
?       github.com/TheDonDope/wits-tui/pkg/cannabis     [no test files]
        github.com/TheDonDope/wits-tui/pkg/service              coverage: 0.0% of statements
        github.com/TheDonDope/wits-tui/pkg/storage              coverage: 0.0% of statements
        github.com/TheDonDope/wits-tui/pkg/tui          coverage: 0.0% of statements
go tool cover -html coverage.out -o coverage.html
```

- Open the results in the browser:

```shell
$ make show-cover
go test -race -v ./... -coverprofile coverage.out
        github.com/TheDonDope/wits-tui/cmd/tui          coverage: 0.0% of statements
?       github.com/TheDonDope/wits-tui/pkg/cannabis     [no test files]
        github.com/TheDonDope/wits-tui/pkg/service              coverage: 0.0% of statements
        github.com/TheDonDope/wits-tui/pkg/storage              coverage: 0.0% of statements
        github.com/TheDonDope/wits-tui/pkg/tui          coverage: 0.0% of statements
go tool cover -html coverage.out -o coverage.html
open coverage.html
<Opens Browser>
```

Both the `coverage.out` as well as the `coverage.html` are explicitly ignored from source control (see [.gitignore](.gitignore)).
