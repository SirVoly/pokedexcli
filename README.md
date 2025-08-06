# Pokedex CLI

A Pokedex in a command-line REPL, using [PokeAPI](https://pokeapi.co).

This project was build with the intention of hands-on practice in JSON, GO, HTTP Requests, CLI tools and caching.

## Technologies Used

*   **GO**: The primary programming language.

## Setup

1. **Clone the repository**
    ```
    git clone https://github.com/SirVoly/pokedexcli.git
    cd pokedexcli
    ```

2. **Build the project**
    ```
    go build
    ```

   This will create an executable file in the current directory (likely called `pokedexcli`).

3. **(Optional) Run Tests**
    ```
    go test ./...
    ```

## Usage

After building, you can run the program from the terminal:
    ```
    ./pokedexcli
    ```

Or, if running directly with Go (for development):
    ```
    go run main.go
    ```

Once running, you’ll be greeted with the REPL prompt, where you can enter commands such as:
- `help` — List available commands
- `explore` — Explore new areas
- `catch <pokemon>` — Attempt to catch a Pokémon
- `pokedex` — View caught Pokémon

# Credit
This project was completed as part of a guided course on [Boot.dev](https://www.boot.dev).
It was build following along with the [Build a Pokedex in Go](https://www.boot.dev/courses/build-pokedex-cli-golang) course.