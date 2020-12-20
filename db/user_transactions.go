package db

import (
	"errors"
	"fmt"
	"math/rand"
	model "shanyraq/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Save user to db
func SaveUser(user model.User) (interface{}, error) {
	isUnique, err := IsUniqueUser(user)
	if !isUnique {
		return nil, err
	}
	db, ctx := MongoDbConnect()

	users := db.Collection("users")

	hashedPassword, err := EncryptPassword([]byte(user.Password))
	if err != nil {
		return nil, err
	}

	res, err := users.InsertOne(ctx, bson.D{
		{Key: "username", Value: user.Username},
		{Key: "password", Value: user.Password},
		{Key: "email", Value: user.Email},
		{Key: "telephone", Value: user.Telephone},
		{Key: "name", Value: user.Name},
		{Key: "surname", Value: user.Surname},
		{Key: "isValidated", Value: false},
		{Key: "password", Value: hashedPassword},
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res.InsertedID, nil
}

func Drop() {
	db, ctx := MongoDbConnect()

	users := db.Collection("users")
	users.Drop(ctx)
}

//Check if users credentials are already present in a database
func IsUniqueUser(user model.User) (bool, error) {
	db, ctx := MongoDbConnect()

	users := db.Collection("users")
	var useR model.User
	err := users.FindOne(ctx, bson.D{{Key: "username", Value: user.Username}}).Decode(&useR)

	if err != mongo.ErrNoDocuments {
		fmt.Println("here", user.Username)
		return false, errors.New("user exists")
	}

	err = users.FindOne(ctx, bson.D{{Key: "email", Value: user.Email}}).Decode(&useR)

	if err != mongo.ErrNoDocuments {
		return false, errors.New("user exists")
	}

	err = users.FindOne(ctx, bson.D{{Key: "telephone", Value: user.Telephone}}).Decode(&useR)

	if err != mongo.ErrNoDocuments {
		return false, errors.New("user exists")
	}

	return true, nil
}

//as its name suggests
func GetUserByUsername(username string) *model.User {
	db, ctx := MongoDbConnect()

	users := db.Collection("users")
	var user model.User
	err := users.FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)

	if err != nil {
		fmt.Println("Couldn't get user: ", err)
	}

	return &user
}

//as its name suggests
func GetUserById(id interface{}) *model.User {
	db, ctx := MongoDbConnect()
	users := db.Collection("users")

	var user model.User
	users.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user)

	return &user
}

func randomUser() *model.User {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	RandomUsername := string(b) + "Test"
	randomEmail := string(b) + "Email"
	randomTel := string(b) + "Tel"
	testUser := &model.User{Username: RandomUsername, Email: randomEmail, Telephone: randomTel, Password: "TestPassword", Name: "TestName", Surname: "TestSurname"}

	return testUser
}

//updates user
func UpdateUser(user model.User) error {
	db, ctx := MongoDbConnect()

	users := db.Collection("users")

	update := bson.D{
		{"$set", bson.D{
			{Key: "username", Value: user.Username},
			{Key: "password", Value: user.Password},
			{Key: "email", Value: user.Email},
			{Key: "telephone", Value: user.Telephone},
			{Key: "name", Value: user.Name},
			{Key: "surname", Value: user.Surname},
			{Key: "isValidated", Value: false},
		}},
	}

	_, err := users.UpdateOne(ctx, bson.D{{Key: "_id", Value: user.ID}}, update)

	return err
}

//Checks if credentials of the user are valid
func IsValidCredentials(user model.User) bool {
	db, ctx := MongoDbConnect()

	users := db.Collection("users")

	var dbUser model.User

	users.FindOne(ctx, bson.D{{Key: "username", Value: user.Username}}).Decode(&dbUser)

	if dbUser.Username != "" {
		val, _ := IsValidPassword(user.Password, dbUser.Password)
		if val {
			return true
		}

	}

	users.FindOne(ctx, bson.D{{Key: "email", Value: user.Email}}).Decode(&dbUser)

	if dbUser.Email != "" {
		val, _ := IsValidPassword(user.Password, dbUser.Password)
		if val {
			return true
		}
	}

	users.FindOne(ctx, bson.D{{Key: "telephone", Value: user.Telephone}}).Decode(&dbUser)

	if dbUser.Telephone != "" {
		val, _ := IsValidPassword(user.Password, dbUser.Password)
		if val {
			return true
		}
	}

	return false
}

func DeleteUserByUsername(username string) bool {
	db, ctx := MongoDbConnect()

	users := db.Collection("users")

	_, err := users.DeleteOne(ctx, bson.D{{Key: "username", Value: username}})
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

//used to validate users email, tel etc
func ValidateUser(username string) error {
	db, ctx := MongoDbConnect()

	user := db.Collection("users")

	update := bson.D{
		{"$set", bson.D{{"isValidated", true}}},
	}

	result, err := user.UpdateOne(ctx, bson.D{{Key: "username", Value: username}}, update)

	if err != nil {
		fmt.Println("here: ", err)
		return err
	}

	fmt.Println("Update Result ", result)

	return nil
}

func InsertToken(user model.User, tokentring string) error {
	db, ctx := MongoDbConnect()

	tokens := db.Collection("tokens")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"username", user.Username}}
	update := bson.D{{"$set", bson.D{{"token", tokentring}}}}

	result, err := tokens.UpdateOne(ctx, filter, update, opts)

	fmt.Println(result)
	return err
}

func DeleteToken(username string, tokenString string) {
	db, ctx := MongoDbConnect()
	tokens := db.Collection("tokens")
	_, err := tokens.DeleteOne(ctx, bson.D{{Key: "username", Value: username}})

	if err != nil {
		fmt.Println("DELETING TOKEN :", err)
	}

}

func GetToken(username string) string {
	db, ctx := MongoDbConnect()

	tokens := db.Collection("tokens")
	user := struct {
		Username string
		Token    string
	}{"", ""}
	filter := bson.D{{Key: "username", Value: username}}
	err := tokens.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return ""
	}

	return user.Token

}
