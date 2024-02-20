# Usar la imagen oficial de Go como base
FROM golang:1.20

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el go.mod y go.sum al directorio de trabajo actual
COPY go.mod go.sum ./

# Descargar todas las dependencias
RUN go mod download

# Copiar el código fuente del host al contenedor
COPY . .

# Compilar la aplicación
RUN go build -o SimonBK_DecoderCo .

# Comando para ejecutar la aplicación
CMD ["./SimonBK_DecoderCo"]
