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
