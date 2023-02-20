---
title: Validación de datos con Gin-Gonic y Validator
date: 2023-02-20
author: Pedro Pérez
tags: ["validación", "gin-gonic", "validator", "go"]
categories: ["Go"]
description: "Además de la validación de datos, también en la
personalización de los mensajes de error para ser mostrados al usuario
de una forma más sencilla."
draft: false
---

### Preámbulo

Este artículo da por hecho que el autor conoce el _framework_ Gin-Gonic
y está habituado al uso, sin embargo necesita un conocimiento mayor
sobre la validación de datos.

La validación de datos enviados desde el cliente es un concepto muy
importante, pues se evita en gran parte los problemas de inyección
de código malicioso. Sin embargo, también es útil para mostrar
información al usuario de que algo no va bien como un inicio de sesión
incorrecto o un campo que requiere sólo números y hay texto.

El _framework_ Gin-Gonic incluye un validador de entrada de datos
proporcionada por la librería [go-playground/validator](https://github.com/go-playground/validator).

### go-playground/Validator

Es una potente librería de validación de datos de _structs_ o de los
atributos de los structs mediante etiquetas (_tags_).

```go
type Login struct {
	User     string `json:"username" binding:"required,gt=0,email"`
	Password  string `json:"password" binding:"required,gt=0"`
}
```

En el _struct_ de arriba verás en una de las etiqueta lo siguiente:

```go
`binding:required,gt=0,email`
```

Esto significa que es un campo obligatorio, debe ser mayor de 0 y debe
ser un correo electrónico. Cualquier elemento que se salga de las
directrices no será valido.

El problema viene a ser el siguiente, que dada una petición HTTP con
datos no válidos, el error que muestra es el siguiente:

```json
{
  "error": "Key: 'Login.User' Error:Field validation for 'User' failed on the 'email' tag"
}
```

Con un mensaje de error así se hace muy difícil tratar los datos para
mostrar los mensajes al usuario de que algo está mal. Arreglemos esto.

### La interfaz `error`

```go
type error interface {
    Error() string
}
```

El funcionamiento es similar a la interfaz `stringer`, donde se puede
crear distintas implementaciones de esta interfaz añadiendo mensajes
personalizdos para facilitar al desarrollador o usuario el error que
está sucediendo.

En otras palabras, cualquier implementación del método `Error()` se
puede transformar cualquier tipo que se defina en un error propio.

Una implementación cualquiera de la interfaz `Error()`:

```go
type DivideByZero{}

func (customErr *DivideByZero) Error() string {
  return "No se puede dividir por 0!"
}
```

Además de ello, Go también proporciona la posibilidad de generar un
nuevo error sin necesidad de definir un `struct` para ello llamando a la
función `New()`, pero esto es un tema para otra publicación.

### ValidationErrors

Para obtener los mensajes de error de la validación tenemos que hacer
uso de la implementación del método `Error()` del tipo
`ValidationErrors`, que es donde contiene toda la información del error.

```go
type ValidationErrors []FieldError

func (ve ValidationErrors) Error() string {

	buff := bytes.NewBufferString("")

	var fe *fieldError

	for i := 0; i < len(ve); i++ {

		fe = ve[i].(*fieldError)
		buff.WriteString(fe.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}
```

El tipo `ValidationErrors` es un _array_ de `FieldError` Un repaso a la
interfaz, que a su vez, esta interfaz llama a la función `Error()` para
ese mismo campo.

```go
type FieldError interface {
	Tag() string
	ActualTag() string
	Namespace() string
	StructNamespace() string
	Field() string
	StructField() string
	Value() interface{}
	Param() string
	Kind() reflect.Kind
	Type() reflect.Type
	Translate(ut ut.Translator) string
	Error() string
}
```

### Dando sentido a todo

Siguiendo esta lógica, lo que habría que hacer es iterar sobre el tipo
`ValidationErrors` para obtener los errores de cada campo. Se define una
variable del tipo y la iteración.

```go
func GetValidationErrors(err error) {
  var verr validator.ValidationErrors
  if errors.As(err, &verr) {
    for _, f := range verr {
      fmt.Printf("field %s, tag %s", f.Field(), f.Tag())
    }
  }
}
```

Se hace uso de la función `As()` del paquete `errors`. Esta función
tiene dos parámetros, el error que devuelve la función `c.ShouldBindJSON(&someVar)`
y un puntero a `verr`, el mecanismo de esta función es buscar el primer
error que coincida con el tipo del puntero, en este caso `ValidationErrors`
y se almacena en `verr`. Finalmente, se imprime por consola los errores
del tipo.

```go
field User, tag required
```

Este mensaje significa que en el campo `User` no se cumple la validación
de la etiqueta `required`, en otras palabras: se está enviando este
campo vacío.

### Haciéndolo JSON

El paso final sería devolver los errors en formato JSON para el
tratamiento en el _frontend_ y mostrar los mensajes correspondientes al
error dado.

Para ello se precisa de hacer una función que recoge como parámetro
el tipo `validator.ValidationErrors` y se itera sobre todos los errores
recogidos para su transformación en pares de clave/valor.

### Fuentes

Este artículo es una reintrerpretación del siguiente artículo:

[Gin Validation Errors Handling](https://blog.depa.do/post/gin-validation-errors-handling)

[FieldError Interface](https://pkg.go.dev/github.com/go-playground/validator/v10#FieldError)
