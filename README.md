# Go Todo API with Postgres

This does not for serious purpose. I created this project because I am burning out and having other issues.
I just need to cope.

## Where tests, where security?

I don't add them because I just want to finish the API as soon as possible.

Note that I haven't even add a service module. I just use repository directly as it is so small.

## Things that I can possibly improve if I consider it to be a long-term project

- Add tests
- Add more endpoints (e.g. endpoint for completion only)
- Add timestamp
- Refactor repetitive code

## Setup

1. Setup `.env` variable for Docker Compose
2. Run `docker compose up -d`
3. Use [golang-migrate](https://github.com/golang-migrate/migrate) to migrate the database.
4. To run the server, use `go run .`
5. To stop the server, simply `Ctrl-C` the `go run .` terminal and `docker compose down`

I do not have intention to build or deploy this stuff. Please find a way yourself.
