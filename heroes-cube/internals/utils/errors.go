package utils

import "errors"

var ClassNotFound = errors.New("classe não encontrada")
var RaceNotFound = errors.New("raça não encontrada")
var PersonNotFound = errors.New("personagem não encontrado")
var ItemNotFoundInInventory = errors.New("item não encontrado no inventario")
var ItemNotCompatible = errors.New("item não compatível com sua classe")
