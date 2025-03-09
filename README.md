# Gematria

This package provides you the ability to transform a string into 6 numbers called Gematria. The three numbers are: 

| Kind                | Example String | Example Score |
|---------------------|----------------|---------------|
| `Gematria.English`  | `Andrei`       | `306`         |
| `Gematria.Jewish`   | `Andrei`       | `139`         |
| `Gematria.Simple`   | `Andrei`       | `51`          |
| `Gematria.Mystery`  | `Andrei`       | `2264`        |
| `Gematria.Majestic` | `Andrei`       | `153`         | 
| `Gematria.Eights`   | `Andrei`       | `338`         | 

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
	"github.com/andreimerlescu/gematria"
)

func main(){
	name := "yah i love yahuah and yah i am saint andrei"
	gematria := gematria.FromString(name)
	fmt.Printf("name = %v ; gematria = %v", name, gematria)
}
```

# Future Proof

This package has no dependencies and will not require to be updated in the future given that Gematria is effectively
a constant as unchangeable as the value of Pi. Therefore, the long term utilization of this package can safely be 
integrated into many types of projects.

# Additional Ciphers Added

The **EIGHTS** Cipher is a new addition to the __Gematria__ package. A begins at 3, B is 5, when combined AB you get 
3+5=8 which is C. Then D is 16 and it increases by 8 until Z.

The **MYSTERY** Cipher is a new addition to the __Gematria__ package. When developing this package, I asked the spirit
to guide me and the numbers I chose was the numbers I chose. They are indeed a mystery. 

The **MAJESTIC** Cipher is a new addition to the __Gematria__ package. Starting at 3, incrementing by 3, it gets 
to Z with a value of 78 (3*26 letters). JKR is 30 33 54 which simplifies to 999. JFK 30 18 33 which simplifies to 999.
Therefore JKR === JFK in the **MAJESTIC** _UNDERSTANDING_ of what these numbers actually mean. Nikola Tesla understood
it, and its worthy of me to program the ciphers using these numbers. They are worthy of studying too.

# Contribution

If you wish to improve this code, please fork the repository and submit your pull request with your changes.

# License

This package is licensed with the Apache 2.0 license so you can use it in your applications while giving proper
credit where credit is due for creating this package in the first place. 