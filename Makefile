api-run:
	@cd api && go run cmd/serverd/main.go

api-vendor:
	@cd api && go mod tidy && go mod vendor

app-run:
	@flutter run flutter_go_websocket
