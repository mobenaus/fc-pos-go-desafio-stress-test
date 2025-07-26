# fc-pos-go-desafio-stress-test

## Descrição do Desafio

```
Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.


O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.


Execução do Teste:

Realizar requests HTTP para a URL especificada.
Distribuir os requests de acordo com o nível de concorrência definido.
Garantir que o número total de requests seja cumprido.
Geração de Relatório:

Apresentar um relatório ao final dos testes contendo:
Tempo total gasto na execução
Quantidade total de requests realizados.
Quantidade de requests com status HTTP 200.
Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
Execução da aplicação:
Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
docker run <sua imagem docker> --url=http://google.com --requests=1000 --concurrency=10
```

## Implementação

O desafio foi implementado como um CLI usando Cobra para recuperar os argumentos da linha de comando.


A execução cria uma instância de uma estrutura **StressTest** passando os argumentos do CLI, executa os testes e apresenta o relatório do resultado.

A execução dos testes, feito pelo método *Execute()*, inicializa uma estrutura **StressRequest** com a URL de testes, e uma estrutura **StressTestResults** para guardar as estatísticas dos testes.

A estrutura **StressTestResults** possui métodos *Start* e *Finish* para computar o tempo decorrido na execução dos testes.

O método *Execute()* inicializa 1 channel simples de requests e um channel com buffer do tamanho das requests que serão chamadas, e 2 WaitGroup para controlar a concorrência da execução de testes e a concorrência do processo para computar os resultados.

A computação dos resultados é feita no método *CountResult(status int)* que recebe o status do resultado da requisisão via channel

O método *Execute()* inicia 1 Go Routine para cada 'concurrency' definido nos parametros, essas rotinas recebem uma requisição via channel, executam uma request e enviam o resultado para o channel de resultados.

O método *Execute()* também inicia uma Go Routine (reportStatusRoutine) para reportar o status do andamento dos testes, mostrando a cada segundo quantas requisições foram executadas.

Finalmente o método *Execute()* faz um loop para utilizar os channel de requisições para acionar as rotinas que executam as requisições.

Na estrutura **StressTestResults** os métodos não tem necessidade de usar técnicas de mutex ou atomic porque é executado no consumo do channel de resultados executando em uma rotina única.

O relatório de resultados mostra o tempo total utilizado para executar as requisições, o total de requisições, total de requisições com sucesso (status 200) e o total de requisições por status.
Finalmente também mostra a quantidade de requests por segundo médio dos testes.

## Execução

Localmente os testes foram utilizados usando um container local que implementa o comportamento de https://httpbin.org/ através do comando:
```
docker run -p 80:80 kennethreitz/httpbin
```

Para executar os testes localmente usando Go utilize o comando:
```
go run main.go --url http://localhost/status/200,202,400,404,500 --requests 500 --concurrency 20
```

Para executar os testes usando docker contra o esse httpbin local pode ser executar o comando:
```
docker run -it --rm --network=host ghcr.io/mobenaus/fc-pos-go-desafio-stress-test:latest --url http://localhost/status/200,202,400,404,500 --requests 500 --concurrency 20
```
Podemos utilizar diretamente no httpbin.org, mas não recomendo por risco de derrubar o serviço.

Testes com sucesso sem problemas foram executado para o endereço do Bing
```
docker run -it --rm ghcr.io/mobenaus/fc-pos-go-desafio-stress-test:latest --url=https://www.bing.com  
```
Resultado:
```
Stress Test executado em 32.880119723s
Total requests........: 1000
Total Success requests: 1000
Requests por segundo 30.413515
```

## Imagem Docker

A imagem foi gerada através do GitHub actions e disponibilizada no Registry do GitHub.
```
ghcr.io/mobenaus/fc-pos-go-desafio-stress-test
```





