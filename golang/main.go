package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_study")
	dsn := "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, gormerr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if gormerr != nil {
		for i := 0; i < 5; i++ {
			db, gormerr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if gormerr == nil {
				fmt.Println("数据库连接成功")
				break
			}

			fmt.Printf("数据库连接失败（尝试 %d）：%v\n", i+1, gormerr)
			time.Sleep(200) // 暂停一段时间后再重试
		}
	}

	if gormerr != nil {
		panic(gormerr)
	}
	//defer db.Close()

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "你好呀！")
	})
	//增
	router.POST("/createUser", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		createrr := _createUser(db, name, email)
		if createrr != nil {
			fmt.Println("Error creating user:", createrr)
			c.String(500, "Error creating user")
			return
		}
		fmt.Println("User created successfully!")
		c.String(200, "User created")
	})
	//查
	router.GET("/getUserByID/:user_id/", func(c *gin.Context) {
		id := c.Param("user_id")
		name, email, usererr := _getUserByID(db, id)
		if usererr != nil {
			fmt.Println("Error getting user:", usererr)
			c.String(500, "Error geettting user")
			return
		}
		fmt.Println("User getted successfully!")
		c.JSON(http.StatusOK, gin.H{"name": name, "email": email})
	})
	//改
	router.POST("/getUserByID/", func(c *gin.Context) {
		id := c.PostForm("id")
		name := c.PostForm("name")
		email := c.PostForm("email")
		fmt.Println("id=", id, "name=", name, "email=", email)
		updateerr := _updateUser(db, id, name, email)
		if updateerr != nil {
			fmt.Println("Error updating user:", updateerr)
			c.String(500, "Error updating user")
			return
		}
		fmt.Println("User updated successfully!")
		c.JSON(http.StatusOK, gin.H{"id": id, "name": name, "email": email})
	})
	//删
	router.DELETE("/deleteUserByID/", func(c *gin.Context) {
		id := c.Query("id")
		deleteerr := _deleteUserByID(db, id)
		if deleteerr != nil {
			fmt.Println("Error deleting user:", deleteerr)
			c.String(500, "Error deleting user")
			return
		}
		fmt.Println("User deleted successfully!")
		c.String(http.StatusOK, "User deleted successfully!")
	})

	router.Run(":8090")
}

type User struct {
	Id    int
	Name  string
	Email string
}

func _createUser(db *gorm.DB, name, email string) error {
	fmt.Println(name, email)
	//stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")

	user := User{
		Name:  name,
		Email: email,
	}

	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	fmt.Println("User created:", user)
	return nil
}

func _getUserByID(db *gorm.DB, id string) (string, string, error) {
	//var name, email string
	//err := db.QueryRow("SELECT name, email FROM users WHERE id = ?", id).Scan(&name, &email)
	var user User
	err := db.Where("id = ?", id).Select("name, email").Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return "", "", err

	}
	return user.Name, user.Email, nil
}

func _updateUser(db *gorm.DB, id, name, email string) error {
	//stmt, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	changes := make(map[string]interface{})
	if name != "" {
		changes["name"] = name
	}
	if email != "" {
		changes["email"] = email
	}
	if len(changes) == 0 {
		return errors.New("no changes to update") // 或者选择返回 nil，这取决于你的业务逻辑
	}

	result := db.Model(&User{}).Where("id = ?", id).Updates(changes)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found or no changes made")
	}
	return nil
}

func _deleteUserByID(db *gorm.DB, id string) error {
	deleaterr := db.Where("id = ?", id).Delete(&User{}).Error
	if deleaterr != nil {
		return deleaterr
	}
	return nil
}
