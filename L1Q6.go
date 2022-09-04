//Refazer a lista 1 pois essa está duplamente encadeada

package main

import (
	"fmt"
)

const TContaBancaria string = "ContaBancaria"
const TContaPoupanca string = "ContaPoupanca"
const TContaFidelidade string = "ContaFidelidade"

type TADConta struct {
	Numero       int
	Saldo        float64
	Bonus        float64
	TipoConta    string
	proximaConta *TADConta
}

type TADListaEncadeadaContas struct {
	cabecaNo   *TADConta
	caudaNo    *TADConta
	tamanho    int
	listaVazia int
}

func (ListaEncadeada *TADListaEncadeadaContas) IniciaLista() {
	ListaEncadeada.tamanho = 0
	ListaEncadeada.listaVazia = 1

}

func (ListaEncadeada *TADListaEncadeadaContas) LiberaLista() {

	if ListaEncadeada.cabecaNo.proximaConta != nil {
		ListaEncadeada.cabecaNo.proximaConta = nil
	}

	if ListaEncadeada.caudaNo.proximaConta != nil {
		ListaEncadeada.caudaNo.proximaConta = nil
	}

	ListaEncadeada.cabecaNo = nil
	ListaEncadeada.caudaNo = nil
	ListaEncadeada.tamanho = 0

}

/*func (ListaEncadeadaContas *TADListaEncadeadaContas) AdicionaOrdenado(ValorNovaConta int) {
	//fmt.Println("Nó a ser adicionado :", ValorNovoNo)

	if ListaEncadeadaContas.tamanho == 0 {
		//fmt.Println("Adiciona primeiro nó, valor novo nó:", ValorNovaConta)
		ListaEncadeadaContas.AdicionaFim(ValorNovaConta)
		ListaEncadeadaContas.listaVazia = 0
		ListaEncadeadaContas.tamanho = ListaEncadeadaContas.tamanho + 1
		return
	}

	//fmt.Println("Vou Adicionar um segundo nó")
	if ValorNovaConta >= ListaEncadeadaContas.ultimaConta().Numero {
		//fmt.Println("Adiciona segundo nó depois da cabeça, valor novo nó:", ValorNovoNo)
		ListaEncadeadaContas.AdicionaFim(ValorNovaConta)
		return
	} else if ValorNovaConta <= ListaEncadeadaContas.caudaNo.Numero {
		//fmt.Println("Adiciona  nó antes da cauda, valor novo nó:", ValorNovoNo)
		ListaEncadeadaContas.AdicionaInicio(ValorNovaConta)
		return

	} else {
		var mContaAnterior, mContaPosterior *TADConta
		//fmt.Println("Tentando inserir que não se encaixa nas regras anteriores :", ValorNovoNo)
		mContaAnterior, mContaPosterior = ListaEncadeadaContas.BuscaNosProximos(ValorNovaConta)
		var mConta = &TADConta{}
		mConta.Numero = ValorNovaConta
		mContaAnterior.proximaConta = mConta
		mConta.proximaConta = mContaAnterior
		ListaEncadeadaContas.tamanho++
		//fmt.Println(mNoAnterior.ValorNo, mNoPosterior.ValorNo)
	}
}*/

func (ListaEncadeadaContas *TADListaEncadeadaContas) BuscaNosProximos(NumeroContaBuscada int) (*TADConta, *TADConta) {
	var mContaAnalisada, mContaAnterior, mContaPosterior *TADConta

	for mContaAnalisada = ListaEncadeadaContas.caudaNo; mContaAnalisada.proximaConta != nil; mContaAnalisada = mContaAnalisada.proximaConta {
		//fmt.Println("BuscaNosProximos, NoAnalisado :", noAnalisado.ValorNo)
		if mContaAnalisada.Numero <= NumeroContaBuscada && mContaAnalisada.proximaConta.Numero > NumeroContaBuscada {
			//fmt.Println("Nó encontrado entr os nós :", noAnalisado.ValorNo, "-", noAnalisado.proximoNo.ValorNo)
			mContaAnterior = mContaAnalisada
			mContaPosterior = mContaAnalisada.proximaConta
			break
		}
	}

	return mContaAnterior, mContaPosterior
}
func (ListaEncadeadaContas *TADListaEncadeadaContas) AdicionaFim(NumeroConta int, Saldo float64, Bonus float64, TipoConta string) {
	var mConta = &TADConta{}
	var ultimaConta = &TADConta{}

	mConta.Numero = NumeroConta
	mConta.Saldo = Saldo
	mConta.TipoConta = TipoConta

	if TipoConta == TContaFidelidade {
		mConta.Bonus = Bonus
	} else {
		mConta.Bonus = 0
	}

	mConta.proximaConta = nil
	ListaEncadeadaContas.listaVazia = 0
	ListaEncadeadaContas.tamanho = ListaEncadeadaContas.tamanho + 1

	/*var ultimoNo = &TADNo{}
	ultimoNo = ListaEncadeada.ultimoNo()
	fmt.Println("Ultimo nó:", ultimoNo.ValorNo)*/

	if ListaEncadeadaContas.caudaNo == nil {
		//fmt.Println("Inserir cauda :", ValorNo)
		ListaEncadeadaContas.caudaNo = mConta
		ListaEncadeadaContas.cabecaNo = mConta
		ListaEncadeadaContas.caudaNo.proximaConta = ListaEncadeadaContas.cabecaNo
		ListaEncadeadaContas.cabecaNo.proximaConta = nil

	} else if ListaEncadeadaContas.tamanho == 2 {
		ListaEncadeadaContas.cabecaNo = mConta
		ListaEncadeadaContas.caudaNo.proximaConta = ListaEncadeadaContas.cabecaNo
		ListaEncadeadaContas.cabecaNo.proximaConta = nil
		//fmt.Println("Inserir nó cabeça :", ValorNo)
	} else {

		ultimaConta = ListaEncadeadaContas.ultimaConta()
		ultimaConta.proximaConta = mConta
		mConta.proximaConta = nil
		//fmt.Println("Imprimir nó normal: ", ValorNo)
	}
}

func (ListaEncadeadaContas *TADListaEncadeadaContas) ultimaConta() *TADConta {
	var mConta *TADConta
	var ultimaConta *TADConta

	for mConta = ListaEncadeadaContas.cabecaNo; mConta != nil; mConta = mConta.proximaConta {
		//fmt.Println("Buscando ultimo nó : ", mNo.ValorNo)
		if mConta.proximaConta == nil {
			ultimaConta = mConta
		}
	}
	return ultimaConta
}

func (ListaEncadeadaContas *TADListaEncadeadaContas) AdicionaInicio(NumeroConta int, Saldo float64, Bonus float64, TipoConta string) {
	var mConta = &TADConta{}

	mConta.Numero = NumeroConta
	mConta.Saldo = Saldo
	mConta.TipoConta = TipoConta

	if TipoConta == TContaFidelidade {
		mConta.Bonus = Bonus
	} else {
		mConta.Bonus = 0
	}

	fmt.Println("Adiciona Inicio :", NumeroConta)

	if ListaEncadeadaContas.listaVazia == 1 {
		//fmt.Println("Inseirir inicio, lista vazia - ", NumeroConta)
		mConta.proximaConta = nil
		ListaEncadeadaContas.cabecaNo = mConta
		ListaEncadeadaContas.caudaNo = mConta
	} else {
		//fmt.Println("Inseirir inicio, lista não vazia - ", NumeroConta)
		mConta.proximaConta = ListaEncadeadaContas.caudaNo
		ListaEncadeadaContas.caudaNo = mConta
	}
	ListaEncadeadaContas.tamanho = ListaEncadeadaContas.tamanho + 1
	ListaEncadeadaContas.listaVazia = 0

}

func ImprimeLista(ListaEncadeadaContas TADListaEncadeadaContas) {

	var mConta *TADConta
	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {

		fmt.Println("Conta:", mConta.Numero, " / Tipo Conta:", mConta.TipoConta, " / Saldo :", mConta.Saldo, " / Bonus :", mConta.Bonus)
	}
	fmt.Println("fim da lista (iterativa)")
}

func BuscaLista(ListaEncadeadaContas TADListaEncadeadaContas, contaBuscada int) int {

	var mConta *TADConta
	var retorno int
	retorno = 0

	//fmt.Println("Numero conta a ser buscada:", contaBuscada)
	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {
		//fmt.Println("Numero conta a ser buscada:", contaBuscada, "Conta Analisada :", mConta.Numero)
		if mConta.Numero == contaBuscada {
			retorno = 1
		}
	}

	return (retorno)
}

func RealizaCredito(ListaEncadeadaContas TADListaEncadeadaContas, contaBuscada int, valorCreditado float64) int {

	var mConta *TADConta
	var contaEncontrada int
	contaEncontrada = 0

	//fmt.Println("Numero conta a ser buscada:", contaBuscada)
	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {
		//fmt.Println("Numero conta a ser buscada:", contaBuscada, "Conta Analisada :", mConta.Numero)
		if mConta.Numero == contaBuscada {
			mConta.Saldo = mConta.Saldo + valorCreditado
			if mConta.TipoConta == TContaFidelidade {
				mConta.Bonus = mConta.Bonus + valorCreditado*0.01
			}
			contaEncontrada = 1
		}
	}

	if contaEncontrada == 1 {
		fmt.Println("Crédito efetuado!")

	} else {

		fmt.Println("Conta não encontrada!")
	}
	return contaEncontrada
}

func RenderBonus(ListaEncadeadaContas TADListaEncadeadaContas, contaBuscada int) {
	var mConta *TADConta
	var contaEncontrada int
	contaEncontrada = 0

	//fmt.Println("Numero conta a ser buscada:", contaBuscada)
	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {
		//fmt.Println("Numero conta a ser buscada:", contaBuscada, "Conta Analisada :", mConta.Numero)
		if mConta.Numero == contaBuscada {
			if mConta.TipoConta == TContaFidelidade {
				mConta.Saldo = mConta.Saldo + mConta.Bonus
				mConta.Bonus = 0
			}
			contaEncontrada = 1
		}
	}

	if contaEncontrada == 1 {
		fmt.Println("Bônus adicionado")
	} else {
		fmt.Println("A conta não foi encontrada ou não era do tipo fidelidade!")
	}

}

func RenderJuros(ListaEncadeadaContas TADListaEncadeadaContas, contaBuscada int, juros float64) {
	var mConta *TADConta
	var contaEncontrada int
	contaEncontrada = 0

	//fmt.Println("Numero conta a ser buscada:", contaBuscada)
	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {
		//fmt.Println("Numero conta a ser buscada:", contaBuscada, "Conta Analisada :", mConta.Numero)
		if mConta.Numero == contaBuscada {
			if mConta.TipoConta == TContaPoupanca {
				mConta.Saldo = mConta.Saldo + (mConta.Saldo * juros)
			}
			contaEncontrada = 1
		}
	}

	if contaEncontrada == 1 {
		fmt.Println("Juros adicionados")
	} else {
		fmt.Println("A conta não foi encontrada ou não era do tipo poupança!")
	}

}

func RealizaTransferenciaEntreContas(ListaEncadeadaContas TADListaEncadeadaContas, contaDebitada int, contaCreditada int, Valor float64) int {
	var contaEncontrada int
	contaEncontrada = 0

	if RealizaDebito(ListaEncadeadaContas, contaDebitada, Valor) == 1 {
		if RealizaCredito(ListaEncadeadaContas, contaCreditada, Valor) == 1 {
			fmt.Println("Transferência efetuada!")
			contaEncontrada = 1
		}
	}
	return contaEncontrada
}

func imprimeSaldo(ListaEncadeadaContas TADListaEncadeadaContas, contaBuscada int) {
	var mConta *TADConta
	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {
		if mConta.Numero == contaBuscada {
			if mConta.TipoConta == TContaFidelidade {
				fmt.Println("Saldo Conta:", mConta.Saldo, " Saldo Bônus:", mConta.Bonus)
			} else {
				fmt.Println("Saldo Conta:", mConta.Saldo)
			}
		} else {
			//fmt.Println("Conta não encontrada!")
		}
	}
}

func RealizaDebito(ListaEncadeadaContas TADListaEncadeadaContas, contaBuscada int, valorDebitado float64) int {

	var mConta *TADConta
	var contaEncontrada int
	contaEncontrada = 0

	//fmt.Println("Numero conta a ser buscada:", contaBuscada)
	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {
		//fmt.Println("Numero conta a ser buscada:", contaBuscada, "Conta Analisada :", mConta.Numero)
		if mConta.Numero == contaBuscada {
			mConta.Saldo = mConta.Saldo - valorDebitado
			contaEncontrada = 1
		}
	}

	if contaEncontrada == 1 {

		fmt.Println("Débito efetuado!")
	} else {

		fmt.Println("Conta não encontrada!")
	}

	return contaEncontrada
}

func comparaDuasListas(ListaEncadeadaContas1 TADListaEncadeadaContas, ListaEncadeadaContas2 TADListaEncadeadaContas) int {
	if ListaEncadeadaContas1.listaVazia == 1 && ListaEncadeadaContas2.listaVazia == 1 {
		return 1
	} else if ListaEncadeadaContas2.tamanho != ListaEncadeadaContas1.tamanho {
		//fmt.Println("Tamanho lista 1 : ", ListaEncadeadaContas1.tamanho)
		//fmt.Println("Tamanho lista 2 : ", ListaEncadeadaContas2.tamanho)
		return 0
	} else {
		var mConta1 *TADConta
		var mConta2 *TADConta

		mConta2 = ListaEncadeadaContas2.caudaNo
		fmt.Println("Executar comparações!")
		for mConta1 = ListaEncadeadaContas1.caudaNo; mConta1 != nil; mConta1 = mConta1.proximaConta {
			if mConta1.Numero != mConta2.Numero && mConta1.TipoConta != mConta2.TipoConta && mConta1.Saldo != mConta2.Saldo && mConta1.Bonus != mConta2.Bonus {
				fmt.Println("Nós diferentes :", mConta1.Numero, "-", mConta2.Numero)
				return 0
			}
			fmt.Println("Nós iguais :", mConta1.Numero, "-", mConta2.Numero)
			mConta2 = mConta2.proximaConta
		}
		return 1
	}
}

func ImprimeListaReversa(ListaEncadeadaContas TADListaEncadeadaContas) {

	var mConta, mContaAtual *TADConta
	var vetorValores = []*TADConta{}

	for mConta = ListaEncadeadaContas.caudaNo; mConta != nil; mConta = mConta.proximaConta {
		vetorValores = append(vetorValores, mConta)
	}
	var i int

	for i = len(vetorValores) - 1; i != -1; i-- {
		mContaAtual = vetorValores[i]
		fmt.Println("Conta:", mContaAtual.Numero, " / Tipo Conta:", mContaAtual.TipoConta, " / Saldo :", mContaAtual.Saldo, " / Bonus :", mContaAtual.Bonus)
		//fmt.Println(mContaAtual.Numero)
	}

	fmt.Println("fim da lista (reversa)")
}

func (ListaEncadeadaContas *TADListaEncadeadaContas) RemoveLista(numeroContaBuscada int) {
	var prev *TADConta
	var noden *TADConta

	fmt.Println("Remover nó da lista : ", numeroContaBuscada)
	fmt.Println("Nó cabeça :", ListaEncadeadaContas.cabecaNo.Numero)
	fmt.Println("Nó seguinte Cabeça:", ListaEncadeadaContas.cabecaNo.proximaConta.Numero)

	if ListaEncadeadaContas.cabecaNo.Numero == numeroContaBuscada {
		fmt.Println("Achei o nó buscado, e ele é o nó cabeça :", numeroContaBuscada)
		ListaEncadeadaContas.cabecaNo = nil //ListaEncadeadaContas.cabecaNo.proximaConta
		ListaEncadeadaContas.tamanho = ListaEncadeadaContas.tamanho - 1
		return
	}

	prev = ListaEncadeadaContas.cabecaNo
	noden = ListaEncadeadaContas.cabecaNo.proximaConta

	for noden != nil {
		fmt.Println("Procurando nó:", noden.Numero)
		if noden.Numero == numeroContaBuscada {
			fmt.Println("Achei o nó buscado :", numeroContaBuscada)
			prev.proximaConta = noden.proximaConta
		}
		prev = noden
		noden = noden.proximaConta
	}

}

func RemoveListaRecursao(mConta *TADConta, numeroContaBuscada int) {
	if mConta != nil {
		var proximaConta *TADConta
		proximaConta = mConta.proximaConta

		if proximaConta != nil {
			if proximaConta.Numero == numeroContaBuscada {
				//fmt.Println("O próximo nó tem o valor que eu quero!")
				//fmt.Println(valorBuscado)
				if proximaConta.proximaConta == nil {
					mConta.proximaConta = nil
				} else {
					mConta.proximaConta = proximaConta.proximaConta
				}
			}
		}

		if proximaConta == nil && mConta.Numero == numeroContaBuscada {
			fmt.Println("Achei o valor e é no último nó :", numeroContaBuscada)
			mConta.proximaConta = nil
		} else {
			fmt.Println("Continua procurando")
			RemoveListaRecursao(mConta.proximaConta, numeroContaBuscada)
		}
	} else {
		fmt.Println("Fim do percurso!")
	}
}

func ImprimeListaRecursao(mConta *TADConta) {
	if mConta != nil {
		fmt.Println("Conta:", mConta.Numero, " / Tipo Conta:", mConta.TipoConta, " / Saldo :", mConta.Saldo, " / Bonus :", mConta.Bonus)
		ImprimeListaRecursao(mConta.proximaConta)
	} else {
		fmt.Println("fim da lista (recursao)")
	}

}

func main() {

	var ListaEncadeadaContas TADListaEncadeadaContas
	var ListaEncadeadaContas2 TADListaEncadeadaContas

	ListaEncadeadaContas = TADListaEncadeadaContas{}
	ListaEncadeadaContas2 = TADListaEncadeadaContas{}

	ListaEncadeadaContas.IniciaLista()
	ListaEncadeadaContas2.IniciaLista()

	ListaEncadeadaContas.AdicionaFim(1, 0, 0, TContaBancaria)
	ListaEncadeadaContas.AdicionaFim(2, 100, 90, TContaFidelidade)
	ListaEncadeadaContas.AdicionaFim(5, 700, 980, TContaPoupanca)

	ImprimeLista(ListaEncadeadaContas)
	RealizaCredito(ListaEncadeadaContas, 2, 123)
	RealizaDebito(ListaEncadeadaContas, 5, 560.5)
	ImprimeLista(ListaEncadeadaContas)
	RenderBonus(ListaEncadeadaContas, 2)
	ImprimeLista(ListaEncadeadaContas)
	RenderJuros(ListaEncadeadaContas, 5, 0.1)
	ImprimeLista(ListaEncadeadaContas)

	RealizaTransferenciaEntreContas(ListaEncadeadaContas, 5, 1, 10)
	ImprimeLista(ListaEncadeadaContas)
	imprimeSaldo(ListaEncadeadaContas, 2)

	ImprimeListaRecursao(ListaEncadeadaContas.caudaNo)
	ImprimeListaReversa(ListaEncadeadaContas)

	var busca int
	/*busca = 2
	fmt.Println("Conta buscada:", busca, " Resultado : ", BuscaLista(ListaEncadeadaContas, busca))
	busca = 70
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeadaContas, busca))*/
	ImprimeLista(ListaEncadeadaContas)
	busca = 5
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeadaContas, busca))
	ListaEncadeadaContas.RemoveLista(busca)
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeadaContas, busca))
	ImprimeLista(ListaEncadeadaContas)
	busca = 2
	RemoveListaRecursao(ListaEncadeadaContas.caudaNo, 2)
	fmt.Println("Elemento buscado:", busca, " Resultado : ", BuscaLista(ListaEncadeadaContas, busca))
	ListaEncadeadaContas.LiberaLista()
	fmt.Println("Lista esvaziada!")
	ImprimeLista(ListaEncadeadaContas)

	ListaEncadeadaContas.AdicionaFim(1, 0, 0, TContaBancaria)
	ListaEncadeadaContas.AdicionaFim(2, 100, 90, TContaFidelidade)
	ListaEncadeadaContas.AdicionaFim(5, 700, 980, TContaPoupanca)

	ListaEncadeadaContas2.AdicionaFim(1, 0, 0, TContaBancaria)
	ListaEncadeadaContas2.AdicionaFim(2, 100, 90, TContaFidelidade)
	ListaEncadeadaContas2.AdicionaFim(5, 700, 980, TContaPoupanca)
	fmt.Println("")
	fmt.Println("Comparação de listas!")
	ImprimeLista(ListaEncadeadaContas)
	ImprimeLista(ListaEncadeadaContas2)

	var compListas int
	compListas = comparaDuasListas(ListaEncadeadaContas, ListaEncadeadaContas2)
	fmt.Println("Resultado comparação:", compListas)

	ListaEncadeadaContas2.AdicionaFim(3, 1, 0, TContaBancaria)
	compListas = comparaDuasListas(ListaEncadeadaContas, ListaEncadeadaContas2)
	fmt.Println("Resultado comparação:", compListas)
}
