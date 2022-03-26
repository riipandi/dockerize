# Dockerize Application Stack

This repo contains the boilerplate of many web applications running within Docker.

## Quick Start

### Start

```sh
# Start the stack
docker-compose up -d

# Stop the stack
docker-compose down

## Watching the logs
docker-compose logs -f

# Updating the stack
docker-compose down --remove-orphans
docker-compose pull
docker-compose up -d
```

### Generate Secret Key

```sh
openssl rand -base64 64 | tr -d '\n' ; echo
```

## License

This project is open-sourced software licensed under the [MIT license](https://aris.mit-license.org).

Copyrights in this project are retained by their contributors.
See the [license file](./license.txt) for more information.
