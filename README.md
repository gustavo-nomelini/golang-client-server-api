# Golang Client-Server API - Cotação do Dólar

Este projeto implementa um sistema cliente-servidor para obter, armazenar e exibir a cotação atual do dólar em relação ao real.

## Estrutura do Projeto

O projeto consiste em dois componentes principais:

1. **server.go**: Servidor HTTP que consome a API de cotações e armazena os dados no SQLite
2. **client.go**: Cliente que consulta o servidor e salva a cotação em um arquivo de texto

## Funcionalidades

### Server

- Expõe um endpoint HTTP `/cotacao` na porta 8080
- Consulta a API externa (economia.awesomeapi.com.br) para obter a cotação atual USD-BRL
- Implementa timeout de 200ms para a requisição à API externa
- Armazena a cotação em um banco de dados SQLite com timeout de 10ms
- Retorna o valor da cotação (campo "bid") para o cliente em formato JSON

### Client

- Realiza uma requisição HTTP ao servidor com timeout de 300ms
- Processa a resposta JSON para extrair o valor da cotação
- Salva a cotação em um arquivo texto "cotacao.txt" no formato "Dólar: {valor}"

## Tecnologias Utilizadas

- Go (Golang) - Linguagem de programação
- SQLite - Banco de dados
- Pacote context - Para gerenciamento de timeouts
- Pacote net/http - Para comunicação HTTP
- Pacote database/sql - Para interação com o banco de dados
- github.com/mattn/go-sqlite3 - Driver SQLite para Go

## Como Executar

1. Certifique-se de ter o Go instalado em sua máquina
2. Clone este repositório
3. Instale as dependências:
   ```
   go mod tidy
   ```
4. Em um terminal, inicie o servidor:
   ```
   go run server.go
   ```
5. Em outro terminal, execute o cliente:
   ```
   go run client.go
   ```

## Tratamento de Erros e Timeouts

O projeto implementa três níveis de timeout, conforme especificado:

1. 200ms para a chamada à API externa no servidor
2. 10ms para a persistência dos dados no banco SQLite
3. 300ms para o cliente receber a resposta do servidor

Se qualquer desses timeouts for excedido, o sistema registrará o erro nos logs.

## Estrutura do Banco de Dados

O servidor cria uma tabela "cotacoes" com a seguinte estrutura:

- id: INTEGER PRIMARY KEY
- bid: TEXT (valor da cotação)
- created_at: DATETIME (momento do registro)

## Arquivo de Saída

O cliente gera um arquivo "cotacao.txt" com o formato:

```
Dólar: {valor}
```
