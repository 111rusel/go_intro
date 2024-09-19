package main

import (
	"fmt"
	"strings"
)

var tickets []string = []string{
	"Орел - Магадан",
	"Рязань - Пенза",
	"Новосибирск - Уфа",
	"Омск - Магнитогорск",
	"Воронеж - Мурманск",
	"Казань - Симферополь",
	"Владивосток - Ставрополь",
	"Калининград - Новосибирск",
	"Мурманск - Орел",
	"Волгоград - Владивосток",
	"Москва - Рязань",
	"Уфа - Воронеж",
	"Магадан - Волгоград",
	"Ставрополь - Омск",
	"Симферополь - Москва",
	"Магнитогорск - Казань",
}

func findCity(t []string) string {
	start := []string{}
	end := map[string]struct{}{}
	for i := range t {
		words := strings.Split(t[i], " - ")
		start = append(start, words[0])
		end[words[1]] = struct{}{}
	}
	for i := range start {
		_, ok := end[start[i]]
		if !ok {
			return start[i]
		}
	}
	return ""
}

func getRoute(startCity string, t []string) []string {
	result := []string{startCity}
	tickets := map[string]string{}
	for i := range t {
		words := strings.Split(t[i], " - ")
		tickets[words[0]] = words[1]
	}
	for startCity != "" {
		city := tickets[startCity]
		result = append(result, city)
		startCity = city
	}
	return result
}

func main() {
	startCity := findCity(tickets)
	route := getRoute(startCity, tickets)
	fmt.Println(startCity, route)
}
