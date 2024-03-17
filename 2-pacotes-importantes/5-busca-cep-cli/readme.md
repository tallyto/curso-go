# Documentação para o Código de Consulta de CEP via ViaCEP em Go

## Introdução

Este é um programa simples em Go que utiliza a API ViaCEP para consultar informações de endereços a partir de CEPs fornecidos como argumentos de linha de comando. O programa realiza uma requisição HTTP para o serviço ViaCEP, analisa a resposta JSON e, em seguida, cria um arquivo de texto com algumas informações do endereço.

## Estrutura do Código

### Struct `ViaCEP`

A struct `ViaCEP` representa a estrutura dos dados recebidos da API ViaCEP. Os campos na struct são mapeados diretamente aos campos no JSON da resposta, usando tags `json`.

```go
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}
```

### Função `main`

A função principal do programa realiza as seguintes etapas:

1. Itera sobre os CEPs fornecidos como argumentos de linha de comando.
2. Faz uma requisição HTTP para a API ViaCEP para cada CEP.
3. Lê e decodifica a resposta JSON em uma instância da struct `ViaCEP`.
4. Cria um arquivo chamado "cidade.txt" e escreve algumas informações do endereço nele.
5. Exibe no console as informações do endereço e uma mensagem indicando o sucesso na criação do arquivo.

## Execução do Programa

O programa espera que os CEPs sejam fornecidos como argumentos de linha de comando. Por exemplo:

```bash
go run main.go 13083842 04101000 20050000
```

Este comando consultará os CEPs 13083-842, 04101-000 e 20050-000 e criará um arquivo "cidade.txt" com informações sobre esses endereços.

## Considerações Finais

Este código serve como um exemplo básico de como interagir com APIs em Go, realizar análise de JSON e manipular arquivos. Certifique-se de respeitar os termos de serviço da API que você está utilizando ao implementar códigos semelhantes em produção. Para melhorar e expandir este código, considere adicionar tratamento de erros mais robusto, suporte para diferentes formatos de saída ou até mesmo integrar mais funcionalidades da API ViaCEP.