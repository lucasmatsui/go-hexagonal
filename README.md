# :fire: Arquitetura Hexagonal em GO

Projeto que manipula um produto, implementando arquitetura hexagonal em **GOlang**.

Aplicando alguns conceitos para proteger a camada de domínio da aplicação.

### Implementação do application
- Foi criando um `ProductServiceInterface` para cuidar do fluxo da aplicação.
- Um `ProductPersistenceInterface`, que é o responsável pela persistência dos dados.
-  E testando essas camadas usando mocks.
