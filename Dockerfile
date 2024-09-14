# Usa una imagen base de Go
FROM golang:1.23

# Establece el directorio de trabajo en /app
WORKDIR /app

ENV HOST 0.0.0.0
ENV PORT 8080
ENV DB_USER auth_bd_brto_user
ENV DB_PASSWORD 2zDlYJVxcvrWXXleEzNGCR825VHCckrC
ENV DB_NAME auth_bd_brto
ENV DB_HOST dpg-cripfejv2p9s738m9m5g-a.oregon-postgres.render.com
ENV DB_PORT 5432

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