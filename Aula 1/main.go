package main

import "fmt"

type Aluno struct {
	RA   string `json:"ra"`
	Nome string `json:"nome"`
}

func (aluno *Aluno) getNome() string {
	return aluno.Nome
}

func main() {
	/*var variavel1 string
	variavel2 := "ola"
	var (
		variavel3 string
		variavel4 int
	)
	const variavel5 = 5

	fmt.Println(variavel1, " ", variavel2, " ", variavel3, " ", variavel4, " ", variavel5)*/

	//Função atribuida a uma variavel
	/*funcao := func() {
		fmt.Println("Hello world")
	}
	funcao()

	resultado, ok := Divisao(3, 0)

	if ok == false {
		fmt.Println("Não é possivel dividir por 0")
	} else {
		fmt.Println(resultado)
	}

	fmt.Println(Somar(1, 5, 5, 34, 2, 1, 2))

	//Função Anonima
	fmt.Println(func(a int, b int) int {
		return a  * b
	}(3, 4))

	fmt.Println(Fibonacci(50))*/

	/*a1 := Aluno{RA: "200126", Nome: "Fábio"}
	fmt.Println("RA: ", a1.RA, "\nNome: ", a1.Nome)

	var a2 Aluno

	a2.RA = "200146"
	a2.Nome = "Pedro"
	fmt.Println(a2.getNome())

	var a3 Aluno
	fmt.Println(a3)

	a4 := Aluno{Nome: "Victor", RA: "200154"}
	//Marshal = Struct -> Json
	if alunoJson, erro := json.Marshal(a4); erro != nil {
		fmt.Println("Erro: ", erro.Error())
		return
	} else {
		stringJson := string(alunoJson)
		fmt.Println(stringJson)

		var alunoNovo Aluno
		//Unmarshal = Json -> Struct
		if newError := json.Unmarshal(alunoJson, &alunoNovo); newError != nil {
			fmt.Println("Erro: ", erro.Error())
			return
		} else {
			fmt.Println(alunoNovo)
		}
	}*/

	/*array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := array[:5]
	fmt.Println(array)
	fmt.Println(slice1)

	dicionario := map[int]string{}

	dicionario[0] = "Zero"
	dicionario[1] = "Zebra"
	dicionario[2] = "Girafa"

	fmt.Println(dicionario)

	dicio2 := map[string]string{
		"1": "um",
		"2": "dois",
	}

	delete(dicio2, "1")

	fmt.Println(dicio2)*/

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Recuperado")
		}
	}()

	if 1 == 1 {
		// Para o programa
		panic("1 é um igual a 1")
	}

}

// Função com retorno nomeado
func Divisao(a int, b int) (c int, ok bool) {
	if b == 0 {
		c = -1
		ok = false
	} else {
		c = a / b
		ok = true
	}

	return
}

// Função com n parametros
func Somar(numeros ...int) (c int) {
	c = 0

	for _, valor := range numeros {
		c += valor
	}
	return
}

// função recursiva
func Fibonacci(number int) int {
	if number == 0 {
		return 0
	}

	if number == 1 {
		return 1
	}

	return Fibonacci(number-1) + Fibonacci(number-2)
}
