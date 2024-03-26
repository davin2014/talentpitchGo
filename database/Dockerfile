# FROM es una instrucción que establece la imagen base para las instrucciones 
# subsiguientes. En este caso, estamos utilizando la imagen de PostgreSQL versión 10.3.
FROM postgres:10.3

# COPY es una instrucción que copia nuevos archivos o directorios desde <src> y 
# los agrega al sistema de archivos del contenedor en la ruta <dest>.
# Aquí, estamos copiando el archivo up.sql al directorio /docker-entrypoint-initdb.d/ 
# con el nombre 1.sql. Este directorio es especial en la imagen de PostgreSQL, 
# cualquier script SQL o de shell en este directorio se ejecutará automáticamente 
# al iniciar el contenedor.
COPY up.sql /docker-entrypoint-initdb.d/1.sql

# CMD proporciona valores predeterminados para un contenedor en ejecución. 
# Estos valores pueden incluir el ejecutable, o pueden omitir el ejecutable, 
# en cuyo caso deberá proporcionar un ENTRYPOINT. 
# Aquí, estamos estableciendo que el contenedor ejecute el servidor PostgreSQL.
CMD ["postgres"]
