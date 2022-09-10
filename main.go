package main

import "fmt"

type TADRegistro struct {
	Elemento int
}

func main() {

	var tamanho int
	var n int
	//var elemento int

	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)
	tamanho = n / 2
	fmt.Println("A tabela será capaz de armazenar:", tamanho, " valores.")

	var registros = []TADRegistro{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	}

	var tabela = TADTabelaHash{}
	var chaves = make([]int, len(registros))
	var i int
	i = 0
	var registro TADRegistro

	for i, registro = range registros {
		println("Adicionar elemento:", registro.Elemento)
		chaves[i] = tabela.Adiciona(registro, tamanho)
		println("elemento:", registro.Elemento, " - com chave:", chaves[i])
		println(registro.Elemento, " - ", chaves[i])

	}

	tabela.Remove(1, tamanho)

	var RegistrosBusca []TADRegistro
	var RegistroBusca TADRegistro
	var chaveBuscada int
	chaveBuscada = 2
	RegistrosBusca = tabela.Recupera(chaveBuscada)
	println("Chave Buscada:", chaveBuscada)
	println("Essa chave possui: ", len(RegistrosBusca), " registros")

	var elementoBuscado = 2
	RegistroBusca = tabela.Busca(elementoBuscado, tamanho)
	println("Elemento Buscado:", RegistroBusca.Elemento)

	/*var chave int
	for _, chave = range chaves {
		registro = tabela.Recupera(chave)

	}*/

	fmt.Println("O tamanho do banco é de :", len(registros))
	fmt.Println("Foram armazenados  :", len(tabela.items), " na tabela hash")

	fmt.Println("Apagar toda tabela!")
	tabela.RemoveTudo()

	//RegistroBusca = tabela.Busca(elementoBuscado, tamanho)
	//println("Elemento Buscado:", RegistroBusca.Elemento)
	fmt.Println("Foram armazenados  :", len(tabela.items), " na tabela hash")

}
