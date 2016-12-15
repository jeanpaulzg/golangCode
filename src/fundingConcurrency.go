package funding

type Fund struct {
    // balance is unexported (private), because it's lowercase
    balance int
    }
    
// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund{
    //We can return a pointer to a new struct without worring about
    // whether it's on the stack or heap: Go figur that out for us
    return &Fund{balance:initialBalance}
    }
    
//Methods sstart with a receiver, in this case a Fund pointer
func (f *Fund) Balance() int{
    return f.balance
    }
    
func (f *Fund) Withdraw(amount int){
    f.balance -= amount
    }
    
