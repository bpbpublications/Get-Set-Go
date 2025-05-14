switch day {
case "Monday":
    fmt.Println("Start of the work week")
case "Friday":
    fmt.Println("End of the work week")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Midweek")
}
