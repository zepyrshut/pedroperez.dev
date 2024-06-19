---
title: Instalando servidor de impresión en Debian
date: 2024-05-02
author: Pedro Pérez
tags: ["servidor", "impresión", "impresora"]
categories: ["Servidores"]
description: "Descripcion aqui."
draft: true
---

### Preámbulo

Un servidor de impresión es muy útil para entornos empresariales o en
oficinas, e incluso en entornos domésticos donde una impresora es
utilizada por varios miembros de la unidad doméstica, la idea es que se
pueda imprimir documentos desde cualquier dispositivo, con la única
condición de que tenga que estar conectado a la misma red que el 
servidor.

> La impresora a configurar es **Samsung M2026**, esta configuración puede
> diferir entre impresoras, consulta en los foros o en documentación
> las configuraciones específicas para cada.

### CUPS

Significa _Common Unix Printing System_ y es el encargado de gestionar
y proporcionar el servicio de impresión en sistemas Linux.

Lo de siempre, actualizar el sistema:

```bash
apt-get update
apt-get upgrade
```

Instalar el servidor CUPS

```bash
apt-get install cups
```

### Configurar el servidor

El servidor tiene un panel de control para ver las colas de impresión y
las impresoras conectadas, 