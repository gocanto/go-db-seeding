### Go DB migrations with easy

Simple Go project that allows for Postgres DB migrations and fake data.

### How do I run this code?

- Check the Makefile
- Call the `db.Migrate()` method to execute migrations against your current schema.
- Run `go run main.go` to fake and populate the users table.

### Aid commands

- Create DB: `make run-db`
- Inspect DB: `make db-bash`
- Create migrations: `make create-migrations`


### DB helper commands
> You need to be in your postgres instance bash command line to 
be able to execute these commands.

- Initiate the DB session: `make db-bash`

- Login: `psql -U postgres`
- Show dbs: `SELECT datname FROM pg_database;`
- Use db: `\c {db-name}`
- Show tables: `\dt`

## Contributing

Please, feel free to fork this repository to contribute to it by submitting a functionalities/bugs-fixes pull request to enhance it.

## License

Please see [License File](https://github.com/gocanto/go-db-seeding/blob/main/LICENSE) for more information.

## How can I thank you?

There are many ways you would be able to support my open source work. There is not a right one to choose, so the choice is yours.

Nevertheless :grinning:, I would propose the following

- :arrow_up: Follow me on [Twitter](https://twitter.com/gocanto).
- :star: Star the repository.
- :handshake: Open a pull request to fix/improve the codebase.
- :writing_hand: Open a pull request to improve the documentation.
- :coffee: Buy me a [coffee](https://github.com/sponsors/gocanto)?

> Thank you for reading this far. :blush:
