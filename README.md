# Challenge de Programación: APIs de Tienda Online

## Descripción

Este repositorio contiene el código de dos APIs para una tienda online de remeras, desarrolladas en Go. Las APIs gestionan el inventario de remeras y el ciclo de vida de las ventas. A continuación se describen sus funcionalidades principales:

### **Shirts API**

La **Shirts API** se encarga de gestionar el inventario de remeras en la tienda. Permite realizar operaciones sobre las remeras disponibles, como:

- **Crear una nueva remera**: Registra una nueva remera con su marca, descripción, precio, unidades disponibles y estado.
- **Actualizar una remera existente**: Modifica los detalles de una remera, como el precio, unidades disponibles o descripción.
- **Obtener una remera por ID**: Recupera la información de una remera específica utilizando su ID.
- **Eliminar una remera**: Marca una remera como eliminada sin borrarla físicamente del sistema.

### **Sales API**

La **Sales API** gestiona las ventas dentro de la tienda, procesando la información de los clientes, los productos comprados y las transacciones. Sus funcionalidades incluyen:

- **Registrar una nueva venta**: Crea un registro de venta con los productos comprados, la cantidad, el precio total y la información del cliente.
- **Procesar una devolución**: Permite devolver una venta, recuperando el stock de las remeras involucradas en la devolución.
- **Obtener detalles de una venta**: Consulta la información de una venta registrada, incluyendo el carrito de productos, el monto total y el estado de la transacción.

---

Este desafío consiste en corregir los errores existentes y mejorar la funcionalidad de ambas APIs para asegurar que el sistema sea robusto, escalable y fácil de usar.

##  Getting Started

Este proyecto está desarrollado en la versión 1.22.3 de Go. Asegúrate de tener esta versión instalada en tu máquina. Puedes comprobar tu versión de Go con el siguiente comando:

```bash
  go version
```

### Pasos para levantar la aplicación

1. **Clonar el repositorio**

   Clona el repositorio a tu máquina local:

   ```bash
   git clone https://github.com/amariniclip/go-challenge.git
   cd go-challenge
   ```

2. **Inicializar los módulos de Go**

   Ejecuta `go mod tidy` para asegurarte de que todos los módulos y dependencias estén correctamente instalados:

   ```bash
   go mod tidy
   ```

3. **Levantar la aplicación**

   Para correr la aplicación, usa el siguiente comando:

   ```bash
   go run cmd/api/main.go
   ```

### Pasos para correr los tests

Para ejecutar los tests de la aplicación y asegurarte de que todo funciona correctamente, puedes usar el siguiente comando:

```bash
    go test -v ./... -race -count=1
```

- v: Muestra la salida detallada de los tests.
- race: Habilita la detección de condiciones de carrera durante las pruebas.
- count=1: Ejecuta los tests solo una vez (por defecto, Go ejecuta los tests varias veces para mejorar la fiabilidad).
