```markdown
# Setup do Ambiente para Desenvolvimento com Go e Air

Este guia detalha como configurar o ambiente para desenvolvimento utilizando [Go](https://golang.org/) e [Air](https://github.com/cosmtrek/air), uma ferramenta de live reload para aplicações Go.

---

## Pré-requisitos

- Sistema operacional: Windows, macOS ou Linux
- Terminal ou prompt de comando com acesso administrativo

---

## Instalação do Go

1. Acesse a página oficial de downloads do Go: [https://golang.org/dl/](https://golang.org/dl/).
2. Baixe o instalador adequado para seu sistema operacional.
3. Siga as instruções do instalador.
4. Após a instalação, verifique se o Go foi instalado corretamente:
   ```bash
   go version
   ```
   O comando acima deve retornar a versão instalada do Go.

5. Adicione o diretório `$GOPATH/bin` ao PATH do sistema (se necessário):
   - Linux/MacOS: Edite o arquivo `~/.bashrc` ou `~/.zshrc`:
     ```bash
     export PATH=$PATH:$(go env GOPATH)/bin
     ```
   - Windows: Adicione `%GOPATH%\bin` às variáveis de ambiente.

---

## Instalação do Air

1. Instale o Air via `go install`:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```
2. Verifique se o Air foi instalado corretamente:
   ```bash
   air -v
   ```
   O comando acima deve retornar a versão do Air instalada.

---

## Rodando o Air no Projeto

1. Certifique-se de estar no diretório raiz do projeto.
2. Navegue até a pasta `cmd/server`:
   ```bash
   cd cmd/server
   ```
3. Execute o Air:
   ```bash
   air
   ```

O servidor será iniciado com suporte a live reload. Qualquer alteração no código-fonte automaticamente reiniciará o servidor.


## Recursos

- [Documentação do Go](https://golang.org/doc/)
- [Repositório do Air](https://github.com/cosmtrek/air)
```

Este arquivo fornece um guia completo para a configuração e execução do ambiente de desenvolvimento com Go e Air. Se precisar de mais personalizações ou esclarecimentos, é só pedir! 😊