# MunayFund

MunayFund es una plataforma de pagos basada en tecnologías blockchain.

## Descripción

MunayFund permite a los usuarios realizar pagos de forma segura y eficiente utilizando tecnología de cadena de bloques. La plataforma utiliza contratos inteligentes para gestionar las transacciones y garantizar la integridad y seguridad de los pagos.

## Características

- Pagos seguros mediante contratos inteligentes.
- Integración con IPFS para almacenamiento descentralizado.
- Interfaz de usuario amigable.
- ...

## Requisitos del Sistema

- Go 1.21
- Docker
- Docker Compose
- IPFS

## Configuración

### Variables de Entorno

Asegúrate de configurar las siguientes variables de entorno antes de ejecutar la aplicación:

- `SECRETKEY`: Clave secreta para la aplicación.
- `PORT`: Puerto en el que se ejecutará la aplicación.
- `MONGODB_CONNECTION_URI`: URI de conexión a la base de datos MongoDB.
- `IPFSURL`: URL de conexión al servicio IPFS.

Ejemplo:

```bash
export SECRETKEY=mysecretkey
export PORT=9090
export MONGODB_CONNECTION_URI=mongodb://user:password@localhost:27017/mydb
export IPFSURL=http://myipfsurl
export ROOT_USERNAME=test
export ROOT_PWD=testtest
export SECRETKEY=testkey
export PORT=8080
export MUNAY_USERNAME=test
export MUNAY_PASSWORD=testtest
```

## Cambios en swagger API
Para construir los cambios de la API podemos ejecutar este comando:
```bash
swag init --parseDependency -g cmd/main.go
```
Con esto actualizaremos la documentacion

## Docker Compose
Puedes ejecutar la aplicación utilizando Docker Compose. Asegúrate de tener Docker y Docker Compose instalados y luego ejecuta:

```bash
docker-compose up
```

## Contribuir
¡Contribuciones son bienvenidas! Si deseas contribuir al proyecto, sigue estos pasos:

1. Haz un fork del repositorio.

2. Crea una nueva rama para tu contribución:
```bash 
git checkout -b feature/nueva-funcionalidad.
```

3. Realiza tus cambios y haz commit:
```bash 
git commit -m "Agrega nueva funcionalidad".
```

4. Haz push a tu rama: 
```bash
git push origin feature/nueva-funcionalidad.
```

5. Abre un Pull Request.

## Licencia
Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.