package post

type Controller interface {
	// TODO - implement methods
}

type controller struct {
	// attributes
}

func NewController(r Repository) Controller {
	return controller{
		// initialized attributes
	}
}

func (c *controller) Method(){
	// method stuff
}

// methods implemented