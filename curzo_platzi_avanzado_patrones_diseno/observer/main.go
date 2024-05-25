package main

import "fmt"

// Topic interface defines methods for managing observers and broadcasting messages.
type Topic interface {
	register(observer Observer) // Add an observer to the topic.
	broadcast()                 // Notify all observers about an event.
}

// Observer interface defines methods for observers that listen to Topic updates.
type Observer interface {
	getId() string            // Returns the ID of the observer.
	updateValue(value string) // Method to update the observer with new information.
}

// Item struct acts as a concrete implementation of a Topic.
// It maintains a list of observers and notifies them when its availability changes.
type Item struct {
	observers []Observer // List of observers interested in updates from this item.
	name      string     // Name of the item.
	available bool       // Availability status of the item.
}

// newItem creates a new Item with the specified name.
func newItem(name string) *Item {
	return &Item{
		name:      name,
		available: false,
	}
}

// updateAvailable marks the item as available and notifies all registered observers.
func (item *Item) updateAvailable() {
	fmt.Printf("Item %s is available\n", item.name)
	item.available = true
	item.broadcast()
}

// broadcast notifies all the observers about the item's availability by calling their updateValue method.
func (item *Item) broadcast() {
	for _, observer := range item.observers {
		observer.updateValue(item.name)
	}
}

// register adds a new observer to the item.
func (item *Item) register(observer Observer) {
	item.observers = append(item.observers, observer)
}

// EmailClient struct represents an observer that receives updates via email.
type EmailClient struct {
	id string // Identifier for the EmailClient.
}

// getId returns the ID of the EmailClient.
func (e *EmailClient) getId() string {
	return e.id
}

// updateValue simulates sending an email to notify the client that an item is available.
func (e *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - EmailClient %s: Item %s is available\n", e.id, value)
}

// main function sets up the observer pattern example with an item and two observers.
func main() {
	item := newItem("iPhone")               // Create a new item.
	firstObserver := &EmailClient{id: "1"}  // Create first observer.
	secondObserver := &EmailClient{id: "2"} // Create second observer.
	item.register(firstObserver)            // Register first observer.
	item.register(secondObserver)           // Register second observer.
	item.updateAvailable()                  // Update item availability and notify observers.

	fmt.Println("Observer pattern")
}
