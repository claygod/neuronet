package domain

// NeuroNet
// Dendrite
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"sync"
)

/*
Dendrite - вход нейрона. Отдельный для каждой связи. Дендрит обладает памятью.
Он может запоминать паттерны, получившие как положительную, так и отрицательную оценку.
*/
type Dendrite struct {
	neuNum int64 // номер под которым этот дендрит числится в агрегаторе
	neuron *Neuron
	memory *dendriteMemory
}

func (d *Dendrite) TransmitSignal(sig *Signal) {
	//TODO: обработка, добавление в память, передача в тело нейрона
}

/*
dendriteMemory - память дендрита, хранит последние отрицательные и положительные
воспоминания и текущий момент.
*/
type dendriteMemory struct {
	m                sync.Mutex
	momentLength     int64
	memoryLength     int64 // максимальное количество хранящихся воспоминаний (это ограничение под вопросом)
	currentEvent     int64
	currentMoment    []int64   //events list
	positiveMemories [][]int64 //mem list
	negativeMemories [][]int64 //mem list
}

func newDendriteMemory(momentLength int64, memoryLength int64) *dendriteMemory {
	return &dendriteMemory{
		//TODO:
		momentLength: momentLength,
		memoryLength: memoryLength,

		currentMoment:    make([]int64, momentLength),
		positiveMemories: make([][]int64, 0, memoryLength),
		negativeMemories: make([][]int64, 0, memoryLength),
	}
}

/*
addEvent - поступил сигнал от предыдущего нейрона, сигнал надо добавить в память момента на случай
положительной или отрицательной обратной связи, проверить, есть ли воспоминания на эту тему и усилить
или наоборот ослабить сигнал.
*/
func (d *dendriteMemory) addEvent(event int64) int64 {
	d.m.Lock()
	defer d.m.Unlock()
	d.addEventToList(event)
	//TODO: обработка сигнала (ищем в памяти позитив и негатив)
	memSum := d.searchInMemory()
	//TODO: отправка сигнала в агрегатор сигналов нейрона
	return memSum //TODO: +/- event как-то надо учитывать входной сигнал
}

func (d *dendriteMemory) searchInMemory() int64 {
	posRate := d.maxRateFromMemory(d.positiveMemories)
	negRate := d.maxRateFromMemory(d.negativeMemories)
	return posRate - negRate
}

func (d *dendriteMemory) maxRateFromMemory(mems [][]int64) int64 {
	var rate int64 = 0
	for i, mem := range mems {
		for u := int64(0); u < d.momentLength; u++ {
			if d.currentMoment[u] != mem[u] {
				rate = u
				k := d.momentLength - int64(i) + u //TODO: формула получения абсолютного значения оценки
				if k > rate {
					rate = k
				}
			}
		}
	}
	return rate
}

func (d *dendriteMemory) reactionToEvent(reaction bool) {
	d.m.Lock()
	defer d.m.Unlock()
	d.addMomentToList(reaction)
}

func (d *dendriteMemory) addMomentToList(reaction bool) {
	// создаём копию текущего момента для добавления в долговременную память
	mem := make([]int64, 0, d.momentLength)
	copy(mem, d.currentMoment)

	// добавляем в начало воспоминаний текущий момент и дописываем остальные
	newCurMom := make([][]int64, 0, d.memoryLength)
	newCurMom = append(newCurMom, mem)
	if reaction {
		newCurMom = append(newCurMom, d.positiveMemories...)
		copy(d.positiveMemories, newCurMom)
	} else {
		newCurMom = append(newCurMom, d.negativeMemories...)
		copy(d.negativeMemories, newCurMom)
	}
}

func (d *dendriteMemory) addEventToList(event int64) {
	newCurMom := make([]int64, 0, d.momentLength)
	newCurMom = append(newCurMom, event)
	newCurMom = append(newCurMom, d.currentMoment...)
	copy(d.currentMoment, newCurMom)
}
