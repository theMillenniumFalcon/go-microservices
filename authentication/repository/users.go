package repository

import (
	"github.com/theMillenniumFalcon/microservices/authentication/models"
	"github.com/theMillenniumFalcon/microservices/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UsersRepository interface{}

type usersRepository struct {
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{c: conn.DB().C(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

func (r *usersRepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"email": email}).One(&user)
	return user, err
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&users)
	return users, err
}
