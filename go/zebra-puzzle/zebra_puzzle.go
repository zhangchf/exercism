package zebra


//There are five houses.
//The Englishman lives in the red house.
//The Spaniard owns the dog.
//Coffee is drunk in the green house.
//The Ukrainian drinks tea.
//The green house is immediately to the right of the ivory house.
//The Old Gold smoker owns snails.
//Kools are smoked in the yellow house.
//Milk is drunk in the middle house.
//The Norwegian lives in the first house.
//The man who smokes Chesterfields lives in the house next to the man with the fox.
//Kools are smoked in the house next to the house where the horse is kept.
//The Lucky Strike smoker drinks orange juice.
//The Japanese smokes Parliaments.
//The Norwegian lives next to the blue house. 



type House struct {
	order int
	color string
	nationality string
	drink string
	cigrette string
	pet string
	next *House
}

type HouseConstraint House

// colors: red, green, yellow, ivory, blue

var constraints = []HouseConstraint {
	{nationality:"Englishman", color: "red"},
	{nationality:"Spaniard", pet: "dog"},
	{drink:"coffee", color: "green"},
	{nationality:"Ukrainian", drink:"tea"},
	//The green house is immediately to the right of the ivory house.
	{cigrette:"Old Gold", pet:"snails"},
	{cigrette:"Kools", color:"yellow"},
	{drink:"Milk", order:3},
	{nationality:"Norwegian", order:1},
	//The man who smokes Chesterfields lives in the house next to the man with the fox.
	//Kools are smoked in the house next to the house where the horse is kept.
	{cigrette:"Lucky Strike", drink:"orange juice"},
	{nationality:"Japanese", cigrette:"Parliaments"},
	//The Norwegian lives next to the blue house.
}

func SolvePuzzle() {
	
}
