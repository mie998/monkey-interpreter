package ast

type ModifierFunc func(Node) Node

func Modify(n Node, f ModifierFunc) Node {
	switch node := n.(type) {
	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, f).(Statement)
		}
	case *ExpressionStatement:
		node.Expression, _ = Modify(node.Expression, f).(Expression)
	case *InfixExpression:
		node.Left, _ = Modify(node.Left, f).(Expression)
		node.Right, _ = Modify(node.Right, f).(Expression)
	case *PrefixExpression:
		node.Right, _ = Modify(node.Right, f).(Expression)
	case *IndexExpression:
		node.Left, _ = Modify(node.Left, f).(Expression)
		node.Index, _ = Modify(node.Index, f).(Expression)
	case *IfExpression:
		node.Condition, _ = Modify(node.Condition, f).(Expression)
		node.Consequence, _ = Modify(node.Consequence, f).(*BlockStatement)
		if node.Alternative != nil {
			node.Alternative, _ = Modify(node.Alternative, f).(*BlockStatement)
		}
	case *BlockStatement:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, f).(Statement)
		}
	case *ReturnStatement:
		if node.ReturnValue != nil {
			node.ReturnValue, _ = Modify(node.ReturnValue, f).(Expression)
		}
	case *LetStatement:
		if node.Value != nil {
			node.Value, _ = Modify(node.Value, f).(Expression)
		}
	}

	return f(n)
}
