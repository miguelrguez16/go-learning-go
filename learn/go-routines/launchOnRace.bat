@echo on
set CGO_ENABLED=1
go run --race main.go
echo FIN