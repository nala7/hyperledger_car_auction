# Subasta de Autos

**Miembros del Equipo (Grupo 512)**
* Nadia González Fernández

* Ariel Antonio Huerta Martín

* Amalia Nilda Ibarra Rodríguez

* José Alejandro Labourdette-Lartigue Soto

* Luis Alejandro Lara Rojas

* Gabriela Bárbara Martínez Giraldo


## Ejercicios

1.	**Identifique y liste cuales son los requerimientos funcionales del negocio.**

2.	**Realice el diseño del chaincode a partir de los requerimientos identificados.**

3.	**Exponga la lógica de Negocio a través de una API Rest documentada con el estándar OpenAPI. Puede realizar esto a través de cualquier método, no necesariamente debe utilizar Iris-Go.**

## Requerimientos funcionales

El proyecto presenta una subasta de autos donde interactúan dos participantes: el comprador y el subastador. El primero realiza ofertas de compra en una subasta, mientras que el segundo es el encargado de registrar el auto a subastar, así como crear, cerrar y verificar subastas.
Se realizará una subasta en sobre cerrado, donde se adjudica el bien subastado a la mejor oferta sin posibilidad de mejorarla, es decir, se ofrecerá un auto al mejor postor y se realiza la compra al monto más alto, sin posibilidad de mejorar la oferta.
Las funciones del producto requeridas se pueden generalizar en:
- Crear una subasta.
- Registrar un comprador en una subasta.
- Iniciar una subasta.
- Pujar por el carro subastado.
- Cerrar subasta.
- Verificar el ganador de la subasta.

## Diseño del chaincode
![HyperledgerSchema drawio](https://user-images.githubusercontent.com/62756227/178613896-c9c9aee0-c27a-4c53-9c5b-404c4728d7bd.svg)
