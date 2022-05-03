# H3 Golang POC
### Sobre que se trata esta POC?
La misma se trata sobre la indexación automatica de algunas tiendas generadas random y luego la busqueda mediante una coordenada(Que simula ser el punto donde esta el usuario). Esto genera un indice en H3 el cual seria el hexagono en donde se encuentra el mismo, luego trae los indices de los anillos vecinos y se busca mediante esos indices los IDs de las tiendas en la base de datos.

#### Referencias
- [Notebook de ejemplo](https://observablehq.com/@nrabinowitz/h3-indexing-order)
- [Binding H3 en Golang](https://github.com/uber/h3-go)
- [Documentación oficial de Uber](https://h3geo.org/)
- [Blog de ayuda](https://betterprogramming.pub/playing-with-ubers-hexagonal-hierarchical-spatial-index-h3-ed8d5cd7739d)
