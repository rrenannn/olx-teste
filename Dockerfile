# Use a imagem oficial do Go como base
FROM golang:alpine

# Defina o diretório de trabalho
WORKDIR /app

# Copie o arquivo go.mod e go.sum
COPY go.mod go.sum ./

# Execute o comando go mod download para baixar as dependências
RUN go mod download

# Copie o restante do código
COPY . .

# Execute o comando go build para compilar o código
RUN go build -o main .

# Exponha a porta 8080
EXPOSE 8080

# Defina o comando para executar o aplicativo
CMD ["./main"]