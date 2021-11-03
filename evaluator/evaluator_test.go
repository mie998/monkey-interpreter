package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"reflect"
	"testing"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, idx int, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("testcase:%v, object is not Integer. got=%T (%+v)", idx, obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("testcase:%v, object has wrong value. got=%d, want=%d", idx, result.Value, expected)
		return false
	}

	return true
}

func testStringObject(t *testing.T, idx int, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("testcase:%v, Evaluated object is not String. got=%T (%+v)", idx, obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("testcase:%v, object has wrong value. got=%s, want=%s", idx, result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, idx int, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, idx int, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("testcase:%v, object is not NULL. got=%T (%+v)", idx, obj, obj)
		return false
	}
	return true
}

func testArrayObject(t *testing.T, idx int, obj object.Object, expected []interface{}) bool {
	result, ok := obj.(*object.Array)
	if !ok {
		t.Errorf("testcase:%v, Evaluated object is not String. got=%T (%+v)", idx, obj, obj)
		return false
	}
	if len(result.Elements) != len(expected) {
		t.Errorf("testcase:%v, Evaluated object length isn't equal. got=%v (%v)", idx, len(result.Elements), len(expected))
		return false
	}
	for i, elem := range expected {
		v := reflect.ValueOf(elem)
		switch v.Kind() {
		case reflect.Int:
			testIntegerObject(t, i, result.Elements[i], elem.(int64))
		case reflect.String:
			testStringObject(t, i, result.Elements[i], elem.(string))
		case reflect.Bool:
			testBooleanObject(t, i, result.Elements[i], elem.(bool))
		case reflect.Array:
			testArrayObject(t, i, result.Elements[i], elem.([]interface{}))
		default:
			t.Errorf("testcase:%v, this valiable type is not suppoerted. got=%v (%+v)", idx, result.Elements[i], result.Elements[i])
			return false
		}
	}

	return true
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"1 + 1 * 1", 2},
		{"1 - 1 + 1", 1},
		{"1 + 3 / 1", 4},
		{"(1 + 1) / 2 + 3", 4},
		{"1 + 4 / 2 - 1", 2},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, i, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 > 2", false},
		{"2 > 1", true},
		{"1 < 2", true},
		{"2 < 1", false},
		{"1 == 1", true},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, i, evaluated, tt.expected)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello, World!"`
	evaluated := testEval(input)
	estimated := "Hello, World!"
	str, ok := evaluated.(*object.String)

	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != estimated {
		t.Errorf("String has wrong value. got=%q, estimated=%q", str.Value, estimated)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + ", " + "World!"`
	evaluated := testEval(input)
	estimated := "Hello, World!"
	str, ok := evaluated.(*object.String)

	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != estimated {
		t.Errorf("String has wrong value. got=%q, estimated=%q", str.Value, estimated)
	}
}

func TestEvalStringLogicalExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`"a" < "b"`, true},
		{`"a" > "b"`, false},
		{`"a" == "b"`, false},
		{`"a" != "b"`, true},
		{`"a" < "a"`, false},
		{`"a" > "a"`, false},
		{`"a" == "a"`, true},
		{`"a" != "a"`, false},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, i, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, i, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 } else { 20 }", 20},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (2 > 1) { if (2 > 1) { return 10; } return 1; }", 10},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, i, evaluated, int64(integer))
		} else {
			testNullObject(t, i, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 2 * 5;", 10},
		{"return 10; 0;", 10},
		{"0; return 10; 0;", 10},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, i, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input         string
		expectMessage string
	}{
		{"5 + true", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if (2 > 1) { if (2 > 1) { return true + false; } return 1; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "identifier not found: foobar"},
		{`"hello" - "world"`, "unknown operator: STRING - STRING"},
		{`"hello" + 1`, "type mismatch: STRING + INTEGER"},
		{`{"name": "Monkey"}[fn(x) { x }];`, "unusable as hash key: FUNCTION"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T (%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectMessage {
			t.Errorf("Got wrong error message. expected=%q got=%q", tt.expectMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for i, tt := range tests {
		testIntegerObject(t, i, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
	}

	for i, tt := range tests {
		testIntegerObject(t, i, testEval(tt.input), tt.expected)
	}
}

func TestEnclosingEnvironments(t *testing.T) {
	input := `
let first = 10;
let second = 10;
let third = 10;

let ourFunction = fn(first) {
  let second = 20;

  first + second + third;
};

ourFunction(20) + first + second;`

	testIntegerObject(t, 0, testEval(input), 70)
}

func TestClosures(t *testing.T) {
	input := `
let newAdder = fn(x) {
	fn(y) { x + y };
};

let addTwo = newAdder(2);
addTwo(2);`

	testIntegerObject(t, 0, testEval(input), 4)
}

func TestBuildinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("unchi")`, 5},
		{`len("")`, 0},
		{`len(1)`, object.Error{Message: "argument to `len` not supported, got INTEGER"}},
		{`len("one", "two")`, object.Error{Message: "wrong number of arguments. got=2, want=1"}},
		{`len([1,2,3])`, 3},
		{`len([])`, 0},
		{`first("one")`, "o"},
		{`first("")`, nil},
		{`first([0,1,2])`, 0},
		{`first([])`, nil},
		{`last("one")`, "e"},
		{`last("")`, nil},
		{`last([0,1,2])`, 2},
		{`last([])`, nil},
		{`rest("one")`, "ne"},
		{`rest("")`, nil},
		{`rest([0,1,2])`, [2]int{1, 2}},
		{`rest([])`, nil},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		if tt.expected == nil && testNullObject(t, i, evaluated) {
			return
		}
		v := reflect.ValueOf(tt.expected)
		switch v.Kind() {
		case reflect.Int:
			testIntegerObject(t, i, evaluated, int64(tt.expected.(int)))
		case reflect.String:
			testStringObject(t, i, evaluated, string(tt.expected.(string)))
		case reflect.Array:
			testArrayObject(t, i, evaluated, tt.expected.([]interface{}))
		default:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("testcase:%v, object is not Error. got=%T (%+v)", i, evaluated, evaluated)
				continue
			}
			if reflect.DeepEqual(tt.expected, errObj) {
				t.Errorf("testcase:%v, wrong error message. expected=%q, got=%q", i, tt.expected, errObj)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, 0, result.Elements[0], 1)
	testIntegerObject(t, 0, result.Elements[1], 4)
	testIntegerObject(t, 0, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"let i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, i, evaluated, int64(integer))
		} else {
			testNullObject(t, i, evaluated)
		}
	}
}

func TestHashLiterals(t *testing.T) {
	input := `let two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, int(expectedValue), pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, i, evaluated, int64(integer))
		} else {
			testNullObject(t, i, evaluated)
		}
	}
}
