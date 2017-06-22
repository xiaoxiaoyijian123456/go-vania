package go_vania

import (
	"github.com/draffensperger/golp"
	"errors"
)

var Error_Invalid_Param = errors.New("Invalid input param")

func FairDistributor(targets, objects []interface{}, weights [][]float64, maximize bool) (map[interface{}][]interface{}, float64, error) {
	n_t := len(targets)
	n_o := len(objects)
	if n_t <= 0 || n_o <= 0 || len(weights) != n_t {
		return nil, 0, Error_Invalid_Param
	}
	objFnVector := []float64{}
	for _, v := range weights {
		if len(v) != n_o {
			return nil, 0, Error_Invalid_Param
		}
		objFnVector = append(objFnVector, v...)
	}

	n := n_t * n_o
	lp := golp.NewLP(0, n)
	v := make([][]float64, n_o)
	for j := 0; j < n_o; j++ {
		v[j] = make([]float64, n)
		for i := 0; i < n_t; i++ {
			v[j][j+i*n_o] = 1
		}
		lp.AddConstraint(v[j], golp.EQ, 1)
	}
	lp.SetObjFn(objFnVector)
	if maximize {
		lp.SetMaximize()
	}

	lp.Solve()
	vars := lp.Variables()
	ret := make(map[interface{}][]interface{})
	for i, target := range targets {
		target_objs := []interface{}{}
		for j, obj := range objects {
			if vars[i*n_o+j] > 0 {
				target_objs = append(target_objs, obj)
			}
		}
		ret[target] = target_objs
	}
	return ret, lp.Objective(), nil
}
