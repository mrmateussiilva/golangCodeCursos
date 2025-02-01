## Estruturas de Controles

### If com init

```go
func main() {
	//A varivel i tem escopo apenas do bloco da condição e das outras condições
	if i := numeroAleatorio(); i > 5 { // Suporte no Switch
		// Consigo acessar a varivel I
		fmt.Println("Ganhou")
	} else {
		// Consigo acessar a varivel I
		fmt.Println("Perdeu")
	}
// Não Consigo acessar a varivel I
	fmt.Println(i)
}
```

obs: A variavel usada no incializador tem apenas escopo do bloco da condição

## Loops

GO por ser minimalista só possui um tipo de loop que no caso é o FOR

Mas isso não significa que ela é limitada por isso, pelo contrario isso torna a linguagem bem simples

Podemos usar ele como uma unica expressão

```go
var c int = 0
	for c <= 10 {
		fmt.Println(c)
		c++
	}
```

Maneira clássica em todas as linguagens 

```go
for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
```

Se precisa de um loop infinito tambem conseguimos

```go
for {
		fmt.Println("Esse código é infinito...")
		time.Sleep(time.Second)
	}
```

Tambem podemos usar as instruções “continue” e “break" dentro do loop for

Tambem podemos iterar sobre estruturas usando o for 

### Switch

O switch não compara uma expressão e sim um valor com outros valores

O switch em go comparado com outras linguagens é melhor pois ele funciona assim se aquele valor macth com outro valor ele automaticamente entra nesse bloco o executa o codigo que tem lá e depois sai do switch 

Em outras linguagens você precisa deixar explicto que depois que ele entra nesse block do case ele precisa sair do switch usando break

se você não usar o break ele vai entrar no case e sair executando todos os outros

Javascript

```jsx
nota = 10
switch (nota){
    case 10:
        console.log("A")
        break
    case 5:
       console.log("B")
        break
    case 2:
        console.log("c")
        break
    default:
        console.log("Nota invalida...")
        break
    
}
```

GO

```go
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // SE der match entra aqui mas não encera o switch
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Invalida"
	}
```

Obs: A palavra 
`fallthrough` 

É usada pra fazer o comportamento do switch em outras linguagens, ou seja, 

entra dentro dese case porem sai e executa os outros cases

```go
time.Now() // Pega a data e hora atual
```
---

### Arrays

O array é um estrutura de dados homogênea e como tamanho fixo, ou seja 

sempre precisamos definir o tipo dos dados armazenados naquele array e a quantidade de dados

- Tamanho fixo
- Dados Homogêneo

Sintaxe:

```go

// Array com 3 posições do tipo float64
var notas [3]float64

// Array com 3 posições do tipo float64
notas := [...]float64{1.5,8.9,9.1}
```

### Obs:

Pode usar a sintaxe de … para deixar o compilador contar quantos elementos possui aquela estrutura e inferir o tamanho durante a compilação.

Caso a gente esqueça esses … o compiladora vai tratar isso como um slice e não com um array

Código da aula

```go
	//	Homogênea (mesmo tipo) e estática (tamanho fixo)
	// O array é um estrutura indexavel inicinaod com indice 0
	var notas [3]float64

	//fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 6.2, 9.5
	//notas[3] = 10 // Erro pois o index 3 não existe no array
	var total float64
	for i := 0; i < len(notas); i++ {
		total = notas[i]
	}
	//	como total é do tipo float64 pra poder relizar a divisão do total
	// pela media eu preciso fazer o casting de tipo do tamanho do array
	media := total / float64(len(notas))
	//fmt.Println(notas)
	fmt.Printf("A media das notas é %.2f\n", media)
```

Pra percorremos um array podemos usar o for range que é meio que um foreach do javascript e de outras langs

Sintaxe:

```go
//
for index,value := range array{}
for index := range array{}
for _,value := range array{}

```

Código da aula:

```go
func main() {
	/*Como estamos usandos a syntax de ... pra falar o tamanho do array
	durante o processo de compilação o compilador vai contar quantos elementos
	existe nesse array
	*/
	// OBS: Se esquecer esses ... go vai tratar isso como um slice e não como array

	numeros := [...]int{1, 2, 3, 4, 5} // size 5
	//Percorrendo o array com for range
	fmt.Println("Pegando tanto o index quanto o value")
	for index, value := range numeros {
		fmt.Printf("Posição %d, valor %d\n", index, value)
	}
	/*Quando usamos o _(underline) assim dizemos pro go,
	"ingora esse valor ou pacote pois não uso ele

	*/
	fmt.Println("Pegando apenas o valor e ignorando o index")
	for _, value := range numeros {
		fmt.Println(value)
	}

	// Tbm podemos pegar apenas o indice da estrutura
	fmt.Println("Pegando apenas o index")
	for index := range numeros {
		fmt.Println(index)
	}
}
```

## Slices

Podemos entender ele como se você uma fatia de um array ou seja temos um array com 10 posições, ai eu vou lá e crio um slice que aponta da posição 3 até a posição 6

- Slice não é um array
- Podemos imaginar um slice como: tamanho e um ptr pra um elemento de um array
- podemos fazer slice de outros slice
- Ele não gera um novo array

O slice normalmente é mais leve que trabalhar diretamente com array pois ele trabalha com fatias

Sintaxe:

```go
array1 := [10]int{1,2,3,4,5,6,7,8,9,10}
slice := array1[2:3] // saida: 3
slice2 := []float32{1,3}
```

Código da aula:

```go
func main() {

	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	//fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [...]int{1, 2, 3, 4, 5}

	// indices 1,2
	s2 := a2[1:3]
	fmt.Println(a2, s2) // novo slice que aponta do indice 1 até o indice 3

	//indice 0,1,2
	s3 := a2[:3] // novo slice do indice 0 até o indice 3
	fmt.Println(a2, s3)

	//indices 0
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
```

anotação da aula:

```go
/*
	O slice não é um array, slice define um pedaço do array
	ou seja podemos ter um array base com 20 posições e podemos ter 2 slices
	que apontam pra regigões diferente do array

	Ele não gera um novo array ele simplesmente cria um ponteiro do
	1° indice da slice até fim do slice

	podemos imaginar o slice como: tamanho e um ptr pra um elemento de um array
	podemos fazer slice de outros slices
*/
```

### Slice com Make

O make é um função bem legal que nos permite criar um slice de maneira fácil, pode debaixo dos panos ele usa array

SIintaxe:

```go
// make(tipo,tamanho)
slice := make([]int,10) // 1° maneira de usar 

// make(tipo,tamanho,capicidade)
slice = make([]int,10,20) // 2° maneira
```

A forma mais básica de usar o make é passando apenas o tamanho do slice,

quando passamos o 3° parametro pra função o make faz isso

“vou criar um slice com n elementos porem o array interno que você vai criar vai n elementos”

Código da aula:

```go
	/*
		make([]tipo,tamanho,capacidade)
		O make internamente vai trabalhar com array
	
	*/
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Printf("Slice %d, \ntamanho: %d, \ncapicidade %d\n", s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 15)
	fmt.Println(s, len(s), cap(s))

```