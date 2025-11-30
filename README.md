# Plataforma de Comunicação Ubíqua - API Distribuída

Experimento de desenvolvimento orientado a especificação com Kiro para a disciplina de Sistemas Distribuídos.

## Visão Geral

Esta é uma plataforma de comunicação distribuída que atua como broker/unificador de mensagens entre múltiplas plataformas externas (WhatsApp, Instagram Direct, Telegram, Messenger) e clientes internos. O sistema suporta comunicação privada e em grupo, persistência no servidor, controle de estados de mensagem, e entrega de arquivos até 2 GB.

## Arquitetura

O sistema segue uma arquitetura de microserviços com os seguintes componentes:

- **Frontend Service**: API REST para clientes
- **Auth Service**: Autenticação e gerenciamento de usuários
- **File Upload Service**: Upload e gerenciamento de arquivos
- **Router Worker**: Roteamento de mensagens entre canais
- **Presence Service**: Rastreamento de status online/offline

### Infraestrutura

- **Kafka**: Message broker para processamento assíncrono
- **MongoDB**: Armazenamento de mensagens e metadados
- **Redis**: Cache e rastreamento de presença
- **MinIO**: Armazenamento de objetos (arquivos)
- **etcd**: Armazenamento de metadados e configuração
- **Prometheus/Grafana**: Métricas e monitoramento
- **Jaeger**: Tracing distribuído

## Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose
- Make (opcional, mas recomendado)

## Instalação

1. Clone o repositório:
```bash
git clone <repository-url>
cd ds-sdd-with-kiro
```

2. Copie o arquivo de configuração de exemplo:
```bash
cp .env.example .env
```

3. Ajuste as variáveis de ambiente no arquivo `.env` conforme necessário.

## Desenvolvimento Local

### Iniciar Infraestrutura

Para iniciar todos os serviços de infraestrutura (Kafka, MongoDB, Redis, MinIO, etcd, Prometheus, Grafana, Jaeger):

```bash
make setup
```

Este comando irá:
- Iniciar todos os containers Docker
- Criar os tópicos Kafka necessários
- Criar o bucket MinIO

### Comandos Disponíveis

```bash
# Ver todos os comandos disponíveis
make help

# Construir todos os serviços
make build

# Construir um serviço específico
make build-frontend

# Executar testes
make test

# Executar testes com relatório de cobertura
make test-coverage

# Executar um serviço específico
make run SERVICE=frontend

# Iniciar infraestrutura Docker
make docker-up

# Parar infraestrutura Docker
make docker-down

# Limpar volumes Docker
make docker-clean

# Ver logs Docker
make docker-logs

# Formatar código
make fmt

# Executar linter
make lint

# Modo desenvolvimento (infraestrutura + serviço)
make dev SERVICE=frontend
```

### Estrutura do Projeto

```
.
├── cmd/                    # Pontos de entrada dos serviços
│   ├── frontend/
│   ├── auth/
│   ├── file-service/
│   ├── router-worker/
│   └── presence/
├── internal/               # Código interno
│   ├── models/            # Modelos de dados
│   └── shared/            # Código compartilhado
│       ├── config/        # Configuração
│       ├── errors/        # Tratamento de erros
│       └── utils/         # Utilitários
├── config/                # Arquivos de configuração
│   ├── prometheus.yml
│   └── grafana/
├── scripts/               # Scripts auxiliares
├── docker-compose.yml     # Configuração Docker
├── Makefile              # Comandos de build e execução
└── .env.example          # Exemplo de variáveis de ambiente
```

## Endpoints da API

### Autenticação
- `POST /v1/auth/register` - Registrar novo usuário
- `POST /v1/auth/token` - Obter token de acesso

### Conversas
- `POST /v1/conversations` - Criar nova conversa
- `GET /v1/conversations` - Listar conversas do usuário
- `GET /v1/conversations/{id}/messages` - Obter mensagens de uma conversa

### Mensagens
- `POST /v1/messages` - Enviar mensagem
- `GET /v1/messages/{id}/status` - Consultar status de mensagem

### Arquivos
- `POST /v1/files/initiate` - Iniciar upload de arquivo
- `POST /v1/files/complete` - Completar upload de arquivo
- `GET /v1/files/{id}` - Obter metadados de arquivo

### Webhooks
- `POST /v1/webhooks` - Registrar webhook
- `DELETE /v1/webhooks/{id}` - Remover webhook

## Monitoramento

- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin)
- **Jaeger UI**: http://localhost:16686
- **MinIO Console**: http://localhost:9001 (minioadmin/minioadmin)

## Testes

```bash
# Executar todos os testes
make test

# Executar testes com cobertura
make test-coverage

# Executar testes de um pacote específico
go test -v ./internal/models/...
```

## Especificações

Este projeto segue desenvolvimento orientado a especificação. As especificações completas estão em:

- `.kiro/specs/distributed-chat-api/requirements.md` - Requisitos
- `.kiro/specs/distributed-chat-api/design.md` - Design
- `.kiro/specs/distributed-chat-api/tasks.md` - Plano de implementação

## Contribuindo

1. Revise as especificações em `.kiro/specs/`
2. Implemente as tarefas conforme o plano em `tasks.md`
3. Escreva testes para novas funcionalidades
4. Execute `make fmt` e `make lint` antes de commitar
5. Garanta que todos os testes passem com `make test`

## Licença

Este é um projeto acadêmico para a disciplina de Sistemas Distribuídos.
