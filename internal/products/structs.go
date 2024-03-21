package products

type (
	Presentation struct {
		code string
		name string
	}

	Product struct {
		id           string
		name         string
		presentation Presentation
	}
)
