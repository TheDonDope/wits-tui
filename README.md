# Wits - The 🥦 Information Tracking System

Wits aims to help cannabis patients and users to manage and monitor their cannabis consumption and inventory.

## Changelog & Roadmap

A detailed changelog can be found in the [CHANGELOG.md](./CHANGELOG.md) and the current development progress is tracked in the [ROADMAP.md](./ROADMAP.md). We do not use GitHub Issues but instead track our features, bugfixes and refactorings there.

## Configuring the Application & Required Environment Variables

Wits can be configured through environment variables, detailed here:

| Environment Variable | Description                                                                 |
| -------------------- | --------------------------------------------------------------------------- |
| `LOG_LEVEL`          | The level at which to log (one of: `DEBUG`, `INFO`, `WARN`, `ERROR`, `OFF`) |
| `LOG_DIR`            | The path to the directory for the application logs                          |
| `LOG_FILE`           | The name of the file for the application logs (within `LOG_DIR`)            |
| `WITS_DIR`           | The directory where the application stores its data (defaults to `.wits`)   |
| `STORAGE_MODE`       | The persistance type to use (either `in-memory` or `yml-file`)              |

A minimum viable `.env` file can be found at [.env.example](.env.example). Simply rename it to `.env` to be able to run the application with a yaml file based storage.

## Building & Running the Application

Building the binary and running it requires only a simple invocation to `make`:

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
