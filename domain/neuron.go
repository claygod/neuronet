package domain

// NeuroNet
// Neuron
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type Neuron struct {
	id   uint64
	axon *Axon
}

func (n *Neuron) GetDendrite(fromId uint64) (*Dendrite, error) {
	return nil, nil //TODO:
}
