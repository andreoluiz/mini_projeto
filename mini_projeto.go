package main

import (
	"fmt"
	"strings"
)

/*criando o tipo Queue para usar posteriormente*/
type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

/*Criando um nó para poder estar sempre atualizando e apontando para o primeiro elemento da fila*/
type Node struct {
	Value1 int
	Next   *Node
}

/* Função para checar tamanho da fila*/
func (q *Queue) tamanho() int {
	return q.Size
}

/* Criando a função de Enqueue sem utilização de for utlizando nó para apontar a posição*/
func (q *Queue) Enqueue(preco int) {
	node := Node{Value1: preco}
	if q.Head == nil {
		q.Head = &node
		q.Tail = &node
	} else {
		q.Tail.Next = &node
		q.Tail = &node

	}
	q.Size++
	return
}

/* Criando função de Dequeue sem utilização de for utilizando nó para atualizar a primeira e a ultima posição da fila*/
func (q *Queue) Dequeue() int {
	if q.Head == nil {
		q.Tail = nil
		return 0
	}
	node := q.Head
	q.Head = q.Head.Next
	q.Size--

	return node.Value1
}

func main() {

	queue := Queue{}
	var escolha string = "a"
	var quant int
	var preco int
	var lucro int = 0

	for escolha != "FIM" {
		fmt.Println("Informe se quer COMPRAR ou VENDER, em seguida a quantidade de ações e o valor.\nCaso deseje sair digite FIM")
		fmt.Scanf("%s%d%d\n", &escolha, &quant, &preco)

		escolha = strings.ToUpper(escolha)

		if escolha == "COMPRAR" {
			for quant != 0 {
				queue.Enqueue(preco)
				quant--
			}
		}
		if escolha == "VENDER" {
			if queue.tamanho() < quant {
				fmt.Println("Quantidade de ações vendidas maior que compradas!")
				continue
			}
			for quant != 0 {
				preco_acao := queue.Dequeue()
				lucro += preco - preco_acao
				quant--

			}
			fmt.Println(lucro)
		}
	}
}
