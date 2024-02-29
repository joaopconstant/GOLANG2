# Estrutura de Dados - AP2

## Grupo
- João Constant
- João Lucas
- Luis Pastura
- Theo Furtado
- Vitor Lobianco

## Como usar o programa 

### Clonando o projeto

```bash
$ git clone https://github.com/lobiancovitor/ap2-golang.git
$ cd ap2-golang
```

### Executando

```bash
# main.go
$ go run main.go

# executável
$ mcronalds.exe
$ ./mcronalds
```

### Interface

```bash
===== PEDIDOS ELETRÔNICOS =====
1. Cadastrar produto
2. Remover produto
3. Buscar produto por id
4. Buscar produto por nome
5. Exibir produtos ordenados por id
6. Adicionar pedido
7. Expedir pedido
8. Exibir métricas do sistema
9. Atualizar preço do produto
10. Exibir produtos ordenados por nome
20. Exibir pedidos em andamento
21. Cadastrar produtos em lote 
100. Sair
```

## TODO

Requisitos Funcionais:
- [x] Atualizar preco do produto por id
- [x] Adicionar métrica: ticket médio (total faturado / pedidos encerrados)
- [x] Exibir os produtos por ordenados pelo nome

Requisitos Não Funcionais:
- [x] Lista de produtos &rarr; lista simplesmente encadeada (refatorar operacões)
  - [x] AdicionarUnico()
  - [x] BuscarId()
  - [x] BuscarNome()
  - [x] ExibirPorId()
  - [x] ExibirPorNome()
  - [x] bubblesort()
  - [x] Excluir()
  - [x] Atualizar()
- [x] Algoritmo de ordenacão para exibir produtos (bubblesort)

Bugs
- [x] Panic ao expedir pedido com lista vazia após adicionar um pedido
- [x] Programa não considerando ids dos produtos pré-carregados


#### &rarr; Testar cada funcionalidade antes de subir.