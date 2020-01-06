package domain

// NeuroNet
// Repository interface
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type NeuronsRepo interface {
	Set(id uint64, neuron NeuronInterface) error
	Get(id uint64) (NeuronInterface, error)
}
