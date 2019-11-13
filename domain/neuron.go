package domain

// NeuroNet
// Neuron
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Neuron - нейрон.
При принятии отрицательного решения по результатам входных сигналов всё равно рассылает через аксон сигналы, но пустышки.
*/
type Neuron struct {
	id   uint64
	axon *Axon
}

func (n *Neuron) GetDendrite(fromId uint64) (*Dendrite, error) {
	return nil, nil //TODO:
}
