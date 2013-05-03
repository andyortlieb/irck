package main

import (
	//"github.com/bmizerany/pq"
	"fmt"
)

func stub(msg string){
	fmt.Println("STUB: ", msg)
}
func erro(msg string){
	fmt.Println("ERROR: ", msg)
}

func initUsers() *[]User{

	stub("Set up some users")
	guyA := User{username: "dingolvrA"}
	idtA1 := guyA.AddIdentity(
		"irc",
		"dv.opasc.net:6667",
		"dingolvrA1",
		true,
	)
	idtA2 := guyA.AddIdentity(
		"irc",
		"dv.opasc.net:6667",
		"dingolvrA2",
		true,
	)
	fmt.Println(20, idtA1.channels, guyA.identities[0].channels)
	idtA1.AddChannel("#dingolove", true)
	fmt.Println(21, idtA1.channels, guyA.identities[0].channels)
	idtA2.AddChannel("#dingolove", true)
	idtA1.Connect()
	idtA2.Connect()

	guyB := User{username: "dingolvrB"}
	idtB1 := guyB.AddIdentity(
		"irc",
		"dv.opasc.net:6667",
		"dingolvrB1",
		true,
	)
	idtB1.AddChannel("#dingolove", true)
	idtB1.Connect()

/*	// Just for now... loop with readstring until we know how to be a real good daemon.

	br := bufio.NewReaderSize(os.Stdin, 512)
	for {
		msg, err := br.ReadString('\n')
		if err != nil {
			break;
		}
		if msg[:5] == "/quit" {
			return
		}
		fmt.Println("#dingolove/dingolvr: ", msg)
		idtA1.connection.Privmsg("AcidTrucks", msg+"\n")
		idtA2.connection.Privmsg("#dingolove", msg+"\n")
		//irccon.Privmsg("#dingolove", msg+"\n")
	}	
*/
	return &[]User{
		guyA,
		guyB,
	}
}

func main(){
	stub("Starting irckd")
	initHttp(initUsers())
}

