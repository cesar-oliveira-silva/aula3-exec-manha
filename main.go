package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// exercicio 1
type produto struct {
	id         int
	nome       string
	preco      float64
	quantidade int
}

type listaProduto struct {
	produtos []produto
}

func main() {

	fmt.Println("func main")

	//exercicio 1
	var loja1 listaProduto
	amaciante := criaProduto(1, "amaciante", 5.0, 100)
	desinfetante := criaProduto(2, "desinfetante", 3.0, 50)
	sabonete := criaProduto(3, "sabonete", 1.0, 70)
	loja1.addProduto(amaciante)
	loja1.addProduto(desinfetante)
	loja1.addProduto(sabonete)

	//debug
	// for _, prod := range loja1.produtos {
	// 	fmt.Println(prod.nome)
	// }

	// cabecalho
	head := []byte("id;nomeProduto;Preco;Quantidade\n")
	errArq := os.WriteFile("./ListaProdutos.txt", head, 0777)
	if errArq != nil {
		fmt.Println("deu erro no WriteFile cabecalho")
	}
	// produtos
	for _, prod := range loja1.produtos {
		//linha := []byte(fmt.Sprint(prod.id, ";", prod.nome, ";", prod.preco, ";", prod.quantidade, "\n"))
		linha := []byte(fmt.Sprintf("%v;%s;%.2f;%v \n", prod.id, prod.nome, prod.preco, prod.quantidade))

		arq, err := os.ReadFile("./ListaProdutos.txt")
		if err != nil {
			fmt.Println("deu erro no ReadFile")
		}
		for _, dado := range linha {
			arq = append(arq, dado)
		}
		//arq = append(arq, linha)
		os.WriteFile("./ListaProdutos.txt", arq, 0777)
		if errArq != nil {
			fmt.Println("deu erro no WriteFile linha")
		}
	}
	// exercicio 2
	fmt.Println("lendo arquivo: \n")
	leeArquivo()

}

// exercicio1
func criaProduto(ide int, name string, price float64, qnt int) produto {
	return produto{ide, name, price, qnt}
}
func (l *listaProduto) addProduto(p produto) {
	fmt.Printf("adicionando o produto %s\n", p.nome)
	l.produtos = append(l.produtos, p)
}

// exercicio 2
func leeArquivo() {

	// esse exercicio ta muito complexo em relacao a aula. nao vimos nenhuma dessas funcoes de scan e tabwriter
	//colando da resolucao do exercicio:

	arquivo, err := os.Open("./ListaProdutos.txt")
	if err != nil {
		fmt.Println("deu erro no Open")
	}
	defer arquivo.Close()

	w := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t', tabwriter.AlignRight)

	scanner := bufio.NewScanner(arquivo)

	scanner.Scan()

	cabecalho := strings.Split(scanner.Text(), ";")

	for _, c := range cabecalho {
		fmt.Fprintf(w, "%s\t", c)
	}

	fmt.Fprintln(w)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ";")
		for _, v := range values {
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
	}

	w.Flush()

	//tentativa de resolucao com o visto em aula
	// arquivo, err := os.ReadFile("./ListaProdutos.txt")
	// if err != nil {
	// 	fmt.Println("deu erro no ReadFile")
	// }

	// arquivoLido := string(arquivo)

	// fmt.Println(string(arquivo))

	// var formatado string
	//  for _, runa := range arquivoLido {
	// 	if runa == ";"{

	// 	}

	//  }

}
