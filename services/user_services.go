package services

import (
    "errors"
    "example/moviecrud/models"
    "golang.org/x/crypto/bcrypt"
    "math/rand"
    "strconv"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

func CreateUser(user models.User) (models.User, error) {
    hashed, err := HashPassword(user.Password)
    if err != nil {
        return user, err
    }

    user.ID = strconv.Itoa(rand.Intn(1000000))
    user.Password = hashed

    models.Users = append(models.Users, user)
    return user, nil
}

func GetAllUsers() []models.User {
    return models.Users
}

func GetUserByID(id string) (models.User, error) {
    for _, u := range models.Users {
        if u.ID == id {
            return u, nil
        }
    }
    return models.User{}, errors.New("user not found")
}

func UpdateUser(id string, newData models.User) (models.User, error) {
    for i, u := range models.Users {
        if u.ID == id {

            // اگر پسورد جدید داده شده بود، هش کن
            if newData.Password != "" {
                hashed, err := HashPassword(newData.Password)
                if err != nil {
                    return newData, err
                }
                newData.Password = hashed
            } else {
                newData.Password = u.Password
            }

            newData.ID = id
            models.Users[i] = newData
            return newData, nil
        }
    }
    return models.User{}, errors.New("user not found")
}

func DeleteUser(id string) error {
    for i, u := range models.Users {
        if u.ID == id {
            models.Users = append(models.Users[:i], models.Users[i+1:]...)
            return nil
        }
    }
    return errors.New("user not found")
}
