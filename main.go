package main

func main() {
	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Buy Bread")
	todos.print()
	todos.delete(0)
	todos.print()

}
