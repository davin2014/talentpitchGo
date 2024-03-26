```markdown
# TalentPitch Go

Este proyecto es una aplicación web construida con Go. Utiliza el paquete `gorilla/mux` para el enrutamiento HTTP y `joho/godotenv` para cargar variables de entorno desde un archivo `.env`.


## Docker

Este proyecto utiliza Docker para contenerizar la base de datos. Puedes construir y ejecutar el contenedor de la base de datos utilizando los siguientes comandos:

Primero, navega al directorio de la base de datos:

```bash
cd database
```

```bash
docker build . -t talen-db
```

```bash
docker run -p 54321:5432 talen-db
```

## Comenzando

Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas.

### Prerrequisitos

Necesitas tener Go instalado en tu máquina. Para verificar si Go está instalado, puedes usar el siguiente comando:

```bash
go version
```

### Instalación

Para instalar el proyecto, primero debes clonarlo. Luego, navega hasta el directorio del proyecto y descarga las dependencias:

```bash
go get -u github.com/usuario/talentpitchGo
```

## Ejecutando la aplicación

Para ejecutar la aplicación, usa el siguiente comando en el directorio del proyecto:

```bash
go run main.go
```

## Pruebas Unitarias

Este proyecto utiliza el marco de pruebas incorporado en Go para realizar pruebas unitarias. Puedes ejecutar las pruebas unitarias utilizando el comando `go test`.

Por ejemplo, para ejecutar las pruebas unitarias para el manejador de usuarios, puedes usar el siguiente comando:

```bash
go test ./handlers/tests/
go test ./handlers/tests/user_test.go

## Rutas de la aplicación

La aplicación tiene las siguientes rutas:

- `/`: Ruta de inicio.
- `/signup`: Ruta para registrarse.
- `/login`: Ruta para iniciar sesión.
- `/deleteUser/{id}`: Ruta para eliminar un usuario.
- `/updateUser/{id}`: Ruta para actualizar un usuario.
- `/users`: Ruta para obtener todos los usuarios.
- `/me`: Ruta para obtener la información del usuario actual.
- `/challenges`: Ruta para crear y listar desafíos.
- `/updateChallenge/{id}`: Ruta para actualizar un desafío.
- `/deleteChallenge/{id}`: Ruta para eliminar un desafío.
- `/challenges/{id}`: Ruta para obtener un desafío específico.

## Contribuyendo

Por favor lee [CONTRIBUTING.md](https://gist.github.com/usuario/talentpitchGo/contributing.md) para detalles de nuestro código de conducta, y el proceso para enviarnos pull requests.


## DevOps

Este proyecto utiliza AWS CodeBuild para la integración continua y la entrega continua (CI/CD). El archivo `buildspec.yml` define las fases y comandos que CodeBuild utilizará para construir y empaquetar la aplicación.

### Fases

El archivo `buildspec.yml` define dos fases: `install` y `build`.

#### Install

En la fase `install`, se especifica la versión de Go que se utilizará para construir la aplicación. En este caso, se utiliza Go 1.15.

#### Build

En la fase `build`, se definen los comandos que se ejecutarán para construir la aplicación. En este caso, se imprime un mensaje en la consola y luego se utiliza el comando `go build -o application` para construir la aplicación.

### Artifacts

Los artefactos son los archivos que se generarán después de que se complete la fase de construcción. En este caso, el archivo `application` que se genera al construir la aplicación se especifica como un artefacto.

### Version

La versión del archivo `buildspec.yml` es 0.2. Esta es la versión del formato del archivo `buildspec.yml` que se está utilizando, no la versión de la aplicación.

## Licencia

Este proyecto está bajo la Licencia (Tu licencia) - mira el archivo [LICENSE.md](LICENSE.md) para detalles
```

Por favor, reemplaza los valores de marcador de posición con la información real de tu proyecto.