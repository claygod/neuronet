package usecases

// NeuroNet
// Net interactor
// Copyright © 2020 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/neuronet/domain"
)

/*
NetInteractor - основной интерактор для работы с нейронами.
*/
type NetInteractor struct {
	neuronsRepo *domain.NeuronsRepo
	// neuroList    []*domain.NeuronInterface
	inputNeurons  map[string]uint64
	outputNeurons map[string]uint64
}

/*
NewNetInteractor - создание интерактора.
*/
func NewNetInteractor() *NetInteractor {
	return &NetInteractor{} //TODO:
}

/*
NextSignals - поступает новый пакет сигналов, на которые сеть далжна дать свои прогнозы.
*/
func (n *NetInteractor) NextSignals(inData map[string]int64) {
	// сгенерировать новые сигналы

	// на основе нового сигнала произвести обучение сети

	// фаза роста связей

	// фаза отмирания связей

	// фаза умирания нейронов

	// отправить сигналы в теперь уже немного подкорректированную сеть

	// результаты работы сети вернуть в ответе

}
