package lasagna

const NoodlePerLayer = 50
const SaucePerLayer = float64(0.2)
const PortionsPerPreparation = 2

func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}

	return len(layers) * preparationTime
}

func Quantities(layers []string) (int, float64) {
	sauce := float64(0)
	noodle := 0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodle++
			break
		case "sauce":
			sauce++
			break
		}
	}

	return int(noodle * NoodlePerLayer), float64(sauce * SaucePerLayer)
}

func AddSecretIngredient(friendsList, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scale := float64(portions) / PortionsPerPreparation
	newQuantities := []float64{}

	for _, quantity := range quantities {
		newQuantities = append(newQuantities, quantity*scale)
	}

	return newQuantities
}
