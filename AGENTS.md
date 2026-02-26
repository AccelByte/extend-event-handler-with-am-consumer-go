# extend-event-handler-with-am-consumer-go

An Extend Event Handler app with **async messaging consumer** written in Go. Consumes messages from AccelByte's async messaging system (not Kafka Connect) and executes custom logic in response.

This is a template project — clone it, replace the sample logic in the service implementation, and deploy.

## Build & Test

```bash
make build                           # Build the project
go test ./...                        # Run unit tests
docker compose up --build            # Run locally with Docker
make proto                           # Regenerate proto code
```

Linting: `golangci-lint run` (config in `.golangci.yml`).

## Architecture

This app consumes messages from AccelByte's async messaging system:

```
AGS Service → Async Messaging → [gRPC] → This App
```

Unlike the standard Event Handler repos which use Kafka Connect, this variant consumes messages directly. **Idempotency is still important** — ensure handlers are safe to run more than once for the same message.

The sample implementation listens for player-joined messages, extracts the player ID from message metadata, and optionally persists the event to CloudSave as a game record.

### Key Files

| Path | Purpose |
|---|---|
| `main.go` | Entry point — starts gRPC server, wires interceptors and observability |
| `pkg/service/asyncMessagingHandler.go` | **Service implementation** — your custom logic goes here |
| `pkg/proto/async_messaging/consumer_service.proto` | Event message definitions (AccelByte-provided, do not modify) |
| `pkg/pb/` | Generated code from proto (do not hand-edit) |
| `pkg/common/` | Auth interceptor, tracing, logging utilities |
| `docker-compose.yaml` | Local development setup |
| `.env.template` | Environment variable template |

### Consumer Notes

Unlike the standard Event Handler repos which use Kafka Connect as a bridge, this variant consumes messages directly from AccelByte's async messaging system. Messages arrive with a topic, body (JSON), and metadata map. The `STORE_MESSAGE_IN_CLOUDSAVE` environment variable controls whether events are persisted to CloudSave. This is a Go-only implementation.

## Rules

See `.agents/rules/` for coding conventions, commit standards, and proto file policies.

## Environment

Copy `.env.template` to `.env` and fill in your credentials.

| Variable | Description |
|---|---|
| `AB_BASE_URL` | AccelByte base URL (e.g. `https://test.accelbyte.io`) |
| `AB_CLIENT_ID` | OAuth client ID |
| `AB_CLIENT_SECRET` | OAuth client secret |
| `AB_NAMESPACE` | Target namespace |
| `STORE_MESSAGE_IN_CLOUDSAVE` | Store event messages in CloudSave |

## Dependencies

- [AccelByte Go SDK](https://github.com/AccelByte/accelbyte-go-sdk) — AGS platform SDK and gRPC plugin utilities
