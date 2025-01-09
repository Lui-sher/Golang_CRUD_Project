Definicion de Rutas:

------------------Peticiones GET---------------------------

:PORT/      
            ->  Solicitud a Pagina Pricipal de conexion con el servidor
                //"¡Bienvenido al servidor Fiber!"

:PORT/db    
            ->  Solicituda a la ruta Pricipal de la base de datos
                //"Aquí irán las operaciones con la base de datos"

:PORT/db/last-record/:table
            ->  solicitud a una tabla especifica (:table) que retorna
                un Json con la informacion del ultimo registro que se 
                halle en la tabla seleccioneada

Ruta TEST
Esta Ruta debe simular:
1. Verificacion de tablas en la base de datos.
2. Creacion de un nuevo registro en la tabla 'users'.
3. Consultar el registro creado.
4. Actualizacion del registro creado anteriormente.
5. Eliminacion del registro creado.

