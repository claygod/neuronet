package usecases

// NeuroScheme
// Ьфз interactor
// Copyright © 2020 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// import (
// 	"github.com/claygod/neuronet/domain"
// )

/*
SchemeInteractor - интерактор для создания схемы нейронной сети.
Схема требуется для сохранения и загрузки нейронной сети.
При этом появляется возможность запускать сеть не только с простой
структурой, но и с уже адаптированной под конкрутный кейс структурой.
Схема, это не копия сети, а именно схема!
*/
type SchemeInteractor struct {
}

/*
NewSchemeInteractor - создание интерактора.
*/
func NewSchemeInteractor() *SchemeInteractor {
	return &SchemeInteractor{} //TODO:
}

func (m *SchemeInteractor) KitFromScheme(sch string) *Kit {
	return nil //TODO:
}

func (m *SchemeInteractor) SchemeFromKit(kt *Kit) string {
	return "" //TODO:
}
