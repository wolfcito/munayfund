FROM golang:1.21-alpine3.19 as builder

# Configuramos las variables de entorno y las dependencias
ENV GO111MODULE=on
RUN apk update && apk --no-cache add ca-certificates bash git gcc g++ pkgconfig build-base zlib-dev pkgconf openssh

# Creamos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos de configuración del módulo Go
COPY go.mod .
COPY go.sum .

# Descargamos las dependencias
RUN go mod download

# Copiamos todo el código fuente
COPY . .

# Construimos la aplicación Go
RUN GOOS=linux go build -a -tags musl -installsuffix cgo -o main ./cmd

# Creamos una imagen mínima para ejecutar la aplicación
FROM alpine:3.15
RUN apk --no-cache add ca-certificates

# Copiamos el ejecutable construido desde la etapa de construcción
COPY --from=builder /app/main .

# Ejecutamos la aplicación al iniciar el contenedor
CMD ["./main"]
