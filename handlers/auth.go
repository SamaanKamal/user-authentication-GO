package handlers

import (
    "net/http"
    "authentication-app/config"
    "authentication-app/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "fmt"
)

var jwtSecret = []byte("your-secret-key")

func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Printf("password for input : %v \n", input.Password)
    fmt.Printf("password for input : %v \n", input.Email)


    // Hash password
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    input.Password = string(hashedPassword)

    result := config.DB.Create(&input)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    config.DB.Where("email = ?", input.Email).First(&user)
    fmt.Printf("password for input : %v \n", input.Password)
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    fmt.Printf("hashed password for input : %v \n", string(hashedPassword))
    fmt.Printf("hashed stored password for  : %v \n", user.Password)

    
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Create JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenString, _ := token.SignedString(jwtSecret)

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
