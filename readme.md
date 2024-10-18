# Rate Limiter Project

## Entrega

### Código-fonte Completo
O código-fonte completo da implementação está disponível neste repositório.

## Documentação

### Como o Rate Limiter Funciona
O rate limiter é um mecanismo que controla o número de solicitações que um cliente pode fazer a um servidor em um determinado período de tempo. Ele utiliza o Redis como armazenamento para manter o estado das solicitações e aplicar as regras de limitação.

### Configuração do Rate Limiter
O rate limiter pode ser configurado através do arquivo de configuração `.env`. As principais configurações incluem:
- **WEBSERVER_HOST**: Host do servidor web.
- **WEBSERVER_PORT**: Porta do servidor web.
- **REDIS_HOST**: Host do Redis.
- **REDIS_PORT**: Porta do Redis.
- **RATE_LIMIT_BY_IP**: Limite de solicitações por IP.
- **RATE_LIMIT_BY_TOKEN**: Limite de solicitações por token.
- **RATE_LIMIT\WINDOW**: Janela de tempo em segundos para a aplicação do limite.
### Testes Automatizados
Os testes automatizados estão incluídos no projeto para demonstrar a eficácia e a robustez do rate limiter. Eles podem ser executados utilizando a ferramenta de testes padrão do Go.

### Docker/Docker-Compose
Para facilitar a execução e os testes da aplicação, utilizamos Docker e Docker-Compose. O arquivo `docker-compose.yml` está configurado para iniciar o serviço Redis e a aplicação web.

### Servidor Web
O servidor web está configurado para responder na porta 8080.

## Como Executar

### Pré-requisitos
- Docker
- Docker-Compose

### Passos para Executar
1. Clone o repositório:
    ```bash
    git clone <URL_DO_REPOSITORIO>
    cd <NOME_DO_REPOSITORIO>
    ```
2. Inicie os serviços com Docker-Compose:
    ```bash
    docker-compose up -d
    ```
3. Acesse a aplicação na porta 8080:
    ```bash
    http://localhost:8080
    ```

### Executando Testes
Para executar os testes automatizados, utilize o comando:
    ```bash
    go test ./...
    ```
```bash
go test ./...
```

Estrutura do Projeto
main.go: Ponto de entrada da aplicação.
config/: Arquivos de configuração.
controller/: Controladores da aplicação.
internal/infra/: Implementações de infraestrutura, incluindo o RedisStore.
middlewares/: Middlewares da aplicação.
services/: Serviços da aplicação, incluindo o RateLimitService.
docker-compose.yml: Configuração do Docker-Compose.