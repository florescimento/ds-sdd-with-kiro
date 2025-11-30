# Quick Start Guide

Este guia rápido ajudará você a configurar e executar o projeto em minutos.

## Pré-requisitos

Certifique-se de ter instalado:
- Go 1.21+ ([Download](https://go.dev/dl/))
- Docker Desktop ([Download](https://www.docker.com/products/docker-desktop))
- Make (geralmente já vem instalado no macOS)

## Passos Rápidos

### 1. Instalar Go (se necessário)

```bash
# macOS
brew install go

# Verificar instalação
go version
```

### 2. Configurar o Projeto

```bash
# Copiar configuração de exemplo
cp .env.example .env

# Baixar dependências Go
go mod download
```

### 3. Iniciar Infraestrutura

```bash
# Iniciar todos os serviços (Kafka, MongoDB, Redis, MinIO, etc.)
make setup
```

⏱️ Aguarde cerca de 30 segundos para todos os serviços iniciarem.

### 4. Construir os Serviços

```bash
# Construir todos os serviços
make build
```

### 5. Executar um Serviço

```bash
# Executar o serviço frontend
make run SERVICE=frontend
```

## Verificar Instalação

### Verificar Infraestrutura

```bash
# Ver status de todos os containers
docker-compose ps
```

Todos devem mostrar status "Up" ou "Up (healthy)".

### Acessar Interfaces Web

- **MinIO Console**: http://localhost:9001 (minioadmin/minioadmin)
- **Grafana**: http://localhost:3000 (admin/admin)
- **Prometheus**: http://localhost:9090
- **Jaeger**: http://localhost:16686

## Comandos Úteis

```bash
# Ver todos os comandos disponíveis
make help

# Executar testes
make test

# Ver logs dos containers
make docker-logs

# Parar infraestrutura
make docker-down

# Limpar tudo (incluindo volumes)
make docker-clean
```

## Próximos Passos

1. 📖 Leia o [README.md](README.md) completo
2. 🔧 Consulte o [Setup Guide](docs/SETUP.md) detalhado
3. 📋 Revise as especificações em `.kiro/specs/distributed-chat-api/`
4. 💻 Comece a implementar as tarefas em `tasks.md`

## Problemas Comuns

### Go não encontrado
```bash
# Instalar Go
brew install go
```

### Portas em uso
```bash
# Verificar porta específica
lsof -i :9092

# Matar processo se necessário
kill -9 <PID>
```

### Docker não está rodando
```bash
# Iniciar Docker Desktop manualmente
open -a Docker
```

### Limpar e recomeçar
```bash
# Limpar tudo
make docker-clean

# Recomeçar do zero
make setup
```

## Suporte

Para mais detalhes, consulte:
- [docs/SETUP.md](docs/SETUP.md) - Guia de configuração completo
- [README.md](README.md) - Documentação do projeto
- `.kiro/specs/distributed-chat-api/design.md` - Arquitetura do sistema
