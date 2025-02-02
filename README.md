# Akatech

## Membuat REST API
### Running

- Up Compose (optional)
```bash
docker compose -f postgresql.yaml up -d
```
- Change .env
```bash
cp .env.example .env
```
- Run Go Application
```bash
go run main.go
```