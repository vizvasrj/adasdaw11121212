// A restaurant keeps a log of (eater_id, foodmenu_id) for all the diners. The eater_id is a unique number for every diner and foodmenu_id is unique for every food item served on the menu. Write a program that reads this log file and returns the top 3 menu items consumed. If you find an eater_id with the same foodmenu_id more than once then show an error.

// Expected output:
// Code that can handle the above problem statement
// Testcases (with example log files) that checks the possible conditions with unit tests.
// Code has to be submitted in your github repository (share the repo link).
// Containerize the application and host the image.

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
)

func get_menu_ids() []string {
	return []string{
		"menu0001", "menu0002", "menu0003", "menu0004",
		"menu0005", "menu0006", "menu0007", "menu0008",
	}
}

type FoodItem struct {
	ID   string
	Freq int
}

func main() {

	results := make(map[string]int)
	foods := []FoodItem{}
	// for _, x := range get_menu_ids() {
	// 	foods = append(foods, FoodItem{
	// 		ID:   x,
	// 		Freq: 0,
	// 	})
	// }

	abs_path, err := filepath.Abs("./logs")
	if err != nil {
		log.Fatalf(err.Error())
	}
	files, err := ioutil.ReadDir(abs_path)
	if err != nil {
		log.Fatalln(err)
	}
	texts := []string{}
	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(abs_path, file.Name()))
		if err != nil {
			log.Fatalf(err.Error())
		}
		text := string(dat)

		for _, x := range strings.Split(text, "\n") {
			texts = append(texts, x)
		}
		// texts = append(texts, text)
	}

	// clean line split from \n
	// for _, x := strings.Split(texts, "\n")

	for _, x := range texts {
		line := strings.Split(x, " ")

		eater := line[0]
		// color.Magenta("\nEater %s", eater)

		// increase dishes count by 1
		var current_users_food []string
		for i := 0; i < len(line)-1; i++ {
			menuID := strings.TrimSpace(fmt.Sprintf("%s", line[len(line)-1-i]))
			// trimMenuID := strings.TrimSpace(menuID)
			// color.Red(menuID)

			// color.Green("%#v\n", current_users_food)
			if ArrayContains(current_users_food, menuID) {
				err := errors.New(fmt.Sprintf("Error here duplication food by user %s for this food %s", eater, menuID))
				fmt.Println(err.Error())
			} else {
				if results[menuID] == 0 {
					results[menuID] = 1
				} else {
					results[menuID]++
				}
			}
			current_users_food = append(current_users_food, menuID)

		}
	}

	// m := map[string]int{"hello": 10, "foo": 20, "bar": 20}
	n := map[int][]string{}
	var a []int
	for k, v := range results {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	count := 0

	// last_res := []map[string]int{}

	for _, k := range a {
		for _, s := range n[k] {
			count += 1
			// fmt.Printf("%s, %d\n", s, k)
			// last_res = append(last_res, map[string]int{
			// 	s: k,
			// })
			foods = append(foods, FoodItem{
				ID:   s,
				Freq: k,
			})
		}

	}

	color.Green("top 3 items are \n")
	for x := 0; x <= 2; x++ {
		color.Red("food_id: %#v, freq: %d", foods[x].ID, foods[x].Freq)
	}

	// color.Red("%#v", results[1])
	// color.Red("%#v", results[2])
	// color.Red("%#v", results[3])

}

func ArrayContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
