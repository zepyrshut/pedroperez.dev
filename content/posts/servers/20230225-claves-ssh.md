---
title: Conectarse mediante claves SSH desde Windows a Linux
date: 2023-02-25
author: Pedro Pérez
tags: ["ssh", "conectarse", "linux", "windows", "putty", "openssh"]
categories: ["Servidores"]
description: "Conectarse a un servidor mediante claves SSH y olvidarse
de las contraseñas, especialmente desde sistemas Windows a sistemas
Linux."
draft: false
---

### Preámbulo

Yendo un poco al grano, en lugar de explicar con detalle las
características del par de claves, vamos a ver cómo se generan y se hace
uso de ellas. Las conexiones se harán desde un equipo Windows 10 a un
servidor Debian 11, y algún que otro matiz.

Lo que está claro y se sabe es que el uso de claves SSH es más seguro
que el uso de contraseñas.

### Claves SSH

Como se hace una conexión desde un equipo Windows a un servidor Linux,
está de la forma que se puede hacer mediante _PuTTY_ o _OpenSSH_. En
ambos casos, se debe tener acceso al servidor, sea por contraseña o por
claves, para poder importar la clave pública.

Es importante conocer las diferencias entre clave pública y clave
privada y el uso del fichero _authorized_keys_.

**Clave pública:** es la clave que se guarda en el servidor y se
comparte con el resto de los clientes que se quiera hacer uso del 
servidor.

**Clave privada:** es la clave que se guarda en el equipo cliente y que 
se utiliza para conectarse al servidor. Debe guardarse en un sitio
seguro.

**_autorized_keys_:** es el fichero que se encuentra en el directorio
_.ssh_ del servidor y que contiene las claves públicas de los clientes
que se quieran conectar al servidor.

**Algoritmmo de cifrado:** es el algoritmo que se utiliza para cifrar
la comunicación entre el cliente y el servidor. Se hará uso del
algoritmo [_EdDSA_](https://es.wikipedia.org/wiki/EdDSA).

### Generando claves SSH mediante PuTTYgen

En la herramienta _PuTTYgen_, se debe seleccionar el algoritmo _EdDSA_,
y le damos al botón _Generate_. Una vez generada la clave, se muestra lo
siguiente:

![PuTTY Key Generator](/image/20230225-puttygen.png)

El texto que está seleccionado es la clave pública, se puede establecer
una contraseña para la clave privada, eso es opcional, y la clave 
privada se obtiene pulsando el botón _Save private key_. Se recomienda
guardarlo en un sitio seguro y no compartirlo con nadie.

[Para guardar la clave pública en el servidor, ir a la sección de guardar la clave pública](#guardando-la-clave-pública-en-el-servidor).

Ahora falta la clave privada, para ello, con la clave privada que se
guardó antes, abrimos PuTTY, seleccionamos la sesión guardada o se crea
una.

![PuTTY selección de sesión](/image/20230225-puttygen_2.png)

Pues en el panel derecho, nos vamos hacia la siguiente ruta:
_Connection/SHH/Auth/Credentials_, y en el panel derecho, buscamos la
clave privada

![Buscando la clave privada](/image/20230225-puttygen_3.png)

Recuerda guardar la sesión antes de dar al botón _Open_, si no, tendrás
que volver a buscar la clave privada.

Abrimos la conexión, introducimos el nombre de usuario, y _et-voilà_.

![Conexión PuTTY](/image/20230225-puttygen_4.png)

### Mediante OpenSSH

El procedimiento es similar al uso de PuTTY, pero en este caso la
generación de clave se hará por consola de Windows.

En la propia terminal:

> Si se generó una clave anteriormente y se hace uso de ella, generar
> una nueva podría ocasionar la pérdida de acceso al servidor. Tenga
> precaución.

```bash
$ ssh-keygen -t ed25519 -C "correo@mail.com"
```

La bandera `-C` significa comentario, normalmente se pone un correo
electrónico asociada a la persona que hace uso de la clave privada.

Una vez generada las claves, se guardan en el directorio _.ssh_ en la
carpeta del usuario. Al ser _ed25519_, los ficheros se llaman 
`id_ed25519`, que es la clave privada e `id_ed25519.pub` que es la clave
pública.

Se recomienda cambiarles de nombre para evitar el reemplazo de estas
claves al generar otras nuevas.

El contenido del fichero `id_ed25519.pub` debe ser copiado al fichero de
_authorized_keys_.

[Para guardar la clave pública en el servidor, ir a la sección de guardar la clave pública](#guardando-la-clave-pública-en-el-servidor).

Mientas que la clave privada se deja donde está, y se procede a editar
el fichero _config_, que se encuentra ubicado en la carpeta _.ssh_ de la
carpeta del usuario. En caso de que no exista dicho fichero, crear uno
con el mismo nombre y **sin extensión**.

```text
Host 192.168.1.57
  HostName 192.168.1.57
  User root
  IdentityFile C:\Users\model\.ssh\vm-debian
```

En el fichero _config_ se debe añadir lo siguiente para hacer uso de la
clave privada, donde `Host` y `HostName` es la dirección IP del
servidor, `User` es el usuario con el que desea entrar e `IdentityFile`
es el fichero donde está la clave privada.



### Guardando la clave pública en el servidor

La clave pública se debe guardarse en el servidor, para ello nos
conectamos al servidor y localizamos el archivo _authorized_keys_ dentro
del directorio _.ssh_.

```bash
$ cd ~/.ssh
$ touch authorized_keys
```

Si no existe el directorio _.ssh_, se puede crear.

```bash
$ cd ~
$ mkdir ~/.ssh && cd .ssh
$ touch authorized_keys
```

Una vez creado el fichero, se añade la clave pública que hemos generado
anteriormente al fichero _authorized_keys_.

```bash
$ echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIO0EV2+vc2DQABB9T5loDEfqbdsv9DKdGgmBn5YiPqQQ eddsa-key-20230225" >> authorized_keys
```

> _Es importante reseñar que el proceso de añadir la clave pública en el
> fichero _authorized_keys_ se está haciendo sobre el directorio de 
> usaurio _root_. Si en lugar de iniciar sesión con _root_, lo prefiere 
> con un usuario estándar se deberá hacer el mismo procedimiento pero 
> en el propio directorio del usuario._

### Eliminar el acceso por contraseña

Una vez que ya se pueda iniciar sesión mediante claves SSH, se puede
quitar la posibilidad de acceso por contraseña. Para ello basta con 
editar el fichero `sshd_config`.

```bash
$ nano /etc/ssh/sshd_config
```

Buscar la línea que pone: `PasswordAuthentication`, no debe estar
comentada y cambiar su valor a `no`. Guardar y reiniciar con:
`service sshd restart`.

### Bonus: Conexión por WinSCP mediante claves SSH

Con la clave privada que se generó en PuTTY, añadir en ajustes avanzados
en la parte de nueva conexión.

![Conexión WinSCP](/image/20230225-puttygen_5.png)

Eso es todo.


### Datos técnicos

- **Windows:** Microsoft Windows [Versión 10.0.19044.2604]
- **Debian:**  Debian GNU/Linux 11.5.0 (bullseye)
- **PuTTY:** Release 0.78
- **OpenSSH en Windows**: OpenSSH_for_Windows_8.1p1, LibreSSL 3.0.2
- **OpenSSH en Debian**: OpenSSH_8.4p1 Debian-5+deb11u1, OpenSSL 1.1.1n  15 Mar 2022