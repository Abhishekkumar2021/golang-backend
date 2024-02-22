package models

import (
	"context"

	"github.com/Abhishekkumar2021/golang-backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollectionRef *mongo.Collection

func init()  {
	userCollectionRef = db.Database.Collection("users")
}

// User model and related functions
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}


func AddUser(user User) (User, error) {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)
	// insert the user
	insertResult, err := userCollectionRef.InsertOne(context.Background(), user)
	if err != nil {
		return user, err
	}
	// Update the user object with the inserted ID, the following line will convert the InsertedID to an ObjectID
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return user, err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := userCollectionRef.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := userCollectionRef.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return user, err
}

func GetUserByID(id string) (User, error) {
	var user User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := userCollectionRef.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	return user, err
}

func GetAllUsers() ([]User, error) {
	var users []User
	cursor, err := userCollectionRef.Find(context.Background(), bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, err
}

func PatchUser(id string, user User) (User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	// Find the current user
	currentUser, err := GetUserByID(id)
	if err != nil {
		return user, err
	}
	
	// If the password is not updated, use the current password
	if user.Password == ""  {
		user.Password = currentUser.Password
	} else {
		// hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return user, err
		}
		user.Password = string(hashedPassword)
	}

	// Now merge the changes
	if user.Username == "" {
		user.Username = currentUser.Username
	}
	if user.Email == "" {
		user.Email = currentUser.Email
	}
	updateResult, err := userCollectionRef.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": user})
	if err != nil {
		return user, err
	}
	if updateResult.ModifiedCount == 0 {
		return user, mongo.ErrNoDocuments
	}
	user.ID = objID
	return user, err
}

func DeleteUser(id string) (User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	// Find the current user
	user, err := GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	deleteResult, err := userCollectionRef.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return User{}, err
	}
	if deleteResult.DeletedCount == 0 {
		return User{}, mongo.ErrNoDocuments
	}
	return user, err
}