package expr

import (
	"fmt"
	"github.com/antonmedv/expr"
	"testing"
)

func TestExpr1(t *testing.T) {
	env := map[string]interface{}{
		"greet":   "Hello, %v!",
		"names":   []string{"world", "you"},
		"sprintf": fmt.Sprintf,
	}

	code := `sprintf(greet, names[0])`

	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

type Tweet struct {
	Len int
}

type Env struct {
	Tweets []Tweet
}

func TestExpr2(t *testing.T) {
	code := `all(Tweets, {.Len <= 240})`

	program, err := expr.Compile(code, expr.Env(Env{}))
	if err != nil {
		panic(err)
	}

	env := Env{
		Tweets: []Tweet{{42}, {98}, {69}},
	}
	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

func TestExpr3(t *testing.T) {
	output, err := expr.Eval("greet >= name", map[string]interface{}{"greet": 1, "name": 2})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("%v\n", output)
}
