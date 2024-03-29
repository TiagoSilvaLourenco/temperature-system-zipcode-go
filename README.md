# temperature-system-zipcode-go

## How to works

- This program works getting temperature in Celsius, Fahrenheit and Kelvin from a zipcode, return in json format.

- The following the format returned by the API: {"tempC":26,"tempF":78.8,"tempK":299.1}

- Where tempc is temperature in Celsius, tempF is temperature in Fahrenheit and tempK is temperature in Kelvin.

## How to test

- git clone project: `git clone git@github.com:TiagoSilvaLourenco/temperature-system-zipcode-go.git`

- In development mode run `docker compose up` to run the program in the port 8080

- For run the program in production run command: `docker compose -f docker-compose.prod.yml up`

- In local machine run `go test` to test the program

- To test the program online, go to: https://temperature-system-goexpert-n5ko4ubijq-uc.a.run.app/{cep}}

- Where {cep} is the zipcode number you want to get the temperature scales, eg.: https://temperature-system-goexpert-n5ko4ubijq-uc.a.run.app/01001000

<!--
    Objetivo: Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

    Requisitos:

    O sistema deve receber um CEP válido de 8 digitos
    O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
    O sistema deve responder adequadamente nos seguintes cenários:
    Em caso de sucesso:
    Código HTTP: 200
    Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
    Em caso de falha, caso o CEP não seja válido (com formato correto):
    Código HTTP: 422
    Mensagem: invalid zipcode
    ​​​Em caso de falha, caso o CEP não seja encontrado:
    Código HTTP: 404
    Mensagem: can not find zipcode
    Deverá ser realizado o deploy no Google Cloud Run.
    Dicas:

    Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
    Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
    Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
    Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
    Sendo F = Fahrenheit
    Sendo C = Celsius
    Sendo K = Kelvin
    Entrega:

    O código-fonte completo da implementação.
    Documentação explicando como rodar o projeto em ambiente dev e production.
    Testes automatizados demonstrando o funcionamento.
    Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
    Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.
 -->
