---
title: Diseñando una API con Go
date: 2023-03-26
author: Pedro Pérez
tags: ["go", "dependency", "injection", "inyeccion", "dependencias"]
categories: ["Go"]
description: "Descripcion aqui."
draft: true
---

### Preámbulo

La inyección de dependencia es un concepto de desarrollo en el que los 
distintos componentes de la aplicación son suministradas hacia un
paquete en lugar de este paquete crear los componentes a utilizar.

Es una forma de proporcionar a los distintos componentes lo que tiene
que hacer sin necesidad de saber el cómo lo tiene que hacer.

Para ello en Go se debe de hacer una envoltura llamada aplicación y
dentro de este tipo proporcionar todos los elementos necesario para su
correcto funcionamiento.

Recuerda que Go favorece la composición en todos los casos.

### 
