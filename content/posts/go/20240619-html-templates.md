---
title: Introducción a los motores de plantillas
date: 2024-06-19
author: Pedro Pérez
tags: ["go", "template engine", "html", "text", "templates"]
categories: ["Go"]
description: "Desarrolladores y unos pocos analistas desconocen el
potencial que tiene el uso de un motor de plantillas en lugar de un framework
front-end completo. Te quitas el mayor problema de un plumazo: añadir más
dependencias al desarrollo."
draft: false
---

### Preámbulo

Tras un periodo desarollando con React y otro periodo con Svelte, cansado ya de
gestionar el back-end junto con el front-end por separado, llega el momento de
dar una oportunidad a desarollar con un motor de plantillas, el protagonista es
el de la librería estándar de Go, que son dos: _html/templates_ y 
_text/templates_. La diferencia principal radica en que la primera generan
ficheros HTML seguro frente a inyección de código malicioso.


### Postámbulo

Si vienes de un _front-end_ como _React_, _Angular_ o _Svelte_, hay que tener en
cuenta que el uso de motores de plantillas es un paradigma completamente
diferente. Cambia la forma de representar los datos, el concepto de _fetch API_
y el paso de datos en formato _JSON_ no aplican, no tal y como estamos
habituados.

El concepto y uso de un motor de plantillas hace que las vistas se acerquen todo
lo posible al _back-end_, mientras que el uso de un _framework front-end_ lo
aleja. Existen muchos motores en el mercado, [Jinja (Python)](https://jinja.palletsprojects.com/en/3.1.x/), 
[Thymeleaf (Java)](https://www.thymeleaf.org/), 
[Blade (PHP)](https://laravel.com/docs/11.x/blade#supercharging-blade-with-livewire)
y muchos más dentro de un mismo lenguaje.

Ademas, no sólo se limita al desarrollo web, los motores también se usan 
para generar documentos de texto como correos electrónicos, facturas, albaranes, 
tickets..., están presente en muchos sitios aunque no nos demos cuenta.

### ¿Qué es un motor de plantillas?

Tan simple como una herramienta de desarrollo que genera texto a partir de una 
plantilla dada. Un ejemplo:

#### Entrada:
```text
¡Hola, {name}! Hoy estamos a {date} y son las {time}.
```
- *name*: Juan
- *date*: `getTodayDate()`
- *time*: `getTimeNow()`

#### Salida:
> ¡Hola, Juan! Hoy estamos a 19/06/2024 y son las 19:44.

Similar a la interpolación de variables en cualquier _framework front-end_,
¿verdad?, de hecho, guardan muchas similitudes.

[Ver código fuente](https://github.com/zepyrshut/go-template-engine/tree/4be9a051fcb057c01e50d983b1bc0f127c41c85b)

### ¿Entonces, diferencias?

La principal: la forma de obtener y representar los datos. La secundaria: la
interactividad y dinamismo que ofrecen en cada forma.

**_Framework front-end_**

Hace una petición hacia una API y éste responde con un texto en formato _JSON_ (principalmente), trabaja y reformula los datos para su representación mediante
el uso de control de flujo desde el lado **cliente**.

![Framework front-end](/image/20240619-framework-front-end.png)

**_Motor de plantillas_**

Tan pronto como se solicite la plantilla, el servidor obtiene los datos, hace su
interpolación de variables mediante control de flujo y deja el documento listo
para ser representado, todo desde el lado **servidor**.

![Template engine](/image/20240619-template-engine.png)

Como ves en los diagramas, el uso de un _framework front-end_ se hace más
independiente del servidor, su interactividad aumenta, pues no necesita el
servidor para seguir trabajando con los datos. Mientras que el uso de motores de 
plantillas la relación con el servidor se estrecha, reduciendo su 
interactividad, tema que se resuelve con un poco de _JavaScript_, pero eso es 
otro para el blog.

### Concluyendo

Si el problema a resolver para el cliente es una particularidad o tiene un uso
limitado, lo que son la mayoría de los proyectos, recomiendo ferviemente el uso 
de motores de plantillas, pues reducirá dependencias y tiempo de desarrollo.
Un ejemplo podría ser una aplicación donde se quiera gestionar unas reservas de 
las distintas zonas de un gimnasio o reservas de un hotel.

Si el cliente pide una aplicación para su tienda _online_, con su panel de 
administración, gestión el catálogo, personalizada y desde cero, quizás encaje
mejor el uso de un _framework front-end_.

### Fuentes

[html/template](https://pkg.go.dev/html/template)

[Código fuente del proyecto](https://github.com/zepyrshut/go-template-engine)

### Datos técnicos

- **Entorno de desarrollo:** VSCode 1.91.0
- **Go:** go1.22.4 windows/amd64

