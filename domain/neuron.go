package domain

// NeuroNet
// Neuron
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sort"
	"sync"
)

type Neuron struct {
	sync.Mutex
	ID      uint64
	neuRepo NeuronsRepo

	roundCounter  int // тут номер ОЖИДАЕМОГО раунда
	roundWaitList []int
	inputsLimit   int64
	//inputsCount  int64
	inputsWeight map[uint64]float64

	weightsBorder float64

	//lastInput  map[uint64]bool
	inputsLog  map[int]map[uint64]bool // map[round]map[nID]bool
	outputList []uint64
}

func (n *Neuron) AddSignal(nID uint64, in bool, round int) error {
	n.Lock()
	defer n.Unlock()

	//TODO: проверка на сигнал от кого не надо
	_, ok := n.inputsWeight[nID]
	if !ok {
		return fmt.Errorf("neuron `%d` sent a signal to neuron `%d`, but it is not connected (no dendrite)", nID, n.ID)
	}

	curInput, ok := n.inputsLog[round]
	if !ok {
		curInput := make(map[uint64]bool, n.inputsLimit)
		n.inputsLog[round] = curInput
	}

	//TODO: проверяем, не пришло ли что-то повторное в round
	if _, ok = curInput[nID]; ok {
		return fmt.Errorf("neuron `%d` was already sending a signal to neuron `%d`", nID, n.ID)
	}

	curInput[nID] = in

	if int64(len(curInput)) == n.inputsLimit {
		n.roundWaitList = append(n.roundWaitList, round)
		if err := n.broadcastList(); err != nil {
			return err
		}
	}
	return nil
}

func (n *Neuron) weightAmount(curInput map[uint64]bool) (float64, error) {
	var errTotal error
	var weightAmount float64

	for num, tru := range curInput {
		if tru {
			weight, ok := n.inputsWeight[num]
			if !ok {
				errTotal = fmt.Errorf("%v unknow neuron number `%d`", errTotal, num)
			} else {
				weightAmount += weight
			}
		}
	}

	return weightAmount, errTotal
}

/*
broadcastList - пытаемся разослать все имеющиеся собранные наборы
*/
func (n *Neuron) broadcastList() error {
	var errTotal error

	sort.Ints(n.roundWaitList)
	newWaitList := make([]int, 0, len(n.roundWaitList))

	for _, round := range n.roundWaitList {
		if round == n.roundCounter {
			n.roundCounter++

			curInput, ok := n.inputsLog[round]
			if !ok {
				errTotal = fmt.Errorf("%v neuron `%d` lost `%d` round", errTotal, n.ID, round)
				continue
			}

			weight, err := n.weightAmount(curInput)
			if err != nil {
				return err
			}

			if err := n.broadcast(weight > n.weightsBorder, round); err != nil { //TODO: с весом сравниваем довольно поздно, возможно, это нужно сделать раньше
				errTotal = fmt.Errorf("%v %v", errTotal, err)
			}
		} else {
			newWaitList = append(newWaitList, round)
		}
	}

	n.roundWaitList = newWaitList

	return errTotal
}

func (n *Neuron) broadcast(in bool, round int) error { // если in == false то всё равно один сигнал стохастически отсылаем положительный
	var errTotal error

	for _, num := range n.outputList {
		neu, err := n.neuRepo.Get(num)

		if err != nil {
			errTotal = fmt.Errorf("%v %v", errTotal, err)
		} else {
			neu.AppendSignal(nil) //TODO: аргумента должно быть три для AddSignal (просто он старый, предыдущей генерации)
		}
	}

	return errTotal
}
