package domain

// NeuroNet
// Signal
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type Signal struct {
	uid    uint64
	level  uint64
	parent *Signal
	owner  *Neuron
}
