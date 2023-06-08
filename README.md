# go_course_web Microservice: Lib Response
URL del curso: https://www.udemy.com/course-dashboard-redirect/?course_id=4809678

Este proyecto forma parte de la transformación del proyecto creado en el repositorio [go_course_web](https://github.com/jousga/go_course_web) pero transformándolo a microservicios

En concreto la intención de este proyecto es gestionar las response de los diferentes microservicios

El conjunto de microservicios que forman parte de este proyecto son:
1. Domain: https://github.com/jousga/go_course_web_microservices_domain
2. Meta: https://github.com/jousga/go_course_web_microservices_meta
3. User: https://github.com/jousga/go_course_web_microservices_user
4. Course: https://github.com/jousga/go_course_web_microservices_course
5. Enrollment: https://github.com/jousga/go_course_web_microservices_enrollment
6. Lib Response: https://github.com/jousga/go_course_web_microservices_lib_response

Para limpiar el proyecto y descargar las dependencias tenemos que usar el comando de GO:

```text
go mod tidy
```

** Recordar que para importar los proyectos Meta y Domain tenemos que usar "go get":

```text
go get github.com/jousga/go_course_web_microservices_meta
go get github.com/jousga/go_course_web_microservices_domain
```

## Servidor:
Como server se ha usado: [https://github.com/gorilla/mux](https://github.com/gorilla/mux)
Para añadirlo a un proyecto nuevo, dentro del directorio de la aplicación, tenemos que usar el comando:

```text
go get github.com/gorilla/mux
```

## Estructura del proyecto:
* internal: Contiene todo el código que no haga referencia a paquetes externos.
* cmd: carpeta donde movemos el fichero main.go.
* pkg: Contiene las partes de la web que se encargan de interaccionar con servicios externos.
    * bootstrap: utilidades comunes a toda la aplicación. Por ejemplo: Conexión a la base de datos o creación del log
    * handler: Contiene el middleware de User y le definimos los endpoints y como deben procesarse y transforma el JSON de las request en un objeto y la respuesta del endpoint en un JSON.
    * meta: para manejar el meta del request, información adicional que enviaremos en el body. Por Ejemplo: cantidad de registros, página actual, total de páginas,...
* .dockers: configuración de los contenedores de docker que utiliza la aplicación
* docker-compose.yml: fichero que se encarga de gestionar los contenedores docker de la aplicación


## Importar proyectos de repositorios privados
Si queremos importar repositorios privados al ejecutar "go get ..." nos dará un error. Para ello primero tenemos que crear
en GitHub un access token en la página de [Github explican como crear un token](https://docs.github.com/es/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).

A continuación un resumen de los pasos a seguir para crear un token del tipo fine-grained:
1. Clicar en nuestra foto de usuario y seleccionar Settings
2. En el menú de la derecha seleccionar la opción "Developer settings"
3. Seleccionar la opción "Personal access tokens"
4. Escoger el token de tipo "Fine-grained tokens"
5. Clicar el botón "Generate new token"
6. Rellenar la información que se solicita:
    1. En Nombre del token, escribe un nombre para el token.
    2. En Expiración, selecciona cuándo expirará el token.
    3. Opcionalmente, en Descripción, agrega una nota para describir el propósito del token.
    4. En Propietario del recurso, selecciona un propietario del recurso. El token solo podrá acceder a los recursos que pertenecen al propietario del recurso seleccionado. Las organizaciones a las que pertenezcas no aparecerán a menos que hayan optado por el uso de un fine-grained personal access token. Para obtener más información, consulta "Establecimiento de una directiva de token de acceso personal en la organización".
    5. Opcionalmente, si el propietario del recurso es una organización que requiere aprobación para el uso de un fine-grained personal access token, escribe una justificación para la solicitud en el cuadro que aparece debajo del propietario del recurso.
    6. En Acceso al repositorio, selecciona los repositorios a los que quieres que acceda el token. Debes elegir el acceso mínimo al repositorio que satisfaga tus necesidades. Los tokens siempre incluyen acceso de solo lectura a todos los repositorios públicos de GitHub.
    7. Si elegiste Solo repositorios seleccionados en el paso anterior, en la lista desplegable Repositorios seleccionados, elige los repositorios a los que quieres que acceda el token.
    8. En Permisos, selecciona los permisos que se concederán al token. En función del propietario del recurso y del acceso al repositorio que hayas especificado, hay permisos de repositorio, de organización y de cuenta. Debes elegir los permisos mínimos que necesites. Para obtener más información sobre los permisos necesarios para cada operación de la API REST, consulta "Permisos necesarios para los tokens de acceso personal específicos".
    9. Haga clic en Generar token.

El siguiente paso es instalar GitHub CLI tal y como se explica en este [artículo](https://github.com/cli/cli#installation) y usando el comando:

```text
winget install --id GitHub.cli
```
Una vez instalado tenemos que el configurar el [gh auth login](https://cli.github.com/manual/gh_auth_login), el cual hará varias preguntas, si queremos usar HTTPS o SSH, escogemos
el primero y como sistema de autentificación seleccionaremos TOKEN, e introduciremos el token que hemos creado en GitHub.
El comando para poder empezar a configurar el login es:

```text
gh auth login
```

Una vez hecho esto, cuando hagamos el "go get ...", automáticamente utilizará HTTPS y el token introducido para descargar el recurso indicado.

### Otros útiles:
Temp keu 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

Administrador credenciales de Git: https://docs.github.com/es/get-started/getting-started-with-git/caching-your-github-credentials-in-git

Para eliminar credenciales wn windows de GIT: https://support.microsoft.com/es-es/windows/obtener-acceso-al-administrador-de-credenciales-1b5c916a-6a16-889f-8581-fc16e8165ac0

Modificar el git config para usar SSH: https://michaelheap.com/golang-how-to-go-get-private-repos/


## Librerías/Packages:

### Go Kit - Toolkit for microservices
[Go kit](https://gokit.io/) Es un paquete que nos permite, entre otras cosas, crear middlewares antes de la capa del endpoint que nos permiten que
estos reciban el JSON de una petición lo transforman en un objeto y el endpoint lo que recibe es el objeto en si y lo mismo
pero al revés, el endpoint envía el objeto y el middleware lo transforma a JSON y envía la respuesta.

https://github.com/go-kit/kit

Para su instalación, dentro del proyecto de Go tenemos que introducir el comando:
```text
go get github.com/go-kit/kit
```


### UUID - ORM for Golang
Es un package de Google y es prácticamente un standard en el creación de UUID.
Este package genera UUIDs basados en RFC 4122 y DCE 1.1: Authentication and Security Services.

[https://github.com/google/uuid](https://github.com/google/uuid)

### GORM - ORM for Golang
Esto es una librería que permite trabajar con la base de datos sin tener que crear las consultas SQL a mano.  La URL del sitio es:

[https://gorm.io/](https://gorm.io/)

Dentro del proyecto de Go tenemos que introducir los comandos:
```text
go get github.com/google/uuid
```

#### Instalar GORM:

Dentro del proyecto de Go tenemos que introducir los comandos:
```text
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```
Con esto ya tenemos las dependencias en el proyecto. Si queremos trabajar con otras bases de datos, podemos ver con las que es compatible en la URL:
[https://gorm.io/docs/connecting_to_the_database.html](https://gorm.io/docs/connecting_to_the_database.html )
#### Hooks

Son funcionalidades que se ejecutan antes o después de realizar una acción en la base de datos. Podemos consultar la documentación completa en:

https://gorm.io/docs/hooks.html

Los hook para cuando se crea un registro son:
* BeforeSave
* BeforeCreate
* AfterCreate
* AfterSave

Los hook para cuando se actualiza un registro son:
* BeforeSave
* BeforeUpdate
* AfterUpdate
* AfterSave

Los hook para cuando se elimina un registro son:
* BeforeDelete
* AfterDelete

El hook para cuando hacemos una query:
* AfterFind

### GoDotEnv
Permite definir variables de ntorno. Podemos ver más información en la página de packages de go [godotenv](https://pkg.go.dev/github.com/joho/godotenv#section-readme)

Por seguridad debemos evitar subir el fichero .env en entornos públicos. Para saber que variables necesitamos definir tenemos
el fichero .env.example que contiene las variables que necesitamos definir en nuestros ficheros .env

El comando para instalarlo como una librería en el proyecto es:
```text
go get github.com/joho/godotenv
```


# GIT
Para este proyecto hemos definido TAGS en GITHUB que nos servirán para gestionar las versiones de nuestros packages y poder
importar diferentes versiones mediante los tags. En git este TAG lo podemos crear dentro de RELEASE, por lo tanto, lo que tenemos que hacer
es:
1. Seleccionar la opción "Create a new release" y
2. Introducimos un nombre para la release con el formato recomendado por Go. Ej: v0.0.1
3. En la misma pantalla seleccionamos "Choose a tag" y le ponemos como tag la misma versión que en el punto anterior y clicamos en "Create new tag: V0.0.1"
4. Pulsamos el botón "Publish release"

Es importante que el formato de la versión/tag sea vx.x.x para que Go lo entienda como versión

También se puede importar mediante commit o branch, pero es mejor utilizar los tags para que quede ordenado.

En la página de Go https://go.dev/blog/publishing-go-modules se explica como recomienda Go el generar las versiones


## Problema con GIT y importar con Go proyectos privados:
Una posible solución puede estar aquí: https://stackoverflow.com/questions/58305567/how-to-set-goprivate-environment-variable

La fuente de este mensaje: https://github.com/fatih/vim-go/issues/3104

## Docker
La aplicación utiliza un contenedor docker para la base de datos MySQL, para ello hemos usado [docker hub mysql](https://hub.docker.com/_/mysql).
La configuración para ejecutar los contenedores se encuentra en el fichero docker-compose.yml. Una página donde podemos
encontrar un ejemplo para el fichero de configuración es esta: https://medium.com/@chrischuck35/how-to-create-a-mysql-instance-with-docker-compose-1598f3cc1bee

Para generar la imagen el comando a ejecutar es:
```text
docker compose build
```
El comando para levantar el contenedor es:
```text
docker compose up
```

Para lanzar el contenedor en segundo plano:
```text
docker compose up -d
```