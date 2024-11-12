# Usa la imagen oficial de Golang 1.23.3 como base
FROM golang:1.23.3 AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo Go al contenedor
COPY main.go .

# Compila la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp main.go

# Usa una imagen más ligera para ejecutar la aplicación
FROM alpine:latest

# Establece el directorio de trabajo
WORKDIR /root/

# Copia el binario de la etapa de construcción
COPY --from=builder /app/myapp .

# Expone el puerto en el que la aplicación escuchará
EXPOSE 8084

# Comando para ejecutar la aplicación
CMD ["./myapp"]