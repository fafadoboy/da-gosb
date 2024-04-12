package utils

import (
	"errors"

	internalModels "github.com/fafadoboy/da-gosb/internal/chapter3/models"
	"github.com/fafadoboy/da-gosb/internal/models"
	"github.com/samber/lo"
)

func cloneAssignment[D models.Clonable](assignment map[string]D) map[string]D {
	copy := make(map[string]D, len(assignment))
	for k, v := range assignment {
		copy[k] = v.Clone().(D)
	}
	return copy
}

type CSP[V models.Hashable, D models.HashableAndClonable] struct {
	variables   []V
	domains     map[string][]D
	constraints map[string][]internalModels.Constraint[V, D]
}

func (o *CSP[V, D]) AddConstraint(constraint internalModels.Constraint[V, D]) error {
	for _, variable := range constraint.Variables() {
		if _, ok := lo.Find[V](o.variables, func(item V) bool { return item.Hash() == variable.Hash() }); !ok {
			return errors.New("LookupError: variable in constrains should have a domain")
		}
		o.constraints[variable.Hash()] = append(o.constraints[variable.Hash()], constraint)
	}
	return nil
}

func (o *CSP[V, D]) Consistent(variable V, assignment map[string]D) bool {
	for _, constraint := range o.constraints[variable.Hash()] {
		if !constraint.Satisfied(assignment) {
			return false
		}
	}
	return true
}

func (o *CSP[V, D]) BacktrackingSearch(assignment map[string]D) map[string]D {
	// The base case is having found a valid assignment for every variable.
	// Stop searching and return
	if len(assignment) == len(o.variables) {
		return assignment
	}

	// To select a new variable whose domain we will explore, we simply go through all the
	// variables and find the first that does not have an assignment.
	unassigned := lo.Filter[V](o.variables, func(item V, index int) bool {
		_, ok := assignment[item.Hash()]
		return !ok
	})
	first := unassigned[0]

	// We try assigning all prossible domain values for that variable, one at a ti,e. The new assignment
	// for each is stored in a local dictionary.
	for _, value := range o.domains[first.Hash()] {
		localAssignment := cloneAssignment(assignment)
		localAssignment[first.Hash()] = value
		// In the new assignment in local assignment is consistent will all constrains, we continue
		// recursively searching with the new assignment in place. If the new assignment turns out
		// to be complete (the base case), we return the nwe assignment up the recursion chain.
		if o.Consistent(first, localAssignment) {
			if result := o.BacktrackingSearch(localAssignment); result != nil {
				return result
			}
		}
	}
	// Finaly, if we have gone through every possible domain value for a particular variable, and there is no solution
	// utilizing the existing set of assignments, we return nil, indicating no solution.
	return nil
}

func NewCPS[V, D models.HashableAndClonable](variables []V, domains map[string][]D) (*CSP[V, D], error) {
	var result *CSP[V, D]

	for _, varialbe := range variables {
		if _, ok := domains[varialbe.Hash()]; ok {
			continue
		}
		return result, errors.New("LookupError: Every variable should have a domain assigned to it")
	}

	result = &CSP[V, D]{
		variables: variables,
		domains:   domains,
		constraints: func() map[string][]internalModels.Constraint[V, D] {
			return make(map[string][]internalModels.Constraint[V, D], 0)
		}(),
	}

	return result, nil
}
