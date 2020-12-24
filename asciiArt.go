package asciiArt

import "fmt"

func asciiArt() {
	asciiArt :=
		`
	 ____      _   _  _____    _   _   _   ____          ____      ____    U  ___ u    _    U _____ u   ____   _____          _       ____      _____   
  U |  _"\ uU |"|u| ||_ " _|  |'| |'| |"| / __"| u     U|  _"\ uU |  _"\ u  \/"_ \/ U |"| u \| ___"|/U /"___| |_ " _|     U  /"\  uU |  _"\ u  |_ " _|  
   \| |_) |/ \| |\| |  | |   /| |_| |\|_|<\___ \/      \| |_) |/ \| |_) |/  | | | |_ \| |/   |  _|"  \| | u     | |        \/ _ \/  \| |_) |/    | |    
	|  _ <    | |_| | /| |\  U|  _  |u    u___) |       |  __/    |  _ <.-,_| |_| | |_| |_,-.| |___   | |/__   /| |\       / ___ \   |  _ <     /| |\   
	|_| \_\  <<\___/ u |_|U   |_| |_|     |____/>>      |_|       |_| \_\\_)-\___/ \___/-(_/ |_____|   \____| u |_|U      /_/   \_\  |_| \_\   u |_|U   
	//   \\_(__) )(  _// \\_  //   \\      )(  (__)     ||>>_     //   \\_    \\    _//      <<   >>  _// \\  _// \\_      \\    >>  //   \\_  _// \\_  
   (__)  (__)   (__)(__) (__)(_") ("_)    (__)         (__)__)   (__)  (__)  (__)  (__)     (__) (__)(__)(__)(__) (__)    (__)  (__)(__)  (__)(__) (__) 
`
	fmt.Println(asciiArt)
	//strings
	var favoriteSnack string

	favoriteSnack = "Hot cheetos"

	fmt.Println("My favorite snack is " + favoriteSnack)

	//assigning variables

	daysOnVacation := 7

	var hoursInDay = 24

	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")

	// Define cupsOfCoffeeConsumed here
	var cupsOfCoffeeConsumed int
	cupsOfCoffeeConsumed = 5
	fmt.Println(cupsOfCoffeeConsumed)

	coolSneakers := 65.99
	niceNecklace := 45.50

	//Tax calculation exercise
	// Define taxCalculation here
	var taxCalculation float64
	// Add coolSneakers to taxCalculation
	taxCalculation += coolSneakers
	// Add niceNecklace to taxCalculation
	taxCalculation += niceNecklace
	// Compute the NYC sales tax
	// 8.875% of the purchase here:
	taxCalculation *= .08875
	// Uncomment this line for a receipt!
	fmt.Println("Purchase of", coolSneakers+niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers+niceNecklace+taxCalculation)

	// Define magicNum and powerLevel below:
	var magicNum, powerLevel int32
	magicNum = 2048
	powerLevel = 9001

	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	// Define amount and unit below:
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")

	//%v = Verb

	animal1 := "cat"
	animal2 := "dog"

	// Add your code below:
	fmt.Printf("Are you a %v or a %v person", animal1, animal2)

	floatExample := 1.75
	// Edit the following Printf for the FIRST step
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// Edit the following Printf for the SECOND step
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// Edit the following Printf for the THIRD step
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)

	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	meditation := fmt.Sprintln(step1, step2)

	fmt.Print(meditation)

	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	// Add your code below:
	wish = fmt.Sprintf(template, pet)

	fmt.Println(wish)

	fmt.Println("What would you like for lunch?")

	// Add your code below:
	var food string
	fmt.Scan(&food)
	fmt.Printf("Sure, we can have %v for lunch.", food)

	heistReady := false
  
	
  if heistReady {
		fmt.Println("Let's go!")
 } else {
    fmt.Println("Act normal.")
  }

}	

vaultAmt := 2356468
	
// Add your code below:
if vaultAmt >= 200000 {
  fmt.Println("We're going to need more bags.")
}

}

rightTime := true
rightPlace := true

// Edit this condition for the FIRST checkpoint
if rightTime && rightPlace {
	fmt.Println("We're outta here!")
} else {
	fmt.Println("Be patient...")
}


enoughRobbers := false
enoughBags := true

// Edit this condition for the SECOND checkpoint
if enoughRobbers || enoughBags {
	fmt.Println("Grab everything!")
} else {
	fmt.Println("Grab whatever you can!")
}
}


readyToGo := true
  

if !readyToGo {
	fmt.Println("Start the car!")
} else {
	fmt.Println("What are we waiting for??")
}
}


amountStolen := 64650

if amountStolen > 1000000 {
	fmt.Println("We've hit the jackpot!")
} else if amountStolen >= 5000 {
fmt.Println("Think of all the candy we can buy!")
} else {
	fmt.Println("Why did we even do this?")
} 

name := "H. J. Simp"

// Add your switch statement below:
switch name {
case "Butch":
fmt.Println("Head to Robbers Roost.")
case "Bonnie":
fmt.Println("Stay put in Joplin")
default:
fmt.Println("Just hide!")
}



if success := true; success {
	fmt.Println("We're rich!")
} else {
	fmt.Println("Where did we go wrong?")
}

amountStolen := 50000
//numOfThieves := 5

switch numOfThieves := 5; numOfThieves {
case 1:
	fmt.Println("I'll take all $", amountStolen)
case 2:
  fmt.Println("Everyone gets $", amountStolen/2)
case 3:
	fmt.Println("Everyone gets $", amountStolen/3)
case 4:
  fmt.Println("Everyone gets $", amountStolen/4)
case 5:
	fmt.Println("Everyone gets $", amountStolen/5)
default:
	fmt.Println("There's not enough to go around...")
}

	// Edit amountLeft below: 
	amountLeft := rand.Intn(10000)

	fmt.Println(amountLeft)  
  
	fmt.Println("amountLeft is: ", amountLeft)
	
	  if amountLeft > 5000 {
		  fmt.Println("What should I spend this on?")
	} else {
	  fmt.Println("Where did all my money go?")
	}

	// Add your code below - changes number each time according to seconds
	rand.Seed(time.Now().UnixNano())
  
	amountLeft := rand.Intn(10000)
	
	fmt.Println("amountLeft is: ", amountLeft)
	
	  if amountLeft > 5000 {
		  fmt.Println("What should I spend this on?")
	} else {
	  fmt.Println("Where did all my money go?")
	}

	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)
	leftSafely := rand.Intn(5)
	
	if eludedGuards >= 50 {
	  fmt.Println("Looks like you've managed to make it past the guards. Good job, but remeber, this is the first step") 
	} else if isHeistOn = false { 
	  fmt.Println("Plan a better disguise next time?") 
	} else {
	  fmt.Println(isHeistOn)
	}
	openedVault := rand.Intn(100)
	if isHeistOn = true && openedVault >= 70 {
	  fmt.Println("Grad and GO!")
	  } else if isHeistOn := false; isHeistOn {
	  fmt.Println("Vault can't be opened")
	}
	
	if isHeistOn == true {
	  switch leftSafely {
		case isHeistOn := false:
		fmt.Println("Heist not on")
		case 1:
		isHeistOn = false {
		  fmt.Println("Turns out vault doors don't open from the inside")
		default:
		fmt.Println("Start the getaway car!")
		}
	if amtStole := 10000 + rand.Intn(1000000); amtStolen
	  }
	}
	}



}
