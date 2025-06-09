package schema

import (
    "errors"
    "time"

    "ai-ops-assistant/internal/auth"
    "ai-ops-assistant/internal/db"
    "ai-ops-assistant/internal/models"
    "github.com/google/uuid"
    "github.com/graphql-go/graphql"
    "golang.org/x/crypto/bcrypt"
)

var LoginField = &graphql.Field{
    Type: graphql.String,
    Args: graphql.FieldConfigArgument{
        "email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
        "password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
    },
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        email := p.Args["email"].(string)
        password := p.Args["password"].(string)

        var user models.User
        if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
            return nil, err
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
            return nil, errors.New("invalid credentials")
        }

        return auth.GenerateJWT(user.ID.String())
    },
}

var SignupField = &graphql.Field{
    Type: graphql.String, // return JWT token
    Args: graphql.FieldConfigArgument{
        "email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
        "password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
    },
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        email := p.Args["email"].(string)
        password := p.Args["password"].(string)

        hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            return nil, err
        }

        user := models.User{
            ID:        uuid.New(),
            Email:     email,
            Password:  string(hashed),
            CreatedAt: time.Now(),
        }

        if err := db.DB.Create(&user).Error; err != nil {
            return nil, err
        }

        return auth.GenerateJWT(user.ID.String())
    },
}

var MeField = &graphql.Field{
    Type: graphql.NewObject(graphql.ObjectConfig{
        Name: "Me",
        Fields: graphql.Fields{
            "id":        &graphql.Field{Type: graphql.String},
            "email":     &graphql.Field{Type: graphql.String},
            "createdAt": &graphql.Field{Type: graphql.String},
        },
    }),
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        userID, ok := p.Context.Value("userID").(string)
        if !ok {
            return nil, errors.New("unauthorized")
        }

        var user models.User
        if err := db.DB.First(&user, "id = ?", userID).Error; err != nil {
            return nil, err
        }

        return user, nil
    },
}