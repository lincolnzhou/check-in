package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Check struct {
	Id_  bson.ObjectId `bson:"_id"`
	UID  int32
	Date time.Time
}
