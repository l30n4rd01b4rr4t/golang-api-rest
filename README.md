# Resumen

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

## Primeros pasos en Go

Este es un ejercicio que busca crear una simple api rest en Go llamando a un par de endpoints y mostrar como se hacen llamadas encadenadas usando `<- chan`.

### ¿Cómo ejecuto la aplicación?

Para correr la aplicación, basta con ejecutar el siguiente comando en la raíz del proyecto:

```
$ go run main.go
```

La salida de la consola debería ser algo como:

```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/handler.PingHandler (3 handlers)
[GIN-debug] GET    /match/sites/:site_id/items/:item_id/step/init --> github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/handler.FormInit (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```
