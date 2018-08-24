package main
import "fmt"
type Data struct {
  value string
  next *Data

}


type List struct {
	length int
	start  *Data
}


func (list *List ) setToTheList(data *Data ){
  if (list.length == 0){
    list.start = data;
  }else{
    exist := list.start;
    for exist.next != nil { //getter method code :==)))
      exist = exist.next;

    }
    exist.next = data
  }
  list.length++
}

func (list *List ) getFromList() { //For Deubgging no IDE :(

  for list.start != nil {
    fmt.Println(list.start.value)
    list.start = list.start.next;
  }

}
