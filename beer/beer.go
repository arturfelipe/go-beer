package beer

// Beer model
type Beer struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Style       string `bson:"style"`
	Abv         int    `bson:"abv"`
}
