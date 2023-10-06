package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "gin_study/elmDemoAPI/docs"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type Dish struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

type User struct {
	gorm.Model
	Username string
	Password string
}

var cart []Dish
var dishes []Dish

var db *gorm.DB

// @title elm
// @version 1.0
// @description 简易系统
// @host 127.0.0.1:8080
// @BasePath /index
func main() {

	initDB()
	//这里没有使用mysql数据库，因为报错，8.0版本之上密码正确还登录不上去
	db.AutoMigrate(&User{})
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 注册页面
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// 处理注册请求
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 检查用户名是否已存在
		var existingUser User
		result := db.Where("username = ?", username).First(&existingUser)
		if result.Error == nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"Error": "用户名已存在",
			})
			return
		}

		// 创建新用户
		newUser := User{
			Username: username,
			Password: password,
		}
		db.Create(&newUser)

		// 注册成功
		c.HTML(http.StatusOK, "success.html", gin.H{
			"Username": username,
		})
	})

	// 显示登录页面
	r.GET("/l", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 处理登录请求
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 查询用户
		var user User
		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"Error": "用户名不存在",
			})
			return
		}

		// 验证密码
		if user.Password != password {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"Error": "密码不正确",
			})
			return
		}

		// 登录成功
		c.HTML(http.StatusOK, "success.html", gin.H{
			"Username": username,
		})

	})

	//显示主菜单
	r.GET("/", showMenu)
	//添加菜品
	r.GET("/add-dish", showAddDishForm)
	r.POST("/add-dish", addDish)
	r.GET("/delete-dish/:id", deleteDish)
	r.GET("/update-dish/:id", showUpdateDishForm)
	r.POST("/update-dish/:id", updateDish)
	r.GET("/cart", showCart)
	r.POST("/add-to-cart", addToCart)
	r.GET("/checkout", checkout)

	// 启动服务器
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server")
	}
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}

	// 迁移数据库
	db.AutoMigrate(&Dish{})
}

//对于界面的话，也不是我管的活了，有点乱，然后我还随意的加一些跳转....
//有的使用html submit 展示按钮能好看点，有的直接html跳转，界面有点丑   不过...应该关系不大

// 显示菜单
func showMenu(c *gin.Context) {
	//var dishes []Dish
	//这里出了一个小问题，真的出问题了是闹心死，先是细致的理解一遍运行的原理，
	//然后一个一个文件的去核对，为什么会出现两列00000000？？然后一度问不起百度了，
	//根本不知道怎么描述，气死我了，急死我了，问一个学长，一直不回我...
	//然后去理解这个关系后，就是发现，这个在菜单界面所创建的数据，为什么索引ID查不到
	//那么导致这个原因的我在想这个两个小模块之间的链接问题
	//然后顺着摸到要用什么载体去传递数据，然后就是定义的结构体变量，我一对比，发现我定义了两个var dishes []Dish
	//一个是局部的   在后文  有一个是全局的
	//然后注释掉局部的，试试，woc太爽了，就好使了，爽死我了
	db.Find(&dishes)
	c.HTML(http.StatusOK, "menu.html", gin.H{
		"dishes": dishes,
	})
}

// 显示添加菜品表单
func showAddDishForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_dish.html", nil)
}

// 添加菜品
func addDish(c *gin.Context) {
	name := c.PostForm("name")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	category := c.PostForm("category")

	dish := Dish{
		Name:     name,
		Price:    price,
		Category: category,
	}

	result := db.Create(&dish)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// 删除菜品
func deleteDish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := db.Delete(&Dish{}, id)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// 显示更新菜品表单
func showUpdateDishForm(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var dish Dish
	result := db.First(&dish, id)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "update_dish.html", gin.H{
		"dish": dish,
	})
}

// 更新菜品
func updateDish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var dish Dish
	result := db.First(&dish, id)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	dish.Name = c.PostForm("name")
	dish.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	dish.Category = c.PostForm("category")

	result = db.Save(&dish)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// 显示购物车
func showCart(c *gin.Context) {
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"cart": cart,
	})
}

// 添加到购物车
func addToCart(c *gin.Context) {
	// 从请求中获取菜品ID
	dishID := c.PostForm("dishID")

	// 根据菜品ID查找菜品
	var selectedDish Dish
	for _, dish := range dishes {
		if strconv.Itoa(int(dish.ID)) == dishID {
			selectedDish = dish
			break
		}
	}

	// 将菜品添加到购物车
	cart = append(cart, selectedDish)

	c.Redirect(http.StatusFound, "/cart")
}

// 结算购物车
func checkout(c *gin.Context) {
	// 计算总价、运费和总计
	totalPrice := 0.0
	for _, item := range cart {
		totalPrice += item.Price
	}
	shippingFee := 5.0
	totalPricePlusShippingFee := totalPrice + shippingFee

	c.HTML(http.StatusOK, "checkout.html", gin.H{
		"cart":                   cart,
		"totalPrice":             totalPrice,
		"shippingFee":            shippingFee,
		"totalPricePlusShipping": totalPricePlusShippingFee,
	})
}
