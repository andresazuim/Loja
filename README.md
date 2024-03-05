# Projeto de Loja em Go


Este projeto consiste em uma aplicação de loja desenvolvida em Go, com o PostgreSQL como banco de dados principal. 
Ele segue o padrão MVC para uma organização eficiente do código.

## Funcionalidades Principais 🚀

**1. Cadastro de Produtos:**
   - Permite a criação de novos produtos, inserindo as informações necessárias.
   - Utiliza um controlador dedicado para processar e salvar os dados no banco de dados.

**2. Consulta de Produtos:**
   - Disponibiliza uma rota para buscar produtos com base em diferentes critérios.
   - O controlador é responsável por acessar o banco de dados e retornar as informações solicitadas.

**3. Edição de Produtos:**
   - Oferece uma rota para atualizar as informações de produtos existentes.
   - O controlador modifica os dados no banco de dados conforme necessário.

**4. Exclusão de Produtos:**
   - Implementa uma rota para deletar produtos existentes.
   - O controlador é encarregado de remover as informações do banco de dados.

## Estrutura do Projeto

O projeto segue a arquitetura MVC:

- **Model:** Manipula e implementa a lógica dos dados, interagindo com o banco de dados.
- **View:** Em um contexto de backend, foca menos na interface do usuário e mais na formatação da resposta.
- **Controller:** Gerencia as requisições do usuário, manipula a lógica de negócios e interage com o modelo.


## Tecnologias Utilizadas 🛠️

- **Go:** Linguagem de programação principal.
- **PostgreSQL:** Banco de dados utilizado para armazenar informações sobre os produtos.
- **MVC:** Padrão de arquitetura adotado para organização do código.

## Pré-requisitos 🌐

Certifique-se de ter um navegador web moderno instalado para uma melhor experiência durante o uso da aplicação.
✨✨✨✨✨✨✨✨
