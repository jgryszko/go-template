# Template Go Project

This is a simple Go project that greets the user with a name provided via the `--name` parameter.

## How to Run

1. Build the project:

    ```bash
    go build -o myapp ./cmd/myapp
    ```

2. Run the app:

    ```bash
    ./myapp --name Joe
    ```

   You can also run without `--name` to use the default value from the config:

    ```bash
    ./myapp
    ```

    Output:

    ```
    Hello, Joe
    ```

    or with the default:

    ```
    Hello, World
    ```

## Running Unit Tests

To run the tests, use:

```bash
go test ./...

# My Go Project

This is a simple Go project that greets the user with a name provided via the `--name` parameter.

## How to Build the Project

You can generate an executable named `greeter` by running the following command:

```bash
go build -o greeter ./cmd/myapp

## Conclsions

Just use as a starting point
