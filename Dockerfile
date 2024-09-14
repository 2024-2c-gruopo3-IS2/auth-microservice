# Usa una imagen base de Go
FROM golang:1.23

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el contenido de la carpeta src a /app
COPY src/ .

# Descarga las dependencias
RUN go mod download
RUN go mod tidy

# Compila la aplicación
RUN go build -o main .

# Expone el puerto que la aplicación usa
EXPOSE 8080

# Ejecuta la aplicación
CMD ["./main"]