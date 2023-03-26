---
title: Interfaces en Go II - Implementación de Unmarshaler
date: 2023-03-25
author: Pedro Pérez
tags: ["interfaces", "go", "duck typing", "tipado de pato", "marshaler", "unmarshaler"]
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

En la primera parte se dio una [introducción a las interfaces en Go](/posts/2023/03/interfaces-en-go-i-introducción/), 
en esta segunda parte se va a ver las implementaciones de las interfaz
`Unmarshaler` de la librería estándar.

En la librería estándar de Go hay muchas interfaces con sus 
implementaciones por defecto, pero que también da la posibilidad de 
crear cualquier implementación personalizada para adaptar a las
necesidades del proyecto. 

Las implementaciones por defecto de la interfaz lo que hacen es
deserializar los tipos de datos de Go a JSON y viceversa a su tipo 
correcto.

### ¿Por qué una implementación personalizada?

`Unmarshaler` implementa el método `UnmarshalJSON([]byte) error`, la 
implementación por defecto de este método lo que hace es deserializar 
el tipo de dato desde JSON.

La necesidad de hacer una implementación personalizada de estas 
interfaces surgen de la necesidad de controlar el formato y los tipos al
transferir entre capas.

#### No siempre es necesario hacerlo...

Hacer una implementación de estas interfaces no siempre es la mejor
opción, es importante analizar y comprender el entorno por el que se
mueve el código fuente y hacer un equilibrio entre rendimiento y
mantenibilidad de código. Pues hacer las implementaciones de estas
interfaces implica una sobrecarga en memoria y procesador, haciendo la
ejecución más lenta, incluso bajando a la mitad, pero trae otras
ventajas que puede compensar. Si fuese analista, me ganaría la vida 
diciendo _depende..._, bromas aparte, analicemos los casos:

1. Importancia en la limpieza y mantenibilidad del código, si se
considera muy importante hacer un código limpio y fácil de mantener,
donde la transformación de JSON a tipo o viceversa requiera estar
en un sitio.

2. Formatos de fechas y horas personalizados, para este caso, basándome
en un caso que tuve en mi experiencia laboral, es importante implementar
esta interfaz para tener un mayor control de esos campos y unificar
campos de fechas y horas hacia el tipo `time.Time`.

3. Consideración de asignar valores por defecto para campos vacíos o
nulos.

4. Estructura que no corresponde al tipo estructurado, esto sería
similar a si se hiciese un _builder_, donde se recibe una estructura
concreta y se requiera transformar a otra estructura diferente 
preservando los datos o viceversa, incluso con capacidades de unir
varios tipos en un solo formato _JSON_.

### Caso 1: La pesadilla de las fechas

Se sabe que el tipo `time.Time` de Go es una fecha que cumple con el
estándar [RFC3339](https://www.rfc-editor.org/rfc/rfc3339), que viene a 
ser de la siguiente manera: `2006-01-02 15:04:05 +0100 CET`, también 
puede venir de otra forma, pero lo importante es que cumple con el 
estándar.

Sin embargo, desde un JSON puede venir la fecha por un lado y la hora 
por otro, fechas en formato `dd/MM/yyyy` o en `yyyy/MM/dd`,  puede venir
en un sinfín de formatos y es importante tener esto bien controlado.

Dado el tipo:

```go
type People struct {
	Name     string    `json:"name"`
	DateTime time.Time
}
```

Con el siguiente _JSON_:

```json
[{
  "name": "Chip",
  "date": "01/07/2022",
  "time": "12:22"
},
// More fields
]
```

Se deduce que no se puede pasar desde el JSON al tipo. Daría el
siguiente error:

```bash
parsing time "01/07/2022" as "2006-01-02T15:04:05Z07:00": cannot parse "01/07/2022" as "2006"
```

Con lo cual ahí surge la necesidad de implementar la interfaz 
`Unmarshaler`.

```go
func (p *People) UnmarshalJSON(data []byte) error {
	type Alias People
	aux := &struct {
		Date string `json:"date"`
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p.DateTime, err = DatesParser(aux.Date, aux.Time)
	return nil
}

// Esta función proporciona mecanismos para manejar datos vacíos en el
// JSON.
func DatesParser(dateStr string, timeStr string) (time.Time, error) {
	var dateString string
	if dateStr == "" {
		dateString = "01/01/0001"
	} else {
		dateString = dateStr
	}

	var timeString string
	if timeStr == "" {
		timeString = "00:00"
	} else {
		timeString = timeStr
	}

	_, err := time.Parse("02/01/2006", dateString)
	if err != nil {
		return time.Time{}, err
	}

	dateTimeString := fmt.Sprintf("%s %s", dateString, timeString)
	return time.Parse("02/01/2006 15:04", dateTimeString)
}

```

Por partes:

1. `type Alias People`

Creación de un nuevo tipo especial `Alias` que se basa en el tipo 
`People`, ambos [tipos son idénticos](https://go.dev/ref/spec#Type_identity).
Se hace esto para evitar la recursión infinita durante la 
deserialización _JSON_, ya que al deserializar, si no incorporamos el
`Alias`, llamaría a sí mismo automáticamente.

2. `aux := &struct { ... }`

En esta estructura anónima se incorpora los campos se necesite trabajar,
en este caso la fecha y hora, que vienen de tipo `string` acompañado de
la etiqueta `json:"date"` y `json:"time"` y un puntero de `Alias`, con
lo cual quedaría algo similar a lo siguiente en términos de campos
aunque su uso y tratamiento sean diferentes:

```go
type A struct {
  Date string `json:"date"`
  Time string `json:"time"`
  Name string `json:"name"`
  DateTime time.Time
}
```

3. `{ Alias: (*Alias)(p) }`

Hace dos cosas, primero convierte el tipo `*Persona p` al tipo `*Alias`
y se asigna al campo `*Alias` de la estructura anónima.

Llegado a este punto, la variable `aux` ya contiene todos los campos
rellenados gracias a `p`, pero faltan los otros campos extra, que son
los _especiales_: la fecha y hora.

4. `json.Unmarshal(data, &aux)`

Transforma todos los elementos del _JSON_ a tipo estructurado de `aux`.
Ahí se puede ver que se hace la deserialización sobre la estructura
anónima, así evitando entrar en una recursión.

5. El final

Justo ahí es donde reside la personalización de la deserialización,
poniendo las conversiones, campos o tratamientos necesarios, en este
caso se consigue combinar la fecha y la hora de tipo `string` y se pasa
a la función `time.Parse("formato esperado", dateTime)`.

Al ejecutar el código, se obtiene lo siguiente:

```text
{Chip 2022-07-01 12:22:00 +0000 UTC}
```

¡Y se obtiene el tipo correcto y bien formado!

Además, si se espera varios formatos de fechas diferentes pero se quiere
convertir a `time.Time` se puede crear una función encargada de las
fechas e invocarla.

### Conclusiones

En este artículo se ha visto cómo hacer una implementación del
`Unmarshaler` con el tema de las fechas, un asunto que puede ser muy
problemático si no se tratan correctamente desde el principio.

Espero que le sea de gran utilidad y si hay dudas no dude en contactar
conmigo por cualquier canal disponible.

### Fuentes

[Custom JSON Marshaler in GO](https://medium.com/picus-security-engineering/custom-json-marshaller-in-go-and-common-pitfalls-c43fa774db05)

[json](https://pkg.go.dev/encoding/json)

[Código fuente del proyecto](https://github.com/zepyrshut/unmarshaler-interface)

### Datos técnicos

- **Entorno de desarrollo:** GoLand 2022.3.2
- **Go:** go1.20.2 windows/amd64
