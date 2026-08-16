package main

func main()  {
	state := config{
		commands: getCommands(),
		next: "https://pokeapi.co/api/v2/location-area/", 
		previous: nil,
	}
	startRepl(&state)
}
