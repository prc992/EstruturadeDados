//Refazer a lista 1 pois essa está duplamente encadeada

package main

import (
	"fmt"
)

type TADNo struct {
	ValorNo   int
	PosicaoNo int
	proximoNo *TADNo
	//anteriorNo *TADNo
}

type TADListaEncadeada struct {
	cabecaNo   *TADNo
	caudaNo    *TADNo
	tamanho    int
	listaVazia int
}

func (ListaEncadeada *TADListaEncadeada) IniciaLista() {
	ListaEncadeada.tamanho = 0
	ListaEncadeada.listaVazia = 1

}

func (ListaEncadeada *TADListaEncadeada) LiberaLista() {
	ListaEncadeada.cabecaNo.proximoNo = nil
	ListaEncadeada.cabecaNo = nil
	ListaEncadeada.caudaNo = nil
}

func (ListaEncadeada *TADListaEncadeada) AdicionaFim(ValorNo int) {
	var mNo = &TADNo{}
	var ultimoNo = &TADNo{}

	mNo.ValorNo = ValorNo
	mNo.proximoNo = nil
	ListaEncadeada.listaVazia = 0
	ListaEncadeada.tamanho = ListaEncadeada.tamanho + 1

	/*var ultimoNo = &TADNo{}
	ultimoNo = ListaEncadeada.ultimoNo()
	fmt.Println("Ultimo nó:", ultimoNo.ValorNo)*/

	if ListaEncadeada.caudaNo == nil {
		//fmt.Println("Inserir cauda :", ValorNo)
		ListaEncadeada.caudaNo = mNo
		ListaEncadeada.cabecaNo = mNo
		ListaEncadeada.caudaNo.proximoNo = ListaEncadeada.cabecaNo
		ListaEncadeada.cabecaNo.proximoNo = nil

	} else if ListaEncadeada.tamanho == 2 {
		ListaEncadeada.cabecaNo = mNo
		ListaEncadeada.caudaNo.proximoNo = ListaEncadeada.cabecaNo
		ListaEncadeada.cabecaNo.proximoNo = nil
		//fmt.Println("Inserir nó cabeça :", ValorNo)
	} else {

		ultimoNo = ListaEncadeada.ultimoNo()
		ultimoNo.proximoNo = mNo
		mNo.proximoNo = nil
		//fmt.Println("Imprimir nó normal: ", ValorNo)
	}
}

func (ListaEncadeada *TADListaEncadeada) ultimoNo() *TADNo {
	var mNo *TADNo
	var ultimoNo *TADNo

	for mNo = ListaEncadeada.cabecaNo; mNo != nil; mNo = mNo.proximoNo {
		//fmt.Println("Buscando ultimo nó : ", mNo.ValorNo)
		if mNo.proximoNo == nil {
			ultimoNo = mNo
		}
	}
	return ultimoNo
}

func (ListaEncadeada *TADListaEncadeada) AdicionaInicio(ValorNo int) {
	var mNo = &TADNo{}

	fmt.Println("Adiciona Inicio :", ValorNo)

	if ListaEncadeada.listaVazia == 1 {
		//fmt.Println("Inseirir inicio, lista vazia - ", ValorNo)
		mNo.ValorNo = ValorNo
		mNo.proximoNo = nil
		ListaEncadeada.cabecaNo = mNo
		ListaEncadeada.caudaNo = mNo
	} else {
		//fmt.Println("Inseirir inicio, lista não vazia - ", ValorNo, "Cauda no:", ListaEncadeada.caudaNo)
		mNo.ValorNo = ValorNo
		mNo.proximoNo = ListaEncadeada.caudaNo
		ListaEncadeada.caudaNo = mNo
	}
	ListaEncadeada.tamanho = ListaEncadeada.tamanho + 1
	ListaEncadeada.listaVazia = 0
}

func ImprimeLista(ListaEncadeada TADListaEncadeada) {

	var mNo *TADNo
	for mNo = ListaEncadeada.caudaNo; mNo != nil; mNo = mNo.proximoNo {

		fmt.Println(mNo.ValorNo)
	}
	fmt.Println("fim da lista (iterativa)")
}

func BuscaLista(ListaEncadeada TADListaEncadeada, valorBuscado int) int {

	var mNo *TADNo
	var retorno int
	retorno = 0

	for mNo = ListaEncadeada.cabecaNo; mNo != nil; mNo = mNo.proximoNo {
		if mNo.ValorNo == valorBuscado {
			retorno = 1
		}
	}

	return (retorno)
}

func ImprimeListaReversa(ListaEncadeada TADListaEncadeada) {

	var mNo *TADNo
	var vetorValores = []int{}

	for mNo = ListaEncadeada.caudaNo; mNo != nil; mNo = mNo.proximoNo {
		vetorValores = append(vetorValores, mNo.ValorNo)
	}
	var i int

	for i = len(vetorValores) - 1; i != -1; i-- {
		fmt.Println(vetorValores[i])
	}

	fmt.Println("fim da lista (reversa)")
}

func (ListaEncadeada *TADListaEncadeada) RemoveLista(valorBuscado int) {
	var prev *TADNo
	var noden *TADNo

	if ListaEncadeada.cabecaNo.ValorNo == valorBuscado {
		ListaEncadeada.cabecaNo = ListaEncadeada.cabecaNo.proximoNo
		ListaEncadeada.tamanho = ListaEncadeada.tamanho - 1
		return
	}

	prev = ListaEncadeada.cabecaNo
	noden = ListaEncadeada.cabecaNo.proximoNo

	for noden != nil {
		if noden.ValorNo == valorBuscado {
			prev.proximoNo = noden.proximoNo
		}
		prev = noden
		noden = noden.proximoNo
	}

}

func RemoveListaRecursao(mNo *TADNo, valorBuscado int) {
	if mNo != nil {
		var proximoNo *TADNo
		proximoNo = mNo.proximoNo

		if proximoNo != nil {
			if proximoNo.ValorNo == valorBuscado {
				//fmt.Println("O próximo nó tem o valor que eu quero!")
				//fmt.Println(valorBuscado)
				if proximoNo.proximoNo == nil {
					mNo.proximoNo = nil
				} else {
					mNo.proximoNo = proximoNo.proximoNo
				}
			}
		}

		if proximoNo == nil && mNo.ValorNo == valorBuscado {
			//fmt.Println("Achei o valor e é no último nó :", valorBuscado)
			mNo.proximoNo = nil
		} else {
			//fmt.Println("Continua procurando")
			RemoveListaRecursao(mNo.proximoNo, valorBuscado)
		}
	} else {
		//fmt.Println("Fim do percurso!")
	}
}

func ImprimeListaRecursao(mNo *TADNo) {
	if mNo != nil {
		fmt.Println(mNo.ValorNo)
		ImprimeListaRecursao(mNo.proximoNo)
	} else {
		fmt.Println("fim da lista (recursao)")
	}

}

func main() {

	var ListaEncadeada TADListaEncadeada

	ListaEncadeada = TADListaEncadeada{}
	ListaEncadeada.IniciaLista()

	ListaEncadeada.AdicionaFim(1)
	ListaEncadeada.AdicionaFim(2)
	ListaEncadeada.AdicionaFim(3)
	ListaEncadeada.AdicionaFim(4)
	ListaEncadeada.AdicionaFim(5)
	ListaEncadeada.AdicionaFim(6)
	ListaEncadeada.AdicionaFim(7)
	ListaEncadeada.AdicionaInicio(0)

	ImprimeLista(ListaEncadeada)
	ImprimeListaRecursao(ListaEncadeada.caudaNo)
	ImprimeListaReversa(ListaEncadeada)

	var busca int
	busca = 7
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeada, busca))
	busca = 70
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeada, busca))
	ListaEncadeada.RemoveLista(7)
	busca = 7
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeada, busca))
	busca = 5
	RemoveListaRecursao(ListaEncadeada.caudaNo, 5)
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeada, busca))
	ListaEncadeada.LiberaLista()
	fmt.Println("Lista esvaziada!")
	ImprimeLista(ListaEncadeada)
}
