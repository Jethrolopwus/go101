package main

import "fmt"


type Validator struct {
    IsActive bool
}


func GetActiveValidators(validators []Validator) []Validator {
    var activeValidators []Validator
    for _, v := range validators {
        if v.IsActive {
            activeValidators = append(activeValidators, v)
        }
    }
    return activeValidators
}

func main() {
    validators := []Validator{
        {IsActive: true},
        {IsActive: false},
        {IsActive: true},
    }

    activeValidators := GetActiveValidators(validators)
    fmt.Printf("Active Validators: %v\n", activeValidators)
}
