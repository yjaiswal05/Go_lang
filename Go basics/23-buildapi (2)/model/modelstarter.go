package model
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct{
	ID 			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`		//whenever we work with the database it provides its own id to make sure its primitive and searchable
	Movie 		string				`json:"movie,omitempty"`
	Watched 	bool				`json:"watched,omitempty"`
}