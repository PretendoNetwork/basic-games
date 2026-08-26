# Basic Games

Game servers that use the common NEX protocol implementations, with no game-specific RMC changes or that only need very minor tweaks. Games with larger changes (replaced RMC methods, custom protocols, etc.) belong in their own repositories.

## Adding a game

Add a file to `games/` and register it in the registry.

## Configuration

Some server configuration data (ports, addresses, etc.) are set during runtime and pulled from environment variables.

| Name                                     | Description                                                                            | Required          |
|------------------------------------------|----------------------------------------------------------------------------------------|-------------------|
| `PN_GAME`                                | Slug of the game to serve                                                              | Yes               |
| `<EnvPrefix>_AUTHENTICATION_SERVER_PORT` | Port for the authentication server                                                     | Yes               |
| `<EnvPrefix>_SECURE_SERVER_HOST`         | Host name for the secure server, usually the same address as the authentication server | Yes               |
| `<EnvPrefix>_SECURE_SERVER_PORT`         | Port for the secure server                                                             | Yes               |
| `<EnvPrefix>_ACCOUNT_GRPC_HOST`          | Host name for the account server gRPC service                                          | Yes               |
| `<EnvPrefix>_ACCOUNT_GRPC_PORT`          | Port for the account server gRPC service                                               | Yes               |
| `<EnvPrefix>_ACCOUNT_GRPC_API_KEY`       | API key for the account server gRPC service                                            | No (assumed open) |
| `<EnvPrefix>_FRIENDS_GRPC_HOST`          | Host name for the friends server gRPC service                                          | Yes               |
| `<EnvPrefix>_FRIENDS_GRPC_PORT`          | Port for the friends server gRPC service                                               | Yes               |
| `<EnvPrefix>_FRIENDS_GRPC_API_KEY`       | API key for the friends server gRPC service                                            | No (assumed open) |
| `<EnvPrefix>_POSTGRES_URI`               | Fully qualified Postgres URI                                                           | Yes               |
| `<EnvPrefix>_HEALTH_CHECK_PORT`          | Port for the UDP health check                                                          | No                |

## Running

Set the `PN_GAME` environment variable to the slug of the game the server should configure itself for, then run the server binary.

```bash
make build
PN_GAME=tetris-axis ./build/server
```
