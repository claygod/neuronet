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

/*
addEvent - поступил сигнал от предыдущего нейрона, сигнал надо добавить в память момента на случай
положительной или отрицательной обратной связи, проверить, есть ли воспоминания на эту тему и усилить
или наоборот ослабить сигнал.
*/
func (d *dendriteMemory) addEvent(event int64) {
	d.m.Lock()
	defer d.m.Unlock()
	d.addEventToList(event)
	//TODO: обработка сигнала (ищем в памяти)
	//TODO: отправка сигнала в агрегатор сигналов нейрона
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
