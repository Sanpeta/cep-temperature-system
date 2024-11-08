# Sistema de Consulta de Temperatura por CEP

Este sistema desenvolvido em Go recebe um CEP válido de 8 dígitos, identifica a cidade correspondente e retorna a temperatura atual formatada em graus Celsius, Fahrenheit e Kelvin. O sistema está configurado para ser executado no Google Cloud Run.

## Funcionalidades

- Recebe um CEP de 8 dígitos via HTTP e verifica se é válido.
- Consulta a cidade associada ao CEP utilizando a API do ViaCEP.
- Consulta a temperatura atual da cidade com a OpenWeatherMap e retorna:
  - Temperatura em Celsius
  - Temperatura em Fahrenheit
  - Temperatura em Kelvin
- Responde com mensagens apropriadas para diferentes cenários de erro, como CEP inválido ou não encontrado.

## Respostas da API

- **Sucesso**:
  ```json
  {
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.65
  }
  ```

- **Erro de CEP inválido**:
  - Código HTTP: `422`
  - Corpo da resposta:
    ```json
    {
      "message": "invalid zipcode"
    }
    ```

- **CEP não encontrado**:
  - Código HTTP: `404`
  - Corpo da resposta:
    ```json
    {
      "message": "can not find zipcode"
    }
    ```

## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal do projeto.
- **Google Cloud Run**: Plataforma para deploy do sistema.
- **API ViaCEP**: Para consulta de cidade com base no CEP.
- **OpenWeatherMap API**: Para consulta da temperatura atual.

## Implementação de Conversão de Temperatura

Para converter as temperaturas, utilizamos as seguintes fórmulas:

- Fahrenheit (F) = Celsius (C) * 1.8 + 32
- Kelvin (K) = Celsius (C) + 273

## Configuração do Ambiente

1. **Clone o repositório**:

   ```bash
   git clone https://github.com/Sanpeta/cep-temperature-system
   cd cep-temperature-service
   ```

2. **Configurar as Variáveis de Ambiente**:

   Crie um arquivo `app.env` na raiz do projeto e adicione as variáveis de ambiente necessárias:

   ```plaintext
   TOKEN_WEATHER_API=<sua-chave-openweathermap>
   ```

## Como Executar o Projeto Localmente

Para rodar o sistema localmente:

1. **Instale as dependências**:

   ```bash
   go mod tidy
   ```

2. **Execute o projeto**:

   ```bash
   go run main.go
   ```

O serviço estará disponível em `http://localhost:8080`.

3. **Teste o Endpoint**:

   Utilize o seguinte comando para enviar uma requisição de teste ao serviço:

   ```bash
   curl -X GET http://localhost:8080/weather?cep=01001010'
   ```

   Caso o CEP seja válido e encontrado, uma resposta contendo a temperatura será retornada.

## Como Fazer o Deploy no Google Cloud Run

Para fazer o deploy no Google Cloud Run, siga os passos abaixo:

1. **Faça login no Google Cloud CLI**:

   ```bash
   gcloud auth login
   ```

2. **Crie o serviço no Google Cloud Run**:

   ```bash
   gcloud run deploy cep-temperature-service \
     --source . \
     --region us-central1 \
     --platform managed \
     --allow-unauthenticated
   ```

3. **Defina as variáveis de ambiente no Google Cloud Run**:

   Após o deploy, configure a variável `TOKEN_WEATHER_API` diretamente no painel do Google Cloud Run para que o serviço possa acessar a OpenWeatherMap API.

4. **Acesse o Endpoint do Serviço**:

   Após a implantação, um URL será fornecido pelo Google Cloud Run. Você pode utilizá-lo para enviar requisições para o serviço, da mesma forma que testou localmente.

### Acesso Público do Serviço

O sistema está disponível publicamente no Google Cloud Run. Você pode acessá-lo em: [https://cep-temperature-system-517639904608.us-central1.run.app/weather?cep=01001010](https://cep-temperature-system-517639904608.us-central1.run.app/weather?cep=01001010)

## Licença

Este projeto está licenciado sob a Licença MIT.
