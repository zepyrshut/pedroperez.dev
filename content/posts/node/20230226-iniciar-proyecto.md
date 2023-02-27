---
title: Iniciar un proyecto con NPM y JavaScript
date: 2023-02-26
author: Pedro Pérez
tags: ["node", "npm", "javascript"]
categories: ["Node"]
description: "Inicializar un proyecto con NPM, el uso de package.json
y los elementos recomendados para un desarrollo colaborativo de calidad,
como el uso de linters y formateadores de código, así como su
comprobación del código fuente antes de subir al repositorio git."
draft: false
---

### Preámbulo

La inicialización de un proyecto basado en un entorno de ejecución Node
con el gestor de paquetes NPM es una cosa muy trivial, pero idealmente
sería la inicialización con un buen _linter_ para frenar en la mayor
medida de lo posible malas prácticas en el código y un formateador en el
que se establece unas directrices aplicables a nivel proyecto donde la
indentación, líneas y formatos siempre es la misma. En este caso, con
un enfoque a JavaScript.

También, además de todo lo anterior, la integración de una serie de
_scripts_ en el que hará todas las tareas automáticamente antes de subir
al repositorio, en caso de incumplirse alguna regla del _linter_ el
desarrollador se verá obligado a corregirla para poder hacer _push_ al
_commit_.

### Lo primero

Iniciarlizar un proyecto, con un comando que todos sabemos: `npm init`.
En caso de querer usarse _Vite_, se puede usar el comando `npm create vite@latest`.

Para este caso, nos centraremos sin Vite.

Tras haber ejecutado el comando de inicialización de un proyecto,
en el espacio de trabajo estará el fichero _package.json_. Similar al
siguiente:

```json
{
  "name": "npm-init",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC"
}
```

Lo primero que toca es el uso de [StandardJS](https://standardjs.com/).

### StandardJS

Es una librería que tiene como objetivo unificar el formato y el uso de
estilos de forma consistente a nivel proyecto. Es una herramienta
todoterreno, apoyada por NodeJS, NPM, GitHub, etc.

Instalación con: `npm i standard -D`.

Se guarda como dependencia de desarrollo, eso significa que esa librería
se usa durante el desarrollo pero es ignorada durante la construcción
del proyecto.

Para usar esa librería, basta con escribir `npx standard`, se nos
mostrará por consola todo los errores que tiene en el proyecto. Dado
este fragmento de código:

```js
function add(a, b) {
  return a + b;
}

let result = add(1, 2)

console.log(result)
```

Nos dirá lo siguiente:

```bash
index.js:1:13: Missing space before function parentheses. (space-before-function-paren)
index.js:2:15: Extra semicolon. (semi)
index.js:5:5: 'result' is never reassigned. Use 'const' instead. (prefer-const)
```

E incluso se puede arreglar automáticamente, en la medida de lo posible
con el comando `npx standard --fix`, y para mostrarse los errores en
el propio editor basta con instalar la extensión con el mismo nombre.

![Reglas violadas StandardJS](/image/20230226-iniciar_1.png)

En la imagen sale marcado en rojo las reglas violadas, y si colocas el
ratón encima, te muestra qué regla violó y si es posible, un acceso
directo a un arreglo rápido.

### Prettier

Para todos los archivos que no sean JavaScript se hace uso de Prettier
como formateador de código, válido para muchos formatos como _json_, 
_HTML_, _CSS_, y muchos más, además tienen algunas cositas interesante
que veremos a lo largo de la sección.

Instalación: `npm i prettier -D` y su uso es con el
comando: `npx prettier --write .`. Tenga preacaución que al ejecutarse
ese comando, escribirá en todos los archivos del proyecto.

Antes que nada, como no interesa que _Prettier_ incida sobre los
archivos JavaScript, se añade el fichero `.prettierignore` y entra todos
los ficheros, directorios o extensiones a ignorar.

```text
*.js
*.cjs
build/
coverage/
```

Además, recomiendo dos _plugins_ de _Prettier_: _organize-attributes_
y si el proyecto usa _TailwindCSS_, el _plugin_
_TailwindCSS x Prettier_.

#### organize-attributes

Este _plugin_ necesita _Prettier_ para funcionar y lo único que hace es
reordenar los atributos HTML para que en todas las etiquetas sigan el
mismo orden.

Instalación: `npm i prettier-plugin-organize-attributes -D`

```html
<img src="#" class="some-class" alt="some alt text" />
```
Tras ejecutar _Prettier_ con el _plugin_.

```html
<img class="some-class" src="#" alt="some alt text" />
```

#### TailwindCSS x Prettier

Es similar al anterior, pero enfocado a las clases TailwindCSS donde su
función es reorganizar las clases para que estén todas agrupadas. Es
oficial de _TailwindLabs_. Es necesario tener _TailwindCSS_ instalado.

Instalación: `npm i prettier-plugin-tailwindcss -D`.

```html
<div class="underline text-3xl font-bold text-red-300">Hello World</div>
```
Tras ejecutar _Prettier_ con el _plugin_.

```html
<div class="text-3xl font-bold text-red-300 underline">Hello World</div>
```

#### ¿Formatear al guardar?, no

En los editores de código y en los entornos de desarrollo integrados es
común que traigan la opción de aplicar el formateo de código al guardar,
no lo recomiendo ya que es posible que se nos distraiga a la hora de
desarrollar.

Es común que, por ejemplo en HTML es dejar mucho espacio por arriba y
por abajo sobre el fragmento que se está trabajando, y al llenarse de
etiquetas, atributos y propiedades, al guardar, se nos juntará todo, así
haciendo más difícil de localizar.

### Husky y lint-staged

Son herramientas que prepara el proyecto para estar todo correctamente
formateado y analizado, impidiendo subir archivos al repositorio si no
cumple con las reglas establecidas, así obligando al desarrollador
arreglar todas estas incidencias.

Husky es un sistema en el que ejecutan acciones según la situación dada,
por ejemplo, si se configura para ejecutar una serie de _scripts_ antes
de hacer _commit_, lo hará dado el momento. Se basan en los [_git hooks_](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks).

lint-staged es un pequeño _plugin_ en el que impide que los _linters_ o
formateadores de código trabajen sobre todo el proyecto y lo hagan solo
sobre los archivos que están confirmados.

Antes que nada, el proyecto debe tener _git_, en caso de que no lo tenga
lo puedes inicializar con `git init`.

Instalación: `npm i husky lint-staged -D`, después ejecutar `npx husky install`.

El segundo comando lo que hace es instalar los _ganchos git_, que son
ejecutados al hacer operaciones _git_ como un _commit_ o un _push_.

Hay que añadir el _script prepare_ al fichero _package.json_: `npm pkg set scripts.prepare="husky install"`. Deberás tener el _script_ en el fichero.

```json
// package.json
"scripts": {
  // some scripts
  "prepare": "husky install"
},
```
El uso de `npm run prepare` es para tener _Husky_ listo para funcionar
en el entorno donde se descargue el proyecto.

En la carpeta _.husky_ debemos crear un fichero llamado _pre-commit_ sin
extensión e introducir lo siguiente:

```sh
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npx lint-staged
```
Y en _package.json_, al final, lo siguiente:

```json
// package.json
"lint-staged": {
    "*.{html,css,scss}": [
      "prettier --write"
    ],
    "*.{js,ts}": [
      "standard --fix"
    ],
  },
```

La combinación del _hook pre-commit_ con la acción `npx lint-staged` es
lo mismo que si el usuario ejecutase el _script_ `npx lint-staged`, 
pero lo hará _Husky_ por nosotros al momento de hacer el _commit_. Si
se quisiera hacer una revisión antes de hacer el _push_, se puede crear
otro fichero con el nombre de _pre-push_.

### Algo de pruebas unitarias

Nos limitaremos a instalar solamente la librería y un caso de prueba,
con la idea de ejecutar todas las pruebas antes de hacer un _push_.

Instalación: `npm i jest -D`

La función sumar:
```js
function add (a, b) {
  return a + b
}
```
Las pruebas:
```js
test('adds some numbers', () => {
  expect(sum(1, 2)).toBe(5);
  expect(sum(3, 9)).toBe(12);
});
```

Y finalmente, crear un fichero en el directorio _.husky_ un archivo
con el nombre de _pre-push_, en él debe de ir lo siguiente:

```bash
#!/usr/bin/env sh
. "$(dirname -- "$0")/_/husky.sh"

npm run test
```

Con esto se ejecutarán todas las pruebas antes de enviar al repositorio.

### Fuentes

prettier

standardjs

npm

node

[organize-attributes](https://github.com/NiklasPor/prettier-plugin-organize-attributes)

[TailwindCSS x Prettier](https://github.com/tailwindlabs/prettier-plugin-tailwindcss)

husky

### Datos técnicos

- **Entorno de desarrollo:** Visual Studio Code 1.73.1
- **Node:** v18.12.1
- **NPM:** 9.3.0 
- **Prettier:** 2.8.4
  - **organize-attributes:** 0.0.5
  - **Tailwind x Prettier:** 
- **StandardJS:** 17.0.0

