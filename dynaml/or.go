package dynaml

import (
	"fmt"
	"reflect"
)

type OrExpr struct {
	A Expression
	B Expression
}

func (e OrExpr) Evaluate(binding Binding, locally bool) (interface{}, EvaluationInfo, bool) {
	a, infoa, ok := e.A.Evaluate(binding, false)
	if ok && !infoa.Undefined {
		if reflect.DeepEqual(a, e.A) {
			return nil, infoa, false
		}
		if isExpression(a) {
			return e, infoa, true
		}
		return a, infoa, true
	}

	b, infob, ok := e.B.Evaluate(binding, false)
	info := infoa.Join(infob)
	info.Undefined = infob.Undefined
	return b, info, ok
}

func (e OrExpr) String() string {
	return fmt.Sprintf("%s || %s", e.A, e.B)
}
