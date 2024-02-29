package produtos

import (
	m "mcronalds/metricas"
	"strings"
)

const maxProdutos = 50
var lista Lista

type No struct {
	produto Produto
	prox    *No
}

type Lista struct {
	cab           *No
	totalProdutos int
}

/*
Auxiliar para Atualizar(id int, novoPreco float64) e Excluir(id int).
Vista em sala.
*/
func buscar(id int) (*No, *No) {
	var ant *No
	no := lista.cab

	if no == nil { return nil, nil }

	for no != nil {
		if no.produto.Id == id { return ant, no }
		if no.produto.Id > id { return ant, nil }

		ant = no
		no = no.prox
	}

	return ant, nil
}

func tentarCriar(nome, descricao string, preco float64, id int) Produto {
	if id != -1 {
		_, idProcurado := BuscarId(id)
		if idProcurado != -1 { return Produto{} }
	}

	return criar(nome, descricao, preco, id)
}

/*
Adiciona um produto com nome, descrição e preço à lista de produtos.
Adiciona o produto primeiro espaço vazio da lista.
Caso já exista um produto com o mesmo id, não adiciona e retorna -3.
Caso já exista um produto com o mesmo nome, não adiciona e retorna erro -2.
Retorna -1 caso a lista esteja cheia, ou o número de produtos cadastrados em caso de sucesso.
*/
func AdicionarUnico(nome, descricao string, preco float64, id int) int {
	if lista.totalProdutos == maxProdutos {
		return -1 // Overflow
	}

	checarNome, _ := BuscarNome(nome)
	if len(checarNome) > 0 {
		return -2
	}

	_, checarId := BuscarId(id)
	if checarId != -1 {
		return -3
	}

	produtoCriado := tentarCriar(nome, descricao, preco, id)

	novoNo := &No{produto: produtoCriado}
	if lista.cab == nil {
        lista.cab = novoNo
    } else {
        ultimo := lista.cab
        for ultimo.prox != nil {
            ultimo = ultimo.prox
        }
        ultimo.prox = novoNo
    }

	lista.totalProdutos++
	m.M.SomaProdutosCadastrados(1)
	return lista.totalProdutos
}

/*
Localiza um produto a partir do seu id.
Retorna o produto encontrado e a sua posição na lista, em caso de sucesso.
Retorna um produto vazio e -1 em caso de erro.
*/
func BuscarId(id int) (*Produto, int) {
	no := lista.cab
	indice := 0

	if no == nil {
		return nil, -1
	}

	for no != nil {
		if no.produto.Id == id {
			return &no.produto, indice
		}
		no = no.prox
		indice++
	}

	return nil, -1
}

/*
Localiza produtos que iniciem com a string passada.
Retorna um slice com todos os produtos encontrados, e o tamanho do slice.
*/
func BuscarNome(nome string) ([]Produto, int) {
	var produtosEncontrados []Produto
	no := lista.cab
	indice := 0

	for no != nil {
		if strings.HasPrefix(no.produto.Nome, nome) {
			produtosEncontrados = append(produtosEncontrados, no.produto)
		}
		no = no.prox
		indice++
	}

	return produtosEncontrados, len(produtosEncontrados)
}

/*
Exibe todos os produtos cadastrados ordenados pelo id.
*/
func ExibirPorId() {
	bubbleSortId()
	no := lista.cab

	for no != nil {
		no.produto.Exibir()
		no = no.prox
	}
}

func bubbleSortNome() {
	trocou := true
	limite := lista.totalProdutos

	for trocou {
		trocou = false
		limite --
		no := lista.cab
		for i := 0; i < limite && no != nil && no.prox != nil; i++ {
			prox := no.prox
			if no.produto.Nome > prox.produto.Nome && no.produto.Nome != "" && prox.produto.Nome != "" {
				no.produto, prox.produto = prox.produto, no.produto
				trocou = true
			}
			no = prox
		}
	}
}

func bubbleSortId() {
	trocou := true
	limite := lista.totalProdutos

	for trocou {
		trocou = false
		limite--
		no := lista.cab
		for i := 0; i < limite && no != nil && no.prox != nil; i++ {
			prox := no.prox
			if no.produto.Id > prox.produto.Id && no.produto.Id != 0 && prox.produto.Id != 0 {
				no.produto, prox.produto = prox.produto, no.produto
				trocou = true
			}
			no = prox
		}
	}
}

/*
Exibe todos os produtos cadastrados ordenados pelo nome.
*/
func ExibirPorNome() {
    bubbleSortNome()

    no := lista.cab
    for no != nil {
        if no.produto == (Produto{}) {
            break
        }
        no.produto.Exibir()
        no = no.prox
    }
}

/*
Remove um produto da lista a partir do seu id.
Retorna -2 caso não haja produtos na lista.
Retorna -1 caso não haja um produto com o id passado, ou 0 em caso de sucesso.
*/
func Excluir(id int) int {
	ant, no := buscar(id)

	if lista.cab == nil { return -2 }
	if no == nil { return -1 }

	if ant == nil {
		lista.cab = no.prox
	} else {
		ant.prox = no.prox
	}

	lista.totalProdutos--
	m.M.SomaProdutosCadastrados(-1)
	return 0
}

/*
Atualiza o preço de um produto da lista a partir do seu id.
Retorna -2 caso não haja produtos na lista.
Retorna -1 caso não haja um produto com o id passado, ou 0 em caso de sucesso.
*/
func Atualizar(id int, novoPreco float64) int {
	_, no := buscar(id)

	if lista.cab == nil { return -2 }
	if no == nil { return -1 }

	no.produto.Preco = novoPreco
	return 0
}
