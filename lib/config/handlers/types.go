package handlers

import "github.com/jinzhu/gorm"

// AppContext contains our local context; our database pool, session store, template
// registry and anything else our handlers need to access. We'll create an instance of it
// in our main() function and then explicitly pass a reference to it for our handlers to access.
type AppContext struct {
	DB *gorm.DB
	// ... and the rest of our globals.
}
