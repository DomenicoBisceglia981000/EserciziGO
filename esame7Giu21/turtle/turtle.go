package main
import(
    "bufio"
    "os"
    "errors"
    "fmt"
    "strconv"
)

func main(){
  scanner := bufio.NewScanner(os.Stdin)
  //CREAZIONE TARTARUGA
  var turtle Turtle
  turtle.x = 0
  turtle.y = 0
  turtle.dir = 'N'
  for scanner.Scan(){
    comandoInserito := scanner.Text()
    if comandoInserito == "stop"{
      break
      //FINE LETTURA COMANDI
    }else{
      var direction byte
      var passi int
      //Dopo aver provato svariate volte a leggere i comandi utilizzando fmt.Sscanf() senza risultati, ho dovuto utilizzare un metodo di lettura "rudimentale" ma efficace
      comando := string(comandoInserito[0])
      if comando == "S"{
        direction = comandoInserito[2]
      }else{
        passi, _ = strconv.Atoi(string(comandoInserito[2:]))
      }
      switch comando{
      case "F":
        //avanza
        forward(&turtle, passi)
      case "B":
        //indietreggia
        backward(&turtle, passi)
      case "S":
        //impostaDir
        setDir(&turtle, direction)
      default :
      //Caso in cui viene inserito un comando non valido
        fmt.Println("invalid command")
      }
    }
    fmt.Println(turtle.String())
  }
}

type Turtle struct{
  x, y int
  dir byte
}

func forward(turtle *Turtle, passi int){
  switch turtle.dir{
  case 'N':
    turtle.y = turtle.y + passi
  case 'S':
    turtle.y = turtle.y - passi
  case 'O':
    turtle.x = turtle.x - passi
  case 'E':
    turtle.x = turtle.x + passi
  }
}

func backward(turtle *Turtle, passi int){
  switch turtle.dir{
  case 'N':
    turtle.y = turtle.y - passi
  case 'S':
    turtle.y = turtle.y + passi
  case 'O':
    turtle.x = turtle.x + passi
  case 'E':
    turtle.x = turtle.x - passi
  }
}

func setDir(turtle *Turtle, dir byte)error{
  err := errors.New("non valid direction")
  if dir == 'E'{
    turtle.dir = 'E'
  }else if dir == 'N'{
    turtle.dir = 'N'
  }else if dir == 'S'{
    turtle.dir = 'S'
  }else if dir == 'O'{
    turtle.dir = 'O'
  }
  return err
}

func (turtle Turtle) String()string{
  stringa := fmt.Sprintf("(%d,%d) dir %s", turtle.x, turtle.y, string(turtle.dir))
  return stringa
}
