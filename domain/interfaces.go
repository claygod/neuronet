package domain

// NeuroNet
// interfaces
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// import (
// 	"unsafe"
// )

/*
NeuronInterface - интерфейс нейрона. Он должен уметь:
- получать сигнал и рассылать результаты
- изменяться после оценки прогноза
*/
type NeuronInterface interface {
	// AppendSignal (учитываем сигналы, накопив комплект переходим к анализу и рассылке результата через аксон
	AppendSignal(sig *Signal)                   // в нейрон добавляем новый сигнал
	ForecastEstimate(weigth int64, sig *Signal) // оценка совпадения прогноза и результата (обучение сети)
	SynapsesInc()                               // фаза роста связей, юзаем SynapseCreate у пред-предыдущих нейронов
	SynapsesDec()                               // фаза отмирания связей, юзаем SynapseRemove
	HealthCheck() int64                         // показатель здоровья
	Died()                                      // смерть нейрона (нейрон обрывает все связи)
	Clone()                                     //TODO: как производить клонирование, ещё надо продумать

	// хэлперы
	SynapseCreate(neuronID int64) // создание новой связи через аксон (снаружи) нейрону
	SynapseRemove(neuronID int64) // удаление старой связи (аксон/дендрит) к умершему (снаружи) нейрону
}

// на основе нового сигнала произвести обучение сети

// фаза роста связей

// фаза отмирания связей

// фаза умирания нейронов

// фаза клонирования (для защиты от переоптимизации)

/*
DendriteInterface - интерфейс дендрита. Он должен уметь:
- получать сигнал и анализировать сигнал
- отдавать результат анализа сигнала
- изменяться после оценки прогноза

По сути дендрит, это хранилище памяти одной конкретной связи нейрон-нейрон и учётчик здоровья этой связи.
*/
type DendriteInterface interface {
	AppendSignal(sig *Signal) int64 // в дендрит добавляем новый сигнал
	ForecastEstimate(weigth int64)  // оценка совпадения прогноза и результата
	HealthCheck() int64             // показатель здоровья
}

/*
DendritsRepoInterfac - интерфейс репозитория дендритов. Он должен уметь:
- добавлять, отдавать и удалять дендрит
- предоставлять список дендритов, их количество
*/
type DendritsRepoInterface interface {
	Get(nID int64) (DendriteInterface, error)
	Set(nID int64, drt DendriteInterface) error
	Del(nID int64) error
	List() []int64
	Count() int64
}

/*
DendriteReactionsAggregateInterface - агрегатор сигналов.
Получив все сигналы дёргает нейрон и возможно, обнуляется для следующего цикла.
! Возможно, тут пригодится заюзать `waitGroup`.
*/

/*
AxonInterface - интерфейс аксона. Он должен уметь:
- подключать и отключать нейроны
- рассылать сигнал по подключенным нейронам
*/
type AxonInterface interface {
	AddNeuron(nID uint64)    // подключаем к рассылке новый нейрон
	RemoveNeuron(nID uint64) // отключаем один из нейронов

	//TODO: рассылку можно объединить в один метод с флагом

	BroadcastTotal(sig *Signal) int64      // рассылка сигнала всем подключенным нейронам
	BroadcastStochastic(sig *Signal) int64 // рассылка пустышки всем и одного случайному нейрону (из подключенных)
}

/*
AnalysisInterface - интерфейс анализатора. За этим интерфейсом могут быть разные анализаторы.
У анализаторов может быть своя память, в которой они сохранят образы, поледовательности и т.д.
Однако это не MemoryInterface , т.е. этот интерфейс скорей всего не подойдёт.
*/
type AnalysisInterface interface {
	MemoryResearch(curMem [][]byte) int64 // исследование текущего состояния памяти (получаемые данные менять нельзя!)
	ForecastEstimate(weigth int64)        // оценка совпадения прогноза и результата
	HealthCheck() int64                   // показатель здоровья
}

/*
NeuronMemoryInterface - интерфейс памяти нейрона.
Хранит воспоминания об удачных паттернах, забывает со временем.
*/
type NeuronMemoryInterface interface {
	Add(nID int64, reaction int64) error
	Summary() int64
}

/*
MemoryInterface - интерфейс хранилища памяти дендрита или нейрона или аксона.
Можно нейрон вообще оставить организатором, а память раскидать по периферийным дендритам и аксону.
*/
type MemoryInterface interface {
	//AppendSignal(sig *Signal) // добавляем новый сигнал
	AppendData(nowData []byte) // добавляем новый сигнал
	Estimate(bool)             // положительная или отрицательная оценка текущего состояния (ситуации)
	ActualMemory() [][]byte    // актуальный срез памяти
	HealthCheck() int64        // показатель здоровья
}
