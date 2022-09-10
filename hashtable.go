package main

import "fmt"

type TADTabelaHash struct {
	items map[int][]TADRegistro
}

func hash(Elemento int, tamanho int) (chave int) {
	chave = Elemento % tamanho
	return chave
}

func (th *TADTabelaHash) Adiciona(registro TADRegistro, tamanho int) int {
	var chave int
	chave = hash(registro.Elemento, tamanho)
	if th.items == nil {
		th.items = make(map[int][]TADRegistro)
	}

	if th.items[chave] == nil {
		fmt.Println("Bucket vazio, pode gravar!")
		th.items[chave] = append(th.items[chave], registro)

	} else {
		fmt.Println("Bucket cheio!, vou armazenar assim mesmo")
		th.items[chave] = append(th.items[chave], registro)
		//chave = -1
	}

	return chave
}

func (th *TADTabelaHash) Remove(elemento int, tamanho int) {
	var chave int
	chave = hash(elemento, tamanho)
	delete(th.items, chave)
}

func (th *TADTabelaHash) RemoveTudo() {

	for k := range th.items {
		delete(th.items, k)
	}
}

func (th *TADTabelaHash) Recupera(chave int) []TADRegistro {
	//println(th.items[chave][0].Elemento)
	return th.items[chave]

}

func (th *TADTabelaHash) Busca(Elemento int, tamanho int) TADRegistro {
	var chave int
	var indice_elemento int
	indice_elemento = -1
	var registros []TADRegistro
	chave = hash(Elemento, tamanho)
	registros = th.items[chave]

	for i, s := range registros {
		if s.Elemento == Elemento {
			indice_elemento = i
			println("Achei o elemento:", registros[i].Elemento)
		}
		//fmt.Println(i, s)
	}

	if indice_elemento == -1 {
		println("Não achei nenhum elemento!")
		indice_elemento = 0

	}

	return registros[indice_elemento]
}
