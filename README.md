# Gematria

This package provides you the ability to transform a string into 3 numbers called Gematria. The three numbers are: 

| Kind                | Example String | Example Score |
|---------------------|----------------|---------------|
| `GemScore.English`  | `Andrei`       | `306`         |
| `GemScore.Jewish`   | `Andrei`       | `139`         |
| `GemScore.Simple`   | `Andrei`       | `51`          |
| `GemScore.Mystery`  | `Andrei`       | `2264`        |
| `GemScore.Majestic` | `Andrei`       | `153`         | 
| `GemScore.Eights`   | `Andrei`       | `338`         | 

When tools like [Gematrix.org](https://gematrix.org?utm_source=projectapario&word=andrei) are used they return
interesting results that share the similar scores with their results. For example: 

| Kind    | Score | Gematrix Entry |
|---------|-------|----------------|
| English | 306   | Michael        |
| English | 306   | BRICS          |
| English | 306   | Batman         |
| English | 306   | GESARA         |
| English | 306   | Rome           | 
| English | 306   | Micheal        |
| English | 306   | Elite          |


| Kind   | Score | Gematrix Entry |
|--------|-------|----------------|
| Jewish | 139   | Google         |
| Jewish | 139   | Sin            |
| Jewish | 139   | Golden Age     |
| Jewish | 139   | Math           |
| Jewish | 139   | Nikolai        |
| Jewish | 139   | Agora          |


| Kind   | Score | Gematrix Entry |
|--------|-------|----------------|
| Simple | 306   | Michael        |
| Simple | 306   | BRICS          |
| Simple | 306   | Batman         |
| Simple | 306   | GESARA         |
| Simple | 306   | Rome           | 
| Simple | 306   | Micheal        |
| Simple | 306   | Elite          |

The word `Andrei` in Gematria is expressed as English `Michael`. The Jewish expression of `Andrei` is the `Golden Age`. 
I am building multi-generational wealth and these utilities are part of how I am doing that. You can either program or
be programmed, and I choose to program. Go is such a powerful programming language because of the memetics that it 
natively supports programmers to use. 

Religious scholars have used Gematria as a mechanism for seeing relationships between words. Their relationships are
the numbers that they share in common with each other. Those numbers are governed by the laws of 3 6 and 9. The 
power of 3 6 and 9 is very real and you are encouraged to use it. Therefore, this package was created... so that you
may use it.

# Package Installation

You will need to download this package to your Go project should you wish to use it. 

```bash
go get -u github.com/andreimerlescu/gematria
```

# Usage

Here is an example program that uses this package.

```go
package main

import (
	"fmt"
	gematria "github.com/andreimerlescu/gematria"
)

func main(){
	name := "Andrei"
	gematria_score := gematria.NewGemScore(name)
	fmt.Printf("name = %v ; gematria = %v", name, gematria_score)
}
```

# Future Proof

This package has no dependencies and will not require to be updated in the future given that Gematria is effectively
a constant as unchangeable as the value of Pi. Therefore, the long term utilization of this package can safely be 
integrated into many types of projects.

# Contribution

If you wish to improve this code, please fork the repository and submit your pull request with your changes.

# License

This package is licensed with the Apache 2.0 license so you can use it in your applications while giving proper
credit where credit is due for creating this package in the first place. 