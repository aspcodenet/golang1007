package main

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	"systementor.se/demo1007/employee"
)

type Player struct {
	Name         string // Peter Forsberg
	Age          int    // 51
	JerseyNumber int    // 21
	Team         string // colorado
}

// STACK - växer och krymper perfekt + inga "hål" - livslängd = dess scope {}
// vi måste veta p förhand hur stora dom är
// String s = "Stefan"; String s = "ewfjkasjfdjaklsdjklafdjklasjklasjklfasdjklfsda";
// int j = 12;
// int r = 13;
// HEAP new /malloc
// player = new Player(); // växer och krymper INTE  perfekt +  "hål"  - kan leva tills ingen använder som längre -  GC fixar
//
//	defragmentera

// func createFoppa() *Player {
// 	return &foppa
// }

// func changeOnePlayer(players [3]Player) { // player BLIR EN KOPIA
// 	players[0].Name = "Zäta the legend"
// }

// SKA VI ÄNDRA PÅ NÅT - använd PEKARE !!!

func changeOnePlayer(players []Player) { // player BLIR EN KOPIA ANTAL|PEKARE
	players[0].Name = "Zäta the legend"
}

// DET SOM ÄR BRA MED NEW -  constructor - anger vi ju parametrar för att säkerställa valid state

// generics i vår bemärkelse finns inte LINQ, streams()
// ansvaret på programmeraren - for each i din lista

func main() {

	stefan := employee.New("Stefan", 52) // Gentlemanna constructor
	//salary := employee.CalculateSalary(stefan) //1000 GO STLW
	salary := stefan.CalculateSalary()
	fmt.Println(salary)

	foppa := Player{
		Name: "Peter Forsberg",
		//Age:          51,
		JerseyNumber: 21,
		Team:         "Colorado",
	}

	zata := Player{
		Name:         "Henrik Zetterberg",
		Age:          44,
		JerseyNumber: 40,
		Team:         "Detroit",
	}
	zata2 := Player{
		Name:         "2",
		Age:          2341,
		JerseyNumber: 2234,
		Team:         "Colorado",
	}
	// var players = [3]Player{zata, zata2, foppa} // ... eller siffra = ARRAY. Konsekutivt minne + STATISK
	// players[0].Name = "Whatever"
	// PLAYER|PLAYER|PLAYER

	// changeOnePlayer(players) // ALLA PARAMNETRAR I GO ÄR COPY BY VALUE

	// SLICE
	var players = []Player{zata, zata2, foppa} // [] slice = fortfarande konsekutivt minne + DYNAMISK
	changeOnePlayer(players)                   // COPY BY VALUE
	// Slices blir "som en referens"
	// ANTAL|PEKARE ->     PLAYER|PLAYER|PLAYER

	// 1. SKapa ny spelare
	newPlayer := Player{
		Name:         "2",
		Age:          2341,
		JerseyNumber: 2234,
		Team:         "Colorado",
	}
	// motsvarande platers = realloc(players, ...)
	players = append(players, newPlayer) // krävs en reallokering?

	// s := "hejsan"
	// s = strings.Replace(s,"ewasd")

	fmt.Println("TILLBAKA")
	t := tabby.New()
	t.AddHeader("NAME", "AGE", "JERSEY", "TEAM")
	for _, player := range players {
		t.AddLine(player.Name, player.Age, player.JerseyNumber, player.Team)
		//fmt.Println(player.Name)
	}
	t.Print()

}
