package synapse

// NeuroNet
// Synapse memory (MemoryInterface implementation)
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"sync"
	// "unsafe"
	// "github.com/claygod/neuronet/domain"
)

type SynapseMemory struct {
	memoryLength int64 // максимальное количество хранящихся воспоминаний (это ограничение под вопросом)
}

// type MemoryInterface interface {
// 	//AppendSignal(sig *Signal) // добавляем новый сигнал
// 	AppendData(nowData []byte) // добавляем новый сигнал
//	Estimate(bool) // положительная или отрицательная оценка текущего состояния (ситуации)
// 	ActualMemory() [][]byte    // актуальный срез памяти
// 	HealthCheck() int64        // показатель здоровья
// }

/*
detailMemory - память дендрита, хранит последние отрицательные и положительные
воспоминания и текущий момент.
*/
type detailMemory struct {
	m                sync.Mutex
	momentLength     int64
	memoryLength     int64 // максимальное количество хранящихся воспоминаний (это ограничение под вопросом)
	currentEvent     int64
	currentMoment    []int64   //events list
	positiveMemories [][]int64 //mem list
	//negativeMemories [][]int64 //mem list
}

func newDendriteMemory(momentLength int64, memoryLength int64) *detailMemory {
	return &detailMemory{
		//TODO:
		momentLength: momentLength,
		memoryLength: memoryLength,

		currentMoment:    make([]int64, momentLength),
		positiveMemories: make([][]int64, 0, memoryLength),
		//negativeMemories: make([][]int64, 0, memoryLength),
	}
}

/*
addEvent - поступил сигнал от предыдущего нейрона, сигнал надо добавить в память момента на случай
положительной или отрицательной обратной связи, проверить, есть ли воспоминания на эту тему и усилить
или наоборот ослабить сигнал.
*/
func (d *detailMemory) addEvent(event int64) int64 {
	d.m.Lock()
	defer d.m.Unlock()
	d.addEventToList(event)
	//TODO: обработка сигнала (ищем в памяти позитив и негатив)
	memSum := d.searchInMemory()
	//TODO: отправка сигнала в агрегатор сигналов нейрона
	return memSum //TODO: +/- event как-то надо учитывать входной сигнал
}

func (d *detailMemory) searchInMemory() int64 {
	posRate := d.maxRateFromMemory(d.positiveMemories)
	//negRate := d.maxRateFromMemory(d.negativeMemories)
	return posRate
}

func (d *detailMemory) maxRateFromMemory(mems [][]int64) int64 {
	var rate int64 = 0
	for i, mem := range mems {
		for u := int64(0); u < d.momentLength; u++ {
			if d.currentMoment[u] != mem[u] && u > int64(i/10) { // вторая проверка для того, чтобы старость воспоминания влияла
				rate += (u - int64(i/10)) * d.currentMoment[d.momentLength] //умножаем на весомость воспоминания (старость тоже влияет)
			}
		}
	}
	return rate
}

func (d *detailMemory) reactionToEvent(weigth int64) {
	d.m.Lock()
	defer d.m.Unlock()
	d.addMomentToList(weigth)
}

func (d *detailMemory) addMomentToList(weigth int64) {
	// создаём копию текущего момента для добавления в долговременную память
	mem := make([]int64, 0, d.momentLength)
	copy(mem, d.currentMoment)
	mem = append(mem, weigth) // последний элемент, это весомость воспоминания

	// добавляем в начало воспоминаний текущий момент и дописываем остальные
	newCurMom := make([][]int64, 0, d.memoryLength)
	newCurMom = append(newCurMom, mem)
	//if reaction {
	newCurMom = append(newCurMom, d.positiveMemories...)
	copy(d.positiveMemories, newCurMom)
	// } else {
	// 	newCurMom = append(newCurMom, d.negativeMemories...)
	// 	copy(d.negativeMemories, newCurMom)
	// }
}

func (d *detailMemory) addEventToList(event int64) {
	newCurMom := make([]int64, 0, d.momentLength)
	newCurMom = append(newCurMom, event)
	newCurMom = append(newCurMom, d.currentMoment...)
	copy(d.currentMoment, newCurMom)
}

/*
TODO: напдо продумать заново логику этого вида памяти!!

simpleMemory - память с упрощённым восприятием сигналов 0/1 длиной 64 бита.
В данносм случае входной сигнал дендрита сюда будет помещён в бинарном варианте.
Т.е. сила входных сигналов не учитывается. Также не учитывается знак последующей
реакции подкрепления-обучения. Таким образом, эта память может служить для получения
коэффициента усиления реакции, т.е. насколько бурно дендрит среагирует.
То, что в этой памяти хранится, можно назвать ритмом. Например два плюса,
два минуса. Обычная память может не заметить системность, т.к. у сигналов может быть
не совсем одинаковая сила, но есть какая-то регулярность.

Можно предположить, что сигналы, на которые не было реакции учителя, сюда не попали, и тогда
функционально всё становится похоже на график крестики-нолики
*/
type simpleMemory struct {
	m         sync.Mutex
	memsCount int64
	curMem    uint64   // текущий момент
	mems      []uint64 // значимые воспоминания

}

/*
addEvent - добавление входного сигнала, нормализованного до бинарного варианта.
*/
func (s *simpleMemory) addEvent(event bool) {
	var x uint64
	if event {
		x = 1
	}
	s.curMem = (s.curMem << 1) + x
}

/*
reactionToEvent - при любой реакции, полученной при обучении-подкреплении это запоминается.
*/
func (s *simpleMemory) reactionToEvent(event bool) {
	memsNew := make([]uint64, 0, s.memsCount)
	memsNew = append(memsNew, s.curMem)
	memsNew = append(memsNew, s.mems...)
	copy(s.mems, memsNew)
}

/*
maxRateFromMemory - поиск максимального совпадения с учётом старости воспоминаний и глубины совпадения.
*/
func (s *simpleMemory) maxRateFromMemory() int64 {
	var rate int64 = 0
	for i, mem64 := range s.mems {
		cur64 := s.curMem
		for u := 0; u < 64; u++ {
			if mem64 == cur64 {
				react := int64(64 - i - u)
				if react > rate {
					rate = react
				}
			}
			cur64 = cur64 >> 1
			mem64 = mem64 >> 1
		}
	}
	return rate
}
