// Get the current hour
currentTime := time.Now()
hour := currentTime.Hour()

// greeting based on the current hour
if hour >= 5 && hour < 12 {
 fmt.Println("Good Morning!")
} else if hour >= 12 && hour < 17 {
 fmt.Println("Good Afternoon!")
} else if hour >= 17 && hour < 21 {
 fmt.Println("Good Evening!")
} else {
 fmt.Println("Good Night!")
}
