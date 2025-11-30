# Requirements Document

## Introduction

Esta especificação define uma plataforma de comunicação ubíqua (API) capaz de rotear mensagens e arquivos entre usuários em múltiplas plataformas externas (WhatsApp, Instagram Direct, Messenger, Telegram) e clientes internos (web/mobile/CLI). O sistema suporta comunicação privada e em grupo, persistência no servidor, controle de estados de mensagem (envio/entrega/leitura), entrega de arquivos até 2 GB e operação em escala com milhões de usuários. A arquitetura é projetada para escalabilidade horizontal e alta disponibilidade.

## Glossary

- **Platform**: O sistema de comunicação ubíqua completo que atua como broker/unificador de mensagens
- **Channel**: Uma plataforma externa de mensageria (ex.: WhatsApp, Instagram Direct, Telegram) ou cliente interno
- **Connector**: Componente adaptador que integra a Platform com um Channel específico
- **Conversation**: Uma sessão de comunicação entre usuários, podendo ser privada (1:1) ou grupo (n membros)
- **Message**: Unidade de comunicação contendo texto ou referência a arquivo
- **Message_ID**: Identificador único universal (UUIDv4) de uma mensagem para garantir idempotência
- **API_Gateway**: Componente de entrada que gerencia autenticação, rate limiting e TLS termination
- **Frontend_Service**: Serviço stateless que expõe endpoints REST/gRPC e valida requisições
- **Message_Broker**: Sistema de streaming durável (ex.: Kafka) para processamento assíncrono de eventos
- **Router_Worker**: Serviço que consome eventos do Message_Broker e executa lógica de roteamento
- **Message_Store**: Banco de dados distribuído para persistência de metadados e texto de mensagens
- **Object_Storage**: Sistema de armazenamento para arquivos grandes (S3-compatible)
- **Presence_Service**: Serviço que rastreia status online/offline dos usuários
- **Username**: Identificador único de usuário na Platform usado para autenticação e roteamento de mensagens

## Requirements

### Requirement 1

**User Story:** Como usuário, quero me registrar e autenticar na plataforma usando um username único, para que o sistema possa identificar-me de forma inequívoca.

#### Acceptance Criteria

1. WHEN um novo usuário solicita registro, THE Platform SHALL validar que o username fornecido é único no sistema
2. THE Platform SHALL rejeitar tentativas de registro com username já existente retornando erro específico
3. WHEN um usuário se autentica, THE Platform SHALL validar as credenciais e retornar um access_token vinculado ao username
4. THE Platform SHALL manter mapeamento persistente entre username e identificadores de Channels externos no metadata store
5. THE Platform SHALL utilizar o username como identificador primário em todas as operações de mensageria (remetente, destinatário)

### Requirement 2

**User Story:** Como operador do sistema, quero que o servidor rastreie quais usuários estão conectados, para que mensagens sejam entregues corretamente apenas aos destinatários autorizados.

#### Acceptance Criteria

1. WHEN um usuário estabelece conexão (websocket ou polling), THE Presence_Service SHALL registrar o username como online com timestamp
2. WHEN um usuário desconecta, THE Presence_Service SHALL atualizar o status do username para offline
3. THE Platform SHALL consultar o Presence_Service antes de rotear mensagens para determinar se entrega em tempo real ou persistência é necessária
4. THE Platform SHALL validar que o destinatário de uma mensagem existe no sistema antes de aceitar a mensagem
5. THE Platform SHALL garantir que mensagens sejam entregues apenas aos usernames especificados como destinatários na Conversation

### Requirement 3

**User Story:** Como usuário da plataforma, quero criar e participar de conversas privadas e em grupo, para que eu possa me comunicar com outros usuários de forma organizada.

#### Acceptance Criteria

1. WHEN um usuário solicita criar uma conversa privada com outro usuário, THE Platform SHALL criar uma Conversation com tipo "private" e exatamente 2 membros
2. WHEN um usuário solicita criar uma conversa em grupo com múltiplos usuários, THE Platform SHALL criar uma Conversation com tipo "group" e n membros onde n é maior ou igual a 2
3. WHEN uma Conversation é criada, THE Platform SHALL atribuir um conversation_id único e persistir os metadados no Message_Store
4. THE Platform SHALL permitir que usuários consultem a lista de Conversations das quais são membros

### Requirement 4

**User Story:** Como usuário, quero enviar mensagens de texto para outros usuários em qualquer plataforma suportada, para que eu possa me comunicar independentemente do canal que o destinatário utiliza.

#### Acceptance Criteria

1. WHEN um usuário envia uma mensagem de texto, THE Platform SHALL aceitar o payload contendo message_id, conversation_id, remetente, destinatários e canais de entrega
2. WHEN uma mensagem é aceita pelo Frontend_Service, THE Platform SHALL retornar status "accepted" com o message_id em menos de 200 milissegundos
3. WHEN o usuário especifica canais de entrega, THE Platform SHALL rotear a mensagem para os Connectors apropriados de cada Channel especificado
4. WHERE o usuário especifica "all" como canal, THE Platform SHALL rotear a mensagem para todos os Channels vinculados ao destinatário
5. THE Platform SHALL garantir que cada Message tenha um message_id único para evitar duplicação

### Requirement 5

**User Story:** Como usuário, quero enviar arquivos de até 2 GB para outros usuários, para que eu possa compartilhar documentos, imagens e vídeos de grande tamanho.

#### Acceptance Criteria

1. WHEN um usuário inicia upload de arquivo, THE Platform SHALL retornar uma upload_url presigned e um file_id único
2. THE Platform SHALL aceitar uploads de arquivos com tamanho até 2 GB usando protocolo chunked e resumable
3. WHEN o upload é completado, THE Platform SHALL validar o checksum fornecido pelo cliente contra o arquivo armazenado
4. THE Platform SHALL armazenar o arquivo no Object_Storage e persistir metadados (file_id, URL, checksum, size) no Message_Store
5. WHEN uma mensagem com arquivo é enviada, THE Platform SHALL incluir a referência file_id no payload da mensagem

### Requirement 6

**User Story:** Como remetente de mensagem, quero receber confirmações de envio, entrega e leitura, para que eu saiba o status da minha comunicação.

#### Acceptance Criteria

1. WHEN o Frontend_Service aceita uma mensagem, THE Platform SHALL atualizar o estado da Message para SENT
2. WHEN a mensagem é entregue ao dispositivo do destinatário, THE Platform SHALL atualizar o estado da Message para DELIVERED
3. WHEN o destinatário lê a mensagem, THE Platform SHALL atualizar o estado da Message para READ
4. THE Platform SHALL permitir que o remetente consulte o histórico de estados de uma Message específica via message_id
5. WHERE o remetente solicita confirmação, THE Platform SHALL enviar callbacks via webhook registrado para cada mudança de estado

### Requirement 7

**User Story:** Como usuário offline, quero receber mensagens enviadas durante minha ausência quando me reconectar, para que eu não perca comunicações importantes.

#### Acceptance Criteria

1. WHEN o Presence_Service indica que um destinatário está offline, THE Platform SHALL persistir a mensagem no Message_Store para entrega posterior
2. WHEN um usuário se reconecta, THE Platform SHALL entregar todas as mensagens pendentes em ordem causal por Conversation
3. THE Platform SHALL manter mensagens persistidas até que sejam entregues com sucesso ao destinatário
4. THE Platform SHALL garantir que mensagens persistidas sejam entregues pelo menos uma vez (at-least-once delivery)

### Requirement 8

**User Story:** Como desenvolvedor de integração, quero utilizar uma API pública bem documentada, para que eu possa integrar meus sistemas com a plataforma de forma eficiente.

#### Acceptance Criteria

1. THE Platform SHALL expor endpoints REST para autenticação, criação de conversas, envio de mensagens, upload de arquivos e consulta de histórico
2. THE Platform SHALL fornecer documentação OpenAPI/Swagger completa de todos os endpoints
3. THE Platform SHALL implementar autenticação via tokens com tempo de expiração configurável
4. THE Platform SHALL suportar versionamento de API (v1, v2) para permitir evolução sem quebrar integrações existentes
5. THE Platform SHALL permitir registro de webhooks para receber callbacks de eventos (delivery, read)

### Requirement 9

**User Story:** Como administrador da plataforma, quero adicionar suporte a novos canais de comunicação sem modificar o núcleo do sistema, para que a plataforma possa evoluir facilmente.

#### Acceptance Criteria

1. THE Platform SHALL definir uma interface padronizada para Connectors contendo métodos connect(), sendMessage(), sendFile() e webhookHandler()
2. WHEN um novo Connector é implementado seguindo a interface, THE Platform SHALL permitir sua integração sem alterações no código do núcleo
3. THE Platform SHALL manter configuração de Connectors disponíveis em metadata store acessível pelos Router_Workers
4. THE Platform SHALL isolar falhas de Connectors individuais para não afetar outros canais ou o sistema como um todo

### Requirement 10

**User Story:** Como operador do sistema, quero que a plataforma suporte milhões de usuários simultâneos, para que o serviço possa crescer conforme a demanda aumenta.

#### Acceptance Criteria

1. THE Platform SHALL processar pelo menos 100.000 mensagens por minuto em configuração de produção
2. WHEN novos nós de Frontend_Service ou Router_Worker são adicionados, THE Platform SHALL distribuir carga automaticamente sem downtime
3. THE Platform SHALL implementar arquitetura stateless nos componentes de API para permitir escalabilidade horizontal ilimitada
4. THE Platform SHALL particionar dados no Message_Broker por conversation_id para preservar ordem causal e permitir paralelização
5. THE Platform SHALL utilizar auto-scaling baseado em métricas de carga (CPU, throughput de mensagens)

### Requirement 11

**User Story:** Como usuário do serviço, quero que a plataforma esteja disponível 99.95% do tempo, para que eu possa confiar na comunicação crítica.

#### Acceptance Criteria

1. THE Platform SHALL implementar failover automático para componentes críticos (API_Gateway, Message_Broker, Message_Store)
2. WHEN um nó de Router_Worker falha, THE Platform SHALL detectar via heartbeat e redistribuir partições para nós saudáveis em menos de 30 segundos
3. THE Platform SHALL replicar dados do Message_Store em pelo menos 3 réplicas para tolerância a falhas
4. THE Platform SHALL manter disponibilidade de leitura mesmo durante falhas parciais de escrita usando consistência eventual
5. THE Platform SHALL implementar circuit breakers em Connectors para isolar falhas de canais externos

### Requirement 12

**User Story:** Como desenvolvedor do sistema, quero garantias de entrega e ordem de mensagens, para que a comunicação seja confiável e compreensível.

#### Acceptance Criteria

1. THE Platform SHALL garantir entrega at-least-once de mensagens com deduplicação baseada em message_id
2. THE Platform SHALL preservar ordem causal de mensagens dentro de uma Conversation usando sequence numbers
3. WHEN uma mensagem duplicada é detectada via message_id, THE Platform SHALL descartar a duplicata e retornar o status da mensagem original
4. THE Platform SHALL permitir replay de eventos do Message_Broker para recuperação de falhas
5. THE Platform SHALL implementar writes idempotentes no Message_Store para suportar retries seguros

### Requirement 13

**User Story:** Como operador do sistema, quero monitorar a saúde e performance da plataforma em tempo real, para que eu possa identificar e resolver problemas rapidamente.

#### Acceptance Criteria

1. THE Platform SHALL expor métricas Prometheus incluindo mensagens por segundo, latência de entrega, taxa de erro de Connectors e utilização de recursos
2. THE Platform SHALL implementar tracing distribuído usando OpenTelemetry com spans em todos os componentes críticos
3. THE Platform SHALL gerar logs estruturados com correlação via trace_id e message_id
4. THE Platform SHALL fornecer dashboards Grafana pré-configurados mostrando métricas chave e alertas
5. WHEN métricas críticas excedem thresholds, THE Platform SHALL disparar alertas via Alertmanager

### Requirement 14

**User Story:** Como usuário da plataforma, quero que minhas mensagens sejam roteadas entre diferentes plataformas externas, para que eu possa alcançar destinatários independentemente do canal que eles usam.

#### Acceptance Criteria

1. THE Platform SHALL manter mapeamento de usuários para múltiplos Channels no metadata store
2. WHEN um usuário envia mensagem especificando múltiplos canais, THE Platform SHALL invocar os Connectors correspondentes em paralelo
3. WHERE um usuário está vinculado a WhatsApp e Instagram, THE Platform SHALL permitir que mensagens enviadas via WhatsApp sejam entregues no Instagram Direct do destinatário
4. THE Platform SHALL implementar retries exponenciais com backoff para falhas temporárias de Connectors
5. WHEN um Connector falha permanentemente, THE Platform SHALL registrar a falha e continuar tentando entregar via outros canais disponíveis
