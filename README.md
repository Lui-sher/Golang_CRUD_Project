# Golang_CRUD_Proyect

Golang CRUD Proyect (Neon-API)

En este proyecto pondremos en practica lo aprendido sobre el lenguaje Go.

Desarrollaremos una REST-API (CRUD), alojando la base de datos online en 
el servidor que nos ofrece NEON
Link de instalacion:

Base de datos:  PosgreSQL alojado en el servidor online Neon
link del servidor: https://console.neon.tech/app/projects

Adicionalmente pondremos a prueba el uso del framework FIBER, que en GO 
funciona como vendria siendo EXPRESS para NODE de Javascript, nos permite 
gestionar facilmente las rutas URL para las quest hechas por el cliente.

por otro lado utilizaremos la herramiente AIR, que vendria siendo como Nodemon
para Javascript que nos reiniciará el servidor cada ves que guardemos un cambio
en nuestro codigo.
Lo iniciaremos con el comando dentro de la carpeta del proyecto: $air

Rutas definidas para interaccion con el cliente desde la web:

------------------Peticiones GET---------------------------

:PORT/      
            ->  Solicitud a Pagina Pricipal de conexion con el servidor
                //"¡Bienvenido al servidor Principal!"

:PORT/db    
            ->  Solicituda a la ruta Pricipal de la base de datos
                //"Aquí irán las operaciones con la base de datos"

:PORT/db/last-record
            ->  Retorna al cliente el ultimo registro almacenado en la
                tabla de usuarios (users)

:PORT/db/find-user/:user_id
            ->  Retorna al cliente todo el registro de un usuario 
                almacenado en la tabla "Users", usando como punto
                de busqueda el user_id, que es un codigo unico
                generado automaticamente al crear un usuario

// ------------------------Peticiones POST ------------------------
:PORT/db/create/user
            ->  Esta ruta necesita que se le envie por Body en formato
                JSON, la estructura:
                {
                    "name" : "fulado",
                    "email" : "fulano@email.com"
                }

                Esta peticion le retorná al cliente un JSON con los datos
                completos generados en la tabla con el siguiente formato:
                {
                    "email": "fulano@email.com",
                    "message": "Data inserted successfully",
                    "name": "fulado",
                    "user_id": "b690835bd8"
                }



Ruta TEST
Esta Ruta simula:
1.  Creacion automatica de 3 nuevos registros en la tabla 'users'
Method: 
        GET

Ruta:
        :PORT/db/post/test

2.  Consulta todos los registros creados por el test anterior y los muestra 
    por consola.
Method:
        GET
Ruta:
        :PORT/db/get/test


4.  Actualiza o edita los registro creados anteriormente.
5.  Elimina los registros creados anteriormente.
