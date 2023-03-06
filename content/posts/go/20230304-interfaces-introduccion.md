---
title: Interfaces en Go I - Introducción
date: 2023-03-04
author: Pedro Pérez
tags: ["interfaces", "go", "duck typing", "tipado de pato"]
categories: ["Go"]
description: "Es muy importante conocer y comprender las interfaces en
Go, su funcionamiento y su uso. El tipado de pato es una característica
muy común en los lenguajes débilmente tipado; sin embargo, en Go
combinan características de tipado estático y tipado dinámico a pesar
de ser un lenguaje fuertemente tipado. Un jaleo, ¿verdad?"
draft: false
---

### Preámbulo

> Para entender esta publicación, debe tener un conocimiento competente
> acerca de los tipos de datos en Go, los tipos estructurados y los
> receptores de función.

Las interfaces de Go se considera _duck typing_, tipado de pato en
español. Eso significa que si un animal camina como un pato, nada como
un pato y grazna como un pato, probablemente sea un pato, pero esto no
significa que sea un pato realmente, habría que hacerle una prueba de
ADN para saberlo, y esto es muy caro.

Esta definición que se explica en los libros o en artículos puede ser
algo confusa si no viene acompañado de algunos ejemplos o explicaciones
alternativas.

### Las interfaces se enseñan mal

El nombre de la sección viene del vídeo que se hizo meme comentando que
el [inglés se enseña mal](https://www.youtube.com/watch?v=vR_b3Mf3b24).

Pues con las interfaces pasa lo mismo, la explicación que siempre dan
es:

> Una interfaz define qué es lo que hay que hacer y no el cómo.

Esta explicación realmente es muy clara y concisa, pero no para una
persona que está aprendiendo a programar. Lo explico de otra forma:

> Una interfaz tiene una serie de métodos que deben ser implementados
> por un tipo.

Mmm... me parece muy vaga también. Otro intento, pero algo más
extendido:

Imagina que estás desarrollando unos métodos para trabajar con una base
de datos, los métodos son: `Create()`, `Update()`, `Delete()` y
`Read()`, también conocido como CRUD.

Pues estos métodos están diciendo qué es lo que tiene que hacer, es 
decir: _crear_, _actualizar_, _borrar_ y _leer_, pero no están diciendo
cómo tienen que hacer esas acciones, pues eso puede variar dependiendo
del sistema gestor de base de datos que se esté usando, el cómo a nadie
le importa.

Pues desde otra capa de la aplicación llaman a las acciones _CRUD_ de la
interfaz y no de sus implementaciones.

Esto da lugar a que si la aplicación crece y se necesita cambiar el
sistema gestor de bases de datos, no haya cambios ninguno en la capa
donde hacen llamadas a las acciones _CRUD_.

### Explicando la interfaz de Go

Las interfaces en Go son un poquito diferentes a las interfaces de Java,
C# y otros lenguajes con tipado estático. En Go las interfaces son
implementadas implícitamente, es decir, no hay que declarar que un tipo
implementa una interfaz, sino que se implementa implícitamente si el
tipo implementa todos los métodos de la interfaz.

> Si una persona saluda y un perro saluda, entonces ambos tipos son
> saludadores.

```go
type Greeter interface {
	Greetings()
}

type Person struct {
	Name  string
	Genre string
}

type Dog struct {
	Name string
	Race string
}

func (p Person) Greetings() {
	fmt.Println("Hola, soy", p.Name)
}

func (p Dog) Greetings() {
	fmt.Println("Guau, soy", p.Name)
}

func main() {
	// Person
	var p Greeter = Person{Name: "Juan", Genre: "Male"}
	p.Greetings()

	// Dog
	var d Greeter = Dog{Name: "Firulais", Race: "German Shepherd"}
	d.Greetings()

	// Dog to Person can be done because both are Greeter
	p = d

	// Demonstration of type assertion, p is a Dog because it was
	// assigned to p in the previous line (p = d).
	dog := p.(Dog)
	fmt.Println(dog)
}
```

Si os fijáis en la función `main` se puede ver que se está creando una
variable del tipo Greeter, pero se le están asignando valores del
tipo Person y del tipo Dog. Esto es posible porque ambos tipos
implementan el método `Greetings()`, por lo tanto, ambos tipos son
Greeters.

Pero además, se puede ver que se puede asignar un valor del tipo Dog al
tipo Persona porque ambos tienen interfaces compatibles.

> Cualquier tipo que satisface una interfaz también es del tipo de la
> interfaz.

Esta frase es muy importante entenderla. Lo puedo decir de otra manera:

> Cualquier tipo que implemente todos los métodos de la interfaz también
> es del tipo de la interfaz.

Dicho de otra manera:

> Dada la interfaz `I` tiene el método `M`, cualquier tipo _struct_ que 
> implemente el método `M` también es del tipo `I`.

### Duck typing

La idea del _duck typing_ es que cualquier implementación de una
interfaz puede ser utilizada en cualquier parte donde se necesite el
tipo de la interfaz.

Es decir, si una función tiene como parámetro un tipo `I`, cualquier
tipo que implemente los métodos de la interfaz `I` puede ser pasado a
esa función.

```go
func SomeFunction(g Greeter) {
    g.Greetings()
}
```

La función `SomeFunction` puede ser utilizada por los tipos `Person` y
`Dog`.
