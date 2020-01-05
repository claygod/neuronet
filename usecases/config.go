package usecases

// NeuroScheme
// Ьфз interactor
// Copyright © 2020 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/neuronet/domain"
)

type Kit struct {
	neuronsRepo *domain.NeuronsRepo
	// neuroList    []*domain.NeuronInterface
	inputNeurons  map[string]uint64
	outputNeurons map[string]uint64
}

/*
Что должно быть в схеме

{
   "inputs":[
      {
         "name":"candle_max",
         "max_size":250
      },
      {
         "name":"candle_min",
         "max_size":250
      }
   ],
   "outputs":[
      {
         "name":"up",
         "rate":1
      },
      {
         "name":"up100",
         "rate":5
      },
      {
         "name":"down",
         "rate":1
      }
   ],
   "dependency":[
      {
         "input":"candle_max",
         "change_type":"plus",
         "outputs":[
            {
               "up":1
            },
            {
               "up100":100
            }
         ]
      },
      {
         "input":"candle_max",
         "change_type":"minus",
         "outputs":[
            {
               "down":-1
            }
         ]
      }
   ]
}

*/
