# Use a imagem oficial do Go
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copia módulos e go.mod
COPY go.mod go.sum ./

# Baixa dependências
RUN go mod download

# Copia todo o projeto
COPY . .

# Expõe a porta da aplicação    
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["go", "run", "cmd/main.go"]
