# Task management service

1. Copy the contents of `.env-example` to `.env`
2. Run `make start` to start the project locally
3. Visit the documentation page: http://localhost:9000/swagger/index.html

## Local development

You can run local development environment in Docker. Make sure you have `.env` and run:

```shell
docker compose up -d --build
```

The backend will start on http://localhost:9000.
