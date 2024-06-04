package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByName(ctx context.Context, username string) (*User, error)
		FindByPage(ctx context.Context, pageNum int64, pageSize int64) ([]*User, int64, int64, error)
		FindPageByGradeAndIdentity(ctx context.Context, identity string, grade string, pageNum int64, pageSize int64) ([]*User, int64, int64, error)
		FindStuById(ctx context.Context, id string) (*User, error)
		FindTeacherById(ctx context.Context, id string) (*User, error)
		ListTeacherOrStuByGrade(ctx context.Context, grade string) ([]*User, error)
		ListByUserids(ctx context.Context, userids []string) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db, collection string) UserModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}

func (m *customUserModel) FindOneByName(ctx context.Context, username string) (*User, error) {
	var data User

	err := m.conn.FindOne(ctx, &data, bson.M{"username": username})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindByPage(ctx context.Context, pageNum int64, pageSize int64) ([]*User, int64, int64, error) {
	var data []*User

	opts := new(options.FindOptions)
	opts.SetSkip((pageNum - 1) * pageSize)
	opts.SetLimit(pageSize)

	// 分页查询列表
	err := m.conn.Find(ctx, &data, bson.M{}, opts)
	// err := m.conn.Find(ctx, &data, bson.M{})
	if err != nil {
		return nil, 0, 0, err
	}

	// 对应filter的记录总数
	cnt, err := m.conn.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, 0, err
	}

	// 转换为总页数
	totalPage := (cnt + pageSize - 1) / pageSize

	return data, totalPage, cnt, nil
}

func (m *customUserModel) FindPageByGradeAndIdentity(ctx context.Context, identity string, grade string, pageNum int64, pageSize int64) ([]*User, int64, int64, error) {
	var data []*User

	filter := bson.M{}
	if identity != "" {
		filter["identity"] = identity
	}

	if grade != "" {
		filter["grade"] = grade
	}

	opts := new(options.FindOptions)
	opts.SetSkip((pageNum - 1) * pageSize)
	opts.SetLimit(pageSize)

	// 分页查询列表
	err := m.conn.Find(ctx, &data, filter, opts)
	if err != nil {
		return nil, 0, 0, err
	}

	// 对应filter的记录总数
	cnt, err := m.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	// 转换为总页数
	totalPage := (cnt + pageSize - 1) / pageSize

	return data, totalPage, cnt, nil
}

func (m *customUserModel) FindStuById(ctx context.Context, id string) (*User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data User

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid, "identity": "学生"})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindTeacherById(ctx context.Context, id string) (*User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data User

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid, "identity": "导师"})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) ListTeacherOrStuByGrade(ctx context.Context, grade string) ([]*User, error) {
	var data []*User

	filter := bson.M{
		"grade": grade,
		"identity": bson.M{
			"$in": []string{"学生", "导师"},
		},
	}

	err := m.conn.Find(ctx, &data, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *customUserModel) ListByUserids(ctx context.Context, userids []string) ([]*User, error) {
	var data []*User

	oids := make([]primitive.ObjectID, len(userids))
	for i, userid := range userids {
		oid, err := primitive.ObjectIDFromHex(userid)
		if err != nil {
			return nil, err
		}

		oids[i] = oid
	}

	filter := bson.M{"_id": bson.M{"$in": oids}}

	err := m.conn.Find(ctx, &data, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}
