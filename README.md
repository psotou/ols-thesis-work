# Generación modelo a través de método propuesto 

Para generar la relación funcional propuesta en la [memoria](https://github.com/psotou/memoria/blob/master/thesis_document/thesis.pdf) descargar el binario y correr lo siguiente:

```bash
./ols -file <path/to/file.csv> -rmax <max_rating>
```

Donde,

+ `<path/to/file.csv>` es la ruta a un archivo csv que tiene la siguientes estructura:

  ```
  crecimiento_costos,madurez_bim
  0.330,1.0
  0.249,1.7
  0.162,1.8
  0.118,2.2
  0.000,4.0
  ```

+ `<max_rating>` es el rating máximo de madurez BIM según la escala que se esté utilizando.