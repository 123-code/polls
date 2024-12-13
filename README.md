# README

## Descripción del Proyecto

Este proyecto es una aplicación backend desarrollada en Go, diseñada para gestionar un sistema de votación y obras públicas. La aplicación permite a los usuarios interactuar con un sistema de votación, registrar candidatos, y analizar el impacto de las obras en las provincias. Utiliza una base de datos PostgreSQL para almacenar la información y se basa en el framework Gin para manejar las solicitudes HTTP.

## Funcionalidades Principales

1. **Gestión de Usuarios**: 
   - Registro de usuarios mediante la creación de un ID único (cédula).
   - Generación de tokens JWT para la autenticación de usuarios.

2. **Sistema de Votación**:
   - Registro de votos por parte de los usuarios.
   - Análisis de los votos por provincia y candidato.
   - Determinación del candidato ganador basado en los votos recibidos.

3. **Gestión de Obras**:
   - Registro de obras públicas asociadas a candidatos.
   - Análisis del impacto de las obras en los votos recibidos por los candidatos.
   - Cálculo de ratios entre obras y votos para evaluar la efectividad de los candidatos.

4. **Interacción con Contratos Inteligentes**:
   - Integración con contratos inteligentes en la blockchain de Ethereum para la creación de billeteras digitales y la gestión de NFTs.

## Estructura del Proyecto

El proyecto está organizado en varias carpetas que contienen diferentes componentes:

- **controllers/**: Contiene los controladores que manejan la lógica de negocio y las interacciones con las solicitudes HTTP.
- **models/**: Define las estructuras de datos que representan las entidades en la base de datos, como `User`, `Obra`, `Cedula`, etc.
- **initializers/**: Contiene la configuración inicial de la base de datos y la sincronización de modelos.
- **util/**: Funciones utilitarias que proporcionan funcionalidades adicionales, como la generación de tokens JWT y la interacción con contratos inteligentes.
- **validators/**: Contiene funciones para validar datos, como la validación de cédulas.

## Instalación

Para instalar y ejecutar el proyecto, sigue estos pasos:

1. **Clonar el Repositorio**:
   ```bash
   git clone https://github.com/tu_usuario/pollsbackend.git
   cd pollsbackend
   ```

2. **Instalar Dependencias**:
   Asegúrate de tener Go instalado y ejecuta:
   ```bash
   go mod tidy
   ```

3. **Configurar la Base de Datos**:
   Modifica la cadena de conexión en `initializers/connectdb.go` para que apunte a tu base de datos PostgreSQL.

4. **Ejecutar la Aplicación**:
   ```bash
   go run main.go
   ```

5. **Acceder a la API**:
   La API estará disponible en `http://localhost:8080`.

## Endpoints de la API

- **/users**: 
  - `POST`: Crear un nuevo usuario.
  - `GET`: Obtener todos los usuarios.
  - `GET /users/:id`: Obtener un usuario específico.
  - `PUT /users/:id`: Actualizar un usuario específico.
  - `DELETE /users/:id`: Eliminar un usuario específico.

- **/vote**: 
  - `POST /vote/:id`: Registrar un voto para un candidato específico.

- **/obras**: 
  - `POST`: Crear una nueva obra.
  - `GET`: Obtener todas las obras.
  - `GET /obras_analysis`: Analizar obras por provincia.

- **/validateid**: 
  - `POST`: Validar la cédula de un usuario.




