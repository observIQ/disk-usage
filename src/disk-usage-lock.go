package main
import (
    "os"
    "fmt"
)


func createLock(x string) {
	var _, err = os.Create(x)
	if err != nil {
		fmt.Println("Failed to create lock file")
		fmt.Println(err.Error())
	}
}


func removeLock(x string)  {
	err := os.Remove(x)
	if err != nil {
		fmt.Println("Failed to remove lockfile")
	} else {
		fmt.Println("Lock removed")
	}
}


func lockExists(x string) bool {
  _, err := os.Stat(x)
  if err == nil {
    return true
  }
  return false
}
