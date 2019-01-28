package main
import (
		"fmt"
		"github.com/gin-contrib/cors"                        // Why do we need this package?
		"github.com/gin-gonic/gin"
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/sqlite"           // If you want to use mysql or any other db, replace this line
       )

var db *gorm.DB                                         // declaring the db globally
var err error
var loguser string
type Login struct {
	Name string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		City string `json:"city"`
}

type Ques struct {
	Genre string `json:"genre"`
		Question string `json:"question"`
		Opt1 string `json:"opt1"`
		Opt2 string `json:"opt2"`
		Opt3 string `json:"opt3"`
}

type Answers struct {
	Ans1 string `json:"ans1"`
		Ans2 string `json:"ans2"`
		Ans3 string `json:"ans3"`
		Ans4 string `json:"ans4"`
		Ans5 string `json:"ans5"`
		Ans6 string `json:"ans6"`
		Ans7 string `json:"ans7"`
		Ans8 string `json:"ans8"`
}
type HighScores struct {
	Username string `json:"username"`
		DN int `json:"dn"`
		CG int `json:"cg"`
		IP int `json:"ip"`
		USP int `json:"usp"`
}

type Score struct {
	Score int `json:"score"`
}

var dnhscore int
var dnscore int
var iphscore int
var ipscore int
var usphscore int
var uspscore int
var cghscore int
var cgscore int
var current string
func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
		if err != nil {
			fmt.Println(err)
		}
	current=""
		defer db.Close()
		db.AutoMigrate(&Login{}, &Ques{} , &Score{}, &HighScores{})
		/*dn:=Ques{Genre: "dn", Question: "What is the price of Shinigami eyes? ",Opt1: "Ownership Of DeathNote",Opt2: "Half Of Remaining Lifespan",Opt3: "A living Soul"}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "dn", Question: "How many times was Misa lifespan cut to half?",Opt1: "Three",Opt2: "Two",Opt3: "One"}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "dn", Question: "What does Teru Mikami say while punishing someone?",Opt1: "Delete",Opt2: "Erase",Opt3: "Destroy"}   //1
		  db.Create(&dn)
		  dn=Ques{Genre: "dn", Question: "Who is the most frequent shinigami in the anime?",Opt1: "Ryuk",Opt2: "Rem",Opt3: "Gelus"}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "dn", Question: "Who are the member's of the taskforce?",Opt1: "Mogi",Opt2: "Matsuda",Opt3: "Aizawa"}   //1 2 3
		  db.Create(&dn)
		  dn=Ques{Genre: "dn", Question: "Who among these are Shinigamis?",Opt1: "Rem" ,Opt2: "Mogi",Opt3: "Ryuk"}   //1 3
		  db.Create(&dn)
		  dn=Ques{Genre: "dn", Question: "Who among these does Lite Yagami date?",Opt1: "Sachiko",Opt2: "Lisa Amane",Opt3: "Kiyomi Takada",}   //2 3
		  db.Create(&dn)
		  dn=Ques{Genre: "dn", Question: "Who all are part of the Yagami family?",Opt1: "Sachiko",Opt2: "Lisa",Opt3: "Sayu"}   //1 3
		  db.Create(&dn)
		  ip:=Ques{Genre: "ip", Question: "Name The First Vice-President Of India? ",Opt1: "Rajendra Awasthi",Opt2: "Rajendra Prasad",Opt3: "Madanmohan Malviya"}   //2
		  db.Create(&ip)
		  ip=Ques{Genre: "ip", Question: "Current CM of Uttar Pradesh?",Opt1: "Akhilesh Yadav",Opt2: "Yogi Adityanath",Opt3: "Mayawati"}   //2
		  db.Create(&ip)
		  ip=Ques{Genre: "ip", Question: "Who is the Founder Of ShivSena?",Opt1: "Bal Thakre",Opt2: "Udhav Thakre",Opt3: "Raj Thakre"}   //1
		  db.Create(&ip)
		  ip=Ques{Genre: "ip", Question: "Current Lok-Sabha Speaker?",Opt1: "Aashish Gogoi",Opt2: " Sumitra Mahajan ",Opt3: "Smriti Irani"}   //2
		  db.Create(&ip)
		  dn=Ques{Genre: "ip", Question: "Which among these have a BJP state-government??",Opt1: " Rajasthan ",Opt2: " Chhattisgarh ",Opt3: "Assam"}   //1 2 3
		  db.Create(&dn)
		  dn=Ques{Genre: "ip", Question: "Who among these is an MP?",Opt1: " Sachin Tendulkar " ,Opt2: " Rahul Dravid ",Opt3: " MaryKom "}   //1 3
		  db.Create(&dn)
		  dn=Ques{Genre: "ip", Question: "Who among these are Central Cabinet Ministers?",Opt1: " Navjot Singh Sidhu ",Opt2: " Arun Jaitley ",Opt3: " Nitin Gadkari ",}   //2 3
		  db.Create(&dn)
		  dn=Ques{Genre: "ip", Question: "Who among these are Congress Leaders?",Opt1: " Rahul Gandhi ",Opt2: " Sushma Swaraj ",Opt3: " Digvijay Singh "}   //1 3
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "Who is zero ? ",Opt1: " Suzaku kururugi ",Opt2: " Lelouch Lamperouge ",Opt3: "Cc"}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "What was the name of the final plan of zero?",Opt1: " Zeros revenge ",Opt2: " Zero requiem ",Opt3: " Zero won "}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "What was the name of lelouch’s sister ?",Opt1: "Nunnally",Opt2: "Annie",Opt3: "Katara"}   //1
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "What was Japan  called ?",Opt1: "Nippon",Opt2: "Area 11",Opt3: "District 13"}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "Which of these are members of the japan liberation front ?",Opt1: "Katase tatewaki ",Opt2: "Josui kusakabe ",Opt3: "Nagisa chiba "}   //1 2 3
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "Whore the members of the Black knight?",Opt1: "Lalouch lamperouge" ,Opt2: "Suzaku kururugi",Opt3: "Kallen kozuki"}   //1 3
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "Which of these belong to the ruling family of the holy britanian empire ?",Opt1: "Suzaku kurungi",Opt2: "Lalouch lamperouge",Opt3: "Charles z Britainia"}   //2 3
		  db.Create(&dn)
		  dn=Ques{Genre: "cg", Question: "Which of these are super powers in the world of code geas ?",Opt1: "Europian union",Opt2: "American empire ",Opt3: "Chinese Federation"}   //1 3
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: " Who drafted the US Declaration of Independence ?",Opt1: " George Washington ",Opt2: " Thomas Jefferson  ",Opt3: " Abraham Lincoln "}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: "What is the nickname for the old regulations requiring racial segregation ?",Opt1: " Old crow laws  ",Opt2: " Jim Crow laws  ",Opt3: " Blue laws  "}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: "Which document does the Fourth of July commemorate ?",Opt1: "Declaration of Independence",Opt2: "Articles of confederation",Opt3: "Gettysburg address "}   //1
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: "What kind of govt does the United States have  ?",Opt1: "Democracy",Opt2: "Republic",Opt3: "Gerontocracy"}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: "Which of the following have been the US vice presidents  ?",Opt1: " Mike Pence ",Opt2: " Joe Biden ",Opt3: " Dick Cheney "}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: "Which of these have been the Us Chief Justice ?",Opt1: " John Roberts ",Opt2: " George Clooney ",Opt3: " John Jay "}   //2
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: " The states in which marijuana is legal ?",Opt1: "Okhlahoma",Opt2: "California",Opt3: "Washington DC "}   //1
		  db.Create(&dn)
		  dn=Ques{Genre: "usp", Question: "Who among these have been US Presidents ?",Opt1: "George Bush",Opt2: "Alama",Opt3: "Barack Obama"}   //2
		  db.Create(&dn)*/

		r := gin.Default()
		r.GET("/view", GetPeople)  
		r.DELETE("/people/:username", DeletePerson) 
		r.POST("/login", LoginFunc)
		r.POST("/register", RegisterFunc)

		r.GET("/LoggedIn/anime",AnimeScore)
		r.GET("/LoggedIn/anime/deathnote",DNFunc)
		r.GET("/LoggedIn/anime/deathnote/score",DNScoreFun)
		r.POST("/LoggedIn/anime/deathnote/score",DNScoreFunc)

		r.GET("/LoggedIn/anime/codegeass",CGFunc)
		r.GET("/LoggedIn/anime/codegeass/score",CGScoreFun)
		r.POST("/LoggedIn/anime/codegeass/score",CGScoreFunc)

		r.GET("/LoggedIn/politics",PoliticsScore)
		r.GET("/LoggedIn/politics/indian",IPFunc)
		r.GET("/LoggedIn/politics/indian/score",IPScoreFun)
		r.POST("/LoggedIn/politics/indian/score",IPScoreFunc)

		r.GET("/LoggedIn/politics/us",USFunc)
		r.GET("/LoggedIn/politics/us/score",USScoreFun)
		r.POST("/LoggedIn/politics/us/score",USScoreFunc)

		r.GET("/scoreboard/anime",AnimeBoard)
		r.GET("/scoreboard/politics",PoliticsBoard)
		r.Use((cors.Default()))
		r.Run(":8080")

}
func DeletePerson(c *gin.Context) {
username := c.Params.ByName("username")
		  var person Login
		  d := db.Where("username = ?", username).Delete(&person)
		  fmt.Println(d)
		  c.Header("access-control-allow-origin", "*")
		  c.JSON(200, gin.H{"username #" + username: "deleted"})
}
func GetPeople(c *gin.Context) {
	var people []Login
		if err := db.Find(&people).Error; err != nil {
			c.AbortWithStatus(404)
				fmt.Println(err)
		} else {
			c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
				c.JSON(200, people)
		}
}

func RegisterFunc(c *gin.Context) {
	var login Login
		c.BindJSON(&login)
		d:=db.Where("username = ?",login.Username).Find(&login).RecordNotFound()
		if d {
high:= HighScores{Username: login.Username, DN: 0, CG: 0, IP:0, USP:0}
     db.Create(&login)
	     db.Create(&high)
	     c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
	     c.JSON(200, login)
		} else {
			c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
				fmt.Println("Already Exits")
				c.AbortWithStatus(404)
		}
}

func LoginFunc(c *gin.Context) {
	var login Login
		c.BindJSON(&login)
		a:=login.Username
		b:=login.Password
		var person Login
		err:=db.Where("Username = ? AND Password = ?",a,b).Find(&person).Error
		if err != nil{
			c.AbortWithStatus(404)
				fmt.Println(err)
		} else {
			current=a
				fmt.Println("Logged In")
				c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
				c.JSON(200,login)
		}
}
func AnimeScore(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a HighScores
		db.Where("Username = ?",current).Find(&a)
		c.JSON(200, a)
}

func PoliticsScore(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a HighScores
		db.Where("Username = ?",current).Find(&a)
		c.JSON(200, a)
}

func DNFunc(c *gin.Context) {
	var dn []Ques
		err:=db.Where("Genre = ?","dn").Find(&dn).Error
		if err != nil {
			c.AbortWithStatus(404)
				fmt.Println(err)
		} else {
			db.Where("Genre = ?","dn").Find(&dn)
				c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
				c.JSON(200, dn)
		}
}
func CGFunc(c *gin.Context) {
	var cg []Ques
		err:=db.Where("Genre = ?","cg").Find(&cg).Error
		if err != nil {
			c.AbortWithStatus(404)
				fmt.Println(err)
		} else {
			db.Where("Genre = ?","cg").Find(&cg)
				c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
				c.JSON(200, cg)
		}
}

func IPFunc(c *gin.Context) {
	var ip []Ques
		err:=db.Where("Genre = ?","ip").Find(&ip).Error
		if err != nil {
			c.AbortWithStatus(404)
				fmt.Println(err)
		} else {
			db.Where("Genre = ?","ip").Find(&ip)
				c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
				c.JSON(200, ip)
		}
}

func USFunc(c *gin.Context) {
	var usp []Ques
		err:=db.Where("Genre = ?","usp").Find(&usp).Error
		if err != nil {
			c.AbortWithStatus(404)
				fmt.Println(err)
		} else {
			db.Where("Genre = ?","usp").Find(&usp)
				c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
				c.JSON(200, usp)
		}
}

func DNScoreFun(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a Score
		a.Score=dnscore
		c.JSON(200, a)
}

func CGScoreFun(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a Score
		a.Score=cgscore
		fmt.Println(a.Score)
		c.JSON(200, a)
}

func IPScoreFun(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a Score
		a.Score=ipscore
		c.JSON(200, a)
}

func USScoreFun(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a Score
		a.Score=uspscore
		c.JSON(200, a)
}

func DNScoreFunc(c *gin.Context) {
	var dn Answers
		c.BindJSON(&dn)
		dnscore=0
		if dn.Ans1 == "2" {
			dnscore++
		}
	if dn.Ans2 == "2" {
		dnscore++
	}
	if dn.Ans3 == "1" {
		dnscore++
	}
	if dn.Ans4 == "2" {
		dnscore++
	}
	if dn.Ans5 == "1 2 3" {
		dnscore++
	}
	if dn.Ans6 == "1 3" {
		dnscore++
	}
	if dn.Ans7 == "2 3" {
		dnscore++
	}
	if dn.Ans8 == "1 3" {
		dnscore++
	}
	var hs HighScores;
	db.Where("Username = ?",current).Find(&hs)
		dnhscore=hs.DN
		if dnscore > dnhscore {
			dnhscore=dnscore
				db.Model(&hs).Where("Username=?",current).Update("dn",dnhscore)
		}
	c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
		c.JSON(200,hs)
}


func CGScoreFunc(c *gin.Context) {
	var dn Answers
		c.BindJSON(&dn)
		cgscore=0
		if dn.Ans1 == "2" {
			cgscore++
		}
	if dn.Ans2 == "2" {
		cgscore++
	}
	if dn.Ans3 == "1" {
		cgscore++
	}
	if dn.Ans4 == "2" {
		cgscore++
	}
	if dn.Ans5 == "1 2 3" {
		cgscore++
	}
	if dn.Ans6 == "1 3" {
		cgscore++
	}
	if dn.Ans7 == "2 3" {
		cgscore++
	}
	if dn.Ans8 == "1 3" {
		cgscore++
	}
	var hs HighScores;
	db.Where("Username = ?",current).Find(&hs)
		cghscore=hs.CG
		if cgscore > cghscore {
			cghscore=cgscore
				db.Model(&hs).Where("Username=?",current).Update("cg",cghscore)
		}
	c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
		c.JSON(200, dn)
}

func IPScoreFunc(c *gin.Context) {
	var dn Answers
		c.BindJSON(&dn)
		ipscore=0
		if dn.Ans1 == "2" {
			ipscore++
		}
	if dn.Ans2 == "2" {
		ipscore++
	}
	if dn.Ans3 == "1" {
		ipscore++
	}
	if dn.Ans4 == "2" {
		ipscore++
	}
	if dn.Ans5 == "1 2 3" {
		ipscore++
	}
	if dn.Ans6 == "1 3" {
		ipscore++
	}
	if dn.Ans7 == "2 3" {
		ipscore++
	}
	if dn.Ans8 == "1 3" {
		ipscore++
	}
	var hs HighScores;
	db.Where("Username = ?",current).Find(&hs)
		iphscore=hs.IP
		if ipscore > iphscore {
			iphscore=ipscore
				db.Model(&hs).Where("Username=?",current).Update("ip",iphscore)
		}
	c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
		c.JSON(200, dn)
}

func USScoreFunc(c *gin.Context) {
	var dn Answers
		c.BindJSON(&dn)
		uspscore=0
		if dn.Ans1 == "2" {
			uspscore++
		}
	if dn.Ans2 == "2" {
		uspscore++
	}
	if dn.Ans3 == "1" {
		uspscore++
	}
	if dn.Ans4 == "2" {
		uspscore++
	}
	if dn.Ans5 == "1 2 3" {
		uspscore++
	}
	if dn.Ans6 == "1 3" {
		uspscore++
	}
	if dn.Ans7 == "2 3" {
		uspscore++
	}
	if dn.Ans8 == "1 3" {
		uspscore++
	}
	var hs HighScores;
	db.Where("Username = ?",current).Find(&hs)
		usphscore=hs.USP
		if uspscore > usphscore {
			usphscore=uspscore
				db.Model(&hs).Where("Username=?",current).Update("usp",usphscore)
		}
	c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
		c.JSON(200, dn)
}

func AnimeBoard(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a []HighScores
		db.Find(&a)
		c.JSON(200, a)
}

func PoliticsBoard(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
		var a []HighScores
		db.Find(&a)
		c.JSON(200, a)
}

