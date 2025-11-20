# Day 1
## Project Setup
Installed Go via Brew: ```brew install go.```

```
mkdir introToGoLang
cd introToGoLang
go version
touch main.go
ll 
go mod init introGoLang
go mod tidy
```

## Running the project
### To run the project use two terminal windows:

One to start server<br/>
```go run main.go```<br/>

Another for requests<br/>
```curl --get localhost:3009```

## First lines of code
### Started the file with a basic structure:
```
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hey Digital Development"))
	})

	log.Printf("Starting server on port %s", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
```

### Replaced it with:
```
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	router := chi.NewRouter()
	router.Get("/", root)

	log.Printf("Starting server on port %s", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}

type Animal struct {
	Color string
	Paws  int
}

func root(w http.ResponseWriter, r *http.Request) {
	animal := Animal{
		Color: "brown",
		Paws:  4,
	}
	a, _ := json.Marshal(animal)
	w.Write(a)
}
```
