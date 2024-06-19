---
title: El motor de plantillas de Go
date: 2024-06-19
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

### Lo más basico












