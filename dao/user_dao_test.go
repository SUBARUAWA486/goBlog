package dao

import (
	"testing"

	"goBlog/config"
	"goBlog/model"

	"github.com/stretchr/testify/assert"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func initTestDB(t *testing.T) {
	var err error
	config.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect test DB: %v", err)
	}

	err = config.DB.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatalf("failed to migrate User model: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	initTestDB(t)

	user := &model.User{
		ID:       "1",
		Account:  "user001",
		Password: "123456",
		Nickname: "测试用户",
		Avatar:   "default.png",
	}

	err := CreateUser(user)
	assert.NoError(t, err, "should create user successfully")

	//重复账号
	user2 := &model.User{
		ID:       "2",
		Account:  "user001", // duplicate
		Password: "123456",
		Nickname: "另一个用户",
		Avatar:   "default.png",
	}
	err = CreateUser(user2)
	assert.Error(t, err, "should fail due to duplicate account")

	//重复昵称
	user3 := &model.User{
		ID:       "3",
		Account:  "user002",
		Password: "123456",
		Nickname: "测试用户", // duplicate
		Avatar:   "default.png",
	}
	err = CreateUser(user3)
	assert.Error(t, err, "should fail due to duplicate nickname")
}

func TestIsNicknameExist(t *testing.T) {
	initTestDB(t)

	user := &model.User{
		ID:       "test-id-124",
		Account:  "acc123457",
		Nickname: "nick_exist",
		Password: "passwd12345",
		Avatar:   "http://example.com/avatar2.jpg",
	}
	_ = CreateUser(user)

	ok, err := IsNicknameExist("nick_exist")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = IsNicknameExist("not_exist")
	assert.NoError(t, err)
	assert.False(t, ok)
}

func TestGetUserByNickname(t *testing.T) {
	initTestDB(t)

	user := &model.User{
		ID:       "test-id-125",
		Account:  "acc123458",
		Nickname: "getnick",
		Password: "passwd12345",
		Avatar:   "http://example.com/avatar3.jpg",
	}
	_ = CreateUser(user)

	u, err := GetUserByNickname("getnick")
	assert.NoError(t, err)
	assert.Equal(t, "acc123458", u.Account)
}

func TestGetUserByAccount(t *testing.T) {
	initTestDB(t)

	user := &model.User{
		ID:       "test-id-126",
		Account:  "acc123459",
		Nickname: "getacc",
		Password: "passwd12345",
		Avatar:   "http://example.com/avatar4.jpg",
	}
	_ = CreateUser(user)

	u, err := GetUserByAccount("acc123459")
	assert.NoError(t, err)
	assert.Equal(t, "getacc", u.Nickname)
}

func TestGetUserByID(t *testing.T) {
	initTestDB(t)

	user := &model.User{
		ID:       "test-id-127",
		Account:  "acc123460",
		Nickname: "getid",
		Password: "passwd12345",
		Avatar:   "http://example.com/avatar5.jpg",
	}
	_ = CreateUser(user)

	u, err := GetUserByID("test-id-127")
	assert.NoError(t, err)
	assert.Equal(t, "acc123460", u.Account)
}
