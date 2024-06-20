---
title: Introducción al motor de plantillas de Go
date: 2024-06-20
author: Pedro Pérez
tags: ["go", "template engine", "html", "text", "templates"]
categories: ["Go"]
description: "Posiblemente si estás acostumbrado a los frameworks front-end, 
este camino te puede resultar algo amargo, pues como comenté en el artículo 
anterior, cambia totalmente la forma de pensar al ser otro paradigma diferente. 
Intentaré ser lo más suave posible."
draft: true
---

### Preámbulo

[Es una continuación de este artículo](https://pedroperez.dev/posts/2024/06/introducci%C3%B3n-a-los-motores-de-plantillas/).

En estas dos últimas semanas me encargaron un proyecto de gestionar las entradas
para un festival de mi localidad, Mazagón (Huelva). La funcionalidad principal
es poder pagar las consumiciones con la misma entrada, previamente recargada de
un cantidad de dinero. Pues en este artículo me voy a basar en ese mismo 
proyecto, con ejemplos y casos de usos reales.

Al ser sencilla no dudé ni un segundo en probar por
primera vez el uso de _html/templates_ en lugar de _Svelte_. Previamente ya
trabajé con motores en un curso de Go, pero en aquel momento (hace más de un 
año) no fui capaz de ver el potencial. Ahora sí que me ha convencido totalmente 
en pro de usar _Svelte_, _Angular_ o _React_.

### Lo más básico

Antes que hay que preparar el servidor, para ello he dejado un [repositorio](https://github.com/zepyrshut/go-template-engine/tree/91ecd48e9086db9512451244b7af7a48ed969228)
con el _commit_ listo para empezar, recomiendo visitarlo y hacer un _clone_. 
Se compone de lo siguiente:

```go
package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// renderTemplate se encarga de obtener la ruta de la plantilla, hacer el
// procesamiento y su ejecución.
func renderTemplate(w http.ResponseWriter, templateName string) {
	path := filepath.Join("templates", templateName)
	tmpl, _ := template.ParseFiles(path)
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})

	http.ListenAndServe(":8080", nil)
}
```
Con ese ejemplo mínimo, si ejecutamos el servidor y vamos a la ruta raíz:
`localhost:8080` veremos exactamante lo mismo que el fichero `index.html`.

### Añadiendo datos a la plantilla

Pero realmente lo que queremos es añadir información dinámica a una plantilla,
en este ejemplo pondré un título y un nombre. Para ello debemos hacer cuatro 
cosas:

1. Crear un _struct_ que se encargue de contener los datos para su 
representación.

```go
type TemplateData struct {
    Title string
    Name string
}
```

2. Añadir un parámetro a la función `renderTemplate(...)`, el de los datos para 
mostrar, y pasar el parámetro a `tmpl.Execute(w, data)`

```go
func renderTemplate(w http.ResponseWriter, templateName string, data TemplateData) {
    ...
}
```

3. Inicializar el _struct_ con los datos que correspondan y pasarlo por el
parámetro.

```go
	templateData := TemplateData{
		Title: "Hola Mundo",
		Name:  "Gopher",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", templateData)
	})
```

4. Modificar la plantilla `index.html` para leer esos datos, **atención** al 
detalle, el _struct_ es de tipo `TemplateData` con dos campos: `Title` y `Name`. 

```html
...
  <title>{{ .Title }}</title>
...
 <h1>¡Hola {{ .Name }}!</h1>
...
 ```

 Se hace la interpolación de las variables directamente en los nombres de los
 campos y no desde `.TemplateData.Title` y `.TemplateData.Name`. Con ello
 deberíamos ver la plantilla ya procesada con las variables introducidas.

```text
¡Hola Gopher!
Ejemplo mínimo de una plantilla procesada desde Go
Este documento proviene de una plantilla llamada index.html.

Ahora tiene variables, pero no tiene estilos ni formato.
```

> Cuando se modifica el código fuente de Go, hay que detener e iniciar el 
> programa en cuestión. Esto no ocurre cuando modificamos las plantillas, pues
> podemos modificarlas y solicitar el recurso al servidor.

Si te sientes perdido, puedes ver el [_commit_](https://github.com/zepyrshut/go-template-engine/tree/4be9a051fcb057c01e50d983b1bc0f127c41c85b).

### Añadiendo más plantillas

Sigamos con ello, ahora lo que veremos es el uso de varias plantillas para
procesarse en una sola, y así usar el mismo código para las distintas partes
como pueden ser el `head`, `nav`, `footer` y cualquier cosa que quieras 
repetir.













