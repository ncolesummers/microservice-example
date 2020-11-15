package persistence

type DatabaseHandler interface {
	AddUser(User) ([]byte, error)
	AddEvent(Event) ([]byte, error)
	AddBookingForUser([]byte, Booking) error
	AddLocation(Location) (Location, error)
	FindUser(string, string) (User, error)
	FindBookingForUser([]byte) (Event, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
	FindLocation(string) (Location, error)
	FindAllLocations() ([]Location, error)
}
