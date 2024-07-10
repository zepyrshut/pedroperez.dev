---
title: Festival Arena
author: Pedro Pérez
---

En Mazagón (Huelva), el 20 de julio se celebra un festival de música urbana con 
artistas como Chris MJ y Chimbala, además de la participación de otros artistas 
locales para darles visibilidad. [Enlace a la noticia](https://huelvabuenasnoticias.com/2024/07/10/mazagon-se-convertira-en-el-epicentro-de-la-musica-urbana-con-el-festival-arena/).

### Problema a resolver

En festivales y eventos multitudinarios es fundamental gestionar de una manera
eficiente y segura las transacciones de bebidas, comidas y merchandising. El uso
de pagos en efectivo genera inseguridad y problemas logísticos en la _"barra"_. 

Es por ello que se evoluciona a un sistema de _tickets_ donde las transacciones 
de dinero por tickets se concentra en un único sitio y en el lugar de las 
consumiciones se entrega el _ticket_ con un valor en cuestión.

El sistema mencionado puede ser interesante pero implica un especial gasto en
papel, también genera mucho residuo, además del problema de seguridad frente a 
robos sigue estando ahí.

### Requisitos impuestos

Tras una reunión con los organizadores del evento para entender lo que buscan,
me cuentan los siguientes requisitos:

1. Cada cliente lleva un código QR, este código debe ser la cartera.
2. El código QR debe poder escanearse desde el móvil.
3. Tres grupos de empleados: administrador, recargador y vendedor.
    1. De estos grupos, en vendedor está dividido en 3 grupos más:
        1. Comidas
        2. Bebidas
        3. Merchandising

    Cada vendedor puede vender las distintas combinaciones asignadas.
4. Consulta de saldo, accesible para todo el mundo, tanto clientes como 
empleados.
5. Panel de administración en el que se pueda añadir nuevos clientes (QR) y
nuevos empleados.
6. Panel de auditación de entrada de dinero y ventas. En tiempo real y con
filtros.

Una vez acordado los requisitos, me pongo manos a la obra.

### Aspectos técnicos

1. #### Tecnologías

Desarrollado en Go, haciendo uso del [motor de plantillas](/tags/template-engine/)
de la librería estándar para las vistas.
Bases de datos en PostgreSQL y empaquetado en un contenedor Docker para ser
desplegado en un servidor privado.

2. #### _Backend_

Las tres dependencias más importantes son [Go Chi](https://github.com/go-chi/chi)
para el router, [PGX](https://github.com/jackc/pgx) para la conexión con la base
de datos y [SCS](https://github.com/alexedwards/scs) para la gestión de sesiones
de usuario. Se han utilizado otras para validación y gestión  de variables de 
entorno.

3. #### _Frontend_

Para facilitar el desarrollo de las plantillas y sus partes repetitivas se ha
hecho con la ayuda de [Gulp](https://gulpjs.com/), para el escáner QR, la
librería [html5-qrcode](https://github.com/mebjas/html5-qrcode) y para añadir un
poco de dinamismo [htmx](https://github.com/bigskysoftware/htmx/). Estilos en
[TailwindCSS](https://tailwindcss.com/) con la ayuda de [Flowbite](https://flowbite.com/)
para acelerar el proceso.

4. #### Otros

En todos los proyectos repito el mismo patrón a nivel gestión como el uso de
_Git_, empaquetado en _Docker_ para ser desplegado en un servidor privado
totalmente gestionado por mi con todo lo que implica. _Make_ para los procesos
de reconstrucción de la base de datos, empaquetado, construcción del proyecto
y alojamiento de la imagen en el servidor.

### Imágenes

|   |   |
|---|---|
| ![](/image/festival-arena/inicio-sesion.png)      | ![](/image/festival-arena/no-existe.png) |
| ![](/image/festival-arena/incorrecto.png)         | ![](/image/festival-arena/no-permitido.png) |
| ![](/image/festival-arena/vender-1.png)           | ![](/image/festival-arena/vender-2.png) |
| ![](/image/festival-arena/venta-realizada.png)    | ![](/image/festival-arena/venta-realizada-2.png) |
| ![](/image/festival-arena/no-existe-codigo-2.png) | ![](/image/festival-arena/no-saldo.png) |
| ![](/image/festival-arena/bebidas.png)            | ![](/image/festival-arena/recargar-1.png) |
| ![](/image/festival-arena/recargar-2.png)         | ![](/image/festival-arena/no-existe-codigo.png) |
| ![](/image/festival-arena/recarga-5.png)          | ![](/image/festival-arena/recarga-100.png) |
| ![](/image/festival-arena/consulta-saldo-1.png)   | ![](/image/festival-arena/consulta-saldo-2.png) |
| ![](/image/festival-arena/nuevo-cliente-1.png)   | ![](/image/festival-arena/nuevo-cliente-2.png) |

![](/image/festival-arena/recargas.png) 

![](/image/festival-arena/ventas.png) 

### Conclusión

Interesante proyecto donde he aprendido el uso de una serie de elementos de
_frontend_ como _Gulp_, _HTMX_, el escáner QR.

Gracias a este proyecto he vuelto a reconsiderar el uso de motores de plantillas
por encima del uso de un _framework_ como _React_ o _Svelte_. Pues con el motor
y el uso de _HTMX_ para darle dinamismo es más que suficiente, especialmente en
los perfiles donde el desarrollador trabaja en solitario; donde se reuqiere más 
dedicación al _backend_ o para equipos de muy bajo presupuesto donde hay que
hacer uso de desarrolladores _fullstack_.

Una de las mayores complicaciones técnicas que se ha tenido que abordar es el 
uso de HTMX. Esta potente librería evita el desarrollo indiscriminado de bloques
JavaScript. Sin embargo, gestionar el estado de la aplicación ha sido desafiante
pues aun no se ha conseguido dominar por completo. A pesar de ello, se pudo 
resolver el problema de rendimiento relacionado con la reutilización de la 
instancia de la cámara.

En este proyecto no solo hay que destacar la importancia de la digitalización de
eventos multitudinarios, sino también resaltar el uso de tecnologías asequibles
para un desarrollo en un marco de tiempo muy reducido. Posiblemente para el
próximo proyecto vuelva a repetir la fórmula de _html/template_ y _HTMX_.
