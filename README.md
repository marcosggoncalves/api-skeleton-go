 
# Setup do Ambiente para Desenvolvimento com Go e Air

Este guia detalha como configurar o ambiente para desenvolvimento utilizando [Go](https://golang.org/) e [Air](https://github.com/cosmtrek/air), uma ferramenta de live reload para aplicações Go.

 

## Pré-requisitos

- Sistema operacional: Windows, macOS ou Linux
- Terminal ou prompt de comando com permissões administrativas

 

## Instalação do Go

1. Acesse a página oficial de downloads do Go: [https://golang.org/dl/](https://golang.org/dl/).
2. Baixe o instalador correspondente ao seu sistema operacional.
3. Siga as instruções do instalador para completar a instalação.
4. Verifique se o Go foi instalado corretamente:
   ```bash
   go version
   ```
 
## Instalação do Air

1. Instale o Air usando o comando `go install`:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```
2. Verifique se o Air foi instalado corretamente:
   ```bash
   air -v
   ```
   O comando acima deve retornar a versão instalada do Air.

## Executando o Air no Projeto

1. Certifique-se de estar no diretório raiz do projeto.
2. Navegue até a pasta `cmd/server` (ou o diretório principal do servidor Go no projeto):
   ```bash
   cd cmd/server
   ```
3. Inicie o Air:
   ```bash
   air
   ```

O servidor será iniciado com suporte a live reload. Qualquer alteração no código-fonte reiniciará o servidor automaticamente.

---

## Recursos

- [Documentação do Go](https://golang.org/doc/)
- [Repositório do Air](https://github.com/cosmtrek/air)
```
 