# Wits - The 🥦 Information Tracking System

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/582a945a5bf24ec79fc6b3894b24544d)](https://app.codacy.com/gh/TheDonDope/wits-tui/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Codecov Badge](https://codecov.io/gh/TheDonDope/wits-tui/graph/badge.svg?token=9sWIVhEeIX)](https://codecov.io/gh/TheDonDope/wits-tui)

Wits aims to help cannabis patients and users to manage and monitor their cannabis consumption and inventory.

![Wits Demo Video](./vhs-output/wits-demo.gif)

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

```sh
make
```

![Wits Make Video](./vhs-output/wits-make.gif)

## Building the Binary for Windows

For windows, the `wits.exe` can be built by invoking the `make build-windows` command:

```sh
make build-windows
```

![Wits Make Windows Video](./vhs-output/wits-make-windows.gif)

## Running Tests

- Run the testsuite with coverage enabled:

```sh
make test
```

![Wits Make Test Video](./vhs-output/wits-make-test.gif)

- Generate the coverage results as html:

```sh
make cover
```

![Wits Make Cover Video](./vhs-output/wits-make-cover.gif)

- Open the results in the browser:

```sh
make show-cover
```

![Wits Make Show Cover Video](./vhs-output/wits-make-show-cover.gif)

Both the `coverage.out` as well as the `coverage.html` are explicitly ignored from source control (see [.gitignore](.gitignore)).
