package main

import (
	"fmt"
)


type filme struct{
	nome	string
	preço	float64
	desc	string
}

var carrinho []filme

// certo, porem redundante
//var lista []filme = []filme{} 

// certo e simplificado
var lista []filme{
	filme{
		nome: "shazam",
		preço: 22.50,
		desc: "um guri q pode ficar adulto e solta raio pela mão",
	},
	filme{
		nome: "homem Aranha",
		preço: 24.00,
		desc: "após o início da pandemia, ninguém saiu de casa",
	},
	filme{
		nome: "homem de ferro",
		preço: 21.00,
		desc: "'Eu sou o homem de ferro', e morreu",
	},
}



func main(){
	erros := map[string]bool{
		"ja_é_meu": false,
	}
	for{
		// limpar a tela
		print("\033[H\033[2J")
		if erros["ja_é_meu"] {
			fmt.Println("\033[91mEsse Item já é seu!\033[m")
			erros["ja_é_meu"] = false
		}
		
		
		for ind, val := range lista {
			var vigia bool
			for _, v_car := range carrinho{
				if v_car == val{
					vigia = true
				}
			}
			if vigia{
				fmt.Println("\033[34m[",ind,"]",val.nome,"\033[m")
				fmt.Printf("->\033[94m\tpreço: R$%.2f\033[32m COMPRADO✓\033[m\t\n",val.preço)
			} else {
				fmt.Println("\033[34m[",ind,"]",val.nome,"\033[m")
				fmt.Printf("->\033[92m\tpreço: R$%.2f\033[m\t\n",val.preço)
			}
			println()
		}
		
		
		if len(carrinho) > 0 {
			fmt.Println("Meu carrinho:\n")
			for _,v := range carrinho{
				fmt.Printf("%s\n",v.nome)
			}
				
		}
		
		
		var opt int
		for{
			fmt.Print("Qual vc deseja comprar: [número]")
			err, _ := fmt.Scan(&opt)
			if err == 1 && opt < len(lista) {
				break
			} else {
				fmt.Println("\033[91mOpção inválida!\033[m")
			}
		}
		
		
		vigia_2 := 0
		for _,v := range carrinho{
			if v != lista[opt]{
				vigia_2++
			}
		}
		if vigia_2 == len(carrinho){
			carrinho = append(carrinho,lista[opt])
		} else {
			erros["ja_é_meu"] = true
		}
		
		
		var exit int 
		for {
			fmt.Println("[ 1 ] sair")
			fmt.Println("[ 2 ] continuar")
			fmt.Print("oq deseja fazer: ")
			err, _ := fmt.Scan(&exit)
			if err == 1 {
				break
			} else {
				fmt.Println("\033[91mOpção inválida!\033[m")
			}
		}
		if exit == 1{
			fmt.Println("flw, até breve!")
			break
		} else {
			// só por boa pratica
			continue
		}
	}
}