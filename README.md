# Go API Template

Template base para proyectos API REST con Go.

## Estructura del Proyecto
```
.
├── api/
│   ├── handlers/      # HTTP handlers
│   ├── helpers/       # Helper functions
│   └── middlewares/   # Middlewares (auth, cors, etc.)
├── auth/              # Autenticación y autorización
├── database/          # Modelos y migraciones
├── types/             # Tipos y estructuras
├── utils/             # Utilidades generales
├── uploads/           # Archivos subidos
└── main.go           # Punto de entrada
```

## Requisitos

- Go 1.21+
- PostgreSQL (opcional)
- Redis (opcional)

## Instalación

1. **Clonar o usar como template**
```bash
   # Si usas GitHub template
   # Click en "Use this template"
   
   # O clonar directamente
   git clone https://github.com/user/go-api-template mi-proyecto
   cd mi-proyecto
```

2. **Actualizar el módulo** (opcional)
```bash
   # Si quieres cambiar el nombre del módulo
   go mod edit -module mi-proyecto
   # O reiniciar go.mod
   rm go.mod go.sum
   go mod init mi-proyecto
```

3. **Instalar dependencias**
```bash
   make deps
```

4. **Configurar variables de entorno**
```bash
   cp .env.example .env
   # Editar .env con tus valores
```

5. **Ejecutar**
```bash
   make run
```

## Comandos Útiles (Makefile)
```bash
make deps         # Instalar/actualizar dependencias
make run          # Ejecutar el servidor
make build        # Compilar binario
make test         # Ejecutar tests
make clean        # Limpiar archivos generados
make dev          # Ejecutar con hot reload (requiere air)
```

## Dependencias Comunes

Agrega según necesites:
```bash
# Router
go get github.com/gin-gonic/gin

# Environment variables
go get github.com/joho/godotenv

# Database
go get gorm.io/gorm
go get gorm.io/driver/postgres

# Redis
go get github.com/redis/go-redis/v9

# JWT
go get github.com/golang-jwt/jwt/v5

# Después de agregar dependencias:
make deps
```

## Inicio Rápido
```bash
# 1. Crear proyecto desde template
git clone https://github.com/user/go-api-template mi-api
cd mi-api

# 2. Configurar
cp .env.example .env
make deps

# 3. Ejecutar
make run
```

## Uso

1. Actualiza el nombre del módulo en `go.mod` (si es necesario)
2. Implementa tus handlers en `api/handlers/`
3. Define tus modelos en `database/`
4. Configura tus rutas
5. ¡Comienza a desarrollar!

## Licencia

MIT