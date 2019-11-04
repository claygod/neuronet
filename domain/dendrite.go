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
	neuron *Neuron
	memory *dendriteMemory
}

/*
dendriteMemory - память дендрита, хранит последние отрицательные и положительные
воспоминания и текущий момент.
*/
type dendriteMemory struct {
	m                sync.Mutex
	momentLength     int64
	memoryLength     int64
	currentEvent     int64
	currentMoment    []int64   //events list
	positiveMemories [][]int64 //mem list
	negativeMemories [][]int64 //mem list
}

func (d *dendriteMemory) addEvent(event int64) {
	d.m.Lock()
	defer d.m.Unlock()
	d.addEventToList(event)
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
