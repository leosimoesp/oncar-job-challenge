# oncar-job-challenge

Mono repo para o desenvolvimento do projeto do oncar-job-challenge

### Arquitetura

Uma possível solução para o problema descrito nesse desafio é a utilização da "Clean Architecture" onde temos
as principais camadas conforme a imagem a seguir:

@TODO criar um diagrama com as camadas

Router - Camada que intercepta diretamente as requisições do clientes web. Neste desafio, considere que o request do usuário
proveniente do frontend tem credenciais de acesso válidas e verificadas.

Controller - Após a rota ter interceptado o request, ela direciona o mesmo para o controller de modo que o mesmo valide
os dados enviados atráves do payload e faz as devidas validações retornando erro quando for o caso. Se tudo estiver correto,
no request a camada de usecase irá realizar a operação.

Usecase - Camada que implementa a lógica de negócio. Essa camada utiliza a camada de repository para efetuar as operações
que incluem o banco de dados(storage).

Repository - Camada que tem uma abstração da interface de banco de dados, ela é totalmente independente da implementação utilizada
e pode ser facilmente trocada por outra implementação.

Domain - Camada que inclui:

Models: objetos/estruturas utilizados no request e response. Muitas vezes podems ser chamados de DTO(Data transfer object).
Entities: objetos/estruturas utilizados nas operações diretas com o banco de dados.
Interfaces: utilizadas pelas camadas de usecases e repositories.

Erro ao executar os testes:

leonardo@workstation:~/git/oncar-job-challenge$ go test ./...
pattern ./...: open /home/leonardo/git/oncar-job-challenge/postgres/data: permission denied

Solução 1: alterar a permissão da pasta postgres (volume local) utilizado pelo docker-compose com o comando

sudo chmod -R 775 postgres

Solução 2: mover a pasta postgres para outro local e criar um link simbólico na raiz do projeto com o nome de postgres
