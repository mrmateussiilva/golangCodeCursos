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
