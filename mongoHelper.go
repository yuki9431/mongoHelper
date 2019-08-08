package mongoHelper

import "gopkg.in/mgo.v2"

type mongoDb struct {
	dial string
	name string
	db   *mgo.Database
}

type MongoDb interface {
	disconnectDb()
	InsertDb(interface{}, string) error
	RemoveDb(interface{}, string) error
	SearchDb(interface{}, string) error
	UpdateDb(interface{}, interface{}, string) error
}

func NewMongo(dial string, name string) (MongoDb, error) {
	session, err := mgo.Dial(dial)
	db := session.DB(name)

	return &mongoDb{
		dial: dial,
		name: name,
		db:   db,
	}, err
}

// mongoDB切断
func (m *mongoDb) disconnectDb() {
	m.db.Session.Close()
}

// mongoDB挿入
func (m *mongoDb) InsertDb(obj interface{}, colectionName string) (err error) {
	defer m.disconnectDb()

	col := m.db.C(colectionName)
	return col.Insert(obj)
}

// mongoDB削除
func (m *mongoDb) RemoveDb(obj interface{}, colectionName string) (err error) {
	defer m.disconnectDb()

	col := m.db.C(colectionName)
	_, err = col.RemoveAll(obj)
	return
}

// mondoDB抽出
func (m *mongoDb) SearchDb(obj interface{}, colectionName string) (err error) {
	defer m.disconnectDb()

	col := m.db.C(colectionName)
	return col.Find(nil).All(obj)
}

// mongoDB更新
func (m *mongoDb) UpdateDb(selector, update interface{}, colectionName string) (err error) {
	defer m.disconnectDb()

	col := m.db.C(colectionName)
	return col.Update(selector, update)
}
