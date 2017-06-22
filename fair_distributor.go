package go_vania

import (
	"github.com/draffensperger/golp"
)

func FairDistributor(targets, objects []interface{}, weights [][]float64, maximize bool) (map[interface{}][]interface{}, float64, error) {
	n := len(targets) * len(objects)
	lp := golp.NewLP(0, n)
	v := make([][]float64, len(objects))
	for j := 0; j < len(objects); j++ {
		v[j] = make([]float64, n)
		for i := 0; i < len(targets); i++ {
			v[j][j+i*len(objects)] = 1
		}
		lp.AddConstraint(v[j], golp.EQ, 1)
	}
	vector := []float64{}
	for _, v := range weights {
		for _, v2 := range v {
			vector = append(vector, v2)
		}
	}
	lp.SetObjFn(vector)
	if maximize {
		lp.SetMaximize()
	}

	lp.Solve()
	vars := lp.Variables()
	ret := make(map[interface{}][]interface{})
	for i, target := range targets {
		target_objs := []interface{}{}
		for j, obj := range objects {
			if vars[i*len(objects)+j] > 0 {
				target_objs = append(target_objs, obj)
			}
		}
		ret[target] = target_objs
	}
	return ret, lp.Objective(), nil
}
