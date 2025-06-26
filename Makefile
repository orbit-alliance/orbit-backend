.PHONY: all deps run

# Alvo principal
all: deps run

# Instala as dependências
deps:
	go mod tidy

# Executa o main.go
run:
	go run ./cmd
