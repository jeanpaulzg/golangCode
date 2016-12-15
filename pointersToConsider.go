package main



type User struct {
  Name string
}

func main() {
  u := &User{Name: "Leto"}
  println(u.Name)
  Modify(&u)
  println(u.Name)
}

func Modify(u **User) {
  *u = &User{Name: "Paul"}//This is still not pass-by-reference. But, since you are able to dereference, you can achieve the same thing.
}