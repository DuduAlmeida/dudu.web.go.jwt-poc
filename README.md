# 🚀 PoC: Autenticação JWT e Refresh Token em Go (Echo Framework)

Este projeto é uma Prova de Conceito (PoC) minimalista para demonstrar a implementação robusta de um sistema de autenticação utilizando **JSON Web Tokens (JWT)** e o mecanismo de **Refresh Token** para gerenciamento de sessões de longa duração.

A arquitetura utiliza o _framework_ **Echo** em Go para roteamento e middlewares.

## ✨ Funcionalidades Principais

O foco desta PoC é testar o fluxo completo de autenticação e autorização, garantindo que o ciclo de vida dos tokens seja gerenciado corretamente:

1. **Geração de Tokens:** No _login_ , são gerados dois tokens distintos:
   - **Access Token:** Curta duração (15 minutos). Usado para acessar recursos protegidos.
   - **Refresh Token:** Longa duração (7 dias). Usado para obter um novo _Access Token_ quando o atual expirar.
2. **Acesso Protegido (Middleware):** A rota `/api/v1/cars` é protegida por um _middleware_ que exige um **Access Token válido** . A requisição falhará (HTTP 401) se o token estiver ausente, malformado ou expirado.
3. **Renovação de Tokens (`/refresh`):** O _endpoint_ `/api/v1/refresh` permite que o cliente envie um _Refresh Token_ válido para receber um **novo par de Access Token e Refresh Token** , sem a necessidade de um novo login (fornecendo credenciais).

## 🧩 Arquitetura de Autenticação

A PoC segue o padrão de segurança recomendado para JWTs .

| **Componente**     | **Função**                                                                                                            |
| ------------------ | --------------------------------------------------------------------------------------------------------------------- |
| **Access Token**   | Autentica e autoriza o acesso à API. É rapidamente expirável para mitigar roubo.                                      |
| **Refresh Token**  | Funciona como uma credencial de longa duração para*re-emitir*um _Access Token_ . Deve ser tratado com alta segurança. |
| **Middleware JWT** | Intercepta requisições para rotas protegidas, valida o**Access Token**e, se inválido, bloqueia o acesso.              |

---

## ⚠️ Nota sobre Regras de Negócio e Implementação

É importante ressaltar que, por se tratar de uma **Prova de Conceito** focada exclusivamente em testar o fluxo de JWT:

> As _controllers_ e _handlers_ deste projeto **incluem a lógica de negócios e persistência (usuários mockados) diretamente** em suas funções. Em uma aplicação de produção, esta lógica seria rigorosamente separada em camadas (Serviços/Repositórios) para garantir escalabilidade, testabilidade e Manutenibilidade.

O objetivo principal é a **demonstração funcional do JWT** , não a arquitetura limpa do Go.

## 🧪 Como Testar os Endpoints

Para facilitar a experimentação e teste do fluxo de tokens, uma coleção completa para o cliente **Bruno** está configurada na pasta `docs`.

### Requisitos

- Go v1.18+
- Cliente HTTP [Bruno](https://www.usebruno.com/) (ou Postman, Insomnia)

### Instruções

1. **Rodar o Servidor:**
   **Bash**

   ```
   go run main.go
   ```

2. **Importar Coleção:**

   - Abra o cliente **Bruno** .
   - Importe a pasta `docs/bruno-collection`.

3. **Executar o Fluxo de Teste:**

   1. Login
   2. Cars
   3. Refresh
